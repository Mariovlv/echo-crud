import axios from "axios";

const BASE_URL = "/api/v1";

const getAlbums = () => {
  return axios.get(`${BASE_URL}/albums`);
};

const getAlbumById = (id) => {
  return axios.get(`${BASE_URL}/albums/${id}`);
};

const getLikedByAlbumID = (albumid) => {
  return axios.get(`${BASE_URL}/likedby/${albumid}`);
};

const getUserByID = (userid) => {
  console.log(userid);
  return axios.get(`${BASE_URL}/users/${userid}`);
};

const insertLikeToAlbumByIDs = (userID, albumID) => {
  return axios.post(`${BASE_URL}/likedby`, {
    user_id: userID,
    album_id: albumID,
  });
};

export {
  getAlbums,
  getAlbumById,
  getLikedByAlbumID,
  getUserByID,
  insertLikeToAlbumByIDs,
};
