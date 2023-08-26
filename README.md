# Stuco2020
This is software I wrote for my school's student council elections when COVID-19 made in-person voting impossible. Each student is given a unique code (can also be a QR code) that is used to cast their vote. This uses a Go backend based on gin-gonic and Vue frontend. This ran my high school's elections smoothly and was able to be recycled a year later.

Since the deadline for producing this software was tight, there are things about it I don't quite like (the frontend view for voting is  monstrous, for example). I'm not sure I would recommend anyone try to put this into production themselves, but it will be here if you'd like to try. 

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