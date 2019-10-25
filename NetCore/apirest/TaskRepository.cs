using MongoDB.Driver;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace apirest
{
    public class TaskRepository: ITaskRepository
    {
        private MongoClient client;

        public TaskRepository()
        {
            client = new MongoClient("mongodb://mongo:27017");
        }

        public async Task<IEnumerable<TaskEntity>> GetTasksAsync()
        {
            var tasks = await this.client.GetDatabase("tasks")
                .GetCollection<TaskEntity>("tasks")
                .FindAsync(new FilterDefinitionBuilder<TaskEntity>().Empty);
            return tasks.ToEnumerable();
        }
    }
}
