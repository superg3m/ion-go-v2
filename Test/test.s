.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $5, -4(%rbp)
	movl -4(%rbp), %eax
	cdq
	movl $3, %r10d
	idiv %r10d
	movl %eax, -4(%rbp)
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
