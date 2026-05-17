.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $112, %rsp
	cmpl $0, $1
	je .L1
.L2:
	movl $1, -112(%rbp)
	jmp .L3
.L1:
	movl $0, -112(%rbp)
.L3:
	movl -112(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
