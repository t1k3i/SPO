. This program shows different kinds of addresing
prg 	START 0

        . primer z EQU
	    .LDCH #97
	    .+STCH screen
	    .LDX #1
	    .LDCH #98
	    .+STCH screen, X
	    .LDX #2
	    .LDCH #99
	    .+STCH screen, X

        . primer z ORG
        +LDB #screen
        BASE screen
        LDCH #97
        STCH screen
        LDX #1
        LDCH #98
        STCH screen, X
        LDX #2
        LDCH #99
        STCH screen, X
        NOBASE

halt 	J halt

.screen  EQU 0xb800

	    ORG 0xb800
screen	RESB 2000

.screen WORD 0xb800 -> to bi z afno

	    END prg