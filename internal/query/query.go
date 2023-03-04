package query

type Resolver struct{}

func (r *Resolver) Hello() string {
	return "Hello, world!"
}
