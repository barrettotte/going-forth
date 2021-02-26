# going-forth
A toy Forth interpreter written to learn about how Forth works
and get more practice with Go.

This is a prototype before I attempt writing a Forth interpreter
in ARM assembly.


## Features
This is missing some features expected of Forth, 
but this is the minimal functionality I wanted.

- Built-in words (see below)
- Comments
  - line starting with ```--``` or ```\```
  - inside ```( )```
- Compile custom words with format ```: doubleme 2 * ;```

Actual examples of features can be found in [examples/](examples/)


TODO: ```cr``` is non-functional unless its written twice??

TODO: error ```1 2dup . . . cr  -- 111  TODO: ???```

TODO: attempt do loop

TODO: attempt if,then,else


### Missing Features
I did say this was a "toy Forth interpreter" so it is missing some things. 
I hope to tackle a more feature complete version in the future.

- ```."``` word for declaring string literals
- Return stack functionality
- Variables, constants


### Built-in Words
TODO: 0 rel ops

- ```:+  ( n1 n2 -- n3 )```
- ```:-  ( n1 n2 -- n3 )```
- ```:*  ( n1 n2 -- n3 )```
- ```:/  ( n1 n2 -- n3 )```
- ```:1+  ( n1 -- n2 )```
- ```:1-  ( n1 -- n2)```
- ```:2*  ( n1 -- n2)```
- ```:2/  ( n1 -- n2)```
- ```:/mod  ( n1 n2 -- n3 n4 )```
- ```:mod  ( n1 n2 -- n3 )```
- ```:=  ( n1 n2 -- flag )```
- ```:<>  ( n1 n2 -- flag )```
- ```:<  ( n1 n2 -- flag )```
- ```:>  ( n1 n2 -- flag )```
- ```:<=  ( n1 n2 -- flag )```
- ```:>=  ( n1 n2 -- flag )```
- ```:true  ( -- flag )```
- ```:false  ( -- flag )```
- ```:dup  ( n -- n n )```
- ```:2dup  ( n1 n2 -- n1 n2 n1 n2 )```
- ```:drop  ( n -- )```
- ```:2drop  ( n1 n2 -- )```
- ```:swap  ( n1 n2 -- n2 n1 )```
- ```:over  ( n1 n2 -- n1 n2 n1 )```
- ```:rot  ( n1 n2 n3 -- n2 n3 n1 )```
- ```:.  ( n -- )```
- ```:.s  ( -- )```
- ```:.r  ( -- )``` (non-functional)
- ```:.cr  ( -- )```
- ```:.emit  ( n -- )```


## References
- [Starting FORTH](https://www.forth.com/starting-forth/)
- [Moving FORTH](http://www.bradrodriguez.com/papers/moving1.htm)
- [Easy Forth](https://skilldrick.github.io/easyforth/)
- [The infamous jonesforth.s](https://github.com/nornagon/jonesforth/blob/master/jonesforth.S)
