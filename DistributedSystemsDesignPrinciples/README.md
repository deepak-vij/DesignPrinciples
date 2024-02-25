# Platform Engineering Design Patterns (Common-Sensical)
Goal is to leverage the following well-proven design patterns as part of the **Platform Engineering** initiative - **Standing on the Shoulders of Giants** by learning from the prior experinces of system developers. Cloud-scale highly available distributed systems are designed to provides key features such as high-availability/failover, scalability/elasticity, security and portability.
- **Availability/Fault-Tolerance**: Cloud platform addresses highly availability and failovers both at application and infrastructure services level.
  - Use probes correctly to detect and automatically recover from failures.
  - Make the application & infrastructure components fail hard (crashing), fast (as soon as a problem occurs), and loudly (with informative error messages in their logs). Doing so prevents data from being stuck in a strange state in the failed application and allows routing of traffic only to healthy component instances, and also provides all the information needed for root cause analysis.
  - **Circuit-Breakers Pattern*: 
  - **Bulkheads Pattern**: Sharded Services.
- **Scalability**: Cloud platform provides an exteremely elastic environment by enabling horizontal/vertical scaling of workloads on the basis of their resource utilizations reaching the certain desired threshold/s. key goal is to enable automatic scaling to ensure overall capacity management.
  - Application & Infrastructure services are well prepare for observability/monitoring​.
  - **Replicated Load-Balanced Services Pattern**: 
  - **Scatter/Gather Pattern: .
- **Security**: Platform addresses security at various levels: cluster, application and network. The API endpoints are secured through transport layer security (TLS/mTLS).
- **Portability**: Platform enables portability in terms of operating system choices, processor architectures (either virtual machines or bare metal), cloud providers, and various container runtimes, besides Docker, can also be added. It also supports workloads across hybrid (private and public cloud) or multi-cloud environments. This, in turn, also supports availability zone fault tolerance within a single cloud provider. 
- **General Design Patterns**
  - **Side-Car**: 
  - **Service Discovery Ambassador Pattern**: 
  - **Adapter**: 
  - **Ownership Election**: 