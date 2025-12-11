import { useEffect, useState } from "react";
import { useCount } from "./hooks/useCount";
import TodoButton from "./components/TodoButton";

export default function App() {
  const { setCount, count, increment, decrement } = useCount();
  const [data, setData] = useState();

  async function getData() {
    const response = await fetch("http://localhost:8080/api/todos");
    const data = await response.json();
    setData(data.data);
  }

  useEffect(() => {
    console.log("component mount ...");
  }, []);

  useEffect(() => {
    getData();
  }, []);

  console.log("data : ", data);

  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center">
      <div className="p-4 mb-3">
        <h1 className="text-3xl font-bold text-blue-600">Todo List Web</h1>
      </div>
      <div className="flex p-4 space-x-5 justify-center items-center">
        <TodoButton onHandler={increment} title={"+"} type={"increment"} />
        <div className="flex flex-col space-y-5 justify-center items-center">
          <h2 className="text-black">{count}</h2>
          <button
            className="p-5 bg-slate-200 text-black rounded-md"
            onClick={() => setCount(0)}
          >
            Refresh
          </button>
        </div>
        <TodoButton onHandler={decrement} title={"-"} type={"decrement"} />
      </div>
    </div>
  );
}
