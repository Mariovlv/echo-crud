import React from "react";
import { Link } from "react-router-dom";

const Unauthorized = () => {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100">
      <div className="bg-white shadow-md rounded-lg p-6 md:p-10 text-center">
        <h2 className="text-2xl md:text-3xl font-semibold text-red-500 mb-4">
          Unauthorized Access
        </h2>
        <p className="text-gray-700 mb-6">
          You are not authorized to view this page. Please log in to continue.
        </p>
        <Link
          to="/login"
          className="inline-block bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded transition duration-200"
        >
          Go to Login
        </Link>
      </div>
    </div>
  );
};

export default Unauthorized;
