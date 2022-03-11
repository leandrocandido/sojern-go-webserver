# sojern-go-webserver

This is my first project in Go, so would be great hear from you what kind of improvements I can do here. I follow at first moment this simple tutorial: https://go.dev/doc/tutorial/web-service-gin 

# Endpoints: 
/median : given list of numbers calculates their median
/avg : given list of numbers calculates their average
/min : given list of numbers and a quantifier (how many) provides min number(s)
/max : given list of numbers and a quantifier (how many) provides max number(s)
/percentile : given list of numbers and quantifier 'q', compute the qth percentile of the list elements

# Postman
this is the postman collection used to test the service : https://www.getpostman.com/collections/4ac4badc6599c5dc1f89.
Ps: if you want to use web postman you have to install postman agent in or machine in order to call http://localhot
https://www.postman.com/downloads/postman-agent/

# Project
This project is comprised of 2 modules, the main one used to hold the endpoint methods(webapi) and the second one to hodl the business logic of this project(calclib).
Calclib project have unit tests to validade its methods. 
