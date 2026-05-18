.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $0, %r10d
	cmpl $0, %r10d
	movl $0, -4(%rbp)
	sete -4(%rbp)
	movl -4(%rbp), %r10d
	cmpl $0, %r10d
	je .L52
	movl $2, -8(%rbp)
	addl $1, -8(%rbp)
	movl -8(%rbp), %r10d
	cmpl $1, %r10d
	movl $0, -12(%rbp)
	setg -12(%rbp)
	movl $3, %r10d
	cmpl -12(%rbp), %r10d
	movl $0, -16(%rbp)
	sete -16(%rbp)
.L53:
	movl -16(%rbp), %r10d
	cmpl $0, %r10d
	je .L52
	movl $1, -20(%rbp)
	jmp .L54
.L52:
	movl $0, -20(%rbp)
.L54:
	movl -20(%rbp), %r10d
	movl %r10d, -24(%rbp)
	addl $1, -24(%rbp)
	movl -24(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
