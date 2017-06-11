/*

Methods and pointer indirection

Comparing the previous two programs, you might notice that functions 
with a pointer argument must take a pointer.

while methods with pointer receivers take either a value 
or a pointer as the receiver when they are called.
For the statement v.Scale(5), even though v is a value and not a pointer, 
the method with the pointer receiver is called automatically. 
That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) 
since the Scale method has a pointer receiver.

*/


/*

Methods and pointer indirection (2)

The equivalent thing happens in the reverse direction.
Functions that take a value argument must take a value of that specific type.
while methods with value receivers take either a value or a pointer as the receiver when they are called.

In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/


/*
Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.
The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. 
This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, 
even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, 
but not a mixture of both. (We'll see why over the next few pages.)
*/