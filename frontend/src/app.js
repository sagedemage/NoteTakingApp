/* Main App */

import './app.css';

import { React } from "react";

import { Route, Routes } from "react-router-dom";

import axios from "axios";

import { Home } from "./components/home";
import { About } from "./components/about";
import { Login } from "./components/login";
import { Register } from "./components/register";
import { Notes } from "./components/view-notes";
import { AuthRoute } from "./components/auth-route";
import { AddNoteForm } from "./components/add-new-note";
import { DeleteNote } from "./components/delete-note";
import { EditNoteForm } from "./components/edit-note";
import { PageNotFound } from "./components/page-not-found";
import { MyNavBar } from "./components/navigation-bar";

import 'bootstrap/dist/css/bootstrap.min.css';

function App() {

    axios.defaults.withCredentials = true;

	return (
		<div id="body">
			<MyNavBar />
			<div className="gaps" id="content">
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
			</div>
			<footer id="copyright">
				<p> &copy; {(new Date().getFullYear())} Salmaan Saeed </p>
			</footer>
		</div>
	);
}

export default App;

