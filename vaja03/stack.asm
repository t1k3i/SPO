prog 	START 0

		JSUB sinit

		LDA #1
		LDB #2

		STA @sp		.damo 2 vrednosti na sklad
		JSUB spush
		STB @sp
		JSUB spush

		JSUB spop	.vzamemo 2 vrednosti iz skalda
        LDB @sp
		JSUB spop
        LDA @sp
		
halt	J halt

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