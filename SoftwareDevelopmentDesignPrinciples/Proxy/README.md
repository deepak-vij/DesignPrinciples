# Proxy
**Proxy** is a structural design pattern that provides an object that acts as a substitute for a real service object used by a client. A proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object.

The proxy object has the same interface as a service, which makes it interchangeable with a real object when passed to a client.

See the companion [code example](/SoftwareDevelopmentDesignPrinciples/Proxy). In this example, a web server such as Nginx can act as a proxy for the application server for providing:
- Controlled access to the application server.
- Do rate limiting.
- Do request caching.
- etc.