const fastify = require('fastify')()
const mongoose = require('mongoose')
const taskRoutes = require('./tasks/taskRoutes')

mongoose
    .connect('mongodb://mongo/tasks')
    .then(() => console.log('MongoDB connected...'))
    .catch(err => console.log(err))

taskRoutes.forEach((route) => {
    fastify.route(route)
})

module.exports = fastify