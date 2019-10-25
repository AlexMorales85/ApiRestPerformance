using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace apirest.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class TasksController : ControllerBase
    {
        private readonly ITaskRepository taskRepository;

        public TasksController(ITaskRepository taskRepository)
        {
            this.taskRepository = taskRepository;
        }

        [HttpGet]
        public Task<IEnumerable<TaskEntity>> GetTasks()
        {
            return this.taskRepository.GetTasksAsync();           
        }
    }
}
