.text
.global main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $11, -4(%rbp)
	movl $12, -8(%rbp)
	movl $0, %r10d
	cmpl $0, %r10d
	jne .L26
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L25
.L26:
	movl $1, -12(%rbp)
	jmp .L27
.L25:
	movl $0, -12(%rbp)
.L27:
	movl -4(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	andl -12(%rbp), %r10d
	movl %r10d, -4(%rbp)
	movl -4(%rbp), %r10d
	cmpl $0, %r10d
	jne .L29
	movl $1, %r10d
	cmpl $0, %r10d
	je .L28
.L29:
	movl $1, -16(%rbp)
	jmp .L30
.L28:
	movl $0, -16(%rbp)
.L30:
	movl -8(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl -8(%rbp), %r10d
	xorl -16(%rbp), %r10d
	movl %r10d, -8(%rbp)
	movl $14, -20(%rbp)
	movl -4(%rbp), %r10d
	cmpl $0, %r10d
	jne .L32
	movl -8(%rbp), %r10d
	cmpl $0, %r10d
	je .L31
.L32:
	movl $1, -24(%rbp)
	jmp .L33
.L31:
	movl $0, -24(%rbp)
.L33:
	movl -20(%rbp), %r10d
	movl %r10d, -20(%rbp)
	movl -20(%rbp), %r10d
	orl -24(%rbp), %r10d
	movl %r10d, -20(%rbp)
	movl $16, -28(%rbp)
	movl -20(%rbp), %r10d
	cmpl $0, %r10d
	jne .L35
	movl -28(%rbp), %r10d
	cmpl $0, %r10d
	je .L34
.L35:
	movl $1, -32(%rbp)
	jmp .L36
.L34:
	movl $0, -32(%rbp)
.L36:
	movl -28(%rbp), %r10d
	movl %r10d, -28(%rbp)
	movl -28(%rbp), %r10d
	sar -32(%rbp), %r10d
	movl %r10d, -28(%rbp)
	movl $18, -36(%rbp)
	movl -20(%rbp), %r10d
	cmpl $0, %r10d
	jne .L38
	movl -28(%rbp), %r10d
	cmpl $0, %r10d
	je .L37
.L38:
	movl $1, -40(%rbp)
	jmp .L39
.L37:
	movl $0, -40(%rbp)
.L39:
	movl -36(%rbp), %r10d
	movl %r10d, -36(%rbp)
	movl -36(%rbp), %r10d
	sal -40(%rbp), %r10d
	movl %r10d, -36(%rbp)
	movl -4(%rbp), %r10d
	cmpl $1, %r10d
	movl $0, -44(%rbp)
	sete -44(%rbp)
	movl -44(%rbp), %r10d
	cmpl $0, %r10d
	je .L40
	movl -8(%rbp), %r10d
	cmpl $13, %r10d
	movl $0, -48(%rbp)
	sete -48(%rbp)
.L41:
	movl -48(%rbp), %r10d
	cmpl $0, %r10d
	je .L40
	movl $1, -52(%rbp)
	jmp .L42
.L40:
	movl $0, -52(%rbp)
.L42:
	movl -52(%rbp), %r10d
	cmpl $0, %r10d
	je .L43
	movl -20(%rbp), %r10d
	cmpl $15, %r10d
	movl $0, -56(%rbp)
	sete -56(%rbp)
.L44:
	movl -56(%rbp), %r10d
	cmpl $0, %r10d
	je .L43
	movl $1, -60(%rbp)
	jmp .L45
.L43:
	movl $0, -60(%rbp)
.L45:
	movl -60(%rbp), %r10d
	cmpl $0, %r10d
	je .L46
	movl -28(%rbp), %r10d
	cmpl $8, %r10d
	movl $0, -64(%rbp)
	sete -64(%rbp)
.L47:
	movl -64(%rbp), %r10d
	cmpl $0, %r10d
	je .L46
	movl $1, -68(%rbp)
	jmp .L48
.L46:
	movl $0, -68(%rbp)
.L48:
	movl -68(%rbp), %r10d
	cmpl $0, %r10d
	je .L49
	movl -36(%rbp), %r10d
	cmpl $36, %r10d
	movl $0, -72(%rbp)
	sete -72(%rbp)
.L50:
	movl -72(%rbp), %r10d
	cmpl $0, %r10d
	je .L49
	movl $1, -76(%rbp)
	jmp .L51
.L49:
	movl $0, -76(%rbp)
.L51:
	movl -76(%rbp), %eax
	movq %rbp, %rsp
	popq %rbp
	ret
