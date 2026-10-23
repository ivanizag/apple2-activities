\ The alphabet, a letter at a time, by the BDOS
start:
    lda #'A'
    sta letter
next:
    lda letter
    ldy #2
    jsr start - 3
    lda letter
    inc letter
    cmp #'Z'
    bne next
    lda #13
    ldy #2
    jsr start - 3
    lda #10
    ldy #2
    jsr start - 3
    rts
letter:
    .byte 0
