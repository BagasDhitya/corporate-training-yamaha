import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "../App";
import Example from "../pages/example";

export default function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/example" element={<Example />} />
      </Routes>
    </BrowserRouter>
  );
}
