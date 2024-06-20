import axios from "axios";

const BASE_URL = "/api/v1";

const getAlbums = () => {
  return axios.get(`${BASE_URL}/albums`);
};

const getAlbumById = (id) => {
  return axios.get(`${BASE_URL}/albums/${id}`);
};

const postLike = (userID, albumID) => {
  return axios.post(`${BASE_URL}/likedby`, {
    user_id: userID,
    album_id: albumID,
  });
};

export { getAlbums, getAlbumById, postLike };
