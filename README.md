# Go-Vasttrafik

This program is a small utility for use with google home projects. It provides a small integration towards the Västtrafiks Resiplanner.
Eventually this will implement the intents from a custom application on Google Home. This is to check the next bus or tram times from a specific location suing Gothenburgs vasttrafik. 

## Prereqs

In order to run this code you need to create a developer account at the vasttrafik website. 
https://developer.vasttrafik.se/portal/#/

Create an account, log in and then create an app, give it a name and you wil get a key and a secret which you can use with the client_credentails grant. 

## Building

`make build`

## Running 

`make run VT_KEY="Nyckel value here" VT_SECRET="Hemlighet value here"`

## TODO

Add support for Trip planner
Add go-chi router for the Google Home dialog flows
Add support for Intents from Google Home