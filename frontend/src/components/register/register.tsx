/* Registration Page */

import { useEffect } from "react";
import { PasswordValidator } from "./password-validator"
import axios from "axios";
import { useState, ChangeEventHandler, FormEventHandler } from "react";

import "./register.css";

export default function Register() {
	useEffect(() => {
		// Start Password Validator
		PasswordValidator();
	}, []);

	const [email, setEmail] = useState('');
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [confirm_pwd, setConfirm] = useState('');
	const [error_status, setErrorStatus] = useState(false);
	const [pwd_error_status, setPwdErrorStatus] = useState(false);
	const [msg_error, setMsgError] = useState('');

	const handleEmailChange: ChangeEventHandler<HTMLInputElement> = event => {
		setEmail(event.target.value);
	};
	const handleUsernameChange: ChangeEventHandler<HTMLInputElement> = event => {
		setUsername(event.target.value);
	};
	const handlePasswordChange: ChangeEventHandler<HTMLInputElement> = event => {
		setPassword(event.target.value);
	};
	const handleConfirmChange: ChangeEventHandler<HTMLInputElement> = event => {
		setConfirm(event.target.value);
	};

	const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
		/* Registration Submission */
		e.preventDefault();

		// reset error messages after submission
		setErrorStatus(false);
		setPwdErrorStatus(false);

		if (password.length < 8) {
			setPwdErrorStatus(true);
			setMsgError("password should be 8 characters or more");
		}
		else if (password !== confirm_pwd) {
			setErrorStatus(true);
			setMsgError("passwords do not match");
		}
		else {
			axios.post(`http://localhost:8080/api/register`, {
				email: email,
				username: username,
				password: password,
				confirm_pwd: confirm_pwd,
			}).then((response) => {
				if (response.data.registered === true) {
					// create new url of the login page
					const url = new URL("/login", "http://localhost:3000");
					// add url parameter
					url.searchParams.append("msg_success", response.data.msg_success);
					// redirect to the login page
					window.location.href = String(url);
				}
				else {
					setErrorStatus(true);
					setMsgError(response.data.msg_error);
				}
			}).catch(e => {
				console.log(e)
			})
		}


	};

	return (
		<div>
			{error_status === true &&
				<div className="alert alert-danger" role="alert">
					{msg_error}
				</div>
			}
			<h2> Register </h2>
			<div className="row">
				<div id="register_form" className="col col-lg-8">
					<form method="post" onSubmit={handleSubmit}>
						<div className="form-group">
							<label htmlFor="exampleInputEmail1">Email address</label>
							<input type="email" className="form-control" id="email"
								name="email" placeholder="Enter email"
								aria-describedby="emailHelp"
								value={email}
								onChange={handleEmailChange}
								required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputUsername1">Username</label>
							<input className="form-control" id="username"
								name="username" placeholder="Enter username"
								value={username}
								onChange={handleUsernameChange}
								required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Password</label>
							<input type="password" className="form-control" id="password"
								name="password" placeholder="Enter password"
								value={password}
								onChange={handlePasswordChange}
								required />
							{pwd_error_status === true &&
								<div className="text-danger">
									<small>{msg_error}</small>
								</div>
							}
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Confirm</label>
							<input type="password" className="form-control" id="confirm"
								name="confirm" placeholder="Enter password again"
								value={confirm_pwd}
								onChange={handleConfirmChange}
								required />
						</div>
						<button type="submit" className="btn btn-primary">Submit</button>
					</form>
				</div>
				<div id="message" className="col-md-auto col-lg-4">
					<div className="progress">
						<div id="p-bar" className="progress-bar bg-success" role="progressbar"></div>
					</div>
					<p id="has_lowercase" className="validator_info"> contains a lowercase letter </p>
					<p id="has_uppercase" className="validator_info"> contains an uppercase letter </p>
					<p id="has_number" className="validator_info"> contains a number </p>
				</div>
			</div>
		</div>
	);
}
