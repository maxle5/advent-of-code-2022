// part 1
var answer1 = File
    .ReadAllText("./input.txt")
    .Split("\r\n\r\n")
    .Select(g => g.Split("\r\n").Where(str => int.TryParse(str, out var _)).Select(str => int.Parse(str)).Sum())
    .Max();

Console.WriteLine(answer1);

// part 2
var answer2 = File
    .ReadAllText("./input.txt")
    .Split("\r\n\r\n")
    .Select(g => g.Split("\r\n").Where(str => int.TryParse(str, out var _)).Select(str => int.Parse(str)).Sum())
    .OrderDescending()
    .Take(3)
    .Sum();

Console.WriteLine(answer2);
