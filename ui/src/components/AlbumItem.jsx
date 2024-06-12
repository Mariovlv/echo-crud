import React from "react";
import { Link } from "react-router-dom";

const AlbumItem = ({ album }) => {
  return (
    <Link to={`/albums/${album.id}`}>
      <li
        className="border-2 hover:bg-stone-100 border-stone-400 p-3 m-2 rounded-md"
        key={album.id}
      >
        <h2 className="font-bold text-lg pb-2">{album.title}</h2>
        <p>{album.artist}</p>
        <p>{album.date}</p>
      </li>
    </Link>
  );
};

export default AlbumItem;
