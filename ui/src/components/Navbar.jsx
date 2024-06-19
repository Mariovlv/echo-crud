// src/components/Navbar.js
import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { getUserByID } from "../api/albums";

function Navbar({ userID }) {
  const [userInfo, setUserInfo] = useState("");

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await getUserByID(userID);
        setUserInfo(response.data);
        console.log(response);
      } catch (error) {
        console.error("Error fetching album:", error);
      }
    };

    fetchUser();
  }, []);

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
        {userInfo && (
          <div className="text-white">
            <span className="mr-4">ID: {userInfo.ID}</span>
            <span>Email: {userInfo.email}</span>
          </div>
        )}
      </div>
    </nav>
  );
}

export default Navbar;
