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
	p_bar_width;
	password_statuses;

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
		if (this.password_statuses[name] === false) {
			this.p_bar_width += value;
			document.getElementById("p-bar").style.width = ConvertToWidthText(this.p_bar_width);
			this.password_statuses[name] = true;
		}
  	}
  
	decrease_bar(name, value) {
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
