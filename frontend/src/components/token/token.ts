/* Token */

import Cookies from 'universal-cookie';

export function getToken() {
	/* Get token */
	const cookies = new Cookies();
	const token = cookies.get("token");
	return token
}

