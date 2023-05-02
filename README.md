# Hexagonal Architecture Deep Dive with PostgreSQL, Redis and Go Practices

![](images/arch.webp)
Source Code for the article:

[Hexagonal Architecture Deep Dive with PostgreSQL, Redis and Go Practices](https://medium.com/towardsdev/hexagonal-architecture-deep-dive-with-postgresql-redis-and-go-practices-4b051f940e93)

## What is Hexagonal Architecture?

Hexagonal Architecture, also known as Ports and Adapters Architecture or Clean Architecture, is a software architecture pattern that emphasizes the separation of concerns between the core business logic of an application and the external systems it interacts with. In this pattern, the core business logic is at the center of the architecture, surrounded by adapters that allow it to interact with the outside world. The adapters can be thought of as the “ports” through which the application communicates with external systems.

## Hexagonal Architecture Components

Here are the components of the Hexagonal Architecture:

### Core Business Logic:

This is the central component of the architecture, and it contains the application’s core domain logic. This component should be independent of any external systems and should not be affected by changes in the infrastructure or the user interface.

### Adapters:

These are the components that connect the core business logic to the external systems. They can be thought of as the “ports” through which the application communicates with the outside world. Adapters can take many forms, including APIs, databases, user interfaces, and messaging systems.

### Primary Adapter:

This is the adapter that handles the application’s primary input/output. For example, in a web application, the primary adapter might be an HTTP server that accepts incoming requests and returns responses. The primary adapter is responsible for translating incoming requests into domain-specific operations that can be processed by the core business logic, and translating the responses back into a format that can be understood by the requesting system.

### Secondary Adapters:

These are the adapters that handle the application’s secondary input/output. They can be thought of as “plugins” that provide additional functionality to the application. For example, a secondary adapter might be a database adapter that stores data for the application.

### Interfaces:

These are the contracts that define the communication between the core business logic and the adapters. They ensure that the adapters provide the necessary functionality to the core business logic, and that the core business logic provides the necessary information to the adapters. Interfaces can be thought of as the “language” that the adapters and the core business logic use to communicate with each other.

### Dependencies:

These are the external libraries or services that the application depends on. They are managed by the adapters, and should not be directly accessed by the core business logic. This allows the core business logic to remain independent of any specific infrastructure or technology choices.

## Application structure

Today, you will learn how to create a messaging backend that allows users to save and read messages. Hexagonal architecture adheres to strict application layout that needs to be implemented. Below is the application layout that you will use. This might look like a lot of work, but it will make sense as we move forward. Go ahead and create the below application structure.

![](images/structure.png)
