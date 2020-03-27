package grpc

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	"google.golang.org/grpc/metadata"

	"github.com/jexia/maestro/instance"
	"github.com/jexia/maestro/logger"
	"github.com/jexia/maestro/refs"
	"github.com/jexia/maestro/schema"
	"github.com/jexia/maestro/specs"
	"github.com/jexia/maestro/specs/trace"
	"github.com/jexia/maestro/transport"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var proxy = &grpc.StreamDesc{
	ServerStreams: false,
	ClientStreams: false,
}

// NewCaller constructs a new HTTP caller
func NewCaller() transport.NewCaller {
	return func(ctx instance.Context) transport.Caller {
		return &Caller{
			ctx: ctx,
		}
	}
}

// Caller represents the caller constructor
type Caller struct {
	ctx instance.Context
}

// Name returns the name of the given caller
func (caller *Caller) Name() string {
	return "grpc"
}

// Dial constructs a new caller for the given host
func (caller *Caller) Dial(schema schema.Service, functions specs.CustomDefinedFunctions, opts schema.Options) (transport.Call, error) {
	logger := caller.ctx.Logger(logger.Transport)

	logger.WithFields(logrus.Fields{
		"service": schema.GetName(),
		"host":    schema.GetHost(),
	}).Info("Constructing new gRPC caller")

	options, err := ParseCallerOptions(opts)
	if err != nil {
		return nil, err
	}

	methods := make(map[string]*Method, len(schema.GetMethods()))

	for _, method := range schema.GetMethods() {
		methods[method.GetName()] = &Method{
			name: method.GetName(),
			fqn:  fmt.Sprintf("/%s.%s/%s", schema.GetPackage(), schema.GetName(), method.GetName()),
		}
	}

	result := &Call{
		ctx:     caller.ctx,
		logger:  logger,
		service: schema.GetName(),
		host:    schema.GetHost(),
		methods: methods,
		options: options,
	}

	return result, nil
}

// Method represents a service method
type Method struct {
	name string
	fqn  string
}

// GetName returns the method name
func (method *Method) GetName() string {
	return method.name
}

// References returns the available method references
func (method *Method) References() []*specs.Property {
	return make([]*specs.Property, 0)
}

// Call represents the HTTP caller implementation
type Call struct {
	ctx     instance.Context
	logger  *logrus.Logger
	service string
	host    string
	methods map[string]*Method
	options *CallerOptions
	mutex   sync.Mutex
	client  *grpc.ClientConn
}

// GetMethods returns the available methods within the HTTP caller
func (call *Call) GetMethods() []transport.Method {
	result := make([]transport.Method, 0, len(call.methods))

	for _, method := range call.methods {
		result = append(result, method)
	}

	return result
}

// GetMethod attempts to return a method matching the given name
func (call *Call) GetMethod(name string) transport.Method {
	for _, method := range call.methods {
		if method.GetName() == name {
			return method
		}
	}

	return nil
}

// Director returns a client connection and a outgoing context for the given method
func (call *Call) Director(ctx context.Context) (*grpc.ClientConn, error) {
	call.mutex.Lock()
	defer call.mutex.Unlock()

	if call.client != nil {
		return call.client, nil
	}

	conn, err := grpc.DialContext(ctx, call.host, grpc.WithCodec(Codec()), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	call.client = conn

	return conn, nil
}

// SendMsg calls the configured host and attempts to call the given endpoint with the given headers and stream
func (call *Call) SendMsg(ctx context.Context, rw transport.ResponseWriter, pr *transport.Request, refs *refs.Store) error {
	method := call.methods[pr.Method.GetName()]
	if method == nil {
		return trace.New(trace.WithMessage("unkown service method %s", pr.Method.GetName()))
	}

	conn, err := call.Director(ctx)
	if err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, CopyMD(pr.Header))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := grpc.NewClientStream(ctx, proxy, conn, method.fqn)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(pr.Body)
	if err != nil {
		return err
	}

	req := &frame{
		payload: body,
	}

	err = stream.SendMsg(req)
	if err != nil {
		return err
	}

	md, err := stream.Header()
	if err != nil {
		return err
	}

	rw.Header().Append(CopyRPCMD(md))

	res := &frame{}
	err = stream.RecvMsg(res)
	if err != nil {
		return err
	}

	_, err = rw.Write(res.payload)
	if err != nil {
		return err
	}

	return nil
}

// Close closes the given caller
func (call *Call) Close() error {
	call.logger.WithField("host", call.host).Info("Closing HTTP caller")
	return nil
}
