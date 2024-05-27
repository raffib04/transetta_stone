# transetta_stone

## Introduction 

Our group comes from a variety of cultural and linguistic backgrounds. John is Italian and German and has traveled extensively due to his family’s military background. Ian is Latin American and Chinese. Rafael is Polish and German. Our linguistic backgrounds are all unique, and we were united in the belief that people should be able to communicate regardless of language. After failing to find an online chat room that does this, we decided to make a websockets-based chat application that would translate messages live. 


## Project Description

Our project is a WebSocket-powered chat application written in Go. Users select a username and language of choice when logging in. Users can create a new chat room or join an existing one. In the chat room, users send and receive messages in their chosen language.

The project uses Go for the backend, a simple React application for the front end, and AWS for translation and database services. Specifically, we use AWS Translate for on-demand, real-time translation, AWS RDS for message storage, and AWS S3 for image and file storage.

## Components and Architecture

The project consists of three main components: the frontend, the backend, and the translation service. The frontend is a React application that uses the WebSocket API to connect to the server. The backend is a Go application that uses the Gorilla WebSocket library to handle WebSocket connections. 

We use Go’s concurrency features to provide a scalable websocket server that can handle multiple chat rooms and users. We also use Go’s channels to communicate between the websocket server and the translation service. The translation service is a separate Go routine that sends messages to AWS Translate and receives the translated messages back. The translated messages are then sent back to the websocket server and broadcasted to the chat room.

The front end is hosted on AWS Amplify, and the backend is hosted on an EC2 instance. 


## Implementation

The chat application is currently running on a web server where users can go to the URL, fill out a username, their language, and their desired chat room, and then continuously send and receive messages in their chosen language. The messages that are sent will be immediately processed through AWS Translate through a concurrent Go subroutine and then outputted in the translated form for all members in the chat room. This is the same for files as they will also be translated for the users once a file is sent. 

All messages are stored in a separate AWS RDS table everytime a new chat room is created, and this way all users will be able to see all past messages in their chosen chat room. Additionally, files will be stored in a bucket in AWS S3 and users will be able to see their uploaded files based on their userid, which is accessed through a user table in AWS RDS as well. 


## Conclusion

By combining our unique cultural and linguistic backgrounds, we were able to create a chat application that allows people to communicate regardless of language. Our project demonstrates the power of technology to bring people together and break down barriers. We hope that our project inspires others to create innovative solutions to real-world problems and promotes cross-cultural understanding and communication. We will continue working on this project to improve its features and make it more user-friendly. We also plan to add more languages and translation services to make the chat application more accessible to a wider audience. We believe that our project has the potential to make a positive impact on society and help people connect and communicate in new and exciting ways.