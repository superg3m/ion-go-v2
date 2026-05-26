# ion-go-v2

A small C-lite toy compiler written in Go that generates x86_64 AT&T syntax assembly.

## Notes

This is probably the hardest "Hello, World!" I have ever done beating out LearnOpenGL's triangle/quad.
For anyone curious about the process: I originally started with an interpreter, then slowly transformed it into a compiler 
by reusing and extending pieces like the lexer, parser, and type system.

One resource I found especially helpful was *Writing a C Compiler* by Nora Sandler. 
It does a great job explaining both the theory and implementation details behind building a compiler.
I also want to be transparent about AI usage in this project. 
I personally do not enjoy using AI, since this project is primarily educational and I genuinely enjoy writing the code myself. 
Because of that, essentially every line of this toy compiler was written by hand at some point in time.

That said, I do occasionally use AI for tedious or peripheral tasks, such as:
- helping write documentation like this README
- explaining assembly behavior: one example I remember is `setl` operates on the first byte not
the whole register, this means that if the higher bytes are not cleared to zero then you can't expect setl to
actually do the thing you would like it to.

---

## Features
Currently supported:
- Integer expressions
- Arithmetic operations
- Control flow
  - Short-circuit evaluation (`&&`, `||`)
  - `return, break, continue`
- Function calls
  - Recursion
  - `putchar`


### Current limitations
- Only supports the `int` type

---

## Compilable Examples
### Hello World

```c
// EXPECT: 0
int putchar(int c);

int main() {
    putchar(72);
    putchar(101);
    putchar(108);
    putchar(108);
    putchar(111);
    putchar(44);
    putchar(32);
    putchar(87);
    putchar(111);
    putchar(114);
    putchar(108);
    putchar(100);
    putchar(33);
    putchar(10);

    return 0;
}
```

Output:

```txt
Hello, World!
```

### Function Calls + Checking Register Clobbering
```c
// EXPECT: 1
int x(int a, int b, int c, int d, int e, int f) {
    return a == 1 &&
           b == 2 &&
           c == 3 &&
           d == 4 &&
           e == 5 &&
           f == 6;
}

int main() {
    int a = 4;
    return x(1, 2, 3, 4, 5, 24 / a);
}
```

## Planned Features (One day!)
- Pointers
- Arrays
- Structs
- Strings
- Type checking
- Better diagnostics/errors
- Register allocation 
- Optimizations
---