/* It's illegal to multiply, divide, or take the modulo of pointers */
int main()
{
    int x = 10;
    int *y = &x;
    (y / 8);
    return 0;
}