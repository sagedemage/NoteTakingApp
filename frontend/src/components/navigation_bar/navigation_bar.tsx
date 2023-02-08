/* Navigation Bar */

import "./navigation_bar.css";
import { Nav, Navbar, Container } from 'react-bootstrap';
import NavDropdown from 'react-bootstrap/NavDropdown';
import axios from "axios";
import { useEffect, useState } from "react";
import { getToken } from "components/token/token";
import { Logout } from "./logout";

export const MyNavBar = () => {
	const [isAuth, setAuth] = useState(false);
	useEffect(() => {
		/* Fetch all the Notes for the Current User */
		const token = getToken();
		if (token !== undefined) {
			axios.post(`http://localhost:8080/api/get-decoded-token`, {
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

	return (
		<div className="mb">
			<Navbar collapseOnSelect id="navbar" expand="lg" bg="myBlue" variant="dark" fixed="top">
				<Container>
					<Navbar.Brand href="/">
						<span className="indent"> Notebook </span>
					</Navbar.Brand>
			
					<Navbar.Toggle aria-controls="responsive-navbar-nav"  />
					<Navbar.Collapse id="responsive-navbar-nav">
						<Nav className="me-auto">
							<Nav.Link href="/about"> About </Nav.Link>
						</Nav>
						{ isAuth === false &&
						<Nav className="ms-auto">
							<Nav.Link href="/login"> Login </Nav.Link>
							<Nav.Link href="/register"> Register </Nav.Link>
						</Nav>
						}
						{ isAuth === true &&
						<Nav className="ms-auto">
							<NavDropdown id="nav-dropdown" title="Account" menuVariant="dark">
								<NavDropdown.Item href="/dashboard">Notes</NavDropdown.Item>
								<NavDropdown.Item onClick={() => Logout() }>
									Logout
								</NavDropdown.Item>
							</NavDropdown>
						</Nav>
						}
					</Navbar.Collapse>
				</Container>
			</Navbar>
		</div>
		
	);
}
