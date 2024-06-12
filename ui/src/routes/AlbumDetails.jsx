import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { getAlbumById } from "../api/albums";

const AlbumDetails = () => {
  const { id } = useParams();
  const [album, setAlbum] = useState(null);

  useEffect(() => {
    const fetchAlbum = async () => {
      try {
        const response = await getAlbumById(id);
        setAlbum(response.data);
      } catch (error) {
        console.error("Error fetching album:", error);
      }
    };

    fetchAlbum();
  }, [id]);

  if (!album) {
    return <div>Loading...</div>;
  }

  return (
    <ul className="flex flex-col m-3 font-mono">
      <li className="text-xl text-stone-700">ID {album.id}</li>
      <li className="text-4xl font-bold">{album.title}</li>
      <li className="text-xl text-stone-800">{album.artist}</li>
    </ul>
  );
};

export default AlbumDetails;
