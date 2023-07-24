import React from "react";
import { useState, useEffect } from "react";
import { nanoid } from "nanoid";
import Todo from "../components/Todo";
import Form from "../components/Form";
import FilterButton from "../components/FilterButton";
import { getCookie } from "../utils/cookie";
import { getTasksByID } from "../api/gettasks";
import { addNewTask } from "../api/addtask";
import { deleteTask } from "../api/deletetask";

const TodoList = () => {
    const FILTER_MAP = {
        All: () => true,
        Active: (task) => !task.completed,
        Completed: (task) => task.completed,
    }
    const FILTER_NAME = Object.keys(FILTER_MAP)

    const [tasks, setTasks] = useState([])
    const [filter, setFilter] = useState("All")

    const toggleTaskCompleted = (id) => {
        const updateTasks = tasks.map((task) => {
            if (id === task.id) {
                return {...task, completed: !task.completed}
            }
            return task
        })
        setTasks(updateTasks)
    }

    const delTask = async (id) => {
        let originTasks = tasks
        const remainTasks = tasks.filter((task) => id !== task.id)
        setTasks(remainTasks)

        const isSuccess = await deleteTask(id)
        if (!isSuccess) {
            setTasks(originTasks)
        }
    }

    const addTask = async (name) => {
        let originTasks = tasks
        const newTask = {id: `todo-${nanoid()}`, name, completed: false}
        setTasks([...tasks, newTask])

        const userID = getCookie("user_id")
        const isSuccess = await addNewTask(userID, name, newTask.id)
        if (!isSuccess) {
            setTasks(originTasks)
        }
    }

    const editTask = (id, newName) => {
        const editedTaskList = tasks.map((task) => {
			if (task.id === id) {
				return {...task, name: newName}
			}
			return task
		})

		setTasks(editedTaskList)
    }

    useEffect(()=>{
        // get tasks by user id
        const userID = getCookie("user_id")
        let gotTasks
        const gettasks = async () => {
        try {
            gotTasks = await getTasksByID(userID)
            
            setTasks(gotTasks)
        } catch(error) {
            console.error('error:', error)
        }
    }
    gettasks()
    }, [])

    const taskList = tasks.filter(FILTER_MAP[filter]).map((task) => {
        return <Todo 
            id={task.id} 
            completed={task.completed} 
            name={task.name}
            key={task.id}
            toggleTaskCompleted={toggleTaskCompleted}
            delTask={delTask}
            editTask={editTask}
        />
    })

    const taskNano = taskList.length !== 1 ? "tasks" : "task"
	const headingText = `${taskList.length} ${taskNano} remaining`

	const filterList = FILTER_NAME.map((name) => {
		
		return	<FilterButton 
				key={name}
				name={name} 
				isPressed={name === filter}
				setFilter={setFilter}
				/>
	})

    return (
        <div className="todoapp stack-large">
          <h1>Todo-List</h1>
    
          <Form addTask={addTask} />
    
          <div className="filters btn-group stack-exception">
            {filterList}
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

export default TodoList