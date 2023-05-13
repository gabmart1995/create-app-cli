package snipets

func GetReactModel() string {
	return (`
import React from 'react';
import ReactDom from 'react-dom/client';

import './style.css';

const root = ReactDom.createRoot(document.getElementById('root'));
root.render(<h1 className="test">Hello world</h1>);
	`)
}

func GetIndexModel() string {
	return (`
<!DOCTYPE html>
<html>
	<head>
		<title>Empty project</title>
		<meta charset="utf-8">
	</head>
	<body>
		<div id="root"></div>
		<script src="bundle.js"></script>
	</body>
</html>
	`)
}

func GetWebpackConfig() string {
	return (`
const path = require('path');

const config = {
  entry: [
    './src/index.js'
  ],
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        use: 'babel-loader',
        exclude: /node_modules/
      },
      {
        test: /\.css$/,
        use: [
            'style-loader',
            'css-loader'
        ],
      },
      // images
      {
          test: /\.(png|jpg|svg|gif)$/,
          use: [
              'file-loader'
          ]
      },
    ]
  },
  devServer: {
    contentBase: './dist'
  }
};

module.exports = config;
	`)
}

func GetBabelConfig() string {
	return (`
{
	presets: [
		[
			'@babel/preset-env',
			{ modules: false }
		],
		'@babel/preset-react'
	],
	plugins: [
	]
}	
	`)
}
