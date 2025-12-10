import { useEffect, useState } from "react";
import { getAllTodos } from "../../api/todo";

export default function Dashboard() {
  const [todos, setTodos] = useState([]);

  const [search, setSearch] = useState("");
  const [category, setCategory] = useState("");
  const [sort, setSort] = useState("desc");

  async function loadTodos() {
    const data = await getAllTodos(search, category, sort);
    setTodos(data);
  }

  useEffect(() => {
    loadTodos();
  }, []);

  return (
    <div className="max-w-4xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">Todo Dashboard</h1>

      {/* filter area */}
      <div className="bg-white p-4 rounded shadow mb-6 grid grid-cols-1 md:grid-cols-4 gap-4">
        {/* search */}
        <div>
          <label className="block text-sm font-medium mb-1">Search</label>
          <input
            type="text"
            className="w-full border rounded p-2"
            placeholder="Search title or description ..."
            onChange={(e) => setSearch(e.target.value)}
          />
        </div>
      </div>

      {/* table */}
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white border">
          <thead className="bg-gray-100">
            <tr>
              <th className="border px-4 py-2">ID</th>
              <th className="border px-4 py-2">Title</th>
              <th className="border px-4 py-2">Description</th>
              <th className="border px-4 py-2">Category</th>
              <th className="border px-4 py-2">Completed</th>
            </tr>
          </thead>
          <tbody>
            {todos.data.length === 0 ? (
              <tr>
                <td className="text-center p-4 text-gray-500">
                  No todos found
                </td>
              </tr>
            ) : (
              todos.data.map((item) => (
                <tr key={item.id}>
                  <td className="border px-4 py-2">{item.id}</td>
                  <td className="border px-4 py-2">{item.title}</td>
                  <td className="border px-4 py-2">{item.description}</td>
                  <td className="border px-4 py-2">{item.category}</td>
                  <td className="border px-4 py-2">
                    {item.isCompleted ? "✅" : "❌"}
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
