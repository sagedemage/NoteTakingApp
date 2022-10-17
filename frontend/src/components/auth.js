import Cookies from 'universal-cookie';


//export const AuthRoute = () => {

export const useAuth = () => {
	const cookies = new Cookies();
	let auth = cookies.get("auth");
	if (auth === undefined) {
		auth = "false";
	}
	return auth;
}

