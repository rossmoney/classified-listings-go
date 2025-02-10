# classified-listings-go

## How to run the project 
- please make sure you have go and yarn installed
- run `make install` this will configure the UI
- run `make run` this will start both the nuxt.js frontend and the go backend
- visit `http://localhost:8080` in the browser or run `make open`

## Viewing API doc
- please make sure you have docker running
- please run `make serve-swagger`
- visit `http://localhost:8000` in the browser

## Assumptions made
- That I would use a Vue UI component library like Vuetify

## What I would do differently in a production environment
- Embed the ui files within the application with go embedFS (I started looking into this)
- Move env vars out into .env files like the app port for example.