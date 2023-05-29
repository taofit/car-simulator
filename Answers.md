1, How long time did you end up spending on this coding test?
I spent around 10 to 12 hours on the task, time includes designing the programm, coding, designing test cases with command, testing, bug fixing, code cleaning up and refactoring. 

2, Explain why you chose the code structure(s) you used in your solution.
It is a relative small program, so I put everything under one folder(package in Golang), I save Object structs in one file car.go and other logic in another file carMove.go. If application gets bigger, there could be a `cmd` folder that contains the main application entry point files and a `pkg` folder that contains all other packages(services) like db, Api, app.
3, What would you add to your solution if you had more time?
I will add more test cases to test the wrong input, like what if the room size contains minus value or number with decimal, same thing applys to the car's input. I would also consider setting the maximum size of the room, so the program wont have a room without size limit.
4, What did you think of this recruitment test?
I think the coding test not only can access a candidate's coding skills, also the coding design, and unit testing, and the implementation of error handling. 