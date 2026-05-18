// Can't perform bitshift operations with void operands
void f(){
    return;
}

int main() {
    int x = 10;
    x << f();
    return 0;
}