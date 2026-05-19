.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $1, -4(%rbp)
	movl $2, -8(%rbp)
	movl $4, -8(%rbp)
	movl -8(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
