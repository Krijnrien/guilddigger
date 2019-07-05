package rpc_resolver

import "google.golang.org/grpc/resolver"

//TODO
// expose ports which can add an address to the address list and then call UpdateState (aka dynamically adding new server locations to already running clients)

// Following is an example name resolver implementation. Read the name
// resolution example to learn more about it.
const (
	exampleScheme      = "example"
	exampleServiceName = "www.example.com"
)

var addrs = []string{"192.168.56.1:50051", "192.168.56.1:50052"}

type ExampleResolverBuilder struct{}

func (*ExampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			exampleServiceName: addrs,
		},
	}
	r.start()
	return r, nil
}
func (*ExampleResolverBuilder) Scheme() string { return exampleScheme }

type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*exampleResolver) Close()                                 {}

func init() {
	resolver.Register(&ExampleResolverBuilder{})
}
