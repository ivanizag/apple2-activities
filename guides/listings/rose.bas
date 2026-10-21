10 REM Four mystic roses: points on a circle, each joined to every other
20 MODE 2
30 PROCrose(320,768,5,1)
40 PROCrose(960,768,7,2)
50 PROCrose(320,256,9,4)
60 PROCrose(960,256,11,5)
70 END
80 DEF PROCrose(X%,Y%,N%,C%)
90 LOCAL I%,J%
100 GCOL 0,C%
110 FOR I%=0 TO N%-2
120   FOR J%=I%+1 TO N%-1
130     MOVE FNx(I%),FNy(I%)
140     DRAW FNx(J%),FNy(J%)
150   NEXT
160 NEXT
170 ENDPROC
180 DEF FNx(K%)=X%+260*SIN(2*PI*K%/N%)
190 DEF FNy(K%)=Y%+230*COS(2*PI*K%/N%)
