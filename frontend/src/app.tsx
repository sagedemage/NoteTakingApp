/* Main App */

import "./app.css";
import axios from "axios";
import { lazy, Suspense } from "react";
import { Route, Routes } from "react-router-dom";

/* Auth Components */
const AuthRoute = lazy(() => import("./components/auth/auth_route/auth_route"));
const Login = lazy(() => import("./components/auth/login/login"));
const Register = lazy(() => import("./components/auth/register/register"));

/* UI Components */
const MyNavBar = lazy(() => import("./components/ui/navigation_bar/navigation_bar"));
const Footer = lazy(() => import("components/ui/footer/footer"));

/* Page Components */
const Home = lazy(() => import("./components/pages/home"));
const About = lazy(() => import("./components/pages/about"));
const Notes = lazy(()=> import("./components/pages/view_notes/view_notes"));
const PageNotFound = lazy(() => import("./components/pages/page_not_found"));


function App() {
	axios.defaults.withCredentials = true;
	return (
		<div id="body">
			<MyNavBar />
			<div className="gaps" id="content">
				<Suspense fallback={<h1>Still Loading...</h1>}>
					<Routes>
						<Route path="/" element={<Home />}> </Route>
						<Route path="about" element={<About />}> </Route>
						<Route path="login" element={<Login />}> </Route>
						<Route path="register" element={<Register />}> </Route>
						<Route path="*" element={<PageNotFound />}> </Route>

						<Route element={<AuthRoute />}>
							<Route path="dashboard" element={<Notes />}> </Route>
						</Route>
					</Routes>
				</Suspense>
			</div>
			<Footer />
		</div>
	);
}

export default App;

