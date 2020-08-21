# Statefun Go SDK

Stateful Functions is an API that simplifies the building of **distributed stateful applications** with
a **runtime built for serverless architectures**. It brings together the benefits of stateful stream
processing, the processing of large datasets with low latency and bounded resource constraints,
along with a runtime for modeling stateful entities that supports location transparency, concurrency,
scaling, and resiliency.

It is designed to work with modern architectures, like cloud-native deployments and popular event-driven FaaS platforms
like AWS Lambda and KNative, and to provide out-of-the-box consistent state and messaging while preserving the serverless
experience and elasticity of these platforms.

Background

The JVM-based Stateful Functions implementation has a RequestReply extension
(a protocol and an implementation) that allows calling into any HTTP endpoint
that implements that protocol. Although it is possible to implement this protocol
independently, this is a minimal library for the Go programing language that:

- Allows users to define and declare their functions in a convenient way.
- Dispatches an invocation request sent from the JVM to the appropriate function previously declared.

See the [Apache Flink Stateful Functions](https://flink.apache.org/stateful-functions.html) website for more
information about the project.

The following shows how to write a simple stateful function and serve it for use in a Statefun deployment.

```go
func (greeter Greeter) Invoke(runtime statefun.StatefulFunctionRuntime, _ *any.Any) error {
	var seen SeenCount
	if err := runtime.Get("seen_count", &seen); err != nil {
		return err
	}

	seen.Seen += 1

	if err := runtime.Set("seen_count", &seen); err != nil {
		return err
	}

	response := computeGreeting(runtime.Self().Id, seen.Seen)

	record := io.KafkaRecord{
		Topic: "greetings",
		Key:   runtime.Self().Id,
		Value: response,
	}

	message, err := record.ToMessage()
	if err != nil {
		return nil
	}

	return runtime.SendEgress(egressId, message)
}
```



This is a simple example that runs a simple stateful function that accepts requests from a Kafka ingress, and then responds by sending greeting responses to a Kafka egress.
It demonstrates the primitive building blocks of a Stateful Functions applications, such as ingresses, handling state in functions, and sending messages to egresses.

See [greeter example](examples/README.md) for an end to end demonstration. 

**WARNING** This library is still in alpha and developers reserve the right to 
make breaking changes to the api. 
