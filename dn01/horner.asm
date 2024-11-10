horner    START 0

first   LDA x
        ADD #2  
        MUL x   
        ADD #3
        MUL x
        ADD #4
        MUL x
        ADD #5
        STA result .(((x+2)x+3)x+4)x+5

halt    J halt

x       WORD 2
result  RESW 1
        END first