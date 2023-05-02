package main

type Ohce struct {
	responseBuilder ResponseBuilderI
	clock           ClockI
}

func (o *Ohce) Greet(name string) string {
	return o.responseBuilder.BuildNameResponse(name, o.clock)
}

func (o *Ohce) Respond(s string) string {
	return o.responseBuilder.BuildNormalResponse(s)
}
