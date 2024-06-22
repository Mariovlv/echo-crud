import axios from "axios";

const BASE_URL = "/api/v1";

const loginUser = (email, password) => {
  return axios
    .post(`${BASE_URL}/login`, { email, password })
    .then((response) => {
      if (response.data) {
        localStorage.setItem("user", JSON.stringify(response.data));
        console.log("login succes", response.data);
        return response.data;
      }
    });
};

const logoutUser = () => {
  localStorage.removeItem("user");
};

const getUserByID = (userid) => {
  console.log(userid);
  return axios.get(`${BASE_URL}/users/${userid}`);
};

export { getUserByID, loginUser, logoutUser };
