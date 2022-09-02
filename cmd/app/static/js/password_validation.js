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

  	constructor() {
		this.p_bar_width = 0;
 	}

	increase_bar(value, cond) {
		if (cond == false) {
			this.p_bar_width += value;
			document.getElementById("p-bar").style.width = ModifyWidthText(this.p_bar_width);
			cond = true;
		}
		return cond;
  	}
  
	decrease_bar(value, cond) {
		if (cond == true) {
			this.p_bar_width -= value;
			document.getElementById("p-bar").style.width = ModifyWidthText(this.p_bar_width);
			cond = false;
		}
		return cond;
  	}
}

let passwordField = document.getElementById("password");

let progress_bar = new ProgressBar();

let lowerCase = false;
let upperCase = false;
let num = false;

passwordField.onkeyup = function() {
	/* Validate the password on key up */

	// contains lowercase letter
    if (passwordField.value.match(lowerCaseLetters)) {
		// Increase the progress bar if there is a lowercase letter
		lowerCase = progress_bar.increase_bar(25, lowerCase);
    }

	else {
    	// Decrease the progress bar if there is no lowercase letter
		lowerCase = progress_bar.decrease_bar(25, lowerCase);
    }

	// contains uppercase letter
    if (passwordField.value.match(upperCaseLetters)) {
		// Increase the progress bar if there is a uppercase letter
		upperCase = progress_bar.increase_bar(25, upperCase);
    }

	else {
    	// Decrease the progress bar if there is no uppercase letter
		upperCase = progress_bar.decrease_bar(25, upperCase);
    }

	// contains number
    if (passwordField.value.match(numbers)) {
		// Increase the progress bar if there is a number
		num = progress_bar.increase_bar(25, num);
    }

	else {
    	// Decrease the progress bar if there is no number
		num = progress_bar.decrease_bar(25, num);
    }

}

