import {useEffect} from "react";
import "./register.css"

export const Register = () => {
	useEffect(() => {
		const lowerCaseLetters = /[a-z]/g;
		const upperCaseLetters = /[A-Z]/g;
		const numeric = /[0-9]/g;

		function ConvertToWidthText(current_width) {
			/* Convert width number to text so that that html 
			 * dom can read the value in the right format  */
			let width_text = current_width.toString();
			let new_width_value = width_text.concat("%");
			
			return new_width_value
		}

		class ProgressBar {
			/* Bootstrap Progress Bar */
			p_bar_width; // the width of the progress bar
			password_statuses; // password statuses for each condition

			constructor() {
				this.p_bar_width = 0;
				this.password_statuses = {
					"lower_case": false,
					"upper_case": false,
					"number": false,
					"good_length": false,
				};

			}

			increase_bar(name, value) {
				// increase the progress if the condition is met
				if (this.password_statuses[name] === false) {
					this.p_bar_width += value;
					document.getElementById("p-bar").style.width = ConvertToWidthText(this.p_bar_width);
					this.password_statuses[name] = true;
				}
			}
		  
			decrease_bar(name, value) {
				// decrease the progress if the condition is met
				if (this.password_statuses[name] === true) {
					this.p_bar_width -= value;
					document.getElementById("p-bar").style.width = ConvertToWidthText(this.p_bar_width);
					this.password_statuses[name] = false;
				}
			}
		}
		function ProgressChangeOnPattern(password_field, progress_bar, pattern, password_status, info_id) {
			/* Progress changes if the pattern matches */
			if (password_field.value.match(pattern)) { // lowerCaseLetters
				// Increase the progress bar if the pattern is met
				progress_bar.increase_bar(password_status, 25); // "lower_case"
				
				document.getElementById(info_id).style.color="green";
				document.getElementById(info_id).style.visibility="visible";
			}

			else {
				// Decrease the progress bar if the pattern is not met
				progress_bar.decrease_bar(password_status, 25);

				document.getElementById(info_id).style.color="red";
				document.getElementById(info_id).style.visibility="visible";
			}
		}

		function ProgressChangeOnPasswordLength(password_field, progress_bar) {
			/* chnage progress bar on the required password length which is 8 charcter */
			if (password_field.value.length >= 8) {
				// Increase the progress bar if the password length is 8 characters or more
				progress_bar.increase_bar("good_length", 25);

				document.getElementById("good_password_length").style.color="green";
				document.getElementById("good_password_length").style.visibility="visible";
			}

			else {
				// Decrease the progress bar if the password length is less than 8 characters
				progress_bar.decrease_bar("good_length", 25);

				document.getElementById("good_password_length").style.color="red";
				document.getElementById("good_password_length").style.visibility="visible";
			}
		}

		function ProgressForPasswordOnKeyPress(password_field, progress_bar) {
			/* Change progress bar for password on key up */
			// contains lowercase letter
			ProgressChangeOnPattern(password_field, progress_bar, lowerCaseLetters, 
				"lower_case", "has_lowercase");
			
			// contains uppercase letter
			ProgressChangeOnPattern(password_field, progress_bar, upperCaseLetters, 
				"upper_case", "has_uppercase");
			
			// contains number
			ProgressChangeOnPattern(password_field, progress_bar, numeric,
				"number", "has_number");

			// password length is 8 or more characters
			ProgressChangeOnPasswordLength(password_field, progress_bar);
		}

		function PasswordValidator() {
			/* Password Validator */
			let password_field = document.getElementById("password");
			let progress_bar = new ProgressBar();

			password_field.onkeyup = function() {
				ProgressForPasswordOnKeyPress(password_field, progress_bar);
			}
		}

		// Start Password Validator
		PasswordValidator();
}, []);

	return (
		<div>
			<h2> Register </h2>
			<div className="row">
				<div className="col-7">
					<form method="post">
						<div className="form-group">
							<label htmlFor="exampleInputEmail1">Email address</label>
							<input type="email" className="form-control" id="email" name="email" aria-describedby="emailHelp" 
					placeholder="Enter email" required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputEmail1">Username</label>
							<input className="form-control" id="username" name="username" placeholder="Enter username" required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Password</label>
							<input type="password" className="form-control" id="password" name="password" 
					placeholder="Enter password" required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Confirm</label>
							<input type="password" className="form-control" id="confirm" name="confirm" 
					placeholder="Enter password again" required />
						</div>
						<button type="submit" className="btn btn-primary">Submit</button>
					</form>
				</div>
				<div id="message" className="col-5">
					<div className="progress">
						<div id="p-bar" className="progress-bar bg-success" role="progressbar" 
							style={{ width: 0 + '%' }} aria-valuenow="25" aria-valuemin="0" aria-valuemax="100">
						</div>
					</div>
					<p id="has_lowercase"> contains a lowercase letter </p>
					<p id="has_uppercase"> contains an uppercase letter </p>
					<p id="has_number"> contains a number </p>
					<p id="good_password_length"> minimum of 8 characters </p>
				</div>
			</div>
		</div>
	);
}
