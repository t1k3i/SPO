. set frequency to 1000000 Hz

prog    START 0

        . sets the starting x and y
        LDA #stax
        STA x
        LDA #stay
        STA y
        CLEAR A

        . reading input to buffer
        LDB #10
        LDS #1
loop    RD #0
        STCH buff, X
        ADDR S, X
        COMPR A, B
        JEQ main
        J loop

        . displaing time
main    JSUB print
        JSUB wait
        JSUB scrcle
        JSUB incs
        J main

halt    J halt

. prints digit given in A starting at x and y
drawdi  STX svX
        STB svB
        STS svS
        CLEAR X
        CLEAR B
        LDS #firstp
        ADDR S, B
        STA dig
        LDA x
        STA svx
        LDA y
        STA svy
        CLEAR A

full    LDA svy
        MUL #cols   
        +ADD #screen
        ADD svx
        ADDR B, A
        STA save
        CLEAR A

        LDA dig 
        ADDR X, A   
        +STA temp
        CLEAR A
        +LDCH @temp
        STA rowdig
row     LDA rowdig
        COMP #0
        JEQ endrow
        AND #1
        COMP #1
        JLT putp
        ADD #251
putp    STCH @save
        LDA save
        SUB #1
        STA save
        LDA rowdig
        SHIFTR A, 1
        STA rowdig
        J row

endrow  TIX #height
        LDA svy
        ADD #1
        STA svy
        JLT full

end1    LDX svX
        LDB svB
        LDS svS
        RSUB

. saved registers for rutin drawdi
svX     WORD 0
svB     WORD 0
svS     WORD 0
svx     WORD 0
svy     WORD 0
dig     WORD 0
temp   	WORD 0
save    WORD 0
rowdig  WORD 0

. clear screen
scrcle  STX svX3
        +LDA #cols
        MUL y
        ADD x
        RMO A, X

        LDA y
        ADD #height
        SUB #1
        MUL #cols
        STA temp3
        LDA #width
        MUL #8
        ADD x
        ADD #7
        ADD temp3
        STA temp3

        LDA #0
loop3   +STCH screen, X
        +TIX temp3        
        JLT loop3
        LDX svX3
        RSUB

. saved registers for rutin srccle
svX3    WORD 0
temp3   WORD 0
zero    WORD 0

. print buff
print   STL svL4
        STX svX4
        LDX #1
        LDA #buff
        SUB #1
        STA sv

loop4   LDA sv
        ADDR X, A
        STA sv
        CLEAR A
        LDCH @sv
        COMP #0x0d
        JEQ end4
        SUB #48
        MUL #height
        +ADD #dig0
        JSUB drawdi

        LDA x
        ADD #width
        ADD #1
        STA x
        J loop4

end4    LDA #stax
        STA x
        LDA #stay
        STA y
        LDL svL4
        LDX svX4
        RSUB

. saved registers for rutin print
svL4    WORD 0
svX4    WORD 0
sv      WORD 0

. incerement seconds in buff
incs    STL svL5
        STX svX5

        LDX #7
        CLEAR A
        LDCH buff, X
        ADD #1
        COMP #58
        JEQ incs2
        STCH buff, X
        J end5

incs2   LDA #48
        STCH buff, X
        LDX #6
        CLEAR A
        LDCH buff, X
        ADD #1
        COMP #54
        JEQ jum1
        STCH buff, X
        J end5

jum1    LDA #48
        STCH buff, X
        JSUB incm

end5    LDL svL5
        LDX svX5
        RSUB

. saved registers for rutin incs
svL5    WORD 0
svX5    WORD 0

. incerement minutes in buff
incm    STL svL6
        STX svX6

        LDX #4
        CLEAR A
        LDCH buff, X
        ADD #1
        COMP #58
        JEQ incm2
        STCH buff, X
        J end6

incm2   LDA #48
        STCH buff, X
        LDX #3
        CLEAR A
        LDCH buff, X
        ADD #1
        COMP #54
        JEQ jum2
        STCH buff, X
        J end6

jum2    LDA #48
        STCH buff, X
        JSUB inch

end6    LDL svL6
        LDX svX6
        RSUB

. saved registers for rutin incm
svL6    WORD 0
svX6    WORD 0

. incerement hours in buff
inch    STL svL7
        STX svX7

        LDX #0
        CLEAR A
        LDCH buff, X
        COMP #50
        JEQ inch3

        LDX #1
        CLEAR A
        LDCH buff, X
        ADD #1
        COMP #58
        JEQ inch2
        STCH buff, X
        J end7

inch2   LDA #48
        STCH buff, X
        LDX #0
        CLEAR A
        LDCH buff, X
        ADD #1
        STCH buff, X
        J end7

inch3   LDX #1
        CLEAR A
        LDCH buff, X
        ADD #1
        COMP #52
        JEQ jum3
        STCH buff, X
        J end7

jum3    LDA #48
        STCH buff, X
        LDX #0
        STCH buff, X

end7    LDL svL7
        LDX svX7
        RSUB

. saved registers for rutin inch
svL7    WORD 0
svX7    WORD 0

. wait for some time
wait    STA svA8
        LDA #0
loop8   ADD #1
        +COMP #waitt
        JLT loop8
        LDA svA8
        RSUB

. saved registers for rutin wait
svA8    WORD 0

. x and y for each digit
x       WORD 0
y       WORD 0

. save time
buff    RESB 12

. wait time
waitt   EQU 350000

. starting x and y
stax    EQU 9
stay    EQU 15

. screen settings
cols    EQU 64
rows    EQU 64
scrlen  EQU 64 * 64
screen  EQU 0x0a000

. digits size
height  EQU 5
width   EQU 5
firstp  EQU width - 1

. digits represented in binary
dig0    BYTE X'1F1111111F'
dig1    BYTE X'040C04041F'
dig2    BYTE X'1F011F101F'
dig3    BYTE X'1F011F011F'
dig4    BYTE X'11111F0101'
dig5    BYTE X'1F101F011F'
dig6    BYTE X'1F101F111F'
dig7    BYTE X'1F01010101'
dig8    BYTE X'1F111F111F'
dig9    BYTE X'1F111F0101'
colon   BYTE X'0004000400'

        END prog