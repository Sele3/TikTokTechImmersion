# Assignment Submission

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

## Setup Instructions
1. Clone the repo using the following command:
   > git clone https://github.com/Sele3/TikTokTechImmersion.git --config core.autocrlf=false

2. In the root directory, run the following command to start the server:
   > docker-compose up -d

3. To test that the server is working, run the following command:
   > curl localhost:8080/ping  
4. If the server is up and running, the following response will be returned:
   > {"message":"pong"}

## Development Environment
To work on this project, make sure you have the following development environment set up:
- Docker 

## Available Endpoints

The application exposes the following endpoints for sending and retrieving messages:

### Send Message

Send a message to a specific chat room.

- **URL:** `POST /api/send`
- **Request Body:**

```json
{
    "chat": "user1:user2",
    "text": "message-text",
    "sender": "user1"
}
```

### Pull Messages

Retrieve messages from a chat room.

- **URL:** `GET /api/pull`
- **Request Parameters:**

  - `chat`: Chat room identifier
  - `cursor`: Starting position of messages 
  - `limit`: Maximum number of messages to retrieve
  - `reverse`: Whether to sort messages in reverse order 

- **Response Body:**

```json
{
    "messages": [
        {
            "chat": "chat-room-id",
            "text": "message-text",
            "sender": "sender-identifier",
            "send_time": 0
        },
        ...
    ],
    "has_more": true,
    "next_cursor": 0
}
```

## Issue: Docker Error due to Line Encoding
When running the command `docker-compose up -d`, an error occurred when executing the line `RUN sh ./build.sh`.

### Root Cause Analysis
This issue is related to line encodings. When CLRF (Carriage Return Line Feed) line ending format is used, it causes compatibility issues with the shell script.

### Resolution
To resolve this issue, follow these steps:
1. Change the line encoding from CLRF to LF (Line Feed) format.

If the issue persists, try the following steps:
1. Run the following command to change the git config input:
   > git config --global core.autocrlf input
2. Delete the local repository.
3. Clone the repository again.

By following these steps, you should be able to resolve the Docker error related to line encoding.
