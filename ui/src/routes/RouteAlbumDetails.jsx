import React, { useState, useEffect } from "react";
import { Link, useParams } from "react-router-dom";
import {
  getAlbumById,
  getLikedByAlbumID,
  insertLikeToAlbumByIDs,
} from "../api/albums";
import Unauthorized from "../components/Unauthorized";
import ButtonLike from "../components/ButtonLike";
import AlbumLikesBy from "../components/AlbumLikesBy";

const RouteAlbumDetails = () => {
  const { id } = useParams();
  const [album, setAlbum] = useState(null);
  const [users, setUsers] = useState();

  useEffect(() => {
    const fetchAlbum = async () => {
      try {
        const response = await getAlbumById(id);
        const respAlbums = await getLikedByAlbumID(id);
        console.log(respAlbums.data);
        setAlbum(response.data);
        setUsers(respAlbums.data);
        console.log(users);
      } catch (error) {
        console.error("Error fetching album:", error);
      }
    };

    fetchAlbum();
  }, [id]);

  if (!album) {
    return <Unauthorized />;
  }

  const likeAlbum = () => {
    console.log("you liked album with id", album.id);
  };

  return (
    <ul className="flex flex-col p-5 m-3 font-mono bg-white rounded-lg shadow-lg">
      <div className="mb-4">
        <li className="text-lg text-gray-600">ID: {album.id}</li>
        <li className="text-2xl font-bold text-gray-900">{album.title}</li>
        <li className="text-lg text-gray-700">{album.artist}</li>
      </div>
      <div className="flex space-x-3 mb-5">
        <ButtonLike
          isLikedByUser={false}
          likeAlbum={insertLikeToAlbumByIDs}
          albumID={id}
          userID={4}
        />
      </div>
      <AlbumLikesBy users={users} />
    </ul>
  );
};

export default RouteAlbumDetails;
