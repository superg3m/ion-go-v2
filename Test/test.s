.text
.global _main
_main:
	pushq %rbp
	movq %rsp, %rbp
	subq $96, %rsp
	movl $1, -4(%rbp)
	movl $5, -8(%rbp)
	pushq %rdi
	movl $3, %edi
	pushq %rsi
	movl $10, %esi
	pushq %rdx
	movl $1, %edx
	pushq %rcx
	movl $2, %ecx
	pushq %r8
	movl $3, %r8d
	pushq %r9
	movl $4, %r9d
	pushq $6
	pushq $5
	call _get_integer
	popq %r9
	popq %r8
	popq %rcx
	popq %rdx
	popq %rsi
	popq %rdi
	addq $16, %rsp
	movl %eax, -12(%rbp)
	movl -12(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -16(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	addl $1, %r10d
	movl %r10d, -4(%rbp)
	movl -16(%rbp), %r10d
	movl %r10d, -8(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $72, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -20(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $101, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -24(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $108, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -28(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $108, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -32(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $111, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -36(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $44, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -40(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $32, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -44(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $87, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -48(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $111, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -52(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $114, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -56(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $108, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -60(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $100, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -64(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $33, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -68(%rbp)
	subq $8, %rsp
	pushq %rdi
	movl $10, %edi
	call _putchar
	popq %rdi
	addq $8, %rsp
	movl %eax, -72(%rbp)
	movl -4(%rbp), %r10d
	cmpl $15, %r10d
	movl $0, -76(%rbp)
	sete -76(%rbp)
	movl -76(%rbp), %r10d
	cmpl $0, %r10d
	je .L245
	movl -8(%rbp), %r10d
	cmpl $14, %r10d
	movl $0, -80(%rbp)
	sete -80(%rbp)
.L246:
	movl -80(%rbp), %r10d
	cmpl $0, %r10d
	je .L245
	movl $1, -84(%rbp)
	jmp .L247
.L245:
	movl $0, -84(%rbp)
.L247:
	movl -84(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
.global _get_integer
_get_integer:
	pushq %rbp
	movq %rsp, %rbp
	subq $16, %rsp
	movl %edi, -4(%rbp)
	movl -4(%rbp), %r10d
	addl 16(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl -8(%rbp), %r10d
	addl 24(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl -8(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
