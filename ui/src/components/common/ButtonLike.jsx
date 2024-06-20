import React from "react";

const LikeButton = ({ likedByUser, onClick }) => {
  const buttonText = likedByUser ? "You Liked This" : "Like";

  return (
    <button
      className={`bg-indigo-500 text-white font-semibold py-2 px-4 rounded focus:outline-none ${
        likedByUser ? "opacity-50 cursor-not-allowed" : ""
      }`}
      onClick={onClick}
      disabled={likedByUser}
    >
      {buttonText}
    </button>
  );
};

export default LikeButton;
