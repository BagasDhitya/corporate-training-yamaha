import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "../App";
import Example from "../pages/example";
import CreateTodo from "../pages/create-todo";

import TodoNavbar from "../components/TodoNavbar";

export default function Router() {
  return (
    <BrowserRouter>
      <TodoNavbar />
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/example" element={<Example />} />
        <Route path="/create-todo" element={<CreateTodo />} />
      </Routes>
    </BrowserRouter>
  );
}
