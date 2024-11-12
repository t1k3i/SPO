prog 	START 0

		EXTDEF sinit, spush, spop, sp

sinit	.nastavi vrednost sp na zacetek stacka
		STA stackA
		LDA #stack
		STA sp
		LDA stackA
		RSUB

spush 	.poveca vrednost sp za eno besedo
		STA stackA
		LDA sp
		ADD #3
		STA sp
		LDA stackA
		RSUB

spop 	.zmanjsa vrednost sp za eno besedo
		STA stackA
		LDA sp
		SUB #3
		STA sp
		LDA stackA
		RSUB

stack 	RESW 1000
sp		WORD 0
stackA	WORD 0

	    END prog