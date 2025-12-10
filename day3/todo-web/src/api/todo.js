const BASE_URL = "http://localhost:8080/api"

export async function getAllTodos(search = '', category = '', sort = 'desc') {
    const params = new URLSearchParams()

    if (search) params.append("search", search)
    if (category) params.append("category", category)
    if (sort) params.append("sort", sort)

    const response = await fetch(`${BASE_URL}/todos?${params.toString()}`)
    return await response.json()
}

export async function getTodoById(id) {
    const response = await fetch(`${BASE_URL}/todos/${id}`)
    return await response.json()
}

export async function createTodo(payload) {
    const response = await fetch(`${BASE_URL}/todos`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
    return await response.json()
}