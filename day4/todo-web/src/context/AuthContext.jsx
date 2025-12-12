import { useState, useEffect, useContext, createContext } from "react"
import { loginUser } from "../api/auth"

const AuthContext = createContext()

export function useAuth() {
    return useContext(AuthContext)
}

export default function AuthProvider({ children }) {
    const [user, setUser] = useState(null)
    const [token, setToken] = useState(localStorage.getItem("token"))
    const [loading, setLoading] = useState(false)

    // untuk simpan data user ketika login
    function saveAuth(token, user) {
        localStorage.setItem("token", token)
        localStorage.setItem("user", JSON.stringify(user))
        setUser(user)
        setToken(token)
    }

    // callback loginUser
    async function login(email, password) {
        setLoading(true)
        const data = await loginUser({ email, password })
        setLoading(false)

        if (data?.token) {
            saveAuth(data.token, data.user)
            return { success: true }
        }

        return { success: false, message: data.message }
    }

    // load user saat refresh
    useEffect(() => {
        const savedUser = localStorage.getItem("user")
        if (savedUser) {
            setUser(JSON.parse(savedUser))
        }
    }, [])

    const value = { user, token, loading, login }

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    )
}
