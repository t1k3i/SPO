prog    START 0

        LDA #66
        JSUB scrfil

        JSUB scrcle

        LDA #65
        JSUB scrfil
        

halt 	J halt

scrcle  LDA #32
scrfil  STX saveX
        LDX #0
loop    +STCH screen, X .to moram vprasat 
        TIX #scrlen        
        JLT loop
        LDX saveX
        RSUB             

saveX   WORD 0 .shrani X za v rutino

sccols  EQU 80
scrows  EQU 25
scrlen  EQU 80 * 25
screen  EQU 0xb800

	    END prog