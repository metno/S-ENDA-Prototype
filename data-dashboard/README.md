# S-ENDA Data dashboard
This project was bootstrapped with [MET React Boilerplate](https://gitlab.met.no/team-frontend/react-boilerplate.git).

## Relevant files for this project
```
src/
Dockerfile
```

## Test dashboard locally
```shell
npm install
npm start
```

## Build and run dashboard docker container
```
docker build -t data-dashboard .
docker run -p 8081:80 data-dashboard
```

### Background information
### Available Scripts

In the project directory, you can run:

#### `npm start`

Runs the app in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br>
You will also see any lint errors in the console.

#### `npm run build`

Builds the app for production to the `build` folder.<br>
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br>
Your app is ready to be deployed!

See the section about [deployment](#deployment) for more information.