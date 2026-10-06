* COLOUR BARS THAT MOVE, IN THE
* LOW RESOLUTION GRAPHICS
         ORG $8000
PTR      EQU $06
SHIFT    EQU $08
COL      EQU $09
START    LDA $C050      ;GRAPHICS
         LDA $C056      ;LOW RESOLUTION
         LDA $C052      ;WHOLE SCREEN
         LDA #0
         STA SHIFT
FRAME    LDX #23        ;EVERY LINE
ROW      LDA LO,X
         STA PTR
         LDA HI,X
         STA PTR+1
         LDY #39        ;EVERY COLUMN
COLUMN   TYA
         LSR            ;TWO COLUMNS A BAR
         CLC
         ADC SHIFT
         AND #$0F
         STA COL
         ASL            ;THE COLOUR IN
         ASL            ;BOTH HALVES
         ASL
         ASL
         ORA COL
         STA (PTR),Y
         DEY
         BPL COLUMN
         DEX
         BPL ROW
         INC SHIFT
         LDA #$60
         JSR $FCA8      ;WAIT
         LDA $C000      ;A KEY?
         BPL FRAME
         STA $C010
         LDA $C051      ;TEXT AGAIN
         JSR $FC58      ;CLEARED
         RTS
* WHERE EACH LINE OF THE SCREEN IS
LO       DFB $00,$80,$00,$80,$00,$80,$00,$80
         DFB $28,$A8,$28,$A8,$28,$A8,$28,$A8
         DFB $50,$D0,$50,$D0,$50,$D0,$50,$D0
HI       DFB $04,$04,$05,$05,$06,$06,$07,$07
         DFB $04,$04,$05,$05,$06,$06,$07,$07
         DFB $04,$04,$05,$05,$06,$06,$07,$07
