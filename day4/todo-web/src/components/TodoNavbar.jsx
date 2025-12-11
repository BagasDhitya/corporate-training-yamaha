import { NavLink } from "react-router-dom";

export default function TodoNavbar() {
  return (
    <div className="w-screen p-5 flex justify-end bg-blue-700 text-white font-semibold">
      <ul className="space-x-5 flex">
        <li>
          <NavLink to={"/"}>Home</NavLink>
        </li>
        <li>
          <NavLink to={"/example"}>Example</NavLink>
        </li>
      </ul>
    </div>
  );
}
