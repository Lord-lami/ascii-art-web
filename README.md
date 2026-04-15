# ASCII Art Web
Ascii Art Fill is a website written in Go Language, HTML and CSS that draws the ASCII art of the ASCII text you pass to its text box. It can also color and align the text. Its output can also be downloaded as a text file.

It uses only the standard libraries of Go language. 

It uses banner files that have the art for each character arranged in the order of the ASCII table and separated by a newline.

It only works for ASCII characters. Unicode characters beyond the [ASCII table](https://www.ascii-code.com/) will cause errors.

Command characters apart from newlines will cause panics.

I hope to be able to make this program the starting point for a deterministic image generator.

## Installation
- `git clone https://github.com/Lord-lami/ascii-art-web.git`

## Usage
- Change directory to the `ascii-art-full` folder
- Run `./asciiartfull --output=<filename> --align=<alignment> --color=<color> <substring to be colored>OPTIONAL <text to draw>` in the terminal. Option (the ones that start with `--` are optional)
- To use a style other than the standard style. 
- - Ensure that the banner file of the style is in the `banners` folder as a `.txt` file.
- - Ensure the banner file is formatted properly as described in the project description.
- - Run `./asciiartfull --align=<alignment>OPTIONAL --color=<color>OPTIONAL <substring to be colored>OPTIONAL <text to draw> <style banner file name without .txt>OPTIONAL`

- The program gives an error when there is no text, the color is invalid or the alignment is invalid.
- The valid colors are: black, red, green, yellow, blue, magenta, cyan, white, default
- The valid alignments are: left, right, center, justify

## Demo
### Basic Usage
![Demonstration of how the program runs with one argument](./demo_images/basic_usage.png)

### With two arguments (text and banner style)
![Demonstration of how the program runs with two arguments. "Text" text first and "shadow" banner style ](./demo_images/two_arguments.png)

### With color flag and text argument
![Demonstration of how the program runs with  color flag = red and a single text argument](./demo_images/color_and_text.png)

### With color flag, substring and text containing the substring
![Demonstration of how the program runs with color flag = red, a substring and a text containing the substring](./demo_images/color_substr_str.png)

### With color flag, substring and text that doesn't contain the substring
![Demonstration of how the program runs with color flag = red, a substring and a text that doesn't contain the substring](./demo_images/color_missingsubstr_str.png)

### With color flag, substring, text containing the substring and banner style, shadow
![Demonstration of how the program runs with color flag = yellow, a substring, a text that contains the substring and a banner style, shadow](./demo_images/color_substr_str_banner.png)

### With color flag, align flag and long text.
![Demonstration of how the program works with a color flag = magneta, align flag = left and a long text](./demo_images/align_color_longtext.png)

### With align flag, multiline text and banner style
![Demonstration of how the program works with an align flag = right, multiline text and banner style, thinkertoy](./demo_images/align_multilinetext_banner.png)

### With align flag, color flag and multiline text
![Demonstration of how the program works with an align flag = center, color flag = blue and multiline text](./demo_images/align_color_multilinetext.png)

### With align flag, color flag, substring and multiline text
![Demonstration of how the program works with an align flag = justify, color flag = green, substring and multiline text](./demo_images/align_color_substr_multilinetext.png)

## Credits
- [Olamide Ifarajimi](https://acad.learn2earn.ng/git/oifaraji)

## License
Copyright © 2026

This Project is [GPL](https://www.gnu.org/licenses/gpl-3.0.en.html) Licensed