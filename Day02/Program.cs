var part1Score = 0;
var part2Score = 0;
var baseScoreMap = new Dictionary<char, int>
{
    ['A'] = 1,
    ['B'] = 2,
    ['C'] = 3,
};

foreach (var line in File.ReadLines("./input.txt"))
{
    var result = line.Split(" ");
    var theirMove = result[0][0];
    var part1Move = ParseMyMovePart1(result[1][0]);
    var part2Move = ParseMyMovePart2(result[1][0], theirMove);

    part1Score += baseScoreMap[part1Move] + Score(theirMove, part1Move);
    part2Score += baseScoreMap[part2Move] + Score(theirMove, part2Move);
}

Console.WriteLine(part1Score);
Console.WriteLine(part2Score);

static char ParseMyMovePart1(char raw)
{
    return raw switch
    {
        'X' => 'A',
        'Y' => 'B',
        'Z' => 'C'
    };
}

static char ParseMyMovePart2(char raw, char theirMove)
{
    var loseMap = new Dictionary<char, char>
    {
        ['A'] = 'C',
        ['B'] = 'A',
        ['C'] = 'B',
    };
    var winMap = new Dictionary<char, char>
    {
        ['A'] = 'B',
        ['B'] = 'C',
        ['C'] = 'A',
    };

    return raw switch
    {
        'X' => loseMap[theirMove],
        'Y' => theirMove,
        'Z' => winMap[theirMove]
    };
}

static int Score(char theirMove, char myMove)
{
    var winningMove = theirMove switch
    {
        'A' => 'B',
        'B' => 'C',
        'C' => 'A'
    };

    if (myMove == winningMove)
    {
        return 6;
    }
    else if (myMove == theirMove)
    {
        return 3;
    }

    return 0;
}