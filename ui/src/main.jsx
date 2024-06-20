import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { UserProvider } from "../context/UserContext"; // Import the UserProvider
import "./index.css";

// Route components
import App from "./App.jsx";
import RouteAlbumDetails from "./routes/RouteAlbumDetails.jsx";
import RouteLoginForm from "./routes/RouteLoginForm.jsx";
import RouteSignUp from "./routes/RouteSignUp.jsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/albums/:id",
    element: <RouteAlbumDetails />,
  },
  {
    path: "/login",
    element: <RouteLoginForm />,
  },
  {
    path: "/signup",
    element: <RouteSignUp />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <UserProvider>
      <RouterProvider router={router} />
    </UserProvider>
  </React.StrictMode>
);
