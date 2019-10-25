using System.Collections.Generic;
using System.Threading.Tasks;

namespace apirest
{
    public interface ITaskRepository
    {
        Task<IEnumerable<TaskEntity>> GetTasksAsync();
    }
}