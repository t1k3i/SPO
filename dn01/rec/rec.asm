main    START 0

        EXTREF sinit,spop,spush,sp,char,nl,string,num

        +JSUB sinit

        CLEAR T
        LDS #10
cle     CLEAR T
        CLEAR A
        CLEAR B
        CLEAR L
        
loop    RD #0xFA
        COMP #0xD
        JEQ print
        SUB #48
        COMP #0
        JEQ prev
addZ    LDT #1
        MULR S, B
        ADDR A, B
        J loop

prev    COMPR T, L
        JEQ halt
        J addZ

print   RD #0xFA
        RMO B, A
        JSUB fak
        +JSUB num
        J cle

halt 	J halt

.rekurzivno izracuna n! (n je v A) (max = 10)
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

	END main