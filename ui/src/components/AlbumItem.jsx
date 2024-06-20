import React from "react";
import { Link } from "react-router-dom";

const AlbumItem = ({ album }) => {
  return (
    <Link to={`/albums/${album.id}`}>
      <li
        className="flex justify-between items-center border-2 hover:bg-stone-100 border-stone-400 p-3 m-2 rounded-md"
        key={album.id}
      >
        <div className="flex flex-col">
          <h2 className="font-bold text-lg">{album.title}</h2>
          <p className="text-sm text-gray-600">{album.artist}</p>
        </div>
        <p className="text-sm text-gray-500">{album.id}</p>
      </li>
    </Link>
  );
};

export default AlbumItem;
