# GIT

git is a project based on solving a series of task and include your feedback

##  Set up git on your local machine 
    install git on your local machine [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Configure git with username and email address
    + Open the command line
    + Set up your username:
        ```bash
            git config --global user.name "EXAMPLE_NAME"
        ```
    + Set up your address:
        ```bash
            git config --global user.email "FULL_NAME@example.com"
        ```
    + Save credentials in plaintext on your local machine
        - Run:
            ```bash
                git config --global credential.helper store
            ```
        - Verify:
            ```bash
                git config --global credential.helper
            ```
            you should get :
                ```bash
                    store
                ```
        - Then:
            ```bash
                git pull
            ```
        - Enter your credentials for the last time and congrats you have saved your credentials.

## Git commits to commit
    create work directory, establish a subdirectory named hello, inside hello create a file called hello.sh and input the following content:
        echo "Hello, World"
    in order to do that we should run the following commands
    ```bash
        mkdir -p work/hello
    ```
    then go to hello Dir
    ```bash
        cd hello
    ```
    ```bash
        touch hello.sh
    ```
    use a text editor to edit hello.sh (Vscode in my example):
    ```bash
        code hello.sh
    ```
    navigate to hello directory and initialize the directory using the command git init:
    ```bash
        git init
    ```
    the ./git Dir is hidden by default the command **git init** display the path of the ./git Dir
    chek status using **git status**:
    ```bash
        git status
    ``` 
    this command will tell you that there is no commits yet becuase the repo is brand new.
    first we need to add hello.sh to track by simply:
    ```bash
        git add hello.sh
    ``` 
    we are going to add a "#!/bin/bash" to our hello.sh and commit changes:
    ```bash
        git commit -m "commit the changes of hello.sh"
    ```
    after trying using **git push** we failed  because we have to configure the push destinationb by the following commands:
    ```bash
        git remote add origin <url>
    ```
    then setting the upstram :
    ```bash
        git push --set-upsteream origin master
    ```
    or :
    ```bash
        git push -u origin master
    ```
    I repeated the same process to commit two separated commits you will discover why on the next session.

## History
    in order to show the history of the current working Dir:
    ```bash
        git log
    ```
    use Ctrl+z to quit
    Show One-Line History for a condensed view showing only commit hashes and messages:
    ```bash
        git log --oneline
    ```
    to customize the log output to show only the last 2 commits :
    ```bash
        git log --oneline -
    ```
    to customize to view the commits made the last 5 minutes :
    ```bash
        git log --oneline -since="5 minutes ago"
    ```
    Personalized Format:
        use the following command ** git log --pretty=format:'%h %ad | %s (%d) [%an]' --date=short **
        ```bash
            git log --pretty=format:'%h %ad | %s (%d) [%an]' --date=short
        ```
        - h for hash
        - ad for date --date=short is the format of the date
        - s for the commit message
        - d for the branch infos 
        - an for author name
# Check it out
    in order to get the hash of the first commit you need to run the command **git rev-list --max-parents=0 HEAD**
        ```bash
            git rev-list --max-parents=0 HEAD
        ```
    after that use **git checkout <commit hash>** in my case [9c879256f](https://learn.zone01oujda.ma/git/mlarbi/git/commit/9c879256f57447028cdefa2c26a05bb9a02d150a):
        ```bash
            git checkout 9c879256f57447028cdefa2c26a05bb9a02d150a
        ```
    