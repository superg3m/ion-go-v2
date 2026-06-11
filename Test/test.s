.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $64, %rsp
	subq $8, %rsp
	pushq %rdi
	movl $72, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -4(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $101, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -8(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $108, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $108, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -16(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $111, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -20(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $44, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -24(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $32, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -28(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $87, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -32(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $111, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -36(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $114, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -40(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $108, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -44(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $100, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -48(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $33, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -52(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $10, %edi
	call putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -56(%rbp)
	movl $0, %eax
	movq %rbp, %rsp
	popq %rbp
	ret
