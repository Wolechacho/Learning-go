## Data types in golang

There are two data types in golang

Basic Data types
string
int
rune
byte
complex number
float
Boolean

Composite Data types

The composite datatype can be further broken down into

Non Reference types
Structs
Array

Reference
slices
map
pointers
channels
functions/methods

Interfaces

Summary
you cannot compare slices,functions,map of the same data types
you can compare struct,array,pointers,channel,interface of the same data types
structs with same fields values are TRUE ,with different value fields are false,these also applies to array

special case for interface
you can compare similar data types
you can compare different data types

Global Variable
A variable will be global within a package if it is declared at the top of a file outside the scope of any function or block.
If this variable name starts with a lowercase letter then it can be accessed from within the the package which contains this variable definition.
If the variable name stats with a uppercase letter then it can be accessed from outside different package other than which it is declared.
Global variable are available throughout the lifetime of a program
