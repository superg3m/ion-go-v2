.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $4, %rsp
.L1:
	movl $1, %r10d
	cmpl $0, %r10d
	je .L2
	movl $0, -4(%rbp)
	jmp .L1
.L2:
	movl $0, %eax
	movq %rbp, %rsp
	popq %rbp
	ret
