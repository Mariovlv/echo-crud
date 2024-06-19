import React from "react";

export default function ButtonLike({
  isLikedByUser,
  likeAlbum,
  albumID,
  userID,
}) {
  return (
    <button
      onClick={() => likeAlbum()}
      className={`rounded-lg p-2 text-sm  ${
        isLikedByUser
          ? "bg-blue-500 text-white"
          : "bg-white text-blue-600 border-2 border-blue-600"
      }`}
    >
      {isLikedByUser ? "Liked album" : "Like album"}
    </button>
  );
}
