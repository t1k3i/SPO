prog 	START 0

		JSUB sinit
        LDA stev
        JSUB fak
        STA rez
        	
halt	J halt

.rekurzivno izracuna n! (n je v A)
fak 	STL @sp
		JSUB spush
		STB @sp
		JSUB spush

		COMP #0
		JEQ put1
		RMO A, B 	.shranimo A .. nekam
		SUB #1		.zmanjsamo A
		JSUB fak
		MULR B, A
        J fakOut

put1    LDA #1
fakOut 	JSUB spop
		LDB @sp
		JSUB spop
		LDL @sp
		RSUB

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

stev    WORD 6
rez     WORD 0

sp		WORD 0
stackA	WORD 0
stack 	RESW 1000

	    END prog