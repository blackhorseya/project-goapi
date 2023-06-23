# Documentation

## Project Structure

### Required

#### `/adapter`

Great! You have designed a top-level folder called `adapter`, which represents the entry points for all applications.
This structure helps organize the code and makes the entry points of the applications more clear and identifiable.

Within the `adapter` folder, you can add different entry points based on the requirements of your applications. For
example, if your application has a web API and a command-line interface (CLI), you can create the following structure
within the `adapter` folder:

```
adapter/
├── api/
│   └── main.go
└── cli/
    └── main.go
```

In the `api` folder, you can place the code for handling web API requests and start the web server in the `main.go`
file.

In the `cli` folder, you can place the code for handling command-line operations and define the respective commands and
command-line options in the `main.go` file.

This structure helps to clearly separate different application entry points, organizes the code, and provides better
scalability. Please adjust this structure according to your specific needs to fit your project.

#### `/entity`

In Clean Architecture, an Entity represents the core domain model or core business entity. They are pure objects that
are independent of any specific technology or framework, used to represent important concepts and rules in the domain.
Entities contain properties and methods related to business logic and are responsible for maintaining their own
consistency and integrity.

Entities should be stateless, meaning their state is not affected by any external operations, and they should adhere to
the business rules of the domain. Entities often include an identity attribute to uniquely identify each instance.

In Clean Architecture, Entities reside in the innermost core layer and do not depend on any external layers or
frameworks. Entities are the core of the entire application, and other layers (such as use cases, interfaces, etc.)
depend on the Entities to implement business logic and manipulate data.

The purpose of Entities is to capture domain knowledge and business logic, providing a clear interface for other layers
to use and manipulate these domain entities.

#### `/pb`

I have designed a top-level folder called `pb` which is specifically intended to store `proto` files.

#### `/pkg`

The `pkg` directory, located at the top level of the `hg-goapi` project, is typically used to store reusable code and
packages. Here are some common uses of the `pkg` directory:

1. Building generic utility classes: Create reusable classes or packages that encapsulate common functions, utility
   methods, or frequently used functionalities. These tools can be referenced and used by other parts of the project.

2. Implementing shared packages: If the project has multiple modules or services that require sharing the same code, you
   can place those shared packages in the `pkg` directory. This ensures easy sharing and reuse of these packages among
   different parts of the project.

3. Encapsulating lower-level packages: If the project needs to use lower-level packages or abstract the underlying
   layers, you can place those packages in the `pkg` directory. This ensures decoupling and flexibility between the
   project's other components and these underlying packages.

4. Creating independent modules or services: If the project consists of multiple independent modules or services, each
   with its own logic and functionalities, you can place the relevant code for each module or service in separate
   subdirectories within the `pkg` directory. This helps maintain clarity and separation of responsibilities between the
   modules or services, making modularity and testing easier.

It is important to maintain clear and consistent organizational structure when using the `pkg` directory and adhere to
the overall project architecture and design principles. Based on the project's needs and structure, you can decide how
to use the `pkg` directory to ensure code readability, maintainability, and reusability.

#### `/docs`

Design and user documents (in addition to your godoc generated documentation).

#### `/scripts`

Scripts to perform various build, install, analysis, etc operations.

#### `/test`

Additional external test apps and test data. Feel free to structure the /test directory anyway you want. For bigger
projects it makes sense to have a data subdirectory. For example, you can have /test/data or /test/testdata if you need
Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "_", so
you have more flexibility in terms of how you name your test data directory.

#### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

Examples:

- https://github.com/kubernetes/kubernetes/tree/master/api
- https://github.com/moby/moby/tree/master/api
