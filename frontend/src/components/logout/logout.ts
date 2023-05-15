/* Logout */
import Cookies from "universal-cookie";

export const Logout = () => {
    /* Logout Action */
    const cookies = new Cookies();
    // set cookie
    cookies.remove("token");

    // redirect to the home page
    window.location.href = "/";
};
