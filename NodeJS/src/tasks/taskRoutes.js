const taskController =  require('./taskController')

const routes = [
    {
        method: 'GET',
        url: '/api/tasks',
        handler: taskController.getTasks
    }
]

module.exports = routes