(*$S+*)
program game2048;
uses applestuff;

const
  left = 1;
  right = 2;
  up = 3;
  down = 4;

type
  four = array[1..4] of integer;

var
  tiles: array[1..4, 1..4] of integer;
  score, moves: integer;
  won, over, quit: boolean;
  key: char;

procedure addtile;
var
  row, col, empty, pick: integer;
begin
  empty := 0;
  for row := 1 to 4 do
    for col := 1 to 4 do
      if tiles[row, col] = 0 then
        empty := empty + 1;
  if empty > 0 then
    begin
      pick := random mod empty;
      for row := 1 to 4 do
        for col := 1 to 4 do
          if tiles[row, col] = 0 then
            begin
              if pick = 0 then
                if random mod 10 = 0 then
                  tiles[row, col] := 4
                else
                  tiles[row, col] := 2;
              pick := pick - 1
            end
    end
end;

procedure newgame;
var
  row, col: integer;
begin
  for row := 1 to 4 do
    for col := 1 to 4 do
      tiles[row, col] := 0;
  score := 0;
  moves := 0;
  won := false;
  over := false;
  addtile;
  addtile
end;

procedure drawboard;
var
  row, col: integer;
begin
  for row := 1 to 4 do
    begin
      gotoxy(20, row * 3 + 1);
      write('+------+------+------+------+');
      gotoxy(20, row * 3 + 2);
      for col := 1 to 4 do
        if tiles[row, col] = 0 then
          write('|      ')
        else
          write('| ', tiles[row, col]:4, ' ');
      write('|');
      gotoxy(20, row * 3 + 3);
      write('|      |      |      |      |')
    end;
  gotoxy(20, 16);
  write('+------+------+------+------+');
  gotoxy(20, 18);
  write('SCORE ', score:6, '    MOVES ', moves:4)
end;

function slideline(var squares: four): boolean;
var
  result: four;
  i, count: integer;
  joined, join: boolean;
begin
  for i := 1 to 4 do
    result[i] := 0;
  count := 0;
  joined := false;
  for i := 1 to 4 do
    if squares[i] <> 0 then
      begin
        (* pascal works out both sides of an and: result[0] does not exist *)
        join := false;
        if (count > 0) and not joined then
          join := result[count] = squares[i];
        if join then
          begin
            result[count] := result[count] * 2;
            score := score + result[count];
            if result[count] = 2048 then
              won := true;
            joined := true
          end
        else
          begin
            count := count + 1;
            result[count] := squares[i];
            joined := false
          end
      end;
  slideline := false;
  for i := 1 to 4 do
    begin
      if result[i] <> squares[i] then
        slideline := true;
      squares[i] := result[i]
    end
end;

function slide(way: integer): boolean;
var
  k, i, row, col: integer;
  squares: four;
  moved: boolean;

  procedure square(k, i: integer; var row, col: integer);
  begin
    case way of
      left:  begin row := k; col := i end;
      right: begin row := k; col := 5 - i end;
      up:    begin row := i; col := k end;
      down:  begin row := 5 - i; col := k end
    end
  end;

begin
  moved := false;
  for k := 1 to 4 do
    begin
      for i := 1 to 4 do
        begin
          square(k, i, row, col);
          squares[i] := tiles[row, col]
        end;
      if slideline(squares) then
        moved := true;
      for i := 1 to 4 do
        begin
          square(k, i, row, col);
          tiles[row, col] := squares[i]
        end
    end;
  slide := moved
end;

function canmove: boolean;
var
  row, col: integer;
begin
  canmove := false;
  for row := 1 to 4 do
    for col := 1 to 4 do
      begin
        if tiles[row, col] = 0 then
          canmove := true;
        if col < 4 then
          if tiles[row, col] = tiles[row, col + 1] then
            canmove := true;
        if row < 4 then
          if tiles[row, col] = tiles[row + 1, col] then
            canmove := true
      end
end;

procedure play(way: integer);
begin
  if slide(way) then
    begin
      moves := moves + 1;
      addtile;
      drawboard;
      over := not canmove
    end
end;

begin
  randomize;
  page(output);
  gotoxy(30, 1);
  write('2 0 4 8');
  gotoxy(10, 20);
  write('THE ARROWS OR I, J, K, L SLIDE THE TILES; Q ENDS THE GAME');
  newgame;
  drawboard;
  quit := false;
  repeat
    read(keyboard, key);
    if ord(key) in [8, 10, 11, 21, 73..76, 81, 105..108, 113] then
      case ord(key) of
        8, 74, 106:   play(left);
        21, 76, 108:  play(right);
        11, 73, 105:  play(up);
        10, 75, 107:  play(down);
        81, 113:      quit := true
      end
  until quit or over;
  gotoxy(20, 22);
  if won then
    write('2048! YOU WIN')
  else
    if over then
      write('NO MORE MOVES')
    else
      write('THE END');
  write(' - PRESS A KEY');
  (* the keys typed ahead, while the board was drawn, are thrown away *)
  while keypress do
    read(keyboard, key);
  read(keyboard, key)
end.
