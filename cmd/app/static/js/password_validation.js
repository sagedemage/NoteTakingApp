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
		if (this.password_statuses[name] == false) {
			this.p_bar_width += value;
			document.getElementById("p-bar").style.width = ModifyWidthText(this.p_bar_width);
			this.password_statuses[name] = true;
		}
  	}
  
	decrease_bar(name, value) {
		if (this.password_statuses[name] == true) {
			this.p_bar_width -= value;
			document.getElementById("p-bar").style.width = ModifyWidthText(this.p_bar_width);
			this.password_statuses[name] = false;
		}
  	}
}

let passwordField = document.getElementById("password");

let progress_bar = new ProgressBar();

passwordField.onkeyup = function() {
	/* Validate the password on key up */

	// contains lowercase letter
    if (passwordField.value.match(lowerCaseLetters)) {
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
    if (passwordField.value.match(upperCaseLetters)) {
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
    if (passwordField.value.match(numbers)) {
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
    if (passwordField.value.length >= 8) {
		// Increase the progress bar if there is a number
		//password_length = progress_bar.increase_bar(25, password_length);
		progress_bar.increase_bar("good_length", 25);

		document.getElementById("good_password_length").style.color="green";
		document.getElementById("good_password_length").style.visibility="visible";
    }

	else {
    	// Decrease the progress bar if there is no number
		//password_length = progress_bar.decrease_bar(25, password_length);
		progress_bar.decrease_bar("good_length", 25);

		document.getElementById("good_password_length").style.color="red";
		document.getElementById("good_password_length").style.visibility="visible";
    }

}

