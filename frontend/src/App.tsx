import { useState } from "react";
import { Button } from "./components/ui/button";
import { Card } from "./components/ui/card";

export function App() {
  const [response, setResponse] = useState("No request sent.");

  function pingBackend() {
    fetch(`${import.meta.env.VITE_API}/`).then(async (res) => {
      try {
        const msg = await res.json();
        setResponse(msg.message + " - " + res.status);
      } catch {
        console.error("Something went wrong");
      }
    });
  }
  return (
    <div className="h-screen w-screen flex items-center justify-center flex-col gap-4">
      <h1 className="text-3xl">Go Dashboard</h1>
      <Button onClick={pingBackend}>Ping Backend</Button>
      <Card className="p-4">{response}</Card>
    </div>
  );
}

export default App;
