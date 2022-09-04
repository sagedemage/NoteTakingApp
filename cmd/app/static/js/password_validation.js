const lowerCaseLetters = /[a-z]/g;
const upperCaseLetters = /[A-Z]/g;
const numbers = /[0-9]/g;

function ModifyWidthText(current_width) {
	let width_text = current_width.toString();
	let new_width_value = width_text.concat("%");
	
	return new_width_value
}

class ProgressBar {
	p_bar_width;
	validate_statuses;

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
			document.getElementById("p-bar").style.width = ModifyWidthText(this.p_bar_width);
			this.password_statuses[name] = true;
		}
  	}
  
	decrease_bar(name, value) {
		if (this.password_statuses[name] === true) {
			this.p_bar_width -= value;
			document.getElementById("p-bar").style.width = ModifyWidthText(this.p_bar_width);
			this.password_statuses[name] = false;
		}
  	}
}

function ProgressForPasswordOnKeyPress(password_field, progress_bar) {
	/* Change progress bar for password on key up */

	// contains lowercase letter
	if (password_field.value.match(lowerCaseLetters)) {
		// Increase the progress bar if there is a lowercase letter
		progress_bar.increase_bar("lower_case", 25);
		
		document.getElementById("has_lowercase").style.color="green";
		document.getElementById("has_lowercase").style.visibility="visible";
	}

	else {
		// Decrease the progress bar if there is no lowercase letter
		progress_bar.decrease_bar("lower_case", 25);

		document.getElementById("has_lowercase").style.color="red";
		document.getElementById("has_lowercase").style.visibility="visible";
	}

	// contains uppercase letter
	if (password_field.value.match(upperCaseLetters)) {
		// Increase the progress bar if there is a uppercase letter
		progress_bar.increase_bar("upper_case", 25);

		document.getElementById("has_uppercase").style.color="green";
		document.getElementById("has_uppercase").style.visibility="visible";
	}

	else {
		// Decrease the progress bar if there is no uppercase letter
		progress_bar.decrease_bar("upper_case", 25);

		document.getElementById("has_uppercase").style.color="red";
		document.getElementById("has_uppercase").style.visibility="visible";
	}

	// contains number
	if (password_field.value.match(numbers)) {
		// Increase the progress bar if there is a number
		progress_bar.increase_bar("number", 25);
		
		document.getElementById("has_number").style.color="green";
		document.getElementById("has_number").style.visibility="visible";
	}

	else {
		// Decrease the progress bar if there is no number
		progress_bar.decrease_bar("number", 25);

		document.getElementById("has_number").style.color="red";
		document.getElementById("has_number").style.visibility="visible";
	}

	// password lenth is 6 more characters
	if (password_field.value.length >= 8) {
		// Increase the progress bar if there is a number
		progress_bar.increase_bar("good_length", 25);

		document.getElementById("good_password_length").style.color="green";
		document.getElementById("good_password_length").style.visibility="visible";
	}

	else {
		// Decrease the progress bar if there is no number
		progress_bar.decrease_bar("good_length", 25);

		document.getElementById("good_password_length").style.color="red";
		document.getElementById("good_password_length").style.visibility="visible";
	}
}

function PasswordValidator() {
	let password_field = document.getElementById("password");
	let progress_bar = new ProgressBar();

	password_field.onkeyup = function() {
		ProgressForPasswordOnKeyPress(password_field, progress_bar);
	}
}

PasswordValidator();
