package models

func GetIndexElectron() string {
	return (`
/* create by create-app-cli */
const { join } = require('path');
const { app, BrowserWindow } = require('electron');

const createWindow = () => {

	const win = new BrowserWindow({
		width: 800,
		height: 600,
		webPreferences: {
			// preload: join( __dirname, 'preload.js' ),
			// contextIsolation: true,
		},
	});

	win.loadFile( join( __dirname, 'frontend', 'index.html' ) );
};

app.whenReady()
	.then(() => {
		
		createWindow();

		app.on('activate', () => {
			if ( BrowserWindow.getAllWindows().length === 0 ) {
				createWindow();
			}
		});
	});

app.on('window-all-closed', () => {
	
	if ( process.platform !== 'darwin' ) {
		app.quit();
	}
});
	`)
}

func GetElectronHTML() string {
	return (`
<!-- created by create-app-cli -->
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Electron app</title>
	</head>
	<body>
		<h1>index works</h1>
	</body>
</html>
	`)
}
