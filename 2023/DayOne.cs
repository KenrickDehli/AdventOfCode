using System.Text.RegularExpressions;

namespace AoCDayOne;

public static class DayOne
{

    public static int GetCalibrationValue()
    {
        var strList = Inputs.InputDayOne.Split('\n').ToList();
        var numList = new List<int>();
        var digitDatabase = new List<string> { "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };
        foreach (var word in strList)
        {
            if (string.IsNullOrEmpty(word) || string.IsNullOrWhiteSpace(word)) continue;

            var dict = new SortedDictionary<int, int>();

            var indexOfNumbers = word.Select((ch, i) =>
            {
                if (char.IsNumber(ch)) return i;
                return -1;
            }).Where(i => i != -1).ToList();

            foreach (var index in indexOfNumbers)
            {
                var num = int.Parse($"{word[index]}");
                dict.Add(index, num);
            }

            for (var i = 0; i < digitDatabase.Count; i++)
                if (word.Contains(digitDatabase[i]))
                {
                    var list = Regex.Matches(word, digitDatabase[i]).Select(m => m.Index).ToList();
                    foreach (var index in list) dict.Add(index, i + 1);
                }

            var values = dict.Values.ToList();
            numList.Add(int.Parse($"{values[0]}{values[^1]}"));
        }

        return numList.Sum();
    }
}