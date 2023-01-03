/* Main App */

import './app.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import axios from "axios";
import { lazy, Suspense } from 'react';
import { Route, Routes } from "react-router-dom";

//import { PageNotFound } from "./components/page_not_found/page_not_found";

import { AuthRoute } from "./components/auth_route/auth_route";
import { MyNavBar } from "./components/navigation_bar/navigation_bar";

/* Page Components */
const Home = lazy(() => import('./components/simple_pages/home'));
const About = lazy(() => import('./components/simple_pages/about'));
const Login = lazy(() => import('./components/login/login'));
const Register = lazy(() => import('./components/register/register'));
const Notes = lazy(()=> import('./components/view_notes/view_notes'));
const AddNoteForm = lazy(()=> import('./components/note_actions/add_new_note'));
const DeleteNote = lazy(()=> import('./components/note_actions/delete_note'));
const EditNoteForm = lazy(() => import('./components/note_actions/edit_note'));
const PageNotFound = lazy(() => import('./components/page_not_found/page_not_found'));


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
							<Route path="add-new-note" element={<AddNoteForm />}> </Route>
							<Route path="delete-note" element={<DeleteNote />}> </Route>
							<Route path="edit-note" element={<EditNoteForm />}> </Route>
						</Route>
					</Routes>
				</Suspense>
			</div>
			<footer id="copyright">
				<p> &copy; {(new Date().getFullYear())} Salmaan Saeed </p>
			</footer>
		</div>
	);
}

export default App;

