import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { getUserByID } from "../api/albums";

function Navbar({ userID }) {
  const [userInfo, setUserInfo] = useState({ id: "", email: "" }); // Initialize userInfo as an object

  useEffect(() => {
    const fetchUser = async () => {
      try {
        console.log(userID);
        const response = await getUserByID(userID);
        setUserInfo({ id: response.data.id, email: response.data.email }); // Update userInfo with the correct property names
        console.log(response);
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };

    if (userID) {
      fetchUser();
    }
  }, [userID]); // Add userID to the dependency array

  return (
    <nav className="bg-gray-800 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <div className="flex space-x-4">
          <Link
            to="/"
            className="text-white text-lg font-semibold hover:text-gray-300"
          >
            Home
          </Link>
          <Link
            to="/login"
            className="text-white text-lg font-semibold hover:text-gray-300"
          >
            Login
          </Link>
        </div>
        {userInfo.id && ( // Check if userInfo.id exists before rendering
          <div className="text-white">
            <span className="mr-4">ID: {userInfo.id}</span>
            <span>Email: {userInfo.email}</span>
          </div>
        )}
      </div>
    </nav>
  );
}

export default Navbar;
