# API REST performance
Compare performance between NodeJS, Golang and NetCore 3 using MongoDB. The test inserts one document in MongoDb and reads it using get operation from API REST.

Steps to run performance test:

- Execute this command:
  ```bash
  docker-compose up
  ```
- Attach shell to LoadTest container and run:
  ```bash
  ./start.sh 1000 1000
  ```
    - Param1: concurrency
    - Param2: requests
    
- When execution completed, png added to the path:
  ```bash
  /LoadTest/scripts/performance.png
  ```
  
  ![Performance image](https://github.com/AlexMorales85/ApiRestPerformance/blob/master/LoadTest/scripts/performance.png)
