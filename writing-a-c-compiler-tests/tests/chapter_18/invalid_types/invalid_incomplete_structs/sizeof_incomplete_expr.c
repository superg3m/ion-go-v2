struct s;
struct s *get_ptr();

int main() {
    struct s *struct_ptr = get_ptr();

    // can't apply sizeof to expression w/ incomplete type
    return sizeof(*struct_ptr);
}