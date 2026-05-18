.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $176, %rsp
	movl $1, %r10d
	cmpl $1, %r10d
	movl $0, -176(%rbp)
	sete -176(%rbp)
	movl -176(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
