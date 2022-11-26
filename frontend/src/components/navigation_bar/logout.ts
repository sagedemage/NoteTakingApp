/* Logout */
import axios from "axios";
import Cookies from 'universal-cookie';

export const Logout = async () => {
	/* Logout Action */
    axios.get(`http://localhost:8080/api/logout`, {
    }).then((response) => {
        if (response.data.auth === false) {
			const cookies = new Cookies();
			// set cookie
			cookies.remove("token");

			// redirect to the home page
            window.location.href = '/';
        }
        console.log(response)
    }).catch(e => {
        console.log(e)
    })
};
