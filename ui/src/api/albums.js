import axios from "axios";

const BASE_URL = "/api/v1";

const getAlbums = () => {
  return axios.get(`${BASE_URL}/albums`);
};

const getAlbumById = (id) => {
  return axios.get(`${BASE_URL}/albums/${id}`);
};

const postLike = (userID, albumID) => {
  console.log("fetching with", userID, albumID);
  return axios.post(`${BASE_URL}/likedby`, {
    user_id: String(userID),
    album_id: albumID,
  });
};

export { getAlbums, getAlbumById, postLike };
