# Hexagonal Architecture Deep Dive with PostgreSQL, Redis and Go Practices

![](images/arch.webp)
This is the source code for the original article:

[Hexagonal Architecture Deep Dive with PostgreSQL, Redis and Go Practices](https://medium.com/towardsdev/hexagonal-architecture-deep-dive-with-postgresql-redis-and-go-practices-4b051f940e93)

Note that the codebase has evolved with more complexity than the article examples.

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

## 🤡 Application structure

Now, let's dive into how to create a messaging backend that allows users to save and read messages. Hexagonal architecture adheres to strict application layout that needs to be implemented. Below is the application layout that we will use. This might look like a lot of work, but it will make sense as we move forward.

![](images/structure.png)

# 👺 To-dos:

- ✅ Finish CRUD process of the messaging service
- ✅ REST API Design
- ✅ Add User
- ✅ Add JWT Authentication
- ⌛️ Alter the whole project with Redis as cache, postgresql as database
- ⌛️ Add Unit Test
- ⌛️ Add a payment service
- ⌛️ Add Distributed system

# Pros and Cons of using GORM in this project

GORM is a popular Object-Relational Mapping (ORM) library for the Go programming language that provides a convenient way to interact with databases, including PostgreSQL.

It provides a high-level, expressive and easy-to-use API for CRUD (Create, Read, Update, Delete) operations and supports several databases, including MySQL, PostgreSQL, SQLite, and others.

Whether GORM is better to use than directly using PostgreSQL depends on the specific use case. If you need a high-level, user-friendly API to interact with your PostgreSQL database, then GORM can be a great choice. On the other hand, if you have specific requirements for your database interactions or need to optimize performance for a large-scale application, then direct interaction with the PostgreSQL database using a lower-level database driver may be more appropriate.

In general, the use of an ORM can simplify and speed up development, especially for CRUD operations. However, it may introduce additional overhead and performance concerns.
