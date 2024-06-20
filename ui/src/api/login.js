import axios from "axios";
import { useContext } from "react";
import UserContext from "../../context/UserContext";

const BASE_URL = "/api/v1";

const login = async ({ email, password }) => {
  const { loginSet } = useContext(UserContext);

  try {
    const response = await axios.post(`${BASE_URL}/login`, { email, password });
    const { userId, userEmail } = response.data; // Assuming the API response contains userId and userEmail
    await loginSet(userId, userEmail); // Set user data in context after successful login
    return response;
  } catch (error) {
    throw error;
  }
};

const signup = (crendentials) => {
  return axios.post(`${BASE_URL}/signup`, crendentials);
};

export { login, signup };
