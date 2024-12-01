using System.Text.RegularExpressions;

namespace AoCDayOne;

public static class DayFour
{
    public static int GetPointsOfCards()
    {
        var sum = 0;

        var input = Inputs.InputDayFour.Split("\n").ToList();
        foreach (var card in input)
        {
            var split = card.Split(":");
            var winningCards = Regex.Matches(split[1].Split("|")[0], @"\d+").Select(number => int.Parse(number.Value))
                .ToList();
            var compareCards = Regex.Matches(split[1].Split("|")[1], @"\d+").Select(number => int.Parse(number.Value))
                .ToList();
            var cardsPoints = 0;
            var factor = 2;
            foreach (var number in winningCards)
                if (compareCards.Contains(number))
                    cardsPoints = cardsPoints > 0 ? cardsPoints * factor : 1;
            sum += cardsPoints;
        }
        return sum;
    }

    public static int GetTotalOfCards()
    {
        var total = 0;
        var input = Inputs.InputDayFour.Split("\n").ToList();
        var cards = new Dictionary<int, Tuple<List<int>, List<int>>>();
        var copies = new Dictionary<int, Tuple<List<int>, List<int>>>();
        foreach (var card in input)
        {
            var split = card.Split(":");
            var id = int.Parse(Regex.Match(split[0], @"\d+").ToString());
            var winningCards = Regex.Matches(split[1].Split("|")[0], @"\d+").Select(number => int.Parse(number.Value))
                .ToList();
            var compareCards = Regex.Matches(split[1].Split("|")[1], @"\d+").Select(number => int.Parse(number.Value))
                .ToList();
            cards.Add(id, new Tuple<List<int>, List<int>>(winningCards, compareCards));
            var cardsPoints = 0;
            foreach (var number in winningCards)
            {
                if (compareCards.Contains(number))
                {
                    cardsPoints++;
                }
            }

            if (cardsPoints > 0)
            {
                for (int i = 1; i < cardsPoints; i++)
                {
                    
                }
            }
        }
        return cards.Count + copies.Count;
    }
}