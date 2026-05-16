.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	movl $2, %eax
	movq %rbp, %rsp
	popq %rbp
	ret
