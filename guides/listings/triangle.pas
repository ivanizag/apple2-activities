program triangle(output);
{ Pascal's triangle, each row made from the one above }
var row: array[0..10] of integer;
    n, i: integer;
begin
  for n := 0 to 10 do
  begin
    row[n] := 1;
    for i := n - 1 downto 1 do
      row[i] := row[i] + row[i - 1];
    write(' ': 2 * (10 - n) + 1);
    for i := 0 to n do
      write(row[i]: 4);
    writeln
  end
end.
