arith   START 0

first   LDA x
        ADD y
        STA sum

        LDA x
        SUB y
        STA diff

        LDA x
        MUL y
        STA prod

        LDA x
        DIV y
        STA quot

        LDA y
        MUL quot
        STA temp
        LDA x
        SUB temp
        STA mod

halt    J halt

x       WORD 24 .vhodni podatki
y       WORD 5

temp    RESW 1

sum     RESW 1  .izhodni podatki
diff    RESW 1
prod    RESW 1
quot    RESW 1
mod     RESW 1
        END first