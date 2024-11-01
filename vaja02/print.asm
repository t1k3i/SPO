print   START 0

	CLEAR X
loop	LDCH txt,X
        JSUB putc
        TIX #len
        JLT loop
        LDCH #10
        JSUB putc
    
halt 	J halt

putc 	WD #170
	RSUB

txt     BYTE C'SIC/XE'
end 	EQU *
len 	EQU end - txt

        END print