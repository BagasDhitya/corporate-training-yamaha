import { useState } from "react";
import TodoAuthForm from "../../../components/TodoAuthForm";
import { registerUser } from "../../../api/auth";

export default function Register() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function handleRegister(e) {
    e.preventDefault();
    const response = await registerUser({ email, password });
    console.log("login result : ", response);
  }

  return (
    <div className="w-screen h-screen flex justify-center items-center">
      <TodoAuthForm
        title={"Register"}
        onSubmit={handleRegister}
        emailValue={email}
        passwordValue={password}
        onEmailChange={(e) => setEmail(e.target.value)}
        onPasswordChange={(e) => setPassword(e.target.value)}
      />
    </div>
  );
}
