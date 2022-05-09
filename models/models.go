package models

import "regexp"

// colores de consola
const (
	ColorBlack  = "\u001b[30m"
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorBlue   = "\u001b[34m"
	ColorReset  = "\u001b[0m"
	ColorCyan   = "\u001B[36m"
	ColorWhite  = "\u001B[37m"
)

func GetHTMLModel(bootstrap bool) string {

	htmlString := (`
<!-- created by create-app-cli -->
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>title</title>
		{{:link_css}}
		<link rel="stylesheet" href="./static/css/index.css">
	</head>
	<body>
		index page
		{{:link_script}}
		<script src="./static/js/index.js"></script>
	</body>
</html>
	`)

	var regex []*regexp.Regexp = []*regexp.Regexp{
		regexp.MustCompile(("[{]{2}:link_css[}]{2}")),
		regexp.MustCompile(("[{]{2}:link_script[}]{2}")),
	}

	var result string

	// check library
	if bootstrap {

		result = regex[0].ReplaceAllString(
			htmlString,
			"<link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC\" crossorigin=\"anonymous\">",
		)

		result = regex[1].ReplaceAllString(
			result,
			"<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM\" crossorigin=\"anonymous\"></script>",
		)

	} else {
		// materialize css
		result = regex[0].ReplaceAllString(
			htmlString,
			"<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css\">",
		)

		result = regex[1].ReplaceAllString(
			result,
			"<script src=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js\"></script>",
		)
	}

	return result
}

func GetCSSModel() string {
	return (`
/** create-app-cli */
:root {
}

body {
}	
	`)
}

func GetJSModel() string {
	return (`
/** create-app-cli */
function init() {
	
}	
	`)
}

func GetPartial(name string) string {
	template := (` <p>{{:name}} works</p>`)
	regex := regexp.MustCompile("[{]{2}:name[}]{2}")

	return regex.ReplaceAllString(template, name)
}
