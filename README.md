# Big Text

create walls of random letters, numbers, or symbols effortlessly.

example:

`./bigtext -letters -numbers -sprinkle "&" -sr 100 -count 10000 -truncate 10000`

options:

`-letters`: generate letters

`-numbers`: generate numbers

`-symbols`: generate symbols

these can be combined! if you want letters and numbers just use both the -letters and -numbers flag.

`-sprinkle`: sprinkle a character in randomly. default: ""

`-sr <int>`: sprinkle rarity ex. 100 picks a random number between 1 and 100 and if the number equals 1, it adds the sprinkle. default: 5

`-count`: how many times to generate a character. default: 6

`-delimeter`: character to be placed in between each character generated. default: ""

`-truncate`: using sprinkles or delimeters can add the generated character length so truncate can chop off the # of characters after a certain amount. default: off (0)

`-nl`: adds a new line after a provided character. default: off (0)

TODO:

program works well for what i want it to do, generate some walls of text to hide flags in for a ctf, but there are definitely some improvements to be made.
- implement the go strings.Builder type to minimize memory copying.
- speed improvements
- usability improvements
