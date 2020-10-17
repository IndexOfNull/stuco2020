# Stuco2020
Just some software for Covid-19 compliant student  council voting. Written with a Go backend and Vue frontend.

## Config
### Backend
Edit the `config/database.go` and change the proper info in the `BuildDBConfig()` function. I know this is awful practice, so it will change when I start distributing binaries. I may also add more command line options or environment variable support (not really a priority right now) to avoid baking in your config.
### Frontend
Edit the `client/src/globalSettings.js` and change the values as needed. As of now the most important one should be `backendUrl`, which is where the frontend can find the backend.
## Deploying
First, make sure everything is configured (see above) and that you're `cd`'ed into the root directory.
### Frontend
```
cd client
npm install //Only do this if you haven't already
npm run build
```
Your files should be placed into the `client/dist` folder; you can deploy those onto a webserver.
### Backend
Make sure you have go 1.13 or higher.
```
go build
```
Your backend binary should be placed in the projects root directory.