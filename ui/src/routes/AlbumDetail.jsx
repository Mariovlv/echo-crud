import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getAlbumById } from "../api/albums";
import AlbumLikesBy from "../components/AlbumLikesBy";

const AlbumDetail = () => {
  const params = useParams(); // Correctly use the useParams hook
  const [albumData, setAlbumData] = useState({
    album: null,
    loading: true,
    error: null,
    users: [],
  });

  useEffect(() => {
    const fetchAlbum = async () => {
      try {
        const response = await getAlbumById(params.id);
        setAlbumData({
          ...albumData,
          album: response.data,
          users: response.data.users,
          loading: false,
        });
      } catch (err) {
        setAlbumData({ ...albumData, error: err, loading: false });
      }
    };

    fetchAlbum();
  }, []);

  const { album, users, loading, error } = albumData;

  if (loading) return <p className="text-center text-gray-500">Loading...</p>;
  if (error)
    return <p className="text-center text-red-500">Error: {error.message}</p>;

  return (
    <div className="max-w-md mx-auto bg-white rounded-xl shadow-lg overflow-hidden md:max-w-2xl mt-10 font-mono">
      <div className="md:flex bg-indigo-500 p-8">
        <h2 className="uppercase tracking-wide text-sm text-white font-semibold">
          Album Details
        </h2>
      </div>
      <div className="p-8">
        {album && (
          <>
            <p className="block mt-1 text-lg leading-tight font-medium text-black">
              ID: {album.id}
            </p>
            <p className="mt-2 text-gray-700">Title: {album.title}</p>
            <p className="mt-2 text-gray-700">Artist: {album.artist}</p>
            <AlbumLikesBy users={users} />
          </>
        )}
      </div>
    </div>
  );
};

export default AlbumDetail;
