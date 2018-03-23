# Channels

From [Distributed Computing with Go](https://www.packtpub.com/application-development/distributed-computing-go)

## Cashier Problem

30 orders, and a cashier can handle no more than 10 a day.

* [Single Cashier][cashier-lacking]
* [Multiple Cashiers without Channels][cashier-sans-channels]
* [Multiple Cashiers with Channels][cashier-with-channels]

[cashier-lacking]:        cashier/cashier-1-lacking.go
[cashier-sans-channels]:  channels/cashier-2-sans-channels.go
[cashier-with-channels]:  channels/cashier-3-with-channels.go
