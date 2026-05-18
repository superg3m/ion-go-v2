.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $10, -4(%rbp)
	movl $12, -8(%rbp)
	movl $0, %r10d
	cmpl $0, %r10d
	jne .L5
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L4
.L5:
	movl $1, -12(%rbp)
	jmp .L6
.L4:
	movl $0, -12(%rbp)
.L6:
	movl -4(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	addl -12(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	cmpl $0, %r10d
	je .L7
.L8:
	movl $0, %r10d
	cmpl $0, %r10d
	je .L7
	movl $1, -16(%rbp)
	jmp .L9
.L7:
	movl $0, -16(%rbp)
.L9:
	movl -8(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl -8(%rbp), %r10d
	imull -16(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl $14, -20(%rbp)
	movl -4(%rbp), %r10d
	cmpl $0, %r10d
	jne .L11
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L10
.L11:
	movl $1, -24(%rbp)
	jmp .L12
.L10:
	movl $0, -24(%rbp)
.L12:
	movl -20(%rbp), %r10d
	movl %r10d, -20(%rbp)
	movl -20(%rbp), %r10d
	subl -24(%rbp), %r10d
	movl %r10d, -20(%rbp)
	movl $16, -28(%rbp)
	movl -20(%rbp), %r10d
	cmpl $0, %r10d
	jne .L14
	movl -28(%rbp), %r10d
	cmpl $0, %r10d
	je .L13
.L14:
	movl $1, -32(%rbp)
	jmp .L15
.L13:
	movl $0, -32(%rbp)
.L15:
	movl -28(%rbp), %eax
	cdq
	movl -32(%rbp), %r10d
	idiv %r10d
	movl %eax, -28(%rbp)
	movl -4(%rbp), %r10d
	cmpl $11, %r10d
	movl $0, -36(%rbp)
	sete -36(%rbp)
	movl -36(%rbp), %r10d
	cmpl $0, %r10d
	je .L16
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	movl $0, -40(%rbp)
	sete -40(%rbp)
.L17:
	movl -40(%rbp), %r10d
	cmpl $0, %r10d
	je .L16
	movl $1, -44(%rbp)
	jmp .L18
.L16:
	movl $0, -44(%rbp)
.L18:
	movl -44(%rbp), %r10d
	cmpl $0, %r10d
	je .L19
	movl -20(%rbp), %r10d
	cmpl $13, %r10d
	movl $0, -48(%rbp)
	sete -48(%rbp)
.L20:
	movl -48(%rbp), %r10d
	cmpl $0, %r10d
	je .L19
	movl $1, -52(%rbp)
	jmp .L21
.L19:
	movl $0, -52(%rbp)
.L21:
	movl -52(%rbp), %r10d
	cmpl $0, %r10d
	je .L22
	movl -28(%rbp), %r10d
	cmpl $16, %r10d
	movl $0, -56(%rbp)
	sete -56(%rbp)
.L23:
	movl -56(%rbp), %r10d
	cmpl $0, %r10d
	je .L22
	movl $1, -60(%rbp)
	jmp .L24
.L22:
	movl $0, -60(%rbp)
.L24:
	movl -60(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
