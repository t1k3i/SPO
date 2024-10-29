arithr   START 0

first   LDT x
        LDS y
        CLEAR A

        ADDR T, A
        ADDR S, A
        STA sum

        CLEAR A
        ADDR T, A
        SUBR S, A
        STA diff

        CLEAR A
        ADDR T, A
        MULR S, A
        STA prod

        CLEAR A
        ADDR T, A
        DIVR S, A
        STA quot

        CLEAR A
        ADDR S, A
        LDX quot
        MULR X, A
        STA temp
        CLEAR A
        LDX temp
        ADDR T, A
        SUBR X, A
        STA mod

halt    J halt

x       WORD 26 .vhodni podatki
y       WORD 5

temp    RESW 1

sum     RESW 1  .izhodni podatki
diff    RESW 1
prod    RESW 1
quot    RESW 1
mod     RESW 1
        END first