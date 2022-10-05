import './App.css';

import { React } from "react";

import { Route, Routes } from "react-router-dom";

import { Nav, Navbar, Container } from 'react-bootstrap';

import NavDropdown from 'react-bootstrap/NavDropdown';

import axios from "axios";

import { Home } from "./components/home";
import { About } from "./components/about";
import { Login } from "./components/login";
import { Register } from "./components/register";
import { Notes } from "./components/view-notes";
import { Logout } from "./components/logout";

import 'bootstrap/dist/css/bootstrap.min.css';

import Cookies from 'universal-cookie';

function App() {

    axios.defaults.withCredentials = true;

	const cookies = new Cookies();

	let auth = cookies.get("auth");

    if (auth === undefined) {
        auth = "false";
    }

	return (
		<div>
			<div>
				<Navbar collapseOnSelect expand="lg" bg="myBlue" variant="dark" fixed="top">
					<Container>
						<Navbar.Brand>
							<span className="indent"> Notebook </span>
						</Navbar.Brand>
				
						<Navbar.Toggle aria-controls="responsive-navbar-nav"  />
						<Navbar.Collapse id="responsive-navbar-nav">
							<Nav className="me-auto">
								<Nav.Link href="/"> Home </Nav.Link>
								<Nav.Link href="/about"> About </Nav.Link>
							</Nav>
							{ auth === "false" &&
							<Nav className="ms-auto">
								<Nav.Link href="/login"> Login </Nav.Link>
								<Nav.Link href="/register"> Register </Nav.Link>
							</Nav>
							}
							{ auth === "true" &&
							<Nav className="ms-auto">
								<NavDropdown
								  id="nav-dropdown"
								  title="Account"
								  menuVariant="dark"
								  variant="dark"
								>
								  	<NavDropdown.Item href="/dashboard">Notes</NavDropdown.Item>
									<NavDropdown.Item onClick={() => Logout() }>Logout</NavDropdown.Item>
								</NavDropdown>
							</Nav>
							}
						</Navbar.Collapse>
					</Container>
				</Navbar>
			</div>
			<br />
			<div className="gaps" id="content">
				<Routes>
					<Route path='/' element={<Home />}> </Route>
					<Route path='about' element={<About />}> </Route>
					<Route path='login' element={<Login />}> </Route>
					<Route path='register' element={<Register />}> </Route>
					<Route path='dashboard' element={<Notes />}> </Route>
				</Routes>
			</div>
			<footer className="gaps" id="bottom">
				<div id="copyright">
					<p id="copyright"> &copy; {(new Date().getFullYear())} Salmaan Saeed </p>
				</div>
			</footer>
		</div>
	);
}

export default App;
