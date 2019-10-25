const app = require('./app')

const start = async () => {
    try {
        app.listen(3000, "0.0.0.0")
        console.log('Server listening on 3000')
    } catch (err) {
        throw err;
    }
}

start()