package function

import (
	"bytes"
	"html/template"
)

const backgroundColor = "green"

// Handle a serverless request
func Handle(req []byte) string {
	return demoPage(backgroundColor)
}

func demoPage(color string) string {
	pageTemplate := `
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Center Flex Item Image</title>
<style>
.box {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

body {
  background: {{ . }};
  margin: 0px;
}
</style>

</head>
<body>

<div class="box">
   <img src="https://cdn-images-1.medium.com/max/1600/1*FoURA3pZUll402Nl_vn1JQ.png">
</div>

</body>
</html>
	`
	t, _ := template.New("page").Parse(pageTemplate)
	buf := new(bytes.Buffer)
	t.Execute(buf, color)
	return buf.String()
}
