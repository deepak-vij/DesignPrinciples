# Singleton
**Singleton** is a creational design pattern, which ensures that only one object of its kind exists and provides a single point of access to it for any other code.

The key goal of the Singleton pattern is to **Ensure that a class/type has just a single instance**. The key usage of this pattern is to cover an instance where there must be a single "broker" controlling access to a resource. The most common reason for this is to control access to some shared resource—for example, shared caching, common logging etc. This Pattern is leveraged to create a cache manager that maintains a single cache instance. This ensures consistent caching strategies throughout the application, avoiding duplication of cache data and optimizing memory usage.

Similarly, a single common Logger manager may be used to manage logging operations throughout the application by ensuring that only one instance of the Logger exists, and all parts of the system can access it easily.

See the companion [code example](/SoftwareDevelopmentDesignPrinciples/Singleton).