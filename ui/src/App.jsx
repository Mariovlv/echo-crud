import React, { useEffect, useState } from "react";
import { getAlbums } from "./api/albums";
import AlbumList from "./components/AlbumList";

const App = () => {
  const [albums, setAlbums] = useState({
    albumList: null,
    loading: true,
    error: null,
  });

  useEffect(() => {
    const fetchAlbums = async () => {
      try {
        let isThereAuthcookie = document.cookie;
        if (isThereAuthcookie) {
          // return an element to ask user to log
          return;
        }
        const response = await getAlbums();
        setAlbums({ ...albums, albumList: response.data, loading: false });
      } catch (err) {
        setAlbums({ ...albums, error: err, loading: false });
      }
    };

    fetchAlbums();
  }, []);

  const { albumList, loading, error } = albums; // Destructure loading and error from albums

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <AlbumList albums={albumList} /> {/* Pass albumList instead of albums */}
    </div>
  );
};

export default App;
