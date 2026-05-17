.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	movl $5, -4(%rbp)
	movl -4(%rbp), %r10d
	imull $2, %r10d
	movl %r10d, -4(%rbp)
	movl %r10d, -8(%rbp)
	movl -4(%rbp), %r10d
	addl $4, -8(%rbp)
	movl -8(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
