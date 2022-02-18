
# ROVER PROBLEM

A squad of robotic rovers are to be landed by NASA on a plateau on Mars. 

This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the
surrounding terrain to send back to Earth.
A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points.
The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left
corner and facing North.

In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90
degrees left or right respectively, without moving from its current spot.
'M' means move forward one grid point, and maintain the same heading.
Assume that the square directly North from (x, y) is (x, y+1).

Output:
The output for each rover should be its final co-ordinates and heading.


### Sample Input and Output
Plateau max X and Y, Starting coordinates, direction and path for two rovers:

5 5

1 2 N

LMLMLMLMM

3 3 E

MMRMMRMRRM

Output and new coordinates:

1 3 N

5 1 E


### I have implemeted the solution in two modes. First one is in CLI MODE and Second one is in RESTFUL Mode

#### 1. CLI MODE
When we start running main.go, On CLI It will ask the user to choose the mode. If user selects 1 then application will run on CLI Mode and if the user selects 2 then application will run on RESTFUL Mode.
#### Sample to demonstrate how CLI looks for the user.
###### WELCOME TO ROVER PROBLEM
###### ENTER YOUR CHOICE
###### PRESS 1 TO CONTINUE IN CLI MODE
###### PRESS 2 TO CONTINUE IN RESTFUL MODE


If the user selects 1 then application will run on CLI Mode.

##### WELCOME TO CLI MODE TO FIND THE ROVER POSITION
##### Enter max position seperated by space:
5 5
##### Enter Rover position seperated by space:
1 2 N
##### Command for rover:
lmlmlmlmm
##### To Continue Inputs Y/N
N
##### Rover Current Position:
1 3 N

#### 2. RESTFUL MODE

##### API END-POINTS:
##### 1.  GETBYID: Get call to fetch the rover position based on the ID
    EndPoint: http://127.0.0.1:3000/<ID>

##### 2. GET: Get call to fetchall the rover positions
    EndPoint: http://127.0.0.1:3000

##### 3. POST: POST Call to find the rover position
    EndPoint: http://127.0.0.1:3000
PayLoad: 

    {
    "max_x": 5,
    "max_y": 5,
    "location": "3 3 E",
    "command": "MMRMMRMR"
    }

Sample-output:

    {
    "id": "6283",
    "max_x": 5,
    "max_y": 5,
    "location": "3 3 E",
    "command": "MMRMMRMR",
    "roverPosition": "4 1 N"
    }



### Generated a .deb package to perform the same operations.

##### Steps to install & run the package.
    1. To install the package. Make sure that you're in a directory where you downloaded the pack.deb file.
       On the CLI type the below command:

       >> sudo dpkg -i pack.deb 

       It will install the downloaded package pack.deb

    2. To run the installed package/Application, Enter the following command on the CLI

       >> Rover

       Once you hit the above command, Your Application will start running on CLI.
    
### Created a Docker File to perform the same operations.

##### Steps to run Dockerfile

    1.  To Pull the image from DockerHub

        >> sudo docker pull malli1998/roverapp


    2.  To run the Application

        >> sudo docker run -it -p 8001:8001 malli1998/roverapp

        Now the Application to start running on the CLI.
     