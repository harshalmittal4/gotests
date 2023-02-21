# Web app

Create a web server where users can track how many games players have won.
GET /players/{name} should return a number indicating the total number of wins
POST /players/{name} should record a win for that name, incrementing for every subsequent POST

In order to test our server, we will need a Request to send in and we'll want to
spy on what our handler writes to the ResponseWriter - i.e implement a type that can look like a ResponseWriter and has a method to be written by the handler
We use http.NewRequest to create a request. The first argument is the request's method and the second is the request's path. The nil argument refers to the request's body, which we don't need to set in this case.
net/http/httptest has a spy already made for us called ResponseRecorder so we can use that. It has many helpful methods to inspect what has been written as a response.
