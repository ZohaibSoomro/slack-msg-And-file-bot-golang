# Slack Message and File Bot (Golang)

This is a simple Slack bot written in Golang that can respond to mentions with predefined messages and send a file when requested. It serves as a basic template for creating your own custom Slack bots.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- Listens for mentions in a Slack channel.
- Responds to mentions with messages like "hi" and "hello" with a default response.
- Responds to "where is file?" with a file attachment.
- Easy-to-customize responses and actions.

## Prerequisites

Before you can use this bot, make sure you have the following:

- A Slack account.
- A Slack app and bot created on your Slack workspace.
- Slack API tokens for your app and bot user.

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/ZohaibSoomro/slack-msg-and-file-bot-golang.git
   ```

2. Change directory to the project folder:

   ```bash
   cd slack-msg-and-file-bot-golang
   ```

3. Build the project:

   ```bash
    go build
   ```

## Configuration

    To configure the bot, you need to provide your Slack app and bot tokens. Create a .env file in the project root directory and add the following lines
    ```bash
    SLACK_APP_TOKEN=your_slack_app_token_here
    SLACK_BOT_TOKEN=your_slack_bot_token_here
    ```
    Replace your_slack_app_token_here and your_slack_bot_token_here with the respective tokens generated for your Slack app and bot user.

## Usage

    1. Run the bot
        ```bash
        ./slack-msg-and-file-bot-golang
        ```
    The bot will start listening for mentions in the Slack channels it has access to.
    2. In your Slack workspace, mention the bot in a channel with one of the following messages:
        "hi"
        "hello"
        "where is file?"
    The bot will respond accordingly.
    3. To customize the bot's responses or actions, edit the main.go file and modify the handleMessage function.

## Contributing

    Contributions are welcome! If you have any ideas or improvements, please open an issue or create a pull request.
