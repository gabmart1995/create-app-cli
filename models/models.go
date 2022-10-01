package models

import "fmt"

// colores de consola UNIX
/* const (
	ColorBlack  = "\u001b[30m"
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorBlue   = "\u001b[34m"
	ColorReset  = "\u001b[0m"
	ColorCyan   = "\u001B[36m"
	ColorWhite  = "\u001B[37m"
) */

func GetHTMLModel() string {

	var htmlString = (`
<!-- created by create-app-cli -->
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>title</title>
		<!-- css library -->
		<link rel="stylesheet" href="./static/css/index.css">
	</head>
	<body>
		<h1>index works</h1>
		<!-- js script -->
		<script src="./static/js/index.js"></script>
	</body>
</html>
	`)
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
*  Plugin Name: Name -- Plugins
*  Plugin URI: 
*  Description: Widget created by create app cli
*  Version: 0.1
*  Author: 
*  Author URI: 
*/

defined('ABSPATH') || exit;
`)
}

/* create the model of web component */
func GetModelComponent(name string) map[string]string {
	modelFiles := make(map[string]string)

	name += "-component"

	// html
	modelFiles[name+".html"] = fmt.Sprintf(`
<!-- /* web-component create-app-cli */ -->
<link rel="stylesheet" href="" />
<div>
	<p>%s works<p>
</div>	
`, name)

	// css
	modelFiles[name+".css"] = (`
/* web-component create-app-cli */	
p {
}
`)
	// js
	modelFiles[name+".js"] = fmt.Sprintf(`
class MyComponent extends HTMLElement {
	constructor() {
		super();
	}

	/* props observed */
	static get observedAttributes() {
		return [];
	}

	connectedCallback() {
	}

	attributeChangedCallback( name, oldValue, newValue ) {
	}

	render() {
	}

	disconnectedCallback() {
	}
}

customElements.define('%s', MyComponent);
	`, name)

	return modelFiles
}
