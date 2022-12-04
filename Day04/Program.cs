// part1
var answer1 = File.ReadLines("./input.txt")
    .Select(line =>
    {
        var pairs = line.Split(',');
        var pair1 = pairs[0];
        var pair2 = pairs[1];

        var start1 = int.Parse(pair1.Split('-')[0].ToString());
        var count1 = int.Parse(pair1.Split('-')[1].ToString()) - start1;

        var start2 = int.Parse(pair2.Split('-')[0].ToString());
        var count2 = int.Parse(pair2.Split('-')[1].ToString()) - start2;

        var range1 = Enumerable.Range(start1, count1 + 1);
        var range2 = Enumerable.Range(start2, count2 + 1);

        return range1.Intersect(range2).Count() == range1.Count() || range2.Intersect(range1).Count() == range2.Count();
    })
    .Where(z => z)
    .Count();

Console.WriteLine($"Answer: {answer1}");

// part2
var answer2 = File.ReadLines("./input.txt")
    .Select(line =>
    {
        var pairs = line.Split(',');
        var pair1 = pairs[0];
        var pair2 = pairs[1];

        var start1 = int.Parse(pair1.Split('-')[0].ToString());
        var count1 = int.Parse(pair1.Split('-')[1].ToString()) - start1;

        var start2 = int.Parse(pair2.Split('-')[0].ToString());
        var count2 = int.Parse(pair2.Split('-')[1].ToString()) - start2;

        var range1 = Enumerable.Range(start1, count1 + 1);
        var range2 = Enumerable.Range(start2, count2 + 1);

        return range1.Intersect(range2).Any();
    })
    .Where(z => z)
    .Count();

Console.WriteLine($"Answer: {answer2}");