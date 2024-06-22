import React, { useState, useEffect } from "react";
import { postLike } from "../api/albums";

export default function AlbumLikesBy({ users, currentUser, albumId }) {
  const [likedUsers, setLikedUsers] = useState(users);
  const [isLiked, setIsLiked] = useState(false);

  useEffect(() => {
    setIsLiked(
      likedUsers.some((user) => user.username === currentUser.username)
    );
  }, [likedUsers, currentUser]);

  const handleLike = async () => {
    try {
      await postLike(currentUser.id, albumId);
      if (!isLiked) {
        setLikedUsers([...likedUsers, currentUser]);
      } else {
        setLikedUsers(
          likedUsers.filter((user) => user.username !== currentUser.username)
        );
      }
      setIsLiked(!isLiked);
    } catch (error) {
      console.error("Error liking/unliking album:", error);
    }
  };

  return (
    <div className="p-4 m-3 bg-white rounded-lg shadow-lg">
      <h3 className="mb-3 text-xl font-semibold text-gray-900">Liked by</h3>
      <ul className="space-y-2">
        {likedUsers.map((user) => (
          <li
            key={user.ID}
            className={`px-3 py-2 rounded-lg ${
              user.username === currentUser.username
                ? "bg-blue-200 text-blue-800"
                : "bg-gray-100 text-gray-700"
            }`}
          >
            {user.username === currentUser.username ? "You" : user.username}
          </li>
        ))}
      </ul>
      <button
        onClick={handleLike}
        className={`mt-4 px-4 py-2 rounded-lg font-semibold ${
          isLiked
            ? "bg-red-500 text-white hover:bg-red-600"
            : "bg-blue-500 text-white hover:bg-blue-600"
        }`}
      >
        {isLiked ? "Unlike" : "Like"}
      </button>
    </div>
  );
}
