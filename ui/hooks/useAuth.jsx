import axios from "axios";
import { useContext } from "react";
import UserContext from "../context/UserContext";

const BASE_URL = "/api/v1";

const useAuth = () => {
  const { login } = useContext(UserContext);

  const loginUser = async ({ email, password }) => {
    try {
      const response = await axios.post(`${BASE_URL}/login`, {
        email,
        password,
      });
      const { userId, userEmail } = response.data; // Assuming the API response contains userId and userEmail
      login(userId, userEmail); // Set user data in context after successful login
      return response;
    } catch (error) {
      throw error;
    }
  };

  const signupUser = (credentials) => {
    return axios.post(`${BASE_URL}/signup`, credentials);
  };

  return { loginUser, signupUser };
};

export default useAuth;
