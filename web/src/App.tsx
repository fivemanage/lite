import { useState } from "react";
import "./App.css";
import { Button } from "./components/ui/button";

function App() {
	const [email, setEmail] = useState<string>("");
	const [password, setPassword] = useState<string>("");

	const handleRegister = async () => {
		const res = await fetch("/api/auth/register", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ email, password }),
		});

		const data = await res.json();
		console.log(data);
	};

	return (
		<main className="bg-background w-full h-screen dark font-medium text-3xl p-4">
			<h1 className="text-foreground">Fivemanage Lite</h1>

			<input
				value={email}
				onChange={(e) => setEmail(e.currentTarget.value)}
				placeholder="Email"
			/>
			<input
				placeholder="Password"
				value={password}
				onChange={(e) => setPassword(e.currentTarget.value)}
			/>

			<Button onClick={handleRegister}>Login</Button>
		</main>
	);
}

export default App;
