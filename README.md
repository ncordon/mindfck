# mindfck

A simple language that transpiles to brainfuck:

```
byte a
byte b
a = 3 + a
a = 33 + a
a = a + 2
byte c
a = a + 21
a = a + 2
b = 10 + a
c = a + b
print c
```

Transpiles to:

```brainfuck
>>[-]+++>>[-]<[-]<<<[>>>+>+<<<<-][-]>>>>[<<<<+>>>>-]>>[-]<<[-]<<[>>+>>+<<<<-][-]>>>>[<<<<+>>>>-]>[-]<<<<[>>>>+<<<+<-][-]>>>>[<<<<+>>>>-]<<[-]<[>+<-]<<<<[-]>>>>>[<<<<<+>>>>>-]>[-]+++++++++++++++++++++++++++++++++<<[-]>>>[-]<<<<<<<[>>>>>>>+<<<+<<<<-][-]>>>>[<<<<+>>>>-]<[-]>>[-]>[<+<<+>>>-][-]<<<[>>>+<<<-]>[-]>>>[<<<+>+>>-][-]<<<[>>>+<<<-]<<[-]>>>[<<<+>>>-]<<<<<[-]>>[<<+>>-]>>[-]<[-]<<<[>>>+>+<<<<-][-]>>>>[<<<<+>>>>-]>>>[-]++<<[-]<<<[-]>[<+>>>+<<-][-]>>[<<+>>-]<[-]>>>[<<<+<<+>>>>>-][-]<<<[>>>+<<<-]>>[-]<<<<[>>>>+<<<<-]<<[-]>>>>>>[<<<<<<+>>>>>>-]<<<<[-]>>[-]<<<<[>>>>+<<+<<-][-]>>[<<+>>-]>[-]+++++++++++++++++++++<[-]>>>>>[-]<<<[>>>+<<<<<+>>-][-]<<[>>+<<-]>>>>>>[-]<<<<<[>>>>>+<+<<<<-][-]>>>>>[<<<<<+>>>>>-]<<[-]>[<+>-]<<<<<<<[-]>>>>>>[<<<<<<+>>>>>>-]>>[-]<<<<<<[-]<<[>>+>>>>>>+<<<<<<<<-][-]>>>>>>>>[<<<<<<<<+>>>>>>>>-]<<<<<[-]++>>>>[-]<[-]<<<<[>>>>+>+<<<<<-][-]>>>>>[<<<<<+>>>>>-]>[-]<<<<<[>>>>>+<<+<<<-][-]>>>>>[<<<<<+>>>>>-]<<<<[-]>>[<<+>>-]<<<<<<[-]>>>>[<<<<+>>>>-]>>>[-]++++++++++<[-]>>[-]<<<<<<<<[>>>>>>>>+<<+<<<<<<-][-]>>>>>>[<<<<<<+>>>>>>-]<<<[-]>[-]>>>[<<<+<+>>>>-][-]<<<<[>>>>+<<<<-]>>>[-]>>[<<+<<+>>>>-][-]<<[>>+<<-]<<<<[-]>>[<<+>>-]<<<[-]>[<+>-]>>>>[-]<<<[-]<<<[>>>+>>>+<<<<<<-][-]>>>>>>[<<<<<<+>>>>>>-]>[-]>[-]<<<<<<<[>>>>>>>+<+<<<<<<-][-]>>>>>>[<<<<<<+>>>>>>-]<<<<<[-]>>>>[-]<<<[>>>+<<<<+>-][-]<[>+<-]>>>>>[-]>[<+<+>>-][-]<[>+<-]<<<[-]>>[<<+>>-]<[-]<[>+<-]>>>[-]<<<<<[-]>>>[<<<+>>>>>+<<-][-]>>[<<+>>-]<<<<<.
```

## FAQ

Q: Is this useful?  
A: No  
Q: Is it fun?  
A: Also no  

## Resources
These resources where used as inspiration / tools for this project

-   https://github.com/LucasMW/Headache
-   https://esolangs.org/wiki/Brainfuck_algorithms
-   https://www.nayuki.io/page/brainfuck-interpreter-javascript
-   https://gist.github.com/roachhd/dce54bec8ba55fb17d3a
