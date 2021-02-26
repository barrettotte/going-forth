\ define a word

: doubleme ( n1 -- n2 ) 2 * ;

8 doubleme . cr cr

: quadme ( n1 -- n2 ) doubleme doubleme ;

8 quadme . cr