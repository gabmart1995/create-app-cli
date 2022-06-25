package models

import (
	"fmt"
)

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

	var htmlString = (`
<!-- created by create-app-cli -->
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>title</title>
		%s <!-- css library -->
		<link rel="stylesheet" href="./static/css/index.css">
	</head>
	<body>
		<h1>index works</h1>
		%s <!-- js script -->
		<script src="./static/js/index.js"></script>
	</body>
</html>
	`)

	// check library
	if bootstrap {

		htmlString = fmt.Sprintf(
			htmlString,
			"<link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC\" crossorigin=\"anonymous\">",
			"<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM\" crossorigin=\"anonymous\"></script>",
		)

	} else if materialize {

		// materialize css
		htmlString = fmt.Sprintf(
			htmlString,
			"<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css\">",
			"<script src=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js\"></script>",
		)

	} else {

		// is basic
		htmlString = (`
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
		<h1>index works</h1>
		<script src="./static/js/index.js"></script>
	</body>
</html>
		`)
	}

	return htmlString
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

/***********************/
/*	wordpress section */
/**********************/

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
	<h1>index works</h1>
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
	Version: 0.1
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

// setup theme
function theme_setup() {
}

add_action('after_setup_theme', 'theme_setup');

// add script and css
function theme_styles() {
}

add_action('wp_enqueue_scripts', 'theme_styles');
`)

	return template
}

func GetModelWidget() string {

	template := (`
<?php
/**
*  Plugin Name: Name -- Widgets
*  Plugin URI: 
*  Description: Widget created by create app cli
*  Version: 0.1
*  Author: 
*  Author URI: 
*  
*/

defined('ABSPATH') || exit;
	
/**
* Adds to widget. more info in https://codex.wordpress.org/Widgets_API
*/
class NameWidget extends WP_Widget {
}	
`)

	return template
}

func GetModelPlugin() string {
	return (`
<?php 
/**
*  Plugin Name: Name -- Widgets
*  Plugin URI: 
*  Description: Widget created by create app cli
*  Version: 0.1
*  Author: 
*  Author URI: 
*/

defined('ABSPATH') || exit;
`)
}
