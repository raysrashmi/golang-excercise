package twelve

var days = []string{
    "first", "second", "third", "fourth", "fifth",
    "sixth", "seventh", "eighth", "ninth",
    "tenth", "eleventh", "twelfth",
}

var gifts = []string{
    "a Partridge in a Pear Tree.",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming",
}

func Verse(i int) string {
    verse := "On the " + days[i-1] + " day of Christmas my true love gave to me: "

    for j := i - 1; j >= 0; j-- {
        if j == 0 && i > 1 {
            verse += "and "
        }
        verse += gifts[j]

        if j > 0 {
            verse += ", "
        }
    }

    return verse
}

func Song() string {
    result := ""

    for i := 1; i <= 12; i++ {
        result += Verse(i)
        if i < 12 {
            result += "\n"
        }
    }

    return result
}
