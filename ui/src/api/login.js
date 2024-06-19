import axios from "axios";
import { useCookies } from "react-cookie";

const BASE_URL = "/api/v1";

const login = (credentials) => {
  return axios.post(`${BASE_URL}/login`, credentials);
};

export { login };
