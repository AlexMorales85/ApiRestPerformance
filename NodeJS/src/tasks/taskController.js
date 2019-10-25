const Task = require('./taskModel')

exports.getTasks = async () => {
    try {
        const tasks = await Task.find()
        return tasks

    } catch (err) {
        throw Error('Failed to get tasks');
    }
}