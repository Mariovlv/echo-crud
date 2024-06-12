import axios from "axios";

const BASE_URL = "/api/v1";

const getAlbums = () => {
  return axios.get(`${BASE_URL}/albums`);
};

const getAlbumById = (id) => {
  return axios.get(`${BASE_URL}/albums/${id}`);
};

const postAlbum = ({ album }) => {
  return axios.post(`${BASE_URL}/albums`, album);
};

export { getAlbums, getAlbumById };
