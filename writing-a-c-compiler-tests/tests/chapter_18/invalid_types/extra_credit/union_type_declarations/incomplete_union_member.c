/* You can't declare a union with an incomplete member type */
struct s;
union u {
    struct s bad_struct;
};

int main(){
    return 0;
}