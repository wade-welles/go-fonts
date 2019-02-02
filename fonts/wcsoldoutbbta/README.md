# wcsoldoutbbta

![wcsoldoutbbta](wcsoldoutbbta.png)

To use this font in your code, simply import it:

```go
import (
  . "github.com/gmlewis/go-fonts/fonts"
  _ "github.com/gmlewis/go-fonts/fonts/wcsoldoutbbta"
)

func main() {
  // ...
  render, err := fonts.Text(xPos, yPos, xScale, yScale, message, "wcsoldoutbbta")
  if err != nil {
    return err
  }
  log.Printf("MBB: (%.2f,%.2f)-(%.2f,%.2f)", render.Xmin, render.Ymin,render.Xmax, render.Ymax)
  for _, poly := range render.Polygons {
    // ...
  }
  // ...
}
```