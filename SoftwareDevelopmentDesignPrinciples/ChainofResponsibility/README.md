# Chain of Responsibility
**Chain of Responsibility** is behavioral design pattern that allows passing request along the chain of potential handlers until one of them handles request. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

With this pattern, the handlers are linked into a chain. Each linked handler has a field for storing a reference to the next handler in the chain. In addition to processing a request, handlers pass the request further along the chain. The request travels along the chain until all handlers have had a chance to process it.

See the companion [code example](/SoftwareDevelopmentDesignPrinciples/ChainofResponsibility).