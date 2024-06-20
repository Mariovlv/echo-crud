import React, { createContext, useState } from "react";

const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [userData, setUserData] = useState({ userId: null, userEmail: "" });

  const loginSet = (userId, userEmail) => {
    setUserData({ userId, userEmail });
  };

  const logout = () => {
    setUserData({ userId: null, userEmail: "" });
  };

  return (
    <UserContext.Provider value={{ userData, loginSet, logout }}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext;
