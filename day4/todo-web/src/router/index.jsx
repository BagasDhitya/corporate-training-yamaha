import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "../App";
import Example from "../pages/example";
import CreateTodo from "../pages/create-todo";
import Dashboard from "../pages/dashboard";
import Login from "../pages/auth/login";
import Register from "../pages/auth/register";

import TodoNavbar from "../components/TodoNavbar";
import ProtectedRoute from "../components/ProtectedRoute";

export default function Router() {
  return (
    <BrowserRouter>
      <TodoNavbar />
      <Routes>
        <Route path="/" element={<App />} />
        <Route
          path="/example"
          element={
            <ProtectedRoute>
              <Example />
            </ProtectedRoute>
          }
        />
        <Route
          path="/create-todo"
          element={
            <ProtectedRoute>
              <CreateTodo />
            </ProtectedRoute>
          }
        />
        <Route
          path="/dashboard"
          element={
            <ProtectedRoute>
              <Dashboard />
            </ProtectedRoute>
          }
        />
        <Route path="/auth/register" element={<Register />} />
        <Route path="/auth/login" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}
