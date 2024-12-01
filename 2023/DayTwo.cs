using System.Text;

namespace AoCDayOne;

public static class DayTwo
{
    private static readonly Dictionary<string, int> Configuration = new()
    {
        { "red", 12 },
        {"blue", 14},
        {"green", 13}
    };


    /*
     * 1. On dictionary level, Find max number for each colour for every subset
     * 2. Multiply them
     * 3. Calculate sum for every game
     */
    public static int GetPowerOfSets()
    {
        var sum = 0;
        var gameDictionary = CreateGameDictionary();
        foreach (var game in gameDictionary)
        {
            var green = -1;
            var blue = -1;
            var red = -1;
            foreach (var subsets in game.Value)
            {
                if (subsets.ContainsKey("green"))
                {
                    green = subsets["green"] > green ? subsets["green"] : green;
                }
                if (subsets.ContainsKey("blue"))
                {
                    blue = subsets["blue"] > blue ? subsets["blue"] : blue;
                }
                if (subsets.ContainsKey("red"))
                {
                    red = subsets["red"] > red ? subsets["red"] : red;
                }
            }
            sum += green * blue * red;
        }
        return sum;
    }
    
    public static int GetSumOfIds()
    {
        /*
         * 1. Filter IDs
         * 2. Filter Subsets of cubes
         * 3. Check for each subset if the configuration is possible -> Match colours and then check if drawn cube count is higher than configuration
         * 4. Take the possible games and sum the Ids.
         * Rules:
         * - Elf puts the cubes back
         * - Not every colour has to be shown
         * - A game is possible if all subsets are possible
         */
        var ret = 0;

        var gameDictionary = CreateGameDictionary();

        foreach (var game in gameDictionary)
        {
            var possible = true;
            foreach (var subsets in game.Value)
            {
                if (!IsPossible(subsets))
                {
                    possible = false; break;
                }
            }
            if (possible)
            {
                ret += game.Key;
            }
        }
        return ret;

    }

    private static Dictionary<int, List<Dictionary<string, int>>> CreateGameDictionary()
    {
        var listOfGames = Inputs.InputDayTwo.Split("\n").ToList();
        var gameDictionary = new Dictionary<int, List<Dictionary<string, int>>>();

        foreach (var game in listOfGames)
        {
            var sp = game.Split(":");
            gameDictionary.Add(ParseNumber(sp[0]), CreateSubsetDictionary(sp[1].Split(";").ToList()));
        }

        return gameDictionary;
    }


    private static bool IsPossible(Dictionary<string, int> gameSubsets)
    {
        foreach (var subset in gameSubsets)
        {
            if (subset.Value > Configuration[subset.Key])
            {
                return false;
            }
        }

        return true;
    }

    private static List<Dictionary<string, int>> CreateSubsetDictionary(List<string> subsets)
    {
        var subsetDictionary = new List<Dictionary<string, int>>();
        foreach (var subset in subsets)
        {
            var tempDict = new Dictionary<string, int>();
            var tempList = subset.Split(",").ToList();
            foreach (var str in tempList)
            {
                var temp = str.Split(" ", StringSplitOptions.RemoveEmptyEntries);
                tempDict.Add(temp[1], Int32.Parse(temp[0]));
            }
            subsetDictionary.Add(tempDict);
        }
        return subsetDictionary;
    }

    private static int ParseNumber(string input)
    {
        var list = input.Select( ch=> 
        {
            if (Char.IsNumber(ch))
            {
                return Int32.Parse(ch.ToString());
            }

            return -1;
        }).Where(i => i!= -1).ToList();

        StringBuilder sb = new StringBuilder();
        foreach (var i in list)
        {
            sb.Append(i);
        }
        return Int32.Parse(sb.ToString());
    }
    
}