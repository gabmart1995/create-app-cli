package snipets

func GetReactModel() string {
	return (`
import React from 'react';
import ReactDom from 'react-dom/client';

import './style.css';

const rootElement = document.getElementById('root');

if (rootElement) {
  const root = ReactDom.createRoot(rootElement);
  root.render(<h1 className="test">Hello world</h1>);
}
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

func GetWebpackConfig(ts bool) string {
	if ts { // flag de configuracion en typescript
		return (`
const path = require('path');

const config = {
  entry: [
    './src/index.tsx'
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
        test: /\.ts(x)?$/,
        loader: 'ts-loader',
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

func GetWebpackConfig5(ts bool) string {
	if ts { // flag de configuracion en typescript
		return (`
const path = require('path');

const config = {
  entry: [
    './src/index.tsx'
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
        test: /\.ts(x)?$/,
        loader: 'ts-loader',
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
    static: {
      directory: './dist'
    }
  }
};

module.exports = config;
        `)
	}

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
      /** discoment this for activate module of typescript */
      /*{
        test: /\.ts(x)?$/,
        loader: 'ts-loader',
        exclude: /node_modules/
      },*/
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
    static: {
      directory: './dist'
    } 
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

/** contenido del archivo con configuracion de typescript */
func GetTypescriptConfig() string {
	return (`
{
  "compilerOptions": {
      "allowSyntheticDefaultImports": true,
      "outDir": "./dist/",
      "sourceMap": true,
      "strict": true,
      "noImplicitReturns": true,
      "noImplicitAny": true,
      "module": "es6",
      "moduleResolution": "node",
      "target": "es5",
      "allowJs": true,
      "jsx": "react-jsx",
      "isolatedModules": true
  },
  "include": [
      "./src/**/*"
  ]
}
  `)
}
