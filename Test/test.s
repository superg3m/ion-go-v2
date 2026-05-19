.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $12, %rsp
	movl $0, -4(%rbp)
.L1:
	movl -4(%rbp), %r10d
	cmpl $5, %r10d
	movl $0, -8(%rbp)
	setl -8(%rbp)
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L2
	movl -4(%rbp), %r10d
	cmpl $3, %r10d
	movl $0, -12(%rbp)
	sete -12(%rbp)
	movl -12(%rbp), %r10d
	cmpl $0, %r10d
	je .L4
	jmp .L2
	jmp .L4
.L4:
	movl -4(%rbp), %r10d
	movl %r10d, -4(%rbp)
	addl $1, -4(%rbp)
	jmp .L1
.L2:
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
