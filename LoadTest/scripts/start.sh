#!/bin/bash

#Feed database
mongo mongodb://mongo:27017/tasks --eval 'var document = { title : "task1", description : "description task 1" }; db.tasks.drop(); db.tasks.insert(document);'

#Run load tests
ab -c $1 -n $2 -r -g nodejs.tsv http://api-nodejs:3000/api/tasks
ab -c $1 -n $2 -r -g golang.tsv http://api-golang:4000/api/tasks
ab -c $1 -n $2 -r -g netcore3.tsv http://api-netcore:5000/api/tasks

#Plot graph
gnuplot performance.p