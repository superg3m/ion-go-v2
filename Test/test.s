.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $5, -4(%rbp)
	movl $11, -4(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl -8(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
