echo    START 0
    
        LDA #txt
        JSUB string

        LDA primer
        JSUB num

halt 	J halt

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

        END echo