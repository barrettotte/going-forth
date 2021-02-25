# going-forth
A toy Forth interpreter written to learn about how Forth works
and get more practice with Go.


## Features
- Built-in words (see below)
- 


## Built-in Words
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
- ```:.r  ( -- )```
- ```:.cr  ( -- )```
- ```:.emit  ( n -- )```


## References
- [Starting FORTH](https://www.forth.com/starting-forth/)
- [Moving FORTH](http://www.bradrodriguez.com/papers/moving1.htm)
- [Easy Forth](https://skilldrick.github.io/easyforth/)
- [The infamous jonesforth.s](https://github.com/nornagon/jonesforth/blob/master/jonesforth.S)