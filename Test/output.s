.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $84, %rsp
	movl $8, -4(%rbp)
	subl $2, -4(%rbp)
	movl $3, -8(%rbp)
	movl -8(%rbp), %r10d
	imull -4(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl $15, -12(%rbp)
	movl -12(%rbp), %r10d
	addl -8(%rbp), %r10d
	movl %r10d, -12(%rbp)
	movl -12(%rbp), %eax
	cdq
	movl $3, %r10d
	idiv %r10d
	movl %eax, -16(%rbp)
	movl $4, -20(%rbp)
	addl $9, -20(%rbp)
	movl $7, -24(%rbp)
	movl -24(%rbp), %r10d
	imull -20(%rbp), %r10d
	movl %r10d, -24(%rbp)
	movl $2, -28(%rbp)
	addl $1, -28(%rbp)
	movl $18, %eax
	cdq
	movl -28(%rbp), %r10d
	idiv %r10d
	movl %eax, -32(%rbp)
	movl -24(%rbp), %r10d
	movl %r10d, -36(%rbp)
	movl -36(%rbp), %r10d
	subl -32(%rbp), %r10d
	movl %r10d, -36(%rbp)
	movl -16(%rbp), %r10d
	movl %r10d, -40(%rbp)
	movl -40(%rbp), %r10d
	addl -36(%rbp), %r10d
	movl %r10d, -40(%rbp)
	movl $2, -44(%rbp)
	addl $1, -44(%rbp)
	movl $6, -48(%rbp)
	movl -48(%rbp), %r10d
	subl -44(%rbp), %r10d
	movl %r10d, -48(%rbp)
	movl $5, -52(%rbp)
	movl -52(%rbp), %r10d
	imull -48(%rbp), %r10d
	movl %r10d, -52(%rbp)
	movl $14, %eax
	cdq
	movl $5, %r10d
	idiv %r10d
	movl %edx, -56(%rbp)
	movl -52(%rbp), %r10d
	movl %r10d, -60(%rbp)
	movl -60(%rbp), %r10d
	addl -56(%rbp), %r10d
	movl %r10d, -60(%rbp)
	movl -40(%rbp), %r10d
	movl %r10d, -64(%rbp)
	movl -64(%rbp), %r10d
	subl -60(%rbp), %r10d
	movl %r10d, -64(%rbp)
	movl $9, -68(%rbp)
	subl $3, -68(%rbp)
	movl $8, %eax
	cdq
	movl $4, %r10d
	idiv %r10d
	movl %eax, -72(%rbp)
	movl $2, -76(%rbp)
	movl -76(%rbp), %r10d
	addl -72(%rbp), %r10d
	movl %r10d, -76(%rbp)
	movl -68(%rbp), %r10d
	movl %r10d, -80(%rbp)
	movl -80(%rbp), %r10d
	imull -76(%rbp), %r10d
	movl %r10d, -80(%rbp)
	movl -64(%rbp), %r10d
	movl %r10d, -84(%rbp)
	movl -84(%rbp), %r10d
	addl -80(%rbp), %r10d
	movl %r10d, -84(%rbp)
	movl -84(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
