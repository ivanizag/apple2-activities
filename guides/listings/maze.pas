program Maze;
{$A-}
{ A maze dug at random, and the way out found, by recursion }
const
  W = 39;
  H = 11;
var
  Open, Seen: array[0..80, 0..24] of Boolean;
  DX, DY: array[0..3] of Integer;
  X, Y, Steps: Integer;
  Found: Boolean;
  Ch: Char;

procedure Put(X, Y: Integer; C: Char);
begin
  GotoXY(X, Y);
  Write(C)
end;

procedure Dig(X, Y: Integer);
var
  Order: array[0..3] of Integer;
  I, J, T, D: Integer;
begin
  Open[X, Y] := True;
  Put(X, Y, ' ');
  for I := 0 to 3 do
    Order[I] := I;
  for I := 3 downto 1 do
  begin
    J := Random(I + 1);
    T := Order[I];
    Order[I] := Order[J];
    Order[J] := T
  end;
  for I := 0 to 3 do
  begin
    D := Order[I];
    if (X + 2 * DX[D] >= 2) and (X + 2 * DX[D] <= 2 * W) and
       (Y + 2 * DY[D] >= 2) and (Y + 2 * DY[D] <= 2 * H) then
      if not Open[X + 2 * DX[D], Y + 2 * DY[D]] then
      begin
        Open[X + DX[D], Y + DY[D]] := True;
        Put(X + DX[D], Y + DY[D], ' ');
        Dig(X + 2 * DX[D], Y + 2 * DY[D])
      end
  end
end;

procedure Solve(X, Y: Integer);
var
  D: Integer;
begin
  Seen[X, Y] := True;
  Put(X, Y, '.');
  Steps := Steps + 1;
  if (X = 2 * W + 1) and (Y = 2 * H) then
    Found := True;
  D := 0;
  while not Found and (D <= 3) do
  begin
    if Open[X + DX[D], Y + DY[D]] and not Seen[X + DX[D], Y + DY[D]] then
      Solve(X + DX[D], Y + DY[D]);
    D := D + 1
  end;
  if not Found then
    Put(X, Y, ' ')
end;

begin
  DX[0] := 1;  DY[0] := 0;
  DX[1] := 0;  DY[1] := 1;
  DX[2] := -1; DY[2] := 0;
  DX[3] := 0;  DY[3] := -1;
  for X := 0 to 80 do
    for Y := 0 to 24 do
    begin
      Open[X, Y] := False;
      Seen[X, Y] := False
    end;
  ClrScr;
  for Y := 1 to 2 * H + 1 do
  begin
    GotoXY(1, Y);
    for X := 1 to 2 * W + 1 do
      Write('#')
  end;
  Randomize;
  Dig(2, 2);
  Open[1, 2] := True;
  Put(1, 2, ' ');
  Open[2 * W + 1, 2 * H] := True;
  Put(2 * W + 1, 2 * H, ' ');
  Found := False;
  Steps := 0;
  Solve(1, 2);
  GotoXY(1, 24);
  Write('Out in ', Steps, ' steps. Press a key.');
  Read(Kbd, Ch)
end.
