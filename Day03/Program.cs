// part 1
var answer1 = File
    .ReadLines("./input.txt") // read file
    .Select(line => new string[] { line[..(line.Length / 2)], line[(line.Length / 2)..] }) // split into halfs
    .Select(sack => sack[0].Intersect(sack[1])) // find common letter
    .Where(item => item.Any()) // filter out non-matches
    .Select(item => item.First()) // select the first match
    .Select(symbol => symbol >= 97 && symbol <= 122 ? symbol - 96 : symbol - 38) // convert letter into priority score
    .Sum(); // calulate sum

Console.WriteLine(answer1);

// part 2
var answer2 = File
    .ReadLines("./input.txt") // read file
    .Chunk(3) // chunk into groups of 3
    .Select(group => group[0].Intersect(group[1]).Intersect(group[2])) // find common letter
    .Where(item => item.Any()) // filter out non-matches
    .Select(item => item.First()) // select the first match
    .Select(symbol => symbol >= 97 && symbol <= 122 ? symbol - 96 : symbol - 38) // convert letter into priority score
    .Sum(); // calulate sum

Console.WriteLine(answer2);
