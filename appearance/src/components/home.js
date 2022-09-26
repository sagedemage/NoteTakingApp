import "./home.css";

import axios from 'axios';
import {useEffect} from "react";
import {useState} from "react";


export const Home = () => {
	const [data1, setData1] = useState([]);
	const [data2, setData2] = useState([]);
	const [data3, setData3] = useState([]);

	useEffect(() => {
		axios.get(`http://localhost:8080/api/test1`)
			.then(response => {
				console.log("Hi", response.data);
				setData1(response.data);
			})
			.catch(err => console.log(err))
	}, [])

	useEffect(() => {
		axios.get(`http://localhost:8080/api/test2`)
			.then(response => {
				console.log("Hi", response.data);
				setData2(response.data);
			})
			.catch(err => console.log(err))
	}, [])

	useEffect(() => {
		axios.get(`http://localhost:8080/api/test3`)
			.then(response => {
				console.log("Hi", response.data);
				setData3(response.data);
			})
			.catch(err => console.log(err))
	}, [])

	return (
		<div id="home-content">
			<h2> Home </h2>
			<p> Welcome to the Notebook app. You can go to the notes tab to view your notes. </p>

			<p> Counting </p>
			<ul>
				<li>{data1.msg}</li>
				<li>{data2.msg}</li>
				<li>{data3.msg}</li>
			</ul>
		</div>
		
	);
}
