subru   START 0

first   LDX zero
loop    LDA in, X
        JSUB sub
        STA results, X .(((x+2)x+3)x+4)x+5
        LDA three
        ADDR A, X
        LDA #len
        COMPR X, A
        JLT loop

halt    J halt

sub     ADD #2  
        MUL in, X   
        ADD #3
        MUL in, X
        ADD #4
        MUL in, X
        ADD #5
        RSUB

in      WORD 0
        WORD 5
        WORD 42
lastin  EQU *
len     EQU lastin - in

results RESW 3

zero    WORD 0
three   WORD 3

        END first