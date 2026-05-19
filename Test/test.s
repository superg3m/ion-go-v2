.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $44, %rsp
	movl $0, -4(%rbp)
	movl $0, -8(%rbp)
	movl $0, -12(%rbp)
	jmp .L25
.L23:
	movl -12(%rbp), %r10d
	movl %r10d, -16(%rbp)
	addl $1, -16(%rbp)
	movl -16(%rbp), %r10d
	movl %r10d, -12(%rbp)
.L25:
	movl -12(%rbp), %r10d
	cmpl $10, %r10d
	movl $0, -20(%rbp)
	setle -20(%rbp)
	movl -20(%rbp), %r10d
	cmpl $0, %r10d
	je .L24
	movl -12(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl -12(%rbp), %eax
	cdq
	movl $2, %r10d
	idiv %r10d
	movl %edx, -24(%rbp)
	movl -24(%rbp), %r10d
	cmpl $0, %r10d
	movl $0, -28(%rbp)
	sete -28(%rbp)
	movl -28(%rbp), %r10d
	cmpl $0, %r10d
	je .L27
	jmp .L23
.L27:
	movl -4(%rbp), %r10d
	movl %r10d, -32(%rbp)
	addl $1, -32(%rbp)
	movl -32(%rbp), %r10d
	movl %r10d, -4(%rbp)
	jmp .L23
.L24:
	movl -4(%rbp), %r10d
	cmpl $5, %r10d
	movl $0, -36(%rbp)
	sete -36(%rbp)
	movl -36(%rbp), %r10d
	cmpl $0, %r10d
	je .L28
	movl -8(%rbp), %r10d
	cmpl $10, %r10d
	movl $0, -40(%rbp)
	sete -40(%rbp)
.L29:
	movl -40(%rbp), %r10d
	cmpl $0, %r10d
	je .L28
	movl $1, -44(%rbp)
	jmp .L30
.L28:
	movl $0, -44(%rbp)
.L30:
	movl -44(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
