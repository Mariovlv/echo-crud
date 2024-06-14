import { useEffect, useState } from "react";
import { getAlbums } from "./api/albums";
import AlbumList from "./components/AlbumList";

function App() {
  const [albums, setAlbums] = useState([]);

  useEffect(() => {
    const fetchAlbums = async () => {
      try {
        const response = await getAlbums();
        console.log(response.data);
        setAlbums(response.data);
      } catch (error) {
        console.error("error fetching albums:", error);
      }
    };

    fetchAlbums();
  }, []);

  return <AlbumList albums={albums} />;
}

export default App;
