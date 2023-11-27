# weatherApp

This project is created to build a Go based application to fetch current weather of a city using an external API from openweathermap.org

Files in the Repo: 
    main.go - It contains basic functions in Golang to fetch the data using an external API 
    main_test.go - It contains test cases for the functions present in main.go file 
    Dockerfile - It contains the docker code to containerize and create image for the application 
    go.mod 
    README.md

To use openweathermap 3.0 API, we read the documentation and found that API 3.0 is One Call API and is only subscription based. That's why we used API 2.5

The note we found on official documentaion page -
    "Please note, that One Call API 3.0 is included in the "One Call by Call" subscription only."

By default, it gives temperature in Kelvin, therefore we used their parameter "units=metric" to change into Celsius 