.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $40, %rsp
	movl $0, %r10d
	cmpl $0, %r10d
	movl $0, -20(%rbp)
	sete -20(%rbp)
	movl -20(%rbp), %r10d
	cmpl $0, %r10d
	je .L7
	movl $2, -24(%rbp)
	addl $1, -24(%rbp)
	movl -24(%rbp), %r10d
	cmpl $1, %r10d
	movl $0, -28(%rbp)
	setg -28(%rbp)
	movl $3, %r10d
	cmpl -28(%rbp), %r10d
	movl $0, -32(%rbp)
	sete -32(%rbp)
.L8:
	movl -32(%rbp), %r10d
	cmpl $0, %r10d
	je .L7
	movl $1, -36(%rbp)
	jmp .L9
.L7:
	movl $0, -36(%rbp)
.L9:
	movl -36(%rbp), %r10d
	movl %r10d, -40(%rbp)
	addl $1, -40(%rbp)
	movl -40(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
