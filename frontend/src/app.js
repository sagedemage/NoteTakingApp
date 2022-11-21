/* Main App */

import './app.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import axios from "axios";
import { React } from "react";
import { Route, Routes } from "react-router-dom";
import { Home } from "./components/simple_pages/home";
import { About } from "./components/simple_pages/about";
import { Login } from "./components/login/login";
import { Register } from "./components/register/register";
import { Notes } from "./components/view_notes/view_notes";
import { AuthRoute } from "./components/auth_route/auth_route";
import { AddNoteForm } from "./components/note_actions/add_new_note";
import { DeleteNote } from "./components/note_actions/delete_note";
import { EditNoteForm } from "./components/note_actions/edit_note";
import { PageNotFound } from "./components/page_not_found/page_not_found";
import { MyNavBar } from "./components/navigation_bar/navigation_bar";

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

