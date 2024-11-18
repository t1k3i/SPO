main 	START 0

		EXTREF sinit,spop,spush,sp

		+JSUB sinit
        LDA stev
        JSUB sumN
        STA rez
        	
halt	J halt

.rekurzivno sesteje 1+....+n (n je podan v A)
sumN 	+STL @sp
		+JSUB spush
		+STB @sp
		+JSUB spush

		COMP #1
		JEQ sumOut
		RMO A, B 	.shranimo A .. nekam
		SUB #1		.zmanjsamo A
		JSUB sumN
		ADDR B, A

sumOut 	+JSUB spop
		+LDB @sp
		+JSUB spop
		+LDL @sp
		RSUB

stev    WORD 5
rez     WORD 0

	    END main