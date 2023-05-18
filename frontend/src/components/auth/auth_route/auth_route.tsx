/* Authorization Routes */

import { Navigate, Outlet, useLocation } from "react-router-dom";
import axios from "axios";
import {useEffect, useState} from "react";
import Cookies from "universal-cookie";

export default function AuthRoute() {
	const origin = useLocation();
	const [isAuth, setAuth] = useState(false);

	useEffect(() => {
		/* Fetch all the Notes for the Current User */
		const cookies = new Cookies();
		const token = cookies.get("token");
		if (token !== undefined) {
			axios.post("http://localhost:8080/api/get-decoded-token", {
				token: token,
			}).then((response) => {
				if (response.data.auth === true) {
					setAuth(true)
				}
				else {
					setAuth(false)
				}
			}).catch(e => {
				console.log(e);
			})
		}
		else {
			setAuth(false)
		}
	}, []);


	if (isAuth !== true) {
		return (
			<div id="home-content">
				<h2> Unauthorized User </h2>
				<p> The User is Unauthorized, <a href="/">click here</a> to go the the home page </p>
			</div>
		);
	}
	
	return isAuth ? <Outlet /> : <Navigate to="/" replace state={{ from: origin }} />;
}
