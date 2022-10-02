import "./home.css"

import { useEffect } from "react";
import axios from "axios";
import {useState} from "react";

/*export const Logout1 = (props) => {

	

    return (
		<div>
			<h2> Logout </h2>
            <p> User logged out </p>
		</div>
	);
}*/

export const Logout = async (e) => {
    axios.get(`http://localhost:8080/api/logout`, {
    }).then((response) => {
        if (response.data.auth === false) {
            localStorage.removeItem("token")
            localStorage.setItem("auth", false)
            window.location.href = '/';
        }
        //console.log(response.data);
        console.log(response)
    }).catch(e => {
        console.log(e)
    })
};
