# Car movement

## Program explanation

The program provides a simulator for radio controlled cars!

This is what the simulator should do:

- Handle various commands to make the cars move.
- Handle various types of radio cars.
- Provide a room for the cars to move in.

Action commands and how they should work:

- A car can move and turn as follows; move forward, move back, turn left and turn right. These commands shall be represented by input commands F, B, L, and R respectively.
- A car can have a direction; North, west, south or east. These commands shall be represented by input commands N, W, S, and E respectively.
- A turn to either the left or the right shall be considered as a 90 degree turn.

Rooms and how it should work:

- The rooms will need two input parameters to provide a scope for a rectangular shaped room.
- Scope input shall represent whole meters and have no decimals.
- The rooms shall be considered to be surrounded by walls.

Cars and how they should work:

- A car should move and turn according to command input.
- A car cannot move through a wall.
- A car can be considered to have a size scope of a 1x1 meter object.

The simulator should work as a console application. It shall run with inputted commands for room, cars and starting position in the room. The sequence for these commands should have the following order:

- Room input.
- Starting position and direction of a car.
- A sequence of action commands.

After a simulation has run, the simulator shall output the result. A result of the simulation can be either successful or unsuccessful.

- Criteria for a successful simulation is that the car moves through the room according to given commands, while not crashing into any wall during the route.
- The output from a successful simulation should consist of the end position of the car as well as the heading of the car.
- Criteria for an unsuccessful simulation is that the car crashes into a wall.
- The output from an unsuccessful simulation should describe an error of what went wrong.

## User Story

As a **user running the simulation**
I want to **provide the simulator with command input**
And get **simulation results based on said command input**
So that **I know if a route is successful or not**.

# Solution & Explanation

Solution is written in Golang

If a room size is w and h, then the valid car's position (x, y) is in the range of (0,0) to (w-1, h-1) inclusive.

If the car moves to North one cell, its y-axis will increment by 1, if the car moves to East on cell, its x-axis will increment by 1, and vice versa.

## Instruction

### Run program in docker:

1, docker build --tag app-car .<br>
2, docker run -it app-car

then follow the instruction in three steps<br>
step 1: enter 2 digit number separated by comma<br>
step 2: enter 2 digit number and a direction separated by comma<br>
setp 3: enter a series of letters(command) separated by comma<br>

### Run program in host machine without docker

make sure Golang is installed on the host machine, then run command: `go run .` under project root folder<br>
and follow the above three steps to run the program

## testing

1, run test using docker, run command: `docker-compose up`<br>
2, run test under host machine: run command: `go test -v`
