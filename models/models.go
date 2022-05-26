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

func GetHTMLModel(bootstrap bool, materialize bool, basic bool) string {

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

	} else if materialize {
		// materialize css
		result = regex[0].ReplaceAllString(
			htmlString,
			"<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css\">",
		)

		result = regex[1].ReplaceAllString(
			result,
			"<script src=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js\"></script>",
		)
	} else {

		// is basic
		result = (`
<!-- created by create-app-cli -->
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>title</title>
		<link rel="stylesheet" href="./static/css/index.css">
	</head>
	<body>
		index page
		<script src="./static/js/index.js"></script>
	</body>
</html>
			`)
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

/** wordpress section */

func GetModelWordpress() string {
	template := (`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
</head>
<body>
	index.php
</body>
</html>
`)
	return template
}

func GetModelStyleWordpress() string {
	template := (` 
/*
	** Change to needs more information in: https://codex.wordpress.org/Theme_Development **

	Theme Name: Create App Template
	Author: create app cli
	Version: 1.0
	Description: "template created by create-app-cli"
	License: GNU General Public License v2 or later
	Tags: basic, responsive-design
*/

:root {
}

body {
}
`)

	return template
}

func GetModelFunctionsWordpress() string {
	template := (`
<?php

/** 
*	communication file with the wordpress api you can write the functions you need to configure 
*	your template and integrate them to the administration panel using the hooks
*/

function theme_setup() {
}

add_action('init', 'theme_setup');
`)

	return template
}
