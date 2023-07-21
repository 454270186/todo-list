const BASE_URL = '/api'

export const loginReq = async (formData) => {
    try {
        console.log(formData)
        const resp = await fetch(`http://127.0.0.1:8080${BASE_URL}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        })

        if (resp.ok) {
            const data = await resp.json()
            return data
        } else {
            throw new Error("Login failed")
        }
    } catch {
        throw new Error("Network error")
    }
}