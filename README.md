# WebSockets 

---

### Table of Contents
- [Description](#description)
---

## Description
#### HomeScreen
This displays two categories namely Student and Teacher. 

#### Client Side
Clicking on the student category will open up the student side where he has to enter the roll number. As soon as he clicks the submit button, a textarea will be displayed which is disabled
by default. It is enabled when the websocket connection is established. when the student starts typing in the textarea, each word is sent to the golang server.

#### ServerSide
The golang server will receive the data sent by the client through a websocket. It then performs the calculations over the data, counts the total number of words, characters, words per minute
and stores them in a key-value pair. Every time the input of the textarea is changed, the data is sent to the server. The summary is also pushed to the redis database .

Further, the server sends back the calculated summary to the teacher(client) where the calculated data is displayed inside a table.

#### issues
Handling multiple connections using gorilla websocket in golang is yet to be implemented.



#### Technologies

-HTML
-CSS
-JavaScript
-GoLang


