prog    START 0
        
        CLEAR A

loo     RD #0xFA
        COMP #10
        JEQ loop
        COMP #48
        JEQ halt
        JSUB fak
        +JSUB num
        +JSUB nl
        J loo  

halt 	J halt

.rekurzivno izracuna n! (n je v A)
fak 	+STL @sp
		+JSUB spush
		+STB @sp
		+JSUB spush

		COMP #0
		JEQ put1
		RMO A, B 	.shranimo A .. nekam
		SUB #1		.zmanjsamo A
		JSUB fak
		MULR B, A
        J fakOut

put1    LDA #1
fakOut 	+JSUB spop
		+LDB @sp
		+JSUB spop
		+LDL @sp
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

char 	WD #1
	RSUB

nl	STA nlA
        STL nlL
	LDCH #10
	JSUB char
        LDL nlL
	LDA nlA
	RSUB

.izpise stringe do NULL znaka
string  STL staL
        STA sta
        STS staS
        LDS one

        LDCH @sta
loop    JSUB char
        LDA sta
        ADDR S, A
        STA sta
        CLEAR A
        LDCH @sta
        COMP zero
        JGT loop
        JSUB nl
        
endstr  LDL staL
        LDA sta
        LDS staS
        RSUB
.konec string rutine

.izpise na zaslon desetisko vrednost stevila podanega v A
num     STA numA
        STA numA2
        STL numL
        STS numS
        STB numB
        LDS #10
        CLEAR X

neg     LDB zero
        COMPR A, B
        JLT min
        J loop2
min     CLEAR A
        LDA #45
        JSUB char
        LDA mask
        SUB numA
        ADD one
        CLEAR B
        STA numA
        STA numA2

loop2   DIVR S, A
        RMO A, B
        MULR S, A
        STA temp
        LDA numA2
        SUB temp
        ADD #48
        STCH numSta, X
        LDA one
        ADDR A, X
        STB numA2
        RMO B, A
        COMP zero
        JEQ print
        J loop2

print   CLEAR A
        LDA one
        SUBR A, X
        LDCH numSta, X
        JSUB char
        LDB zero
        COMPR B, X
        JEQ endnum
        J print

endnum  JSUB nl
        LDA numA
        LDL numL
        LDS numS
        LDB numB
        RSUB
.konec num rutine

primer  WORD 000920790

numL    WORD 0
numA    WORD 0
numA2   WORD 0
numS    WORD 0
numB    WORD 0
temp    WORD 0
numSta  RESB 10

nlA 	WORD 0
nlL     WORD 0

staL    WORD 0
staS    WORD 0
staT    WORD 0
sta     WORD 0

txt     BYTE C'hello'
        BYTE 0

one     WORD 1
zero    WORD 0
mask    WORD -1

stack 	RESW 1000
sp		WORD 0
stackA	WORD 0

	    END prog