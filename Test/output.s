.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	movl $2, -4(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl %eax, -8(%rbp)
	movq %rbp, %rsp
	popq %rbp
	ret
