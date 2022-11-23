/* Delete Note Page */

import { FormEventHandler } from "react";

import axios from "axios";

export const DeleteNote = () => {
	let url = new URL(window.location.href);
	const note_id = parseInt(url.searchParams.get("note_id"));

	if (note_id === null) {
		window.location.href = '/';
	}

	const GoBack = () => {
		window.location.href='/dashboard';
	}
	const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
		/* Delete Note Submission */
		e.preventDefault();
		axios.post(`http://localhost:8080/api/delete-note`, {
			note_id: note_id,
		}).then((response) => {
			// redirect to the dashboard
            window.location.href = '/dashboard';
            console.log(response);
		}).catch(e => {
            console.log(e);
        })
	};

	return (		
		<div>
			<h2> Delete Note </h2>
			<p> You sure you want to delete the note? </p>
			<form method="post" onSubmit={handleSubmit}>
				<button type="submit" className="btn btn-danger">
					Delete
				</button>
				<button type="button" className="btn btn-secondary" 
					onClick={ GoBack }>
					Back
				</button>
			</form>
		</div>
	);
}
