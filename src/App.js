import Todo from "./components/Todo";
import Form from "./components/Form";
import FilterButton from "./components/FilterButton";
import { useState } from "react";
import { nanoid } from "nanoid";

function App(props) {
	const [tasks, setTasks] = useState(props.tasks)

	function toggleTaskCompleted(id) {
		const updateTasks = tasks.map((task) => {
			if (id == task.id) {
				return {...task, completed: !task.completed}
			}
			return task
		})
		setTasks(updateTasks)
	}

	function delTask(id) {
		const remainTasks = tasks.filter((task) => id !== task.id)
		setTasks(remainTasks)
	}

	const taskList = tasks?.map((task) => (
		<Todo 
			id={task.id} 
			completed={task.completed} 
			name={task.name}
			key={task.id}
			toggleTaskCompleted={toggleTaskCompleted}
			delTask={delTask}
			/>
	))

	function addTask(name) {
		const newTask = {id: `todo-${nanoid()}`, name, completed: false}
		setTasks([...taskList, newTask])
	}

	const taskNano = taskList.length !== 1 ? "tasks" : "task"
	const headingText = `${taskList.length} ${taskNano} remaining`

  return (
    <div className="todoapp stack-large">
      <h1>TodoMatic</h1>

      <Form addTask={addTask} />

      <div className="filters btn-group stack-exception">
        <FilterButton />
        <FilterButton />
        <FilterButton />
      </div>
      <h2 id="list-heading">{headingText}</h2>
      <ul
        role="list"
        className="todo-list stack-large stack-exception"
        aria-labelledby="list-heading">
				{taskList}
      </ul>
    </div>
  );
}

export default App;
