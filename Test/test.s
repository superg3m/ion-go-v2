.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $5, -4(%rbp)
	movl -4(%rbp), %r10d
	cmpl $2, %r10d
	movl $0, -8(%rbp)
	setl -8(%rbp)
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L1
	movl $9, %eax
	movq %rbp, %rsp
	popq %rbp
	ret
	jmp .L1
.L1:
	movl $1, %eax
	movq %rbp, %rsp
	popq %rbp
	ret
.L2:
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
