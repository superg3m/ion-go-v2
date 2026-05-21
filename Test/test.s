.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $-16, %rsp
	movl $1, -4(%rbp)
	movl $5, -8(%rbp)
	movl $3, %edi
	movl $10, %esi
	movl $1, %edx
	movl $2, %ecx
	movl $3, %r8d
	movl $4, %r9d
	pushq $6
	pushq $5
	call get_integer
	addq $16, %rsp
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
.global get_integer
get_integer:
	pushq %rbp
	movq %rsp, %rbp
	subq $0, %rsp
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
