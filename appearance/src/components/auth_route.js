import { Navigate, Outlet, useLocation } from 'react-router-dom';

import { useAuth } from "./auth";

export const AuthRoute = () => {
	
	const isAuth = useAuth();
	const location = useLocation();

	if (isAuth !== "true") {
		return (
			<div id="home-content">
				<h2> Unauthorized User </h2>
				<p> The User is Unauthorized, <a href="/">click here</a> to go the the home page </p>
			</div>
		);
	}
	
	return isAuth ? <Outlet /> : <Navigate to="/" replace state={{ from: location }} />;
}
