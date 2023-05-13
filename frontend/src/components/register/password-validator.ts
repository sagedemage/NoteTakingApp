/* Password Validator */

const lowerCaseLetters: RegExp = /[a-z]/g;
const upperCaseLetters: RegExp = /[A-Z]/g;
const numeric: RegExp = /[0-9]/g;

function ConvertToWidthText(current_width: number) {
	/* Convert width number to text so that that html 
	 * dom can read the value in the right format  */
	let width_text = current_width.toString();
	let new_width_value = width_text.concat("%");
	
	return new_width_value
}

class ProgressBar {
	/* Bootstrap Progress Bar */
	p_bar_width: number; // the width of the progress bar
	password_statuses: Map<string, boolean>; // password statuses for each condition

	constructor() {
		this.p_bar_width = 0;
		this.password_statuses = new Map<string, boolean>();
		this.password_statuses.set("lower_case", false);
		this.password_statuses.set("upper_case", false);
		this.password_statuses.set("number", false);
		this.password_statuses.set("good_length", false);
	}

	increase_bar(name: string, value: number) {
		if (this.password_statuses.get(name) === false) {
			this.p_bar_width += value;
			document.getElementById("p-bar")!.style.width = ConvertToWidthText(this.p_bar_width);
			this.password_statuses.set(name, true);
		}
	}

	decrease_bar(name: string, value: number) {
		if (this.password_statuses.get(name) === true) {
			this.p_bar_width -= value;
			document.getElementById("p-bar")!.style.width = ConvertToWidthText(this.p_bar_width);
			this.password_statuses.set(name, false);
		}
	}
}

function ProgressChangeOnPattern(password_field: HTMLInputElement, progress_bar: ProgressBar, pattern: RegExp, password_status: string, info_id: string) {
	/* Progress changes if the pattern matches */
	if (password_field.value.match(pattern)) {
		// Increase the progress bar if the pattern is met
		progress_bar.increase_bar(password_status, 33.33);
		// show green text
		document.getElementById(info_id)!.style.color="green";
	}

	else {
		// Decrease the progress bar if the pattern is not met
		progress_bar.decrease_bar(password_status, 33.33);
		// show dark red text
		document.getElementById(info_id)!.style.color="darkred";
	}
}

function ProgressForPasswordOnKeyPress(password_field: HTMLInputElement, progress_bar: ProgressBar) {
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
}

export function PasswordValidator() {
	/* Password Validator */
	let password_field: HTMLInputElement = (document.getElementById("password") as HTMLInputElement);
	let progress_bar = new ProgressBar();

	password_field!.onkeyup = function() {
		ProgressForPasswordOnKeyPress(password_field, progress_bar);
	};
}

