.text
.global _fib
_fib:
	pushq %rbp
	movq %rsp, %rbp
	subq $32, %rsp
	movl %edi, %r10d
	cmpl $0, %r10d
	movl $0, -4(%rbp)
	sete -4(%rbp)
	movl -4(%rbp), %r10d
	cmpl $0, %r10d
	jne .L4
	movl %edi, %r10d
	cmpl $1, %r10d
	movl $0, -8(%rbp)
	sete -8(%rbp)
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L3
.L4:
	movl $1, -12(%rbp)
	jmp .L5
.L3:
	movl $0, -12(%rbp)
.L5:
	movl -12(%rbp), %r10d
	cmpl $0, %r10d
	je .L1
	movl %edi, %eax
	movq %rbp, %rsp
	popq %rbp
	ret
	jmp .L2
.L1:
	movl %edi, -16(%rbp)
	movl -16(%rbp), %r10d
	subl $1, %r10d
	movl %r10d, -16(%rbp)
	pushq %rdi
	movl -16(%rbp), %edi
	call _fib
	movl %eax, -16(%rbp)
	popq %rdi
	movl %edi, -20(%rbp)
	movl -20(%rbp), %r10d
	subl $2, %r10d
	movl %r10d, -20(%rbp)
	pushq %rdi
	movl -20(%rbp), %edi
	call _fib
	movl %eax, -20(%rbp)
	popq %rdi
	movl -16(%rbp), %r10d
	movl %r10d, -20(%rbp)
	movl -20(%rbp), %r10d
	addl -20(%rbp), %r10d
	movl %r10d, -20(%rbp)
	movl -20(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
.L2:
.global _main
_main:
	pushq %rbp
	movq %rsp, %rbp
	subq $16, %rsp
	movl $6, -4(%rbp)
	pushq %rdi
	movl -4(%rbp), %edi
	call _fib
	movl %eax, -4(%rbp)
	popq %rdi
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
