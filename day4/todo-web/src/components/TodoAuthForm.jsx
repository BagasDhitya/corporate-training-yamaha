export default function TodoAuthForm({
  onSubmit,
  emailValue,
  passwordValue,
  onEmailChange,
  onPasswordChange,
  title,
}) {
  return (
    <form
      onSubmit={onSubmit}
      className="w-full h-96 p-5 rounded-md shadow-md bg-slate-100"
    >
      <h2 className="text-xl font-bold mb-5">{title}</h2>
      <div className="mb-5">
        <label>Email</label>
        <input
          className="p-3 border text-black"
          type="email"
          value={emailValue}
          onChange={onEmailChange}
        />
      </div>
      <div className="mb-5">
        <label>Password</label>
        <input
          className="p-3 border text-black"
          type="email"
          value={passwordValue}
          onChange={onPasswordChange}
        />
      </div>
      <button type="submit" className="p-3 rounded-md text-white bg-blue-500">
        Submit
      </button>
    </form>
  );
}
