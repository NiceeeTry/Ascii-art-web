
# Ascii-Art-Web

It is a web server for outputting the text entered by user in a graphic representation of ASCII.

## To execute the program run the following code: go run main.go
To see the web page insert http://localhost:8080/ in a browser.
## Description
Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of a program that outputs text in a graphical representation of ASCII (ascii-art project)

Webpage must allow the use of the different banners:
- shadow
- standard
- thinkertoy

Implement the following HTTP endpoints:
- GET /: Sends HTML response, the main page.
- 1.1. GET Tip: go templates to receive and display data from the server.
- POST /ascii-art: that sends data to Go server (text and a banner)
- POST Tip: use form and other types of tags to make the post request.\
### The main page must have:
- text input
- radio buttons, select object or anything else to switch between banners
- button, which sends a POST request to '/ascii-art' and outputs the result on the page.

## Usage
After cloning the repository, simply do a [go run main.go.] After that, connect to the localhost:8080 through your browser and start using the static website.
## 