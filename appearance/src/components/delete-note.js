import axios from "axios";
import {useState} from "react";
import Cookies from 'universal-cookie';

export const DeleteNote = () => {

	let url = new URL(window.location.href);
	const note_id = parseInt(url.searchParams.get("note_id"));
	console.log(note_id);

	if (note_id === null) {
		window.location.href = '/';
	}

	const GoBack = () => {
		window.location.href='/dashboard';
	}

	const handleSubmit = async (e) => {
		e.preventDefault();
		//const cookies = new Cookies();
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
				<button type="submit" className="btn btn-danger">Delete</button>
				<button type="button" className="btn btn-secondary" onClick={ GoBack }>Back</button>
			</form>
		</div>
	);
}
