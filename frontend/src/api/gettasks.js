const BASE_URL = '/api/task'

// Get task list by user ID
export const getTasksByID = async (userID) => {
    const resp = await fetch(`http://127.0.0.1:8080${BASE_URL}/`+userID, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })

    const data = await resp.json()
    if (resp.ok && data.status_code === 0) {
        return data.tasks
    } else {
        throw new Error(data.msg)
    }
}