// Can't use void controlling expression in switch statement
void f() {
    return;
}

int main() {
    switch(f()) {
        default: return 0;
    }
}