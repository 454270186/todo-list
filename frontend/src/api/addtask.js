const BASE_URL = '/api/task/add'

export const addNewTask = async (userID, name, taskID) => {
    const resp = await fetch(`http://127.0.0.1:8080${BASE_URL}/?user_id=`+userID+'&name='+name+'&task_id='+taskID, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    })

    const data = await resp.json()
    if (resp.ok && data.status_code === 0) {
        return true
    } else {
        return false
    }
}