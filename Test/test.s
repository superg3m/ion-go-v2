.text
.global _main
_main:
	pushq %rbp
	movq %rsp, %rbp
	subq $16, %rsp
	movl $3, -4(%rbp)
	movl $1, %edi
	movl $2, %esi
	movl $3, %edx
	movl $4, %ecx
	movl $5, %r8d
	movl $6, %r9d
	pushq $8
	pushq $7
	call _even_arguments
	addq $16, %rsp
	movl %eax, -4(%rbp)
	subq $8, %rsp
	movl $1, %edi
	movl $2, %esi
	movl $3, %edx
	movl $4, %ecx
	movl $5, %r8d
	movl $6, %r9d
	pushq $9
	pushq $8
	pushq $7
	call _odd_arguments
	addq $32, %rsp
	movl %eax, -4(%rbp)
	movl -4(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
