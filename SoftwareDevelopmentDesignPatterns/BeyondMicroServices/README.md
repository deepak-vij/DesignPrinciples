# Towards Modern Development of Cloud Applications – beyond “Microservices”
Google recently unveiled a new programming model called “Service Weaver”. This is quite a profound development as it addresses all the issues software development community is currently mired in while dealing with development & management of “Microservices” architecture-based software development process. The following provides a description of all this.

In the last several decades of software development journey, there has been profound changes in the overall software development architecture from initially Monolithic -> SOA (Service Oriented Architecture) -> Microservices.

Monolithic architecture is where the software program is designed to be self-contained, and all software program components are tightly coupled rather than loosely coupled, each component and its associated components must all be present for code to be executed or compiled and for the software to run. In monolithic architecture we have a single code base which couples all together the business logic of the application. Key advantages of such an architecture are the following.
-  Easy deployment
-  Simplified testing
-  Easy debugging
-  High performance

However, monolithic architecture has the following key disadvantages:
-  Scalability: Scalability becomes a tedious task as we expand and onboard more users.
-  Deployment: A change in a monolithic application will require a new deployment of the entire monolith.
-  Reliability: This is less reliable since a failure in any module will cause the downtime of the whole application.
-  Flexibility: Any fluctuation or change in the technology or stack or updates in dependencies of the project will affect the whole project. This makes changes very expensive and time-consuming and also leads to a constraint in technology usage.

Due to the limitations of monolithic architecture, in the last couple of decades or so there has been a drastic push towards service oriented architecture – initially “SOA” and currently “Microservices”. Such an architecture has huge advantages in the overall Agility; Flexible scaling; Maintenance, Testing and Troubleshooting; Independent deployment; and finally Technology flexibility.

However, Service Oriented architecture has several disadvantages, specifically Microservices:

-  Development issues: Due to the number of different services, we can face some issues in managing the large number of teams. Using microservices adds more complexity as compared to a monolith.

-  Debugging challenges: Each microservice has its own set of logs, which makes debugging more complicated. Plus, a single business process can run across multiple machines, further complicating debugging.
-  Communication challenges

One of the fundamental problems with the Microservices architecture is the fact it conflates the logical boundaries (how code is written) with physical boundaries (how code is deployed) – resulting in a massive hairball of microservices architecture.

To address this issue, Google folks are proposing a different programming methodology that reduces application latency by up to 15x and cost by up to 9x --- decoupling development from distributed deployment. Their open source implementation is available at [Service Weaver](github.com/ServiceWeaver). Service Weaver consists of two core pieces:
-  A set of programming libraries, which let you write your application as a single modular binary, using only native data structures and method calls.
-  A set of deployers, which let you configure the runtime topology of your application and deploy it as a set of microservices, either locally or on the cloud of your choosing.

[Research Paper](https://lnkd.in/g7qpMiDR)

Berkeley "Ray" framework already does similar distributed processing for ML/AI applications. UW Sapphire research project had similar aspirations too (Irene Zhang, Prof. Arvind Krishnamurthy – University of Washington).

**Motivation for Building Service Weaver**
<br>
While writing microservices-based applications, we found that the overhead of maintaining multiple different microservice binaries—with their own configuration files, network endpoints, and serializable data formats—significantly slowed our development velocity.

More importantly, microservices severely impacted our ability to make cross-binary changes. It made us do things like flag-gate new features in each binary, evolve our data formats carefully, and maintain intimate knowledge of our rollout processes. Finally, having a predetermined number of specific microservices effectively froze our APIs; they became so difficult to change that it was easier to squeeze all of our changes into the existing APIs rather than evolve them.

As a result, we wished we had a single monolithic binary to work with. Monolithic binaries are easy to write: they use only language-native types and method calls. They are also easy to update: just edit the source code and re-deploy. They are easy to run locally or in a VM: simply execute the binary.

Service Weaver, is a framework that has the best of both worlds: the development velocity of a monolith, with the scalability, security, and fault-tolerance of microservices.