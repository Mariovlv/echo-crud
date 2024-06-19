import { useEffect, useState } from "react";
import { getAlbums } from "./api/albums";
import AlbumList from "./components/AlbumList";
import Unauthorized from "./components/Unauthorized";
import Navbar from "./components/Navbar";

function App() {
  const [albums, setAlbums] = useState([]);
  const [isUnauthorized, setIsUnauthorized] = useState(false);

  useEffect(() => {
    const fetchAlbums = async () => {
      try {
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
  }, []);

  if (isUnauthorized) {
    return <Unauthorized />;
  }

  return (
    <div>
      <Navbar userID={4} />
      <AlbumList albums={albums} />
    </div>
  );
}

export default App;
