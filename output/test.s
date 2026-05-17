.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $108, %rsp
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
	movl -16(%rbp), %r10d
	movl %r10d, -20(%rbp)
	negl -20(%rbp)
	movl -20(%rbp), %r10d
	movl %r10d, -24(%rbp)
	negl -24(%rbp)
	movl $4, -28(%rbp)
	addl $9, -28(%rbp)
	movl $7, -32(%rbp)
	movl -32(%rbp), %r10d
	imull -28(%rbp), %r10d
	movl %r10d, -32(%rbp)
	movl $2, -36(%rbp)
	addl $1, -36(%rbp)
	movl $18, %eax
	cdq
	movl -36(%rbp), %r10d
	idiv %r10d
	movl %eax, -40(%rbp)
	movl -40(%rbp), %r10d
	movl %r10d, -44(%rbp)
	negl -44(%rbp)
	movl -32(%rbp), %r10d
	movl %r10d, -48(%rbp)
	movl -48(%rbp), %r10d
	subl -44(%rbp), %r10d
	movl %r10d, -48(%rbp)
	movl -24(%rbp), %r10d
	movl %r10d, -52(%rbp)
	movl -52(%rbp), %r10d
	addl -48(%rbp), %r10d
	movl %r10d, -52(%rbp)
	movl $2, -56(%rbp)
	addl $1, -56(%rbp)
	movl $6, -60(%rbp)
	movl -60(%rbp), %r10d
	subl -56(%rbp), %r10d
	movl %r10d, -60(%rbp)
	movl $5, -64(%rbp)
	movl -64(%rbp), %r10d
	imull -60(%rbp), %r10d
	movl %r10d, -64(%rbp)
	movl $14, %eax
	cdq
	movl $5, %r10d
	idiv %r10d
	movl %edx, -68(%rbp)
	movl -68(%rbp), %r10d
	movl %r10d, -72(%rbp)
	movl -64(%rbp), %r10d
	movl %r10d, -76(%rbp)
	movl -76(%rbp), %r10d
	addl -72(%rbp), %r10d
	movl %r10d, -76(%rbp)
	movl -52(%rbp), %r10d
	movl %r10d, -80(%rbp)
	movl -80(%rbp), %r10d
	subl -76(%rbp), %r10d
	movl %r10d, -80(%rbp)
	movl $9, -84(%rbp)
	subl $3, -84(%rbp)
	movl $8, %eax
	cdq
	movl $4, %r10d
	idiv %r10d
	movl %eax, -88(%rbp)
	movl $2, -92(%rbp)
	movl -92(%rbp), %r10d
	addl -88(%rbp), %r10d
	movl %r10d, -92(%rbp)
	movl -84(%rbp), %r10d
	movl %r10d, -96(%rbp)
	movl -96(%rbp), %r10d
	imull -92(%rbp), %r10d
	movl %r10d, -96(%rbp)
	movl -80(%rbp), %r10d
	movl %r10d, -100(%rbp)
	movl -100(%rbp), %r10d
	addl -96(%rbp), %r10d
	movl %r10d, -100(%rbp)
	movl $3, -104(%rbp)
	notl -104(%rbp)
	movl -100(%rbp), %r10d
	movl %r10d, -108(%rbp)
	movl -108(%rbp), %r10d
	subl -104(%rbp), %r10d
	movl %r10d, -108(%rbp)
	movl -108(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
