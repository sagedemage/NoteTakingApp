/* Authentication Status */

import Cookies from 'universal-cookie';


export const DecodeToken = () => {
	/* Get the User Authentication Status */
	const cookies = new Cookies();
	let token = cookies.get("token");
	let auth = "true";
	if (token === undefined) {
		auth = "false";
	}
	return auth;
}

