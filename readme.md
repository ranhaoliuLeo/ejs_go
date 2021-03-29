## ejs_go

```golang
TransTo(OriginPath string, DestinationPath string, FileName string, TemplateMap[string]string)
```

### for Example:
- index.ejs
```html
<html>
<head><%= title =></head>
<%= body =>
</html>
```
- main.go
```golang
package main
import("...")

func main() {
templateMap := map[string][string]{
    "title": "ranhao",
    "body": "world",
}
TransTo("./index.ejs", "./", "index.html", templateMap)
```
