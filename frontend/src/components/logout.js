/* Logout Action */

import axios from "axios";
import Cookies from 'universal-cookie';

export const Logout = async () => {
    axios.get(`http://localhost:8080/api/logout`, {
    }).then((response) => {
        if (response.data.auth === false) {
			const cookies = new Cookies();
			// set cookie
			//cookies.set("auth", false);
			cookies.remove("token");

			// redirect to the home page
            window.location.href = '/';
        }
        console.log(response)
    }).catch(e => {
        console.log(e)
    })
};
