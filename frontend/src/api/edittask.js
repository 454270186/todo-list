const BASE_URL = '/api/task/edit'

export const editName = async (taskID, newName) => {
    const resp = await fetch(`http://127.0.0.1:8080${BASE_URL}/`+'?task_id='+taskID+'&new_name='+newName+'&action=1', {
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

export const editCompleted = async (taskID) => {
    const resp = await fetch(`http://127.0.0.1:8080${BASE_URL}/`+'?task_id='+taskID+'&action=2', {
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