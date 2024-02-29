# Composite
**Composite** is a structural design pattern that allows you to combine objects from the same “family” in a tree-like structure and work with it as if it were a single object. Allowing access to each of the objects to execute their methods.

This pattern is commonly used to solve the cases when we need to handle tree structures because is easy to iterate each one of its child/inner objects. To define this pattern, we can use a concrete type (the struct type) or we can use the interface type to define the child objects.

See the companion [code example](/SoftwareDevelopmentDesignPrinciples/Composite). In this example, imagine that we have a file system and there we have folders and files. Now Imagine that you need to develop a code to search in the file system by a given string, there is when the “Composite Pattern” comes to action. The pattern will help us to organize the data structure and help to handle the folder and files iterations.