.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $12, %rsp
	movl $0, -4(%rbp)
	movl $0, -8(%rbp)
	movl -4(%rbp), %r10d
	cmpl $0, %r10d
	je .L89
	movl $2, -4(%rbp)
	jmp .L90
.L89:
	movl $3, -4(%rbp)
.L90:
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L91
	movl $4, -8(%rbp)
	jmp .L92
.L91:
	movl $5, -8(%rbp)
.L92:
	movl -4(%rbp), %r10d
	movl %r10d, -12(%rbp)
	movl -12(%rbp), %r10d
	addl -8(%rbp), %r10d
	movl %r10d, -12(%rbp)
	movl -12(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
