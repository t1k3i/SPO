poly    START 0

first   LDA x
        MUL x
        MUL x
        MUL x   
        STA result  .x^4

        LDA x
        MUL x
        MUL x
        MUL #2
        ADD result
        STA result  .x^4 + 2*x^3

        LDA x
        MUL x
        MUL #3
        ADD result
        STA result  .x^4 + 2*x^3 + 3*x^2

        LDA x
        MUL #4
        ADD result
        ADD #5
        STA result  .x^4 + 2*x^3 + 3*x^2 + 4*x + 5

halt    J halt

x       WORD 2
result  RESW 1
        END first