import { useState } from "react";
import TodoAuthForm from "../../../components/TodoAuthForm";
import { loginUser } from "../../../api/auth";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function handleLogin(e) {
    e.preventDefault();
    const response = await loginUser({ email, password });
    console.log("login result : ", response);
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
