# going-forth
A toy Forth interpreter written to learn about how Forth works
and get more practice with Go.

This is a prototype before I attempt writing a Forth interpreter
in ARM assembly.


## Features
This is missing a lot of features expected of a typical Forth, 
but this is the minimal functionality I wanted to play around with.

- Built-in words (see below)
- Comments
  - line comments starting with ```--``` or ```\```
  - inline comments within ```( )```
- Compile custom words with format ```: doubleme 2 * ;```


### Example
```forth
\ examples/define.fth
\
\ example of compiling words

: doubleme ( n1 -- n2 ) 2 * ;
: quadme ( n1 -- n2 ) doubleme doubleme ;

8 doubleme . cr  \ 16
8 quadme . cr    \ 32

: hello 79 76 76 69 72 emit emit emit emit emit ;

cr hello cr  \ HELLO
```

More examples of features can be found in [forth-pgms/](forth-pgms/)


### Missing Features
I did say this was a "toy Forth interpreter" so it is missing a lot. 
I hope to tackle a more feature complete version in the future.

- Zero-relational operators - ```0=  0<>  0<  0>  0<=  0>=```
- String literals - ```."```
- do loop
- if, then, else
- Return stack functionality - ```r>```, ```r@```, ```.r```
- Variables, constants
- File includes


### Built-in Words

- ```: +  ( n1 n2 -- n3 )```
- ```: -  ( n1 n2 -- n3 )```
- ```: *  ( n1 n2 -- n3 )```
- ```: /  ( n1 n2 -- n3 )```
- ```: 1+  ( n1 -- n2 )```
- ```: 1-  ( n1 -- n2)```
- ```: 2*  ( n1 -- n2)```
- ```: 2/  ( n1 -- n2)```
- ```: /mod  ( n1 n2 -- n3 n4 )```
- ```: mod  ( n1 n2 -- n3 )```
- ```: =  ( n1 n2 -- flag )```
- ```: <>  ( n1 n2 -- flag )```
- ```: <  ( n1 n2 -- flag )```
- ```: >  ( n1 n2 -- flag )```
- ```: <=  ( n1 n2 -- flag )```
- ```: >=  ( n1 n2 -- flag )```
- ```: true  ( -- flag )```
- ```: false  ( -- flag )```
- ```: dup  ( n -- n n )```
- ```: 2dup  ( n1 n2 -- n1 n2 n1 n2 )```
- ```: drop  ( n -- )```
- ```: 2drop  ( n1 n2 -- )```
- ```: swap  ( n1 n2 -- n2 n1 )```
- ```: over  ( n1 n2 -- n1 n2 n1 )```
- ```: rot  ( n1 n2 n3 -- n2 n3 n1 )```
- ```: .  ( n -- )```
- ```: .s  ( -- )```
- ```: .r  ( -- )``` (non-functional)
- ```: .cr  ( -- )```
- ```: .emit  ( n -- )```


## References
- [Starting FORTH](https://www.forth.com/starting-forth/)
- [Moving FORTH](http://www.bradrodriguez.com/papers/moving1.htm)
- [The infamous jonesforth.S](https://github.com/nornagon/jonesforth/blob/master/jonesforth.S)
- [Easy Forth](https://skilldrick.github.io/easyforth/)
