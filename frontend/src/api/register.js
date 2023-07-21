const BASE_URL = '/api'

export const registerReq = async (formData) => {
    try {
        const resp = await fetch(`http://127.0.0.1:8080${BASE_URL}/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        })

        const data = await resp.json()
        if (resp.ok) {
            console.log(data)
            if (data.status_code !== 0) {
                throw new Error(data.msg)
            }
        } else {
            throw new Error(data.msg)
        }
    } catch(error){
        throw new Error(error.message)
    }
}