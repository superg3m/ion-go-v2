.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $24, %rsp
	movl $1, -4(%rbp)
	movl $5, -8(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -12(%rbp)
	addl $1, -12(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl -4(%rbp), %r10d
	cmpl $2, %r10d
	movl $0, -16(%rbp)
	sete -16(%rbp)
	movl -16(%rbp), %r10d
	cmpl $0, %r10d
	je .L1
	movl -8(%rbp), %r10d
	cmpl $1, %r10d
	movl $0, -20(%rbp)
	sete -20(%rbp)
.L2:
	movl -20(%rbp), %r10d
	cmpl $0, %r10d
	je .L1
	movl $1, -24(%rbp)
	jmp .L3
.L1:
	movl $0, -24(%rbp)
.L3:
	movl -24(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
.global get_integer
get_integer:
	pushq %rbp
	movq %rsp, %rbp
	subq $0, %rsp
	movl $4, %eax
	movq %rbp, %rsp
	popq %rbp
	ret
