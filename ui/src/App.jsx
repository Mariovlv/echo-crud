import React, { useEffect, useState, useContext } from "react";
import { getAlbums } from "./api/albums";
import AlbumList from "./components/AlbumList";
import Unauthorized from "./components/Unauthorized";
import Navbar from "./components/Navbar";
import UserContext from "../context/UserContext"; // Check the correct path to UserContext

function App() {
  const [albums, setAlbums] = useState([]);
  const [isUnauthorized, setIsUnauthorized] = useState(false);
  const { userId } = useContext(UserContext);

  useEffect(() => {
    const fetchAlbums = async () => {
      try {
        console.log(userId);
        const response = await getAlbums();
        if (response.status === 401) {
          setIsUnauthorized(true);
        } else {
          setAlbums(response.data);
        }
      } catch (error) {
        if (error.response && error.response.status === 401) {
          setIsUnauthorized(true);
        } else {
          console.error("Failed to fetch albums", error);
        }
      }
    };

    fetchAlbums();
  }, [userId]); // Add userId to the dependency array

  if (isUnauthorized) {
    return <Unauthorized />;
  }

  return (
    <div>
      <Navbar userID={userId} />
      <AlbumList albums={albums} />
    </div>
  );
}

export default App;
