# Internship-Task-1
This is a repository that contains sample tasks particularly in GoLang


## Installation steps

Step 1: Clone the Repository

Open your terminal or command prompt.
Use the following command to clone the repository:

```
git clone https://github.com/OnkarOjha/Internship-Task-1
```

Step 2: Navigate to the Repository Root

Change your current working directory to the root of the cloned repository:

```
cd Internship-Task-1
```
Step 3: Execute the Batch Script

To start the server, run the run_app.bat batch script in the terminal:

```
./run_app.bat
``` 

## Docker steps

Step 1 : Build the Docker image using the following command (don't forget the . at the end, which indicates the current directory

```
docker build -t task-1 .
```

Step 2 : Once the image is successfully built, you can run the application inside a Docker container using the following command:

```
docker run -p 8000:8000 task-1
```
Happy coding!
