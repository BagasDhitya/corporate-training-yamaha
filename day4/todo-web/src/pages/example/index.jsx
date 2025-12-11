export default function Example() {
  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center">
      <h1>Tailwind Responsivity</h1>
      <div className="grid grid-cols-1 lg:grid-cols-6 md:grid-cols-2 sm:grid-cols-1">
        <div className="p-5 bg-blue-500 rounded-md text-white">Item 1</div>
        <div className="p-5 bg-blue-500 rounded-md text-white">Item 2</div>
        <div className="p-5 bg-blue-500 rounded-md text-white">Item 3</div>
        <div className="p-5 bg-blue-500 rounded-md text-white">Item 4</div>
        <div className="p-5 bg-blue-500 rounded-md text-white">Item 5</div>
        <div className="p-5 bg-blue-500 rounded-md text-white">Item 6</div>
      </div>
    </div>
  );
}
