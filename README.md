# 🤖 Intelligent Chatbot System for Streamlined Ticketing Processes 🎫

This project presents an intelligent chatbot system designed to streamline ticketing processes in large organizations. The system follows the CRUD (Create, Read, Update, Delete) model and utilizes a RESTful architecture for efficient customer management.

## 📋 Table of Contents
1. [Requirements](#requirements)
2. [Installation](#installation)
3. [Usage](#-usage)
4. [Project Structure](#project-structure)
5. [Contributions](#contributions)

## Requirements <a name="requirements"></a>
To run the intelligent chatbot system, you need to have the following software installed:
- Docker
- Docker Compose

## 🚀 Installation <a name="installation"></a>
Follow these steps to set up the project:

1. Clone the repository:
  ```bash Copy code
  git clone https://github.com/your-username/intelligent-chatbot-system.git
  ```

Navigate to the project directory:
```bash Copy code
cd Final_year_project
```

Download the required models:

Mistral-7b-OpenOrca model:

Visit the GPT4All website and choose the OpenOrca option.
Download the model file and rename it to mistral-7b-openorca.gguf2.Q4_0.gguf.


LLaMA 2 model:

Visit the LLaMA 2 model repository on Hugging Face.
Download the model file and rename it to llama-2-7b.Q4_0.gguf.




Create a folder named Gpt4all inside the project/chatbot directory.
Place the downloaded model files (mistral-7b-openorca.gguf2.Q4_0.gguf and llama-2-7b.Q4_0.gguf) inside the Gpt4all folder.
Build and run the Docker containers:
bashCopy codedocker-compose up --build
This command will build the necessary Docker images and start the containers. The process may take a few minutes to complete.
Once the containers are up and running, the intelligent chatbot system should be ready to use! 🎉

## 💬 Usage <a name="usage"></a>

To interact with the chatbot system, follow these steps:

1.  Open your web browser and navigate to http://localhost:3000.

2. Create an admin acount

3. you will be now able to ecess the chatbot, but beore use genereate some tikcekts so the chatbot has it as information


4. You will be presented with a user-friendly interface where you can engage in conversations with the chatbot. 🗨️

5. Type your queries or requests related to ticketing processes, and the chatbot will provide intelligent and relevant responses. ✨

```bash
📂 Project Structure
The project structure is organized as follows:
Copy codeintelligent-chatbot-system/
├── backend/
│   ├── Dockerfile
│   └── ...
├── chatbot/
│   ├── Dockerfile
│   ├── Gpt4all/
│   │   ├── mistral-7b-openorca.gguf2.Q4_0.gguf
│   │   └── llama-2-7b.Q4_0.gguf
│   └── ...
├── frontend/
│   ├── Dockerfile
│   └── ...
├── docker-compose.yml
└── ...
```
The backend directory contains the backend server code and its associated Dockerfile.
The chatbot directory houses the chatbot implementation, including the Gpt4all folder where the downloaded models should be placed.

The frontend directory contains the frontend user interface code and its associated Dockerfile.

The docker-compose.yml file defines the services and their configurations for running the system using Docker Compose.

## 🤝 Contributions <a name="contributions"></a>

Contributions to the intelligent chatbot system are welcome!
If you encounter any issues, have suggestions for improvements, or would like to add new features, please submit a pull request or open an issue on the GitHub repository. 😊

Let's collaborate to make this chatbot system even better and revolutionize the way organizations handle their ticketing processes! 🌟
Copy code
This README file includes emojis to add visual appeal and enhance the readability of the content. Feel free to customize the emojis or add more to suit your preferences!