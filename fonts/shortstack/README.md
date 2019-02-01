# shortstack

![shortstack](shortstack.png)

To use this font in your code, simply import it:

```go
import (
	. "github.com/gmlewis/go-fonts/fonts"
	_ "github.com/gmlewis/go-fonts/fonts/shortstack"
)

func main() {
	// ...
	render, err := Text(x, y, xs, ys, message, "shortstack"),
	// ...
}
```