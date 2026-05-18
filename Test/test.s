.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $248, %rsp
	movl $12, -112(%rbp)
	movl -112(%rbp), %r10d
	imull $3, %r10d
	movl %r10d, -112(%rbp)
	movl $20, -116(%rbp)
	addl $5, -116(%rbp)
	movl -112(%rbp), %r10d
	cmpl -116(%rbp), %r10d
	movl $0, -120(%rbp)
	setg -120(%rbp)
	movl -120(%rbp), %r10d
	movl %r10d, -124(%rbp)
	movl -124(%rbp), %r10d
	imull $7, %r10d
	movl %r10d, -124(%rbp)
	movl $8, -128(%rbp)
	subl $3, -128(%rbp)
	movl -128(%rbp), %r10d
	cmpl $6, %r10d
	movl $0, -132(%rbp)
	sete -132(%rbp)
	movl $0, %r10d
	cmpl -132(%rbp), %r10d
	movl $0, -136(%rbp)
	sete -136(%rbp)
	movl -136(%rbp), %r10d
	movl %r10d, -140(%rbp)
	movl -140(%rbp), %r10d
	imull $5, %r10d
	movl %r10d, -140(%rbp)
	movl -124(%rbp), %r10d
	movl %r10d, -144(%rbp)
	movl -144(%rbp), %r10d
	addl -140(%rbp), %r10d
	movl %r10d, -144(%rbp)
	movl $15, %eax
	cdq
	movl $4, %r10d
	idiv %r10d
	movl %edx, -148(%rbp)
	movl -148(%rbp), %r10d
	cmpl $2, %r10d
	movl $0, -152(%rbp)
	setl -152(%rbp)
	movl -152(%rbp), %r10d
	movl %r10d, -156(%rbp)
	movl -156(%rbp), %r10d
	imull $9, %r10d
	movl %r10d, -156(%rbp)
	movl -144(%rbp), %r10d
	movl %r10d, -160(%rbp)
	movl -160(%rbp), %r10d
	subl -156(%rbp), %r10d
	movl %r10d, -160(%rbp)
	movl $10, -164(%rbp)
	negl -164(%rbp)
	movl -164(%rbp), %r10d
	movl %r10d, -168(%rbp)
	negl -168(%rbp)
	movl -168(%rbp), %eax
	cdq
	movl $2, %r10d
	idiv %r10d
	movl %eax, -172(%rbp)
	movl -172(%rbp), %r10d
	cmpl $5, %r10d
	movl $0, -176(%rbp)
	sete -176(%rbp)
	movl -176(%rbp), %r10d
	movl %r10d, -180(%rbp)
	movl -180(%rbp), %r10d
	imull $11, %r10d
	movl %r10d, -180(%rbp)
	movl -160(%rbp), %r10d
	movl %r10d, -184(%rbp)
	movl -184(%rbp), %r10d
	addl -180(%rbp), %r10d
	movl %r10d, -184(%rbp)
	movl $3, %r10d
	cmpl $3, %r10d
	movl $0, -188(%rbp)
	setle -188(%rbp)
	movl $0, %r10d
	cmpl -188(%rbp), %r10d
	movl $0, -192(%rbp)
	sete -192(%rbp)
	movl $0, %r10d
	cmpl -192(%rbp), %r10d
	movl $0, -196(%rbp)
	sete -196(%rbp)
	movl -196(%rbp), %r10d
	movl %r10d, -200(%rbp)
	movl -200(%rbp), %r10d
	imull $13, %r10d
	movl %r10d, -200(%rbp)
	movl -184(%rbp), %r10d
	movl %r10d, -204(%rbp)
	movl -204(%rbp), %r10d
	addl -200(%rbp), %r10d
	movl %r10d, -204(%rbp)
	movl $7, -208(%rbp)
	addl $8, -208(%rbp)
	movl $3, -212(%rbp)
	movl -212(%rbp), %r10d
	imull $5, %r10d
	movl %r10d, -212(%rbp)
	movl -208(%rbp), %r10d
	cmpl -212(%rbp), %r10d
	movl $0, -216(%rbp)
	setne -216(%rbp)
	movl -216(%rbp), %r10d
	movl %r10d, -220(%rbp)
	movl -220(%rbp), %r10d
	imull $4, %r10d
	movl %r10d, -220(%rbp)
	movl -204(%rbp), %r10d
	movl %r10d, -224(%rbp)
	movl -224(%rbp), %r10d
	subl -220(%rbp), %r10d
	movl %r10d, -224(%rbp)
	movl $5, -228(%rbp)
	movl -228(%rbp), %r10d
	imull $4, %r10d
	movl %r10d, -228(%rbp)
	movl $20, -232(%rbp)
	movl -232(%rbp), %r10d
	subl -228(%rbp), %r10d
	movl %r10d, -232(%rbp)
	movl $1, -236(%rbp)
	negl -236(%rbp)
	movl -232(%rbp), %r10d
	cmpl -236(%rbp), %r10d
	movl $0, -240(%rbp)
	setge -240(%rbp)
	movl -240(%rbp), %r10d
	movl %r10d, -244(%rbp)
	movl -244(%rbp), %r10d
	imull $3, %r10d
	movl %r10d, -244(%rbp)
	movl -224(%rbp), %r10d
	movl %r10d, -248(%rbp)
	movl -248(%rbp), %r10d
	addl -244(%rbp), %r10d
	movl %r10d, -248(%rbp)
	movl -248(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
