const BASE_URL = import.meta.env.VITE_BASE_URL

export async function registerUser(payload) {
    const response = await fetch(`${BASE_URL}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    })
    const result = await response.json()
    return result
}

export async function loginUser(payload) {
    const response = await fetch(`${BASE_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    })
    const result = await response.json()
    return result
}