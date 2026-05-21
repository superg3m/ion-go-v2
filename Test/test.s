.text
.global _main
_main:
	pushq %rbp
	movq %rsp, %rbp
	subq $32, %rsp
	movl $1, -4(%rbp)
	movl $5, -8(%rbp)
	subq $8, %rsp
	movl $3, %edi
	movl $10, %esi
	movl $1, %edx
	movl $2, %ecx
	movl $3, %r8d
	movl $4, %r9d
	pushq $6
	pushq $5
	call _get_integer
	addq $24, %rsp
	movl %eax, -8(%rbp)
	movl -8(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -12(%rbp)
	movl -4(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	addl $1, %r10d
	movl %r10d, -4(%rbp)
	movl -12(%rbp), %r10d
	movl %r10d, -8(%rbp)
	subq $8, %rsp
	movl $72, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $101, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $108, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $108, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $111, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $44, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $32, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $87, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $111, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $114, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $108, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $100, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $33, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	subq $8, %rsp
	movl $10, %edi
	call _putchar
	addq $8, %rsp
	movl %eax, -12(%rbp)
	movl -4(%rbp), %r10d
	cmpl $15, %r10d
	movl $0, -16(%rbp)
	sete -16(%rbp)
	movl -16(%rbp), %r10d
	cmpl $0, %r10d
	je .L1
	movl -8(%rbp), %r10d
	cmpl $14, %r10d
	movl $0, -20(%rbp)
	sete -20(%rbp)
.L2:
	movl -20(%rbp), %r10d
	cmpl $0, %r10d
	je .L1
	movl $1, -24(%rbp)
	jmp .L3
.L1:
	movl $0, -24(%rbp)
.L3:
	movl -24(%rbp), %eax
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
