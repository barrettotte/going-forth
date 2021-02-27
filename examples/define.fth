\ example of compiling words

: doubleme ( n1 -- n2 ) 2 * ;
: quadme ( n1 -- n2 ) doubleme doubleme ;

8 doubleme . cr  \ 16
8 quadme . cr  \ 32

: hello 79 76 76 69 72 emit emit emit emit emit ;

cr hello cr  \ HELLO