import { useCount } from "./hooks/useCount";
import TodoButton from "./components/TodoButton";

export default function App() {
  const { count, increment, decrement } = useCount();

  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center">
      <div className="p-4 mb-3">
        <h1 className="text-3xl font-bold text-blue-600">Todo List Web</h1>
      </div>
      <div className="flex p-4 space-x-5 justify-center items-center">
        <TodoButton onHandler={increment} title={"+"} type={"increment"} />
        <h2 className="text-black">{count}</h2>
        <button
          onClick={decrement}
          className="p-3 bg-red-500 text-white w-40 rounded-md"
        >
          -
        </button>
      </div>
    </div>
  );
}
