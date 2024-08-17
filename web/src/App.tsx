import "./App.css";
import { Button } from "./components/ui/button";

function App() {
	const handleLogin = async () => {
		console.log("Login");

		const res = await fetch("http://localhost:8080/api/auth/login", {
			method: "GET",
		});

		const data = await res.json();
		console.log(data);
	};

	return (
		<main className="bg-background w-full h-screen dark font-medium text-3xl p-4">
			<h1 className="text-foreground">Fivemanage Lite</h1>

			<Button onClick={handleLogin}>Login</Button>
		</main>
	);
}

export default App;
