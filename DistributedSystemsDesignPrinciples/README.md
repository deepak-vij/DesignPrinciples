# Platform Engineering Key Design Patterns (Common-Sensical)
Goal is to leverage the following well-proven design patterns as part of the **Platform Engineering** initiative - **Standing on the Shoulders of Giants** by learning from the prior experinces of system developers. Cloud-scale highly available distributed systems are designed to provides key features such as high-availability/failover, scalability/elasticity, security and portability.
- **Availability/Fault-Tolerance**: Cloud platform addresses highly availability and failovers both at application and infrastructure services level.
  - Use probes correctly to detect and automatically recover from failures.
  - Make the application & infrastructure components fail hard (crashing), fast (as soon as a problem occurs), and loudly (with informative error messages in their logs). Doing so prevents data from being stuck in a strange state in the failed application and allows routing of traffic only to healthy component instances, and also provides all the information needed for root cause analysis.
  - **Circuit-Breakers Pattern**: Similar to electrical fuses that prevent fires when a circuit that is connected to the electrical grid starts drawing a high amount of power which causes the wires to heat up and combust, the Similar to electrical fuses, circuit breaker design pattern is a fail-first mechanism that shuts down the circuit, request/response relationship or a service in order to prevent bigger failures. A simplistic implementation of this design pattern may be a simple counter that records success and failure states of a circuit along with a timestamp and calculates the consecutive number of failures. One can simply apply a wrapper component around potentially dangerous operations (typically outbound/egress) to circumvent calls when the system is not healthy. This is different versus retries as circuit brakers exist to prevent operations rather than re-execute them. 
  - **Bulkheads Pattern**: Bulkhead pattern exhibits a principle of damage containment by partitioning the overall system. With the replicated stateless services, each replica is entirely homogeneous and capable of serving every request. In contrast to replicated services, with sharded services, each replica, or shard, is only capable of serving a subset of all requests. A load-balancing node, or root, is responsible for examining each request and distributing each request to the appropriate shard or shards for processing.
- **Scalability**: Cloud platform provides an exteremely elastic environment by enabling horizontal/vertical scaling of workloads on the basis of their resource utilizations reaching the certain desired threshold/s. key goal is to enable automatic scaling to ensure overall capacity management.
  - **Observability/Monitoring**: ​Application & Infrastructure services are well prepared for observability/monitoring​, possibly using Ambassador design pattern.
  - **Replicated Load-Balanced Services Pattern**: In such a service, every server is identical to every other server (Stateless Services) and all are capable of supporting traffic. Stateless services are ones that don’t require saved state to operate correctly. In the simplest stateless applications, even individual requests may be routed to separate instances of the service. Examples of stateless services include things like static content servers and complex middleware systems that receive and aggregate responses from numerous different backend systems. This pattern consists of a scalable number of servers with a load balancer in front of them. The load balancer is typically either completely round-robin or uses some form of session stickiness.
  - **Scatter/Gather Pattern**: .
- **Security**: Platform addresses security at various levels: cluster, application and network. The API endpoints are secured through transport layer security (TLS/mTLS).
- **Portability**: Platform enables portability in terms of operating system choices, processor architectures (either virtual machines or bare metal), cloud providers, and various container runtimes, besides Docker, can also be added. It also supports workloads across hybrid (private and public cloud) or multi-cloud environments. This, in turn, also supports availability zone fault tolerance within a single cloud provider. 
- **General Design Patterns**
  - **Side-Car**: In order to provide Modularity and Reusability. in the sidecar pattern, the secondary container enhances the main application (typically containerized) by providing new functionality. The sidecar shares the same life cycle as the parent application, being created and retired alongside the parent. The sidecar pattern is typically made up of two containers. The first is the application container. It contains the core logic for the application. Without this container, the application would not exist. In addition to the application container, there is a sidecar container. The role of the sidecar is to augment and improve the application container, often without the application container’s knowledge. In its simplest form, a sidecar container can be used to add functionality to a container that might otherwise be difficult to improve. Sidecar containers are co-scheduled onto the same machine via an atomic container group. In addition to being scheduled on the same machine, the application container and sidecar container share a number of resources, including parts of the filesystem, hostname and network, and many other namespaces. Typical usage of this design pattern is the following:
    - Security/SSL-Termination
    - Obserability
    - Logging
    - Debugging
  - **Ambassador Pattern**: Ambassador design pattern is typically used for proxying outbound communication to and from a main container. For example, application maybe speaking the memcache protocol using a twemproxy ambassador. The application believes that it is simply talking to a single memcache on localhost, but actually twemproxy is sharding the requests across a distributed installation of multiple memcache nodes elsewhere in the cluster. Typical usage of this design pattern is the following:
    - Service Discovery/Brokering 
  - **Adapter**: Real-world application development is a heterogeneous, hybrid exercise. Some parts of your application might be written from scratch by your team, some supplied by vendors, and some might consist entirely of off-the-shelf open source or proprietary software that you consume as precompiled binary. The net effect of this heterogeneity is that any real-world application you deploy will have been written in a variety of languages, with a variety of conventions for logging, monitoring, and other common services. Yet, to effectively monitor and operate your application, you need common interfaces. When each application provides metrics using a different format and interface, it is very difficult to collect all of those metrics in a single place for visualization and alerting. This is where the adapter pattern is relevant. The adapter pattern uses a secondary container to standardise external interfaces. In contrast to the ambassador pattern, which presents an application with a simplified view of the outside world, adapters present the outside world with a simplified, homogenized view of an application.They do this by standardizing output and interfaces across multiple containers.Typical usage of this design pattern is the following:
    - Standardise log formats
    - Provide common metrics for applications
  - **Ownership Election**:...Typical usage of this design pattern is the following: