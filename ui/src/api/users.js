import axios from "axios";

const BASE_URL = "/api/v1";

const loginUser = (email, password) => {
  return axios.post(`${BASE_URL}/login`, { email, password });
};

const getUserByID = (userid) => {
  console.log(userid);
  return axios.get(`${BASE_URL}/users/${userid}`);
};

export { getUserByID, loginUser };
