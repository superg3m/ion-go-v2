.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $12, %rsp
	movl $5, -4(%rbp)
	addl $54, -4(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -8(%rbp)
	addl $1, -8(%rbp)
	movl -8(%rbp), %r10d
	movl %r10d, -12(%rbp)
	subl $5, -12(%rbp)
	movl -12(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
