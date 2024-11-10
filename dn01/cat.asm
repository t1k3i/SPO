cat     START 0

        CLEAR X
        LDB #10
        LDS #1
        
sta     CLEAR X
loop    RD #0
        STCH buf, X
        ADDR S, X
        COMPR A, B
        JEQ print
        J loop

print   STX len
        LDT len
        CLEAR X
        
pri     LDCH buf, X
        WD #1
        ADDR S, X
        COMPR X, T
        JLT pri
        J sta
        
halt 	J halt

len     WORD 0
buf     RESB 1000

        END cat