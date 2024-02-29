# Dependency Inversion Principle
The Dependency Inversion Principle is a fundamental principle of the SOLID design principles that guides the development of flexible and maintainable software. Dependency Inversion Principle encourages the use of abstractions and promotes loose coupling between modules or components.

The Dependency Inversion Principle states that high-level modules should not depend on low-level modules. Instead, both should depend on abstractions. This principle aims to invert the traditional dependency hierarchy, where higher-level modules depend on lower-level modules. By relying on abstractions, Dependency Inversion Principle promotes modular design, decoupling, and the ability to replace implementations without affecting the higher-level modules.

**Benefits of Dependency Inversion**:
-  Decoupling: By depending on abstractions, rather than concrete implementations, Dependency Inversion Principle reduces direct dependencies between modules. This loose coupling allows for independent development and modification of individual modules, making the system more maintainable and adaptable.

-  Modularity: Dependency Inversion Principle promotes the creation of well-defined interfaces or abstractions that represent the behaviors and contracts between components. This modularity facilitates code organization, enhances code reusability, and enables the composition of different modules to build complex systems.

-  Testability: The use of abstractions in Dependency Inversion Principle leads to code that is easily testable. By depending on interfaces or abstractions, you can create mock implementations or stubs during testing, isolating and verifying the behavior of individual modules or components.

This principle states that high-level modules should not depend on low-level modules, but rather both should depend on abstractions. This helps to reduce the coupling between components and make the code more flexible and maintainable.

In our code example, we have a struct Worker that represents a worker in a company, and a struct Supervisor that represents a supervisor. According to the Dependency Inversion Principle, high-level modules should not depend on low-level modules. Instead, both should depend on abstractions. In this case, the dependency direction is inverted, interfaces are determined at a higher level and classes in their same level depend on them, therefore they depend on abstractions. Additionally, lower-level class implementations depend on the interfaces defined at a higher level, therefore details now depend on abstractions.

This demonstrates the Dependency Inversion Principle, as the Department struct depends on an abstraction (the Employee interface), rather than on a specific implementation (the Worker or Supervisor struct). If both client and service depend on an abstraction, then you essentially have an agreement that, if respected, will allow for the following:
-  The client will be agnostic of who the dependency is, and instead rely on what it does
-  The dependency will be guaranteed to behave in a way that is determined by the high-level policy
-  The client can be reused in other contexts safely, trusting that its dependencies will respect the contract
-  The dependency can always be replaced by another implementation that implements the same contract

This makes the code more flexible and easier to maintain, as changes to the implementation of workers and supervisors will not affect the Department struct. Applying the Dependency Inversion Principle and consequently depending on the abstractions rather than concrete implementations, promotes flexibility, modularity, maintainability, extensibility, and testability.