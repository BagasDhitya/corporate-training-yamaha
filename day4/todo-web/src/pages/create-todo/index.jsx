import { useState } from "react";
import { createTodo } from "../../api/todo";

export default function CreateTodo() {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("programming");
  const [isCompleted, setIsCompleted] = useState(false);

  async function handleCreate(e) {
    e.preventDefault();

    const payload = {
      title,
      description,
      category,
      isCompleted,
    };

    await createTodo(payload);

    alert("Todo created!");
  }

  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center gap-10">
      {/* 📝 FORM CREATE TODO */}
      <form
        onSubmit={handleCreate}
        className="p-5 rounded-md shadow-md border w-96"
      >
        <div className="flex flex-col mb-5">
          <label>Title</label>
          <input
            type="text"
            className="p-3 text-black rounded-md border"
            onChange={(e) => setTitle(e.target.value)}
          />
        </div>

        <div className="flex flex-col mb-5">
          <label>Description</label>
          <textarea
            className="p-3 text-black rounded-md border"
            onChange={(e) => setDescription(e.target.value)}
          />
        </div>

        <div className="flex flex-col mb-5">
          <label>Category</label>

          <select
            className="p-3 text-black rounded-md border"
            onChange={(e) => setCategory(e.target.value)}
            value={category}
          >
            <option value="programming">Programming</option>
            <option value="health">Health</option>
            <option value="food">Food</option>
          </select>
        </div>

        <button
          type="submit"
          className="w-full p-3 bg-blue-600 rounded-md text-white"
        >
          Create Todo
        </button>
      </form>
    </div>
  );
}
