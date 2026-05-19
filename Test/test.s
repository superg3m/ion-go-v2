.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $12, %rsp
	movl $0, -4(%rbp)
.L82:
	movl -4(%rbp), %r10d
	cmpl $5, %r10d
	movl $0, -8(%rbp)
	setl -8(%rbp)
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L83
	movl -4(%rbp), %r10d
	movl %r10d, -12(%rbp)
	addl $2, -12(%rbp)
	movl -12(%rbp), %r10d
	movl %r10d, -4(%rbp)
	jmp .L82
.L83:
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
