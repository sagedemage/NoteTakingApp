/* Authentication Status */

import Cookies from 'universal-cookie';



export const useAuth = () => {
	/* Get the User Authentication Status */
	const cookies = new Cookies();
	let auth = cookies.get("auth");
	if (auth === undefined) {
		auth = "false";
	}
	return auth;
}


export const DecodeToken = () => {
	/* Get the User Authentication Status */
	const cookies = new Cookies();
	let token = cookies.get("token");
	if (token === undefined) {
		auth = "false";
	}
	return auth;
}

