# dtype
Package dtype provides a definition of a Dtype, which is a part of the type system that Gorgonia uses.

The main type that this package provides is `Dtype{}`. `Dtype` is used within the Gorgonia family of libraries in its type system.

Further, this package also provides some definitions for type classes. These type classes are also used in the Gorgonia family of libraries.

# Type Classes #

What is a type class? A type class is a set of types that all have a particular feature.

Consider for example, a `string` and an `int`. Both these types allow a notion of "addition" - e.g.

```
// string "addition".
s := "hello"
s += "world"

// int addition.
i := 1
i += 1
```

So they are to be found in a type class called `Addable`.

Note that this library DOES NOT enforce the features. Instead, think of this library as holding a collection of known types that does a particular thing.


The list of type classes are:

| Type Class | What It Does/For |
|--|--|
| All | A collection of known types in Gorgonia's type system. |
| Specialized | A collection of types that receives specialized support/operations within Gorgonia's engines. |
| Addable | Types whose values can be "added" together. |
| Number | Types whose values are numbers. |
| Ord | Types whose values can be sorted in order. |
| Eq | Types whose values can be compared for equality. |
| Unsigned | Types whose values are unsigned numbers. |
| Signed | Types whose values are signed numbers. |
| SignedNonComplex | Types whose values are signed numbers, and are not complex numbers. |
| Floats | Types whose values are represented by floating point numbers. |
| Complexes | Types whose values are represented by floating point complex numbers. |
| FloatComplex | Types whose values are floating point numbers or complex numbers made up of floating point numbers. |
| NonComplexNumber | Types whose values are not complex numbers. |
| Generatable | Types whose values are generatable by a random number generator (strings are also generatable) |
