# Channels in Go

From [Distributed Computing with Go](https://www.packtpub.com/application-development/distributed-computing-go)

## Cashier Problem

30 orders, and a cashier can handle no more than 10 a day.

* [Single Cashier][cashier-lacking]
* [Multiple Cashiers without Channels][cashier-sans-channels]
* [Multiple Cashiers with Channels][cashier-with-channels]

## Channels in General

* [Channel Basic Demo][channel-demo]
* [Channel Function Demo][channel-with-functions]
* [Channel Being Closed][channel-being-closed]

## Multiplexing

* [Naive Multiplexing][multiplexing-naive]
* [Proper Multiplexing - Attempt 1][multiplexing-proper-1]
* [Proper Multiplexing - Attempt 2][multiplexing-proper-2]

[cashier-lacking]:         cashier/cashier-1-lacking.go
[cashier-sans-channels]:   cashier/cashier-2-sans-channels.go
[cashier-with-channels]:   cashier/cashier-3-with-channels.go
[channel-demo]:            channels/channel-demo.go
[channel-with-functions]:  channels/channel-with-functions-demo.go
[channel-being-closed]:    channels/channel-being-closed-demo.go
[multiplexing-naive]:      multiplexing/naive-multiplexing.go
[multiplexing-proper-1]:   multiplexing/proper-multiplexing-attempt-1.go
[multiplexing-proper-2]:   multiplexing/proper-multiplexing-attempt-2.go
