import { useEffect, useState } from "react";
import { getAlbums } from "./api/albums";
import AlbumList from "./components/AlbumList";

const App = () => {
  const [albums, setAlbums] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchAlbums = async () => {
      try {
        const response = await getAlbums();
        setAlbums(response.data);
        setLoading(false);
      } catch (err) {
        setError(err);
        setLoading(false);
      }
    };

    fetchAlbums();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <AlbumList albums={albums} />
    </div>
  );
};

export default App;
