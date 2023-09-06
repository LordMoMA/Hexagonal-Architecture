# 🌌 Hexagonal Architecture with Go - A Thorough Exploration of Backend Engineering and Distributed System

![](images/arch.webp)
This is the source code for the original article:

[Hexagonal Architecture Deep Dive with PostgreSQL, Redis and Go Practices](https://medium.com/towardsdev/hexagonal-architecture-deep-dive-with-postgresql-redis-and-go-practices-4b051f940e93)

Note that the codebase has evolved with more complexity than the article's example. The article is just a starting point for the project.

## 🏡 What is Hexagonal Architecture?

Hexagonal Architecture, also known as Ports and Adapters Architecture or Clean Architecture, is a software architecture pattern that promotes loose coupling between the application core (business logic) and external components such as user interface, database, and external services.

In Hexagonal Architecture, the core of the application is isolated from external components and is instead accessed through a set of well-defined interfaces or ports. Adapters are then used to implement the required interfaces and integrate with the external components.

## 🔮 Hexagonal Architecture Components

Here are the components of the Hexagonal Architecture:

### 🎖 Core Business Logic:

The Core Business Logic is responsible for the main functionality of the application. This component represents the heart of the application and should be designed to be independent of any external dependencies. In Hexagonal Architecture, the Core Business Logic is implemented as a set of use cases that encapsulate the behavior of the application.

For example, if we are building a banking application, the Core Business Logic would include use cases such as creating an account, transferring funds, and checking account balance.

### 👯 Adapters:

The Adapters are responsible for connecting the Core Business Logic to the external world. Adapters can be of two types: Primary and Secondary.

#### 🕺 Primary Adapter:

The Primary Adapter is responsible for handling incoming requests from the external world and sending them to the Core Business Logic. In Hexagonal Architecture, the Primary Adapter is typically an HTTP server, which receives HTTP requests from clients and converts them into requests that can be understood by the Core Business Logic.

For example, in a banking application, the Primary Adapter would be an HTTP server that listens for incoming requests from clients, such as transferring funds or checking account balances, and then converts them into use cases that can be understood by the Core Business Logic.

#### 🥁 Secondary Adapters:

The Secondary Adapters are responsible for interfacing with external dependencies that the Core Business Logic relies on. These dependencies can be databases, message queues, or third-party APIs. Secondary Adapters implement the ports defined by the Core Business Logic.

In a banking application, the Secondary Adapters would include database adapters that interface with the Core Business Logic to store and retrieve data about accounts, transactions, and other related information.

### 😈 Interfaces:

In software architecture, an interface refers to a contract or an agreement between two software components. It defines a set of rules or protocols that a component must follow in order to communicate with another component.

In the context of hexagonal architecture, interfaces play a critical role as they define the boundaries of the core business logic and the adapters. The core business logic only interacts with the adapters through their interfaces. This allows for easy replacement of adapters without affecting the core business logic.

For example, let's say you have an online shopping application that needs to process payments. You can define an interface for the payment gateway adapter, which outlines the methods that the core business logic can use to interact with the payment gateway.

You can then have multiple payment gateway adapters that implement this interface, such as PayPal, Stripe, and Braintree. The core business logic only interacts with the payment gateway adapters through their defined interface, allowing for easy replacement or addition of payment gateways without affecting the core business logic.

### 👨‍👦‍👦 Dependencies:

These are the external libraries or services that the application depends on. They are managed by the adapters, and should not be directly accessed by the core business logic. This allows the core business logic to remain independent of any specific infrastructure or technology choices.

## 🤡 Application structure (to be updated)

Now, let's dive into how to create a messaging backend that allows users to save and read messages. Hexagonal architecture adheres to strict application layout that needs to be implemented. Below is the application layout that we will use. This might look like a lot of work, but it will make sense as we move forward.

# Structure of the Project

```
└── Hexagonal Architecture
   ├── cmd
   │   └── main.go
   ├── .env
   ├── images
   ├── go.mod
   ├── go.sum
   └── internal
       ├── adapters
       │   ├── cache
       │   │   └── cache.go
       │   │
       │   ├── handler
       │   │   ├── error_handler.go
       │   │   ├── login_handler.go
       │   │   ├── message_handler.go
       │   │   ├── stripe_handler.go
       │   │   ├── user_handler.go
       │   │   └── webhook_handler.go
       │   │
       │   ├── repository
       │   │   ├── apiCfg.go
       │   │   ├── db.go
       │   │   ├── message.go
       │   │   ├── payment.go
       │   │   └── user.go
       │   └── tests
       │       ├── integration
       │       └── unit
       │
       ├── config
       │    ├── config.go
       │    └── nginx.conf
       ├── core
       │   ├── domain
       │   │   └── model.go
       │   ├── ports
       │   │   ├── ports.go
       │   │   └── ports.go
       │   ├── services
       │   │   ├── message.go
       │   │   ├── payment.go
       │   │   └── user.go
       └── web
```

# 👺 To-dos:

- ✅ Finish CRUD process of the messaging service
- ✅ REST API Design with Gin
- ✅ Add User service
- ✅ Add JWT Authentication and Authorisation
- ✅ Optimise error handling with clean code
- ✅ Add Webhook to update membership status (idempotent)
- ✅ Add a payment service
- ✅ Work with Stripe API
- ✅ postgreSQL as database
- ✅ Redis as cache on users to improve performance
- ✅ Add a new server for v2/payments endpoint
- ✅ Add load balancer for server cluster
- 🥷🏻 Fix tests to use DB migration and pass the CI 
- ⌛️ Add telemetry to APIs 
- ⌛️ Add observability and monitoring to the /users/:id endpointd
- ⌛️ Design wallet service
- ⌛️ Design payment event service
- ⌛️ Design a double-entry ledger system
- ⌛️ Add Unit Test
- ⌛️ Add Distributed services
- ⌛️ Add URL Queries

## How to keep the Test DB running and test the exact func with the triangle button in an IDE?

```bash
chmod +x testDB.sh
```
Now you can use the following command to start the test database:
```bash
./testDB.sh -t start
```
And use the following command to run tests and stop the test database afterwards:
```bash
./testDB.sh -t unit    # For unit tests
./testDB.sh -t integration    # For integration tests
./testDB.sh    # For all tests
./testDB.sh -t stop # stop the test database
```
![badge](https://github.com/LordMoMA/Hexagonal-Architecture/.github/workflows/go.yml/badge.svg)

# 🚀 Pros and Cons of using GORM in this project

GORM is a popular Object-Relational Mapping (ORM) library for the Go programming language that provides a convenient way to interact with databases, including PostgreSQL.

It provides a high-level, expressive and easy-to-use API for CRUD (Create, Read, Update, Delete) operations and supports several databases, including MySQL, PostgreSQL, SQLite, and others.

Whether GORM is better to use than directly using PostgreSQL depends on the specific use case. If you need a high-level, user-friendly API to interact with your PostgreSQL database, then GORM can be a great choice. On the other hand, if you have specific requirements for your database interactions or need to optimize performance for a large-scale application, then direct interaction with the PostgreSQL database using a lower-level database driver may be more appropriate.

In general, the use of an ORM can simplify and speed up development, especially for CRUD operations. However, it may introduce additional overhead and performance concerns.

please show me step by step of how payment service can work with Stripe API based on the following payment structure of hexagonal architecture, no need to use code.

# 🧠 Thoughts on the Payment Service and Stripe API Integration

If you already have an API endpoint that interacts with the Stripe API, you may not need a payment service in the Hexagonal Architecture. However, if you want to store payment data in your local database for future reference or analysis, you can create a payment service to handle this.

To get the payment data from the Stripe API endpoint, you can use webhooks to receive events from Stripe when a payment is made. You can then parse the webhook data and store the relevant payment information in your local database.

Alternatively, if you are using Stripe's checkout feature, you can use the client_secret that is returned when you create a PaymentIntent to confirm the payment after it is made. Once the payment is confirmed, you can retrieve the payment data from Stripe using the PaymentIntent ID and store it in your local database.

Overall, the payment service in the Hexagonal Architecture would be responsible for storing and retrieving payment data from the local database, and potentially for processing payments and interacting with the Stripe API via webhooks or other methods.

# 🌈 The Stripe's Checkout and PaymentIntent Confusion

When you are working with Stripe API for the first time, you might be confused about the difference between Stripe's Checkout and PaymentIntent. This is because they are both used to accept payments, but they serve different purposes and have different capabilities.

Stripe's Checkout and PaymentIntent are both features that allow you to accept payments through Stripe, but they serve different purposes and have different capabilities.

Stripe Checkout is a pre-built payment page that handles the payment process on behalf of the merchant. It allows merchants to quickly and easily integrate a payment flow into their website without having to build their own payment form. Stripe Checkout also supports many payment methods, including credit and debit cards, Apple Pay, and Google Pay.

PaymentIntent, on the other hand, is a flexible API that allows merchants to create and manage payment transactions programmatically. With PaymentIntent, merchants have more control over the payment process, including the ability to handle complex payment scenarios, such as partial payments, deferred payments, and payments with multiple payment methods. In other words, the PaymentIntent API is a low-level API that allows you to create and manage payment transactions programmatically. It is not a pre-built payment form like Stripe Checkout.

In summary, Stripe Checkout is a pre-built payment form that makes it easy for merchants to get started with Stripe payments, while PaymentIntent provides a more flexible and powerful API for handling payment transactions programmatically.

# 🌱 On Redis Parameters

In the Get method of RedisCache, the value parameter is defined as interface{} because it can take any type of value that is stored in the cache. The Get method is used to retrieve a value from the cache by providing the key. However, since the type of the value stored in the cache is unknown, it is specified as an empty interface interface{} which is a type that can hold any value.

# 🚇 On Redis Cache and PostgreSQL DB Dada Consistency

To maintain consistency between Redis Cache and PostgreSQL DB, you can implement a write-through or write-behind caching strategy.

In the write-through caching strategy, when data is updated in the PostgreSQL DB, it is also updated in the Redis Cache. This ensures that the data in the Redis Cache is always up-to-date with the latest data in the PostgreSQL DB. However, this approach can result in slower write performance due to the additional overhead of updating the cache.

In the write-behind caching strategy, data is first updated in the Redis Cache and then asynchronously updated in the PostgreSQL DB. This approach can improve write performance as data is first updated in the faster Redis Cache and then updated in the slower PostgreSQL DB. However, this approach can result in a temporary inconsistency between the Redis Cache and PostgreSQL DB.

Additionally, you can use a combination of database transactions and cache invalidation to ensure consistency. When a transaction is committed to the PostgreSQL DB, the cache is invalidated, and the next read from the cache will result in the latest data from the PostgreSQL DB.

Please read my article for more information on Cache Invalidation:
[The Hard Thing in Computer Science: Cache Invalidation](https://medium.com/@lordmoma/the-hard-thing-in-computer-science-cache-invalidation-11ca0da2dba4)

It's also important to ensure that the TTL (Time-to-Live) of the cached data is set appropriately. This ensures that the cached data is not stale and remains consistent with the data in the PostgreSQL DB.

However, in this project I want to keep this problem simple and easy to handle:

I will delete the cache data everytime the database is updated.

For example: when a user's email is updated in the database, I can delete the corresponding user's cache in Redis, so that the next time the user data is requested, it will be fetched from the database and cached again with the updated email. This ensures that the cache data remains consistent with the database data.

To achieve this, I added a cache invalidation logic in the code that detects changes in the database and deletes the corresponding cache data. This can be done using database triggers, which are special stored procedures that automatically execute in response to certain database events, such as an update or delete operation on a table.

🕺 Note: I have improved user query speed by 11.37 (10.438294ms / 918.226µs) times. (1 ms millisecond = 1000 µs microseconds).

# 🔭 On Observability and Monitoring to the `/users/:id` Endpoint

### Add instrumentation to the endpoint code:

Use a tracing library like OpenTelemetry or OpenTracing to add instrumentation to the code that handles the `/users/:id` endpoint. This will allow you to track the duration of the request, as well as any errors that occur during processing.

OpenTelemetry span is a way to track an operation's progress through a distributed system. A span represents a single operation, which could be a function call or a network request, and contains metadata about that operation such as the start and end times, any attributes, events, and links associated with the operation.

In the case of the ReadUser function, using OpenTelemetry span would allow you to track the progress of the function call and collect relevant metadata for the operation. For example, you could create a span to represent the ReadUser function call, add attributes to the span such as the user ID being read, and record any events associated with the operation, such as when the database query was executed.

Using OpenTelemetry span in this way would allow you to gather valuable data about the ReadUser operation, such as the time it takes to execute, any errors that occur, and the performance of the underlying systems involved. This data could be used to diagnose issues, optimize performance, and improve the overall reliability of the system.

### Use a metrics library to collect metrics:

Use a metrics library like Prometheus or StatsD to collect metrics about your application. You can instrument your code to emit metrics related to the /users/:id endpoint, such as the number of requests received or the latency of each request.

### Use a logging library to log important events:

Use a logging library like Logrus or Zap to log important events related to the /users/:id endpoint. For example, you can log when a request is received, when it is processed, and when it is completed.

### Use a monitoring tool to visualize the data:

Use a monitoring tool like Grafana or Kibana to visualize the data collected by your metrics and logging libraries. This will allow you to identify trends, spot anomalies, and diagnose issues.

# 🏃🏻‍♀️ How to Improve the Maximum TPS of the v2/payments API Service

To improve the maximum TPS (transactions per second) of the v2/payments API service, there are several strategies that can be employed:

### Optimize database queries:

One of the most common bottlenecks in a high-traffic API service is the database. By optimizing queries, indexing tables, and caching frequently-accessed data, the response time can be improved, leading to a higher TPS.

### Load balancing:

Distributing the incoming traffic across multiple servers can help increase the TPS of the API service. Load balancing can be done using hardware load balancers or software load balancers such as Nginx or HAProxy.

### Caching:

Caching frequently-accessed data can reduce the number of database queries required and improve the TPS. Memcached or Redis can be used for caching.

### Asynchronous processing:

By using asynchronous processing for time-consuming tasks, such as sending emails or processing images, the API service can handle more requests per second.

### Horizontal scaling:

Adding more servers to the server cluster can help increase the TPS. Kubernetes or Docker Swarm can be used for container orchestration to manage the scaling of the API service.

By implementing these strategies, the maximum TPS of the `v1/payments` API service can be improved.

# About Nginx as a Load Balancer

In the context of the hexagonal architecture, Nginx can be used as a load balancer to distribute the incoming traffic across multiple servers. This can help increase the TPS (transactions per second) of the API service.

Take a look at `internal/config/nignx.conf` file:

    ```nginx
        upstream myapp {
      server localhost:5000 weight=3 max_fails=3 fail_timeout=30s;
    }
    ```

In this configuration, Nginx is acting as a reverse proxy and load balancer. It receives HTTP requests from clients and forwards them to one of the backend servers specified in the upstream block, in this case localhost:5000. The weight parameter specifies how much traffic each server should receive relative to the others. The max_fails and fail_timeout parameters specify how Nginx should handle failures on a backend server.

So, with this configuration, Nginx is distributing the load across multiple backend servers, which can help improve the performance and availability of your application.

# 👽 About the Payment Platform's Architecture

Zooming in our payment architecture, we can distinguish several key components (see figure 1 below):

1. API, providing a uniform interface to the payment functionality,
2. Risk Engine, making decisions about payment-related risks,
3. Payment Profile Service, providing the details of payment mechanisms,
4. User Profile Service, providing details about user payment and other settings,
5. Payment Auth Service, providing services for authentication of payments,
6. PSP Gateways, implementing the integration with payment service providers (PSPs),
7. Order store, storing data about orders, and
8. Account store, storing data about the accounts of payment parties.

   ![](images/payment.png)
   figure 1

## Implementation: Distributed Stream-Processing

At a deeper level, our Payments platform is implemented as a collection of microservices organized as a stream-processing architecture. Streaming data refers to data that is continuously generated, typically in high volumes and at high velocity. Uber handles dozens of millions of transactions daily, making a streaming-based architecture a natural choice.

### Key Technology: Apache Kafka

The key technology used by our Payments Platform is Apache Kafka: an open-source stream-processing software platform

![](images/kafka.png)
figure 2

Kafka has several key capabilities, could be inherited by the our Payments platform:

- Publishing and subscribing to streams of records, similar to message queues or enterprise messaging systems.

- Storing streams of records in a fault-tolerant durable way.

- Asynchronously processing streams of records as they occur. Asynchronous processing maps well on the transactions in the payments domain: payment processing requires high reliability, but can afford to be implemented asynchronously (within a time-bound).

- Horizontal scaling to handle changing load.

Nodes connected via Kafka are typically microservices, can be built in Go and Java, NodeJS or Python.

In addition, Kafka supports well high performance and scalability requirements. Kafka is horizontally scalable, fault-tolerant, and optimized for speed, running as a cluster on one or more servers that can span multiple data centers (For example: Uber uses use a combination of third-party cloud computing services and co-located data centers).

## Performance and Scalability

Let's look at the Uber's example to see how the Payments platform is implemented in practice.

One of the key technical challenges that Uber faces in the implementation of payments platform is the scale of its operations. For illustration, here are some recent stats:

65 countries, 600 cities,
75 million Uber passengers,
3.9 million Uber drivers,
14 million Uber trips per day (well over 10 billion trips have been completed worldwide).
In addition to the worldwide scale, the load is not uniform and may have unexpected spikes.

While details are not publically available, the technical presentations provide some insights in mechanisms used by Uber for handling performance and scalability requirements, such as:

- Extensive parallelization of processing with the [competing consumers pattern](https://www.enterpriseintegrationpatterns.com/patterns/messaging/CompetingConsumers.html), by having multiple parallelly running (micro)service instances

- Independent scaling of processing components, to more flexibly manage needed capacity

- Using optimistic locking, to avoid the need for complex distributed locking mechanisms.

## Reliability

Implementing a reliable streaming-based payment system comes with several challenges:

- System failures (a failure may occur midway through processing)
- Poison pill (an inbound message cannot be consumed)
- Functional bugs (no technical errors, but results are invalid)

Key mechanisms to deal with reliability requirements include:

- Redundancy of all services, including the messaging infrastructure, enables resilience during internal system failures,
- Implementation of the guaranteed delivery pattern pattern, by using Kafka capability to persist messages so that they are not lost even if the messaging system crashes,
- Implemention of timeouts, both in integration with external systems, as well as internal services to prevent long-term system overloading,
- Retrying operations, based on a defined error strategy (see figure 3), or move messages to a dead letter queue, so that messages are never lost,
- Implementation of idempotent message handling for service operations. An [idempotent operation](https://stackoverflow.com/questions/1077412/what-is-an-idempotent-operation) is one that has no additional effect if it is called more than once with the same input parameters. Apache Kafka implements the “at least once” message delivery strategy, implying subscribers may receive the same message multiple times, so subscribers that manage state and cause side effects should implement idempotent message handling.
- Load-smoothing through queuing, to avoid overloading of services, and
- Validation of processing results based on side-effects recording

![](images/error.png)
Figure 3: Error handling requires an error strategy. An error can lead to retry of an operation, of its achiving the dead message queue (DMQ).

![](images/validation.png)
Figure 4: Each complex operation will lead to some side effects. A validator can them at some moment check if actual side effects match the expected once.

## Implementation: Integration with External Systems

The Payments Platform interacts with payment service providers (PSPs) and banks to execute payment transactions.

Each integration with PSPs and banks is different, we can distinguish two integrations styles (Figure 5):

- API-based integrations with modern PSP integrations, with REST-based APIs, exchanging data in JSON, one transaction at a time, near-real time.

- Legacy batch integration with banks, where integrations are done by exchanging files via SFTP, with relatively low frequency (day or hours).

![](images/external_integrations.png)
Figure 5: Two integrations styles for integration with external systems: API-based, and file-based.

Idempotency is an essential theme in integration with external payment systems. A good thing about PSP and banking systems is that they are normally implementing their services as idempotent message processors. Idempotency is essential for payment systems for two reasons:

- It helps to prevent double charging
- It improves reliability and simplifies system architecture.

When a failure occurs (e.g., a network error), it may be challenging to determine if some operation succeeded or failed and in which state the system is. Without idempotency, for instance, retying operations may be risky, as you may execute the same operation twice (e.g. charging a customer twice for the same service).

With idempotency, you can repeat the failed operation without such worries. Figure 6 illustrates how idempotency (in the context of integration with external systems) works in an ideal scenario.

![](images/idempotency.png)
Figure 6: Idempotent message processing systems will not process the same message twice.

The good news is that I have developed our own webhook handler to solve this problem. We can also adopt Stripe API's webhook to solve this problem.

### Challenges with Idempotency

Idempotency works well if you repeat the request against the same system, with the same operation ID. The operations ID needs to be provided by the application calling an idempotent service so that the service knows if it is getting the new request (not previously processed ID) or a repeated operation (already processed ID).

One challenge of implementing idempotency when interacting with external systems relates to the IDs used for idempotent operations. Legacy payments systems accept a more limited range of values for IDs. Careful rotation and timing of such IDs are essential to avoid the external system rejecting the payment request.

Another challenge is multiplexing PSPs:

Payments operations use several PSPs in a complex arrangement, and another PSP may be used if a payment fails with the originally selected one. Such practice may improve collection rate, but naively retrying a failed operation on another PSP may lead to double charging, as illustrated in Figure 7.

![](images/multiplexing.png)
Figure 7: The incorrect way to retry operations in the case of network failures when working with multiple PSPs. Network error does not necessarily mean that the operation has failed, and retrying the operation on a different PSP may thus lead to double charging.

### Solutions to this problem:

Using dedicated request storage when a retry needs to be performed, to ensure that retry goes back to an original service (Figure 8).

![](images/request_storage.png)
Figure 8: The correct way to retry operations in the case of network failures when working with multiple PSPs. Using dedicated request storage to ensure that retry goes back to an original service.

# 👾 On Test

## The Myths of Integration Testing and Unit Testing

In the context of software testing, integration testing and unit testing are two different types of tests that serve different purposes.

Unit testing focuses on testing individual units of code in isolation, typically at the function or method level. The goal of unit testing is to ensure that each unit of code works correctly on its own, without dependencies on other parts of the system. Unit tests are usually automated and can be run frequently as part of a continuous integration process.

Integration testing, on the other hand, tests the interactions and dependencies between different parts of the system. Integration tests may involve multiple units of code, subsystems, or external systems. The goal of integration testing is to ensure that all the pieces of the system work correctly together as a whole.

In the context of the hexagonal architecture, the unit tests would typically test the behavior of the core domain logic in isolation, while the integration tests would test the interactions and dependencies between the core logic and the adapters (such as the database or external APIs).

In the provided structure, the unit folder contains the user_service_test.go file, which likely contains tests for the UserService functions at the core level, testing their functionality in isolation from other parts of the system.

The integration folder contains the user_integration_test.go file, which likely contains tests that simulate the interaction between the UserService and the adapters, such as the UserRepository. These tests may use a real database or external API, and aim to test the behavior of the system as a whole.

## Benchmarking

Please read my article on [6 Tips on High Performance Go — Advanced Go Topics](https://medium.com/@lordmoma/6-tips-on-high-performance-go-advanced-go-topics-37b601fa329d) for more information.

We have implemented a benchmarking on `createUser_test.go` to make sure that
the performance of our code is not degraded.

```bash
go test -bench=. -benchmem
```

Output:

```bash
goos: darwin
goarch: amd64
pkg: github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/tests/benchmark
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
BenchmarkCreateUser-4                 16          70744288 ns/op           35311 B/op        594 allocs/op
PASS
ok      github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/tests/benchmark    3.321s
```

Analysis:

```bash
goos: the operating system the benchmark was run on.
goarch: the architecture of the processor the benchmark was run on.
pkg: the package being benchmarked.
cpu: the processor being used.
BenchmarkCreateUser-4: the name of the benchmark.
The "-4" indicates that the benchmark was run with 4 CPUs.
16: the number of iterations run in the benchmark.
70744288 ns/op: the average time it took to run one iteration of the benchmark, measured in nanoseconds.
35311 B/op: the average number of bytes allocated per iteration of the benchmark.
594 allocs/op: the average number of allocations per iteration of the benchmark.
```

In this case, the BenchmarkCreateUser benchmark was run with 16 iterations, and each iteration took an average of 70,744,288 nanoseconds (or about 70.7 milliseconds) to complete. During each iteration, an average of 35,311 bytes were allocated, and an average of 594 allocations were made.

## Profiling

Please read my article on [6 Tips on High Performance Go — Advanced Go Topics](https://medium.com/@lordmoma/6-tips-on-high-performance-go-advanced-go-topics-37b601fa329d) for more information.

Go has built-in profiling tools that can help you gain insight into what your code is doing. The most common profiling tool is the CPU profiler, which can be enabled by adding the -cpuprofile flag to the go test command.

```bash
go test -cpuprofile=prof.out
```

Output:

```bash
testing: warning: no tests to run
PASS
ok      github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/tests/benchmark    1.381s
```

```bash
go tool pprof prof.out
```

Output:

```bash
Type: cpu
Time: May 11, 2023 at 8:04pm (CST)
Duration: 202.62ms, Total samples = 0
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
```

After entering interactive mode in the go tool pprof, you can use the following commands to analyze and interact with the CPU profile:

```bash
top: Show top entries in the profile.
list [function]: Show source code of a function, or list functions in the profile.
web: Open a graphical representation of the profile in your default web browser.
focus [function]: Focus on a specific function in the profile, hiding everything else.
unfocus: Remove focus on a function, showing everything again.
help: Show a list of available commands.
quit: Exit the interactive mode.
```

### Introduction to gin-contrib/pprof

This is gin middleware pprof. You can find it at https://github.com/gin-contrib/pprof.

Integration of pprof:

First of all, install pprof:

```bash
$ go get github.com/gin-contrib/pprof
```

Then integrate pprof into the gin router:

```go
httpRouter := gin.Default()
pprof.Register(httpRouter)
```

After start the server, go to `http://localhost:5000/debug/pprof/`

you will see:

![](images/pprof.png)

Goroutine: stack traces of all current Goroutines
CPU: stack traces of CPU returned by the runtime
Heap: a sampling of memory allocations of live objects
Allocation: a sampling of all past memory allocations
Thread: stack traces that led to the creation of new OS threads
Block: stack traces that led to blocking on synchronization primitives
Mutex: stack traces of holders of contended mutexes

## Measuring the performance

We are going to measure how many requests per second the microservice is able to handle. This can be done using the HTTP load generators.

Install hey

```bash
brew install hey
```

To check the performance of our app let’s run our app with go run command i.e. go run cmd/main.go

Then let’s generate load to our web app as below:

```bash
hey -n 10000000 -c 8 http://localhost:5000/v1/users
```

This will generate 10000000 of requests to /api/user with maximum 8 number of workers to run concurrently. By default, hey sets 50 number of workers.

output:

```bash
Summary:
  Total:        199.9452 secs
  Slowest:      0.8245 secs
  Fastest:      0.0004 secs
  Average:      0.0114 secs
  Requests/sec: 703.3378


Response time histogram:
  0.000 [1]     |
  0.083 [137442]|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.165 [2552]  |■
  0.248 [396]   |
  0.330 [137]   |
  0.412 [47]    |
  0.495 [16]    |
  0.577 [17]    |
  0.660 [15]    |
  0.742 [4]     |
  0.824 [2]     |


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0020 secs
  50% in 0.0036 secs
  75% in 0.0066 secs
  90% in 0.0329 secs
  95% in 0.0536 secs
  99% in 0.1218 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0004 secs, 0.8245 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0056 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0022 secs
  resp wait:    0.0111 secs, 0.0004 secs, 0.8243 secs
  resp read:    0.0002 secs, 0.0000 secs, 0.2291 secs

Status code distribution:
  [200] 140629 responses
```

The analysis provides information about the performance of a system based on the given summary, response time histogram, latency distribution, and details. Let's break down the information:

Summary:

- Total time: 199.9452 seconds: This is the total duration of the performance test.
- Slowest response: 0.8245 seconds: The slowest individual response recorded during the test.
- Fastest response: 0.0004 seconds: The fastest individual response recorded during the test.
- Average response time: 0.0114 seconds: The average response time across all requests.
- Requests per second: 703.3378: The number of requests processed per second.

Response time histogram:

The histogram displays the distribution of response times in different ranges. The number of requests falling within each range is represented by the vertical bars.

Latency distribution:

This section shows the distribution of response times based on percentiles.

For example, 10% of the requests had a response time of 0.0012 seconds or lower.
90% of the requests had a response time of 0.0329 seconds or lower.
99% of the requests had a response time of 0.1218 seconds or lower.

Details:

The details section provides average, fastest, and slowest times for different stages of the request-response cycle.

- DNS+dialup: The time taken for DNS resolution and establishing a connection with the server.
- DNS-lookup: The time taken for DNS resolution only.
- req write: The time taken to write the request to the server.
- resp wait: The time spent waiting for the server's response.
- resp read: The time taken to read the response from the server.

Status code distribution:

The number of responses for each status code is provided. In this case, there were 140,629 responses with a status code of 200 (OK).

Overall, this analysis gives insights into the performance characteristics of the system, including the distribution of response times, latency percentiles, and details about different stages of the request-response cycle. It helps identify areas that may require optimization or further investigation to improve the system's performance.

## Generating a report

Make sure your app is running!

### CPU profile

The CPU profiler runs for 30 seconds by default. It uses sampling to determine which functions spend most of the CPU time. The Go runtime stops the execution every 10 milliseconds and records the current call stack of all running goroutines.

```bash
go tool pprof http://localhost:5000/debug/pprof/profile
```

After 30 seconds, you will see something like this:

```bash
(base) lifuyis-MacBook-Pro:Hexagonal-Architecture davidlee$ go tool pprof http://localhost:5000/debug/pprof/profile
Fetching profile over HTTP from http://localhost:5000/debug/pprof/profile
Saved profile in /Users/davidlee/pprof/pprof.samples.cpu.001.pb.gz
Type: cpu
Time: May 23, 2023 at 12:27am (CST)
Duration: 30s, Total samples = 0
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

When pprof enters the interactive mode, type top, the command will show a list of functions that appeared most in the collected samples. In our case these are all runtime and standard library functions, which is not very useful:

```bash
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 0, 0% of 0 total
      flat  flat%   sum%        cum   cum%
(pprof)
```

Now let’s load some requests to the server:

```bash
hey -n 10000000 http://localhost:5000/v1/users
```

Again, when fetching the profile and looking at the top:

```bash

```

Let’s check this in the graphical chart; we do this by using -http flag.

```bash
go tool pprof -http=:5001 http://localhost:5000/debug/pprof/profile
```

Browse to `http://localhost:5001`

![](images/graph.png)

### How to read graph?

I would suggest you to go through [this section](https://github.com/google/pprof/blob/master/doc/README.md#interpreting-the-callgraph) and know how you can read more clearly the graph.

### Heap profile

Run the heap profiler:

```bash
go tool pprof http://localhost:5000/debug/pprof/heap
```

In graphical chart:

```bash
go tool pprof -http=:5003 http://localhost:5000/debug/pprof/heap
```

![](images/heap.png)

But we are more interested in the number of allocated objects. Call pprof with -alloc_objects option:

```bash
go tool pprof -alloc_objects -http=:5002 http://localhost:5000/debug/pprof/heap
```

### Block profile

Blocking profile shows function calls that led to blocking on synchronization primitives like mutexes and channels.

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

In this way, you can check other profiles of your web app.

### Optimization tips

Avoid unnecessary heap allocations.

For big structures, it might be cheaper to pass a pointer than to copy the whole structure. But, prefer values over pointers for not big structures.

Go compiler is smart enough to turn some dynamic allocations into stack allocations. Things get worse for example when you start dealing with interfaces. So, preallocate maps and slices if you know the size beforehand.
Don’t log if you don’t have to.

Use buffered I/O if you do many sequential reads or writes.
If your application extensively uses JSON, consider utilizing parser/serializer generators.

Sometimes the bottleneck maybe not be what you are expecting — profiling is the best and sometimes the only way to understand the real performance of your application.

# 🥊 Adding `tcpdump` for Network Analysis

While Gin provides built-in logging functionality to measure and log the Round Trip Time (RTT) of requests, there may be situations where you would want to use tcpdump for network analysis. Here are some scenarios where tcpdump can be useful:

Network troubleshooting: tcpdump can be used to capture network packets and analyze their contents. If you encounter issues related to network connectivity, packet loss, or unexpected behavior, tcpdump can help you inspect the network traffic to identify potential problems.

Performance analysis: While Gin's built-in logging provides an overview of request processing time, tcpdump allows you to inspect the actual network packets exchanged between the client and server. This can provide more detailed insights into network performance, including packet latency, retransmissions, and other network-level metrics.

Security analysis: tcpdump can be used to capture and analyze network traffic for security purposes. It allows you to inspect packet payloads, detect potential vulnerabilities, or investigate suspicious network activity.

Protocol analysis: If you're working with custom protocols or need to debug protocol-level issues, tcpdump can help you capture and analyze the protocol-specific packets to understand the communication flow and identify any anomalies.

It's important to note that tcpdump operates at a lower level of the network stack compared to Gin's logging. It captures all network traffic, not just the traffic related to your specific application. This can provide a broader perspective on network behavior but may require additional analysis and filtering to focus on the relevant traffic.

In summary, while Gin's built-in logging is typically sufficient for measuring the RTT of your application's requests, tcpdump can be a valuable tool for in-depth network analysis, troubleshooting, performance analysis, and security assessment.

## To use tcpdump for network troubleshooting, you can follow these steps:

### Install tcpdump:

Ensure that tcpdump is installed on your system. The installation process may vary depending on your operating system.

```bash
brew install tcpdump
```

### Capture network packets:

Run tcpdump with appropriate options to capture network packets. For example, to capture all packets on a specific network interface (e.g., eth0), you can use the following command:

```bash
sudo tcpdump -i eth0
```

This will start capturing network packets on the specified interface. If it does not work, Use the `ifconfig` or `ip addr` command to list the available network interfaces on your system. Look for the interface you want to capture packets on. The interface name may differ depending on your system and network configuration.

### Filter captured packets:

By default, tcpdump captures all packets on the specified interface. However, you can use filters to narrow down the captured packets based on specific criteria. Filters can be based on source/destination IP addresses, port numbers, protocols, etc. For example, to capture only HTTP packets, you can use the following command:

```bash
sudo tcpdump -i eth0 port 80
```

This will capture only packets with the destination or source port set to 80 (HTTP).

### Analyze captured packets:

Once tcpdump is capturing packets, it will display information about each packet in real-time. This includes details such as source/destination IP addresses, port numbers, packet size, packet flags, etc. Analyze this information to identify any anomalies or issues.

## Packet Loss:

Packet loss refers to the situation where one or more network packets fail to reach their destination. It can occur due to various reasons such as network congestion, faulty network equipment, high latency, or other network issues. Packet loss can negatively impact network performance and result in degraded application performance.

Here's an example of using tcpdump to capture packets on a specific network interface (eth0) and filter based on source IP address:

```bash
sudo tcpdump -i eth0 src 192.168.1.100
```

This command will capture all packets coming from the source IP address 192.168.1.100 on the eth0 interface. You can modify the filter based on your specific requirements.

Remember to run tcpdump with appropriate permissions (e.g., using sudo) to capture packets effectively. Additionally, tcpdump provides numerous options and filters to customize the capture process further. You can refer to the tcpdump documentation or use the man tcpdump command in the terminal for more information on its usage and available options.

When you run sudo tcpdump -i en0, tcpdump starts capturing packets on the en0 network interface. By default, tcpdump will capture packets continuously until you stop it manually. To stop tcpdump, you can press Ctrl+C in the terminal where it's running.

The summary information at the end of the logs provides statistics about the captured packets:

```bash
606 packets captured
606 packets received by filter
0 packets dropped by kernel
```

- "606 packets captured" indicates the total number of packets captured by tcpdump.
- "606 packets received by filter" indicates the number of packets that passed the filter criteria specified in the tcpdump command.
- "0 packets dropped by kernel" indicates that no packets were dropped by the network stack/kernel during the capture.

## TCPDump Tutorial

[TCPDump-Tutorial](https://github.com/Samsar4/Ethical-Hacking-Labs/blob/master/11-Bonus/TCPDump-Tutorial.md)

# 🍕 Thoughts Collection on Recent Amazon Prime Video's Dump of its AWS Distributed Serverless Architecture and Move to “Monolith”

I think it is important for a software engineer to constantly keep track of the software architecture, so I brought this topic up to discussion with
Amazon Prime Video's case as a retrospection on our current architecture.

Following are the thoughts I collected from the internet:

- The main scaling bottleneck in the architecture was the orchestration management that was implemented using AWS Step Functions. Our service performed multiple state transitions for every second of the stream, so we quickly reached account limits. Besides that, AWS Step Functions charges users per state transition. The second cost problem we discovered was about the way we were passing video frames (images) around different components. To reduce computationally expensive video conversion jobs, we built a microservice that splits videos into frames and temporarily uploads images to an… S3 bucket. Defect detectors (where each of them also runs as a separate microservice) then downloaded images and processed it concurrently using AWS Lambda. However, the high number of Tier-1 calls to the S3 bucket was expensive.

- To some observers, “the design in the PV [Prime Video] article is problematic. Misusing services doesn’t fix architecture issues, it exposes them” as Lambros Petrou, a senior software engineer at DataDog added on Twitter; a view to some degree shared by former CTO Steve Chambers, who said: “Basically, they now use the same architecture but condense components into containers so they don’t have expensive calls and data transfers across a network between different cloud services… it's kind of an obvious optimization!

- One engineer added in a Reddit debate on the shift: “Microservices have overhead. What used to be a simple inter-process communication or even an in-memory call between two small parts of a system becomes a full HTTPS, OAuth, JSON encoding/decoding exercise every time one of those short conversations needs to happen. When your system is blown apart into 500,000 pieces and each communication requires that setup, AND you’re being billed for each transaction, the cost and complexity adds up. The reaction against monoliths was the need to replace the entire application in one shot, meaning developers would actually need to test stuff. DevOps means there’s no more testing and we fail forward in production, and the only way you can do that is by having tiny functional pieces so you can find/fix stuff fast. I don’t think there’s anything wrong with saying these super-chatty parts of the application belong together without the need to open millions of connections all the time…”

- “The relentless drumbeat of a distributed, microservices-based platform that decouples everything from data, network endpoints to segregated UX with various protocols was maddening without context” commented one global CTO on LinkedIn after reading the post, which was originally filed in March but just attracted attention across the engineering and broadly technology community this month, adding drily: “I wonder if cloud providers are now going to patternize and sell full stack monoliths on their platform.”
