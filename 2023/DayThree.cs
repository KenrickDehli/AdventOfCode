using System.Text;

namespace AoCDayOne;

public static class DayThree
{
    public static int GetSumOfPartNumbers()
    {
        var sum = 0;
        var stringList = Inputs.InputDayThree.Split('\n').ToList();
        for (var i = 0; i < stringList.Count; i++)
        {
            var j = 0;
            while (j < stringList[i].Length)
                if (char.IsNumber(stringList[i][j]))
                {
                    var lengthOfNumber = 0;
                    var sb = new StringBuilder();
                    var k = j;
                    while (char.IsNumber(stringList[i][k]))
                    {
                        sb.Append(stringList[i][k]);
                        lengthOfNumber++;
                        k++;
                        if (k == stringList[i].Length) break;
                    }

                    var numberStartIndex = j;
                    var numberEndIndex = j + lengthOfNumber;
                    var number = int.Parse(stringList[i].Substring(j, lengthOfNumber));

                    var indexLeft = numberStartIndex - 1 < 0 ? 0 : numberStartIndex - 1;
                    var indexRight = numberEndIndex == stringList[i].Length ? numberEndIndex - 1 : numberEndIndex;
                    var indexAbove = i - 1 < 0 ? 0 : i - 1;
                    var indexBelow = i + 1 >= stringList.Count ? i : i + 1;

                    var charLeft = stringList[i][indexLeft];
                    var charRight = stringList[i][indexRight];
                    var charAboveLeft = stringList[indexAbove][indexLeft];
                    var charAboveRight = stringList[indexAbove][indexRight];
                    var charBelowLeft = stringList[indexBelow][indexLeft];
                    var charBelowRight = stringList[indexBelow][indexRight];

                    if (
                        (charLeft != '.' && !char.IsNumber(charLeft)) ||
                        (charRight != '.' && !char.IsNumber(charRight)) ||
                        (charAboveLeft != '.' && !char.IsNumber(charAboveLeft)) ||
                        (charAboveRight != '.' && !char.IsNumber(charAboveRight)) ||
                        (charBelowLeft != '.' && !char.IsNumber(charBelowLeft)) ||
                        (charBelowRight != '.' && !char.IsNumber(charBelowRight))
                    )
                        sum += number;

                    for (var l = numberStartIndex; l < numberEndIndex; l++)
                    {
                        var charAbove = stringList[indexAbove][l];
                        var charBelow = stringList[indexBelow][l];
                        if ((charAbove != '.' && !char.IsNumber(charAbove)) ||
                            (charBelow != '.' && !char.IsNumber(charBelow)))
                        {
                            sum += number;
                            break;
                        }
                    }

                    j += lengthOfNumber;
                }
                else
                {
                    j++;
                }
        }
        return sum;
    }
}