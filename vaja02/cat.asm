cat     START 0

        CLEAR X
        LDB #10
        LDS #1

loop    RD #0
        COMPR A, B
        JEQ print
        STCH buf, X
        ADDR S, X
        J loop

print   STX len
        LDT len
        CLEAR X
        
pri     LDCH buf, X
        WD #1
        ADDR S, X
        COMPR X, T
        JLT pri
        
halt 	J halt

len     WORD 0
buf     RESB 1000

        END cat