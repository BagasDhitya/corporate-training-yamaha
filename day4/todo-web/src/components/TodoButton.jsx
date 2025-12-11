export default function TodoButton({ onHandler, title, type = 'increment' | 'decrement' }) {
  return (
    <button
      className={`p-3 ${type === "increment" ? 'bg-green-500' : 'bg-red-500'} text-white w-40 rounded-md`}
      onClick={onHandler}
    >
      {title}
    </button>
  );
}
