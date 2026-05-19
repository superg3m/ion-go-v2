.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $5, -4(%rbp)
	movl $4, -8(%rbp)
	movl -8(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
