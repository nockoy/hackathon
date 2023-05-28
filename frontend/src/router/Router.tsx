import { BrowserRouter, Route, Routes } from "react-router-dom";
import SignUp from "../pages/SignUp";
import Login from "../pages/Login";
import Logout from "../pages/Logout";
import Mypage from "../pages/Mypage";
import NotFound from "../pages/NotFound";

export const Router = () => {
    return (
      <div className="container">
        <BrowserRouter>
          <Routes>
            <Route path={`/`} element={<Mypage />} />
            <Route path={`/signup/`} element={<SignUp />} />
            <Route path={`/login/`} element={<Login />} />
            <Route path={`/logout/`} element={<Logout />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </BrowserRouter>
      </div>
    );
  };