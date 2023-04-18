# Operator Foundation

[Operator](https://operatorfoundation.org) makes usable tools to help people around the world with censorship, security, and privacy.

# Optimizer

Optimizer is a pluggable transport that uses one of several possible “Strategies” to choose between the transports you provide to create a connection. It is not a standalone transport, but is rather a mechanism for choosing between various transports in order to find the one best suited for the user’s needs. For more information about pluggable transports, please refer to [pluggabletransports.info](https://www.pluggabletransports.info/).

Here is a list of the currently available Optimizer strategies:

**Rotate Strategy**: This strategy simply rotates through the list of provided transports and tries the next one in the list each time a connection is needed.

**Choose Random Strategy**: A transport is selected at random from the list for each connection request.

**Track Strategy**: A strategy that  attempts to connect with each of the provided transports. It keeps track of which transports are connecting successfully and favors using those.

**Minimize Dial Strategy**: The transport is chosen based on which has been shown to connect the fastest.

## Shapeshifter

The Shapeshifter project provides network protocol shapeshifting technology (also sometimes referred to as obfuscation). The purpose of this technology is to change the characteristics of network traffic so that it is not identified and subsequently blocked by network filtering devices.

There are two components to Shapeshifter: transports and the dispatcher. Each transport provides a different approach to shapeshifting. These transports are provided as a Go library which can be integrated directly into applications. The dispatcher is a command line tool which provides a proxy that wraps the transport library. It has several different proxy modes and can proxy both TCP and UDP network traffic.

If you are an application developer working in the Go programming language, then you probably want to use the transports library directly in your application. If you are an end user that is trying to circumvent filtering on your network or you are a developer that wants to add pluggable transports to an existing tool that is not written in the Go programming language, then you probably want the dispatcher. Please note that familiarity with executing programs on the command line is necessary to use this tool. You can find Shapeshifter Dispatcher here: <https://github.com/OperatorFoundation/shapeshifter-dispatcher>

If you are looking for a complete, easy-to-use VPN that incorporates shapeshifting technology and has a graphical user interface, consider [Moonbounce](https://github.com/OperatorFoundation/Moonbounce), an application for macOS which incorporates Shapeshifter without the need to write code or use the command line.

### Shapeshifter Transports
The transports implement the Go API from the [Pluggable Transports 3.0 specification](https://github.com/Pluggable-Transports/Pluggable-Transports-spec/blob/main/releases/PTSpecV3.0/Pluggable%20Transport%20Specification%20v3.0%20-%20Go%20Transport%20API%20v3.0.md).

The purpose of the transport library is to provide a set of different transports. Each transport implements a different method of shapeshifting network traffic. The goal is for application traffic to be sent over the network in a shapeshifted form that bypasses network filtering, allowing the application to work on networks where it would otherwise be blocked or heavily throttled.

## Installation
Optimizer is written in the Go programming language. To compile it you need
to install Go:

<https://golang.org/doc/install>

If you already have Go installed, make sure it is a compatible version:

    go version

The version should be 1.17 or higher.

If you get the error "go: command not found", then trying exiting your terminal
and starting a new one.

In order to use Optimizer in your project, you must have Go modules enabled in your project. How to do this is beyond the scope of this document. You can find more information about Go modules here: <https://blog.golang.org/using-go-modules>

To use in your project, simply import:

    import "github.com/OperatorFoundation/Optimizer-go/Optimizer/v3"
    
Your go build tools should automatically add this module to your go.mod and go.sum files. Otherwise, you can add it to the go.mod file directly. See the official Go modules guide for more information on this.    

Please note that the import path includes "/v3" to indicate that you want to use the version of the module compatible with the PT v3.0 specification. This is required by the Go modules guide.

When you build your project, it should automatically fetch the correct version of the transport module.

## Using Optimizer

1. First you will need to initialize the transports you would like Optimizer to use:
    `dialer := proxy.Direct`
	`shadowTransport := shadow.Transport{"InsertPasswordHere", "InsertCipherNameHere", "InsertAddressHere"}`
	
2. Create an array with these transports:
    `transports := []Transport{shadowTransport}`
    
3. Initialize the strategy of your choice using the array of transports you created:
    `strategy := NewMinimizeDialDuration(transports)`
    
4. Create an instance of OptimizerConnectionFactory using your new Strategy instance:
    `optimizerTransport := NewOptimizerClient(transports, strategy)`
    
5. Call Dial on optimizerTransport:
    `_, err := optimizerTransport.Dial()`
