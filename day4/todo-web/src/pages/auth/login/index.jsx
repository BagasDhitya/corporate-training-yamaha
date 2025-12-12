import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../../context/AuthContext";
import TodoAuthForm from "../../../components/TodoAuthForm";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { login } = useAuth();
  const navigate = useNavigate();

  async function handleLogin(e) {
    e.preventDefault();
    const data = await login(email, password);

    // if (data.success) {
    //   navigate("/dashboard");
    // } else {
    //   alert(data.message);
    // }
  }

  return (
    <div className="w-screen h-screen flex justify-center items-center">
      <TodoAuthForm
        title={"Login"}
        onSubmit={handleLogin}
        emailValue={email}
        passwordValue={password}
        onEmailChange={(e) => setEmail(e.target.value)}
        onPasswordChange={(e) => setPassword(e.target.value)}
      />
    </div>
  );
}
