import { useState } from "react";
import { Button } from "./components/ui/button";
import { Card } from "./components/ui/card";
import { Input } from "./components/ui/input";

export function App() {
  const [id, setId] = useState("0");
  const [response, setResponse] = useState("No request sent.");

  function pingBackend() {
    if (id) {
      getPayment(id);
      return;
    }

    fetch(`${import.meta.env.VITE_API}/`).then(async (res) => {
      try {
        const msg = await res.json();
        setResponse(msg.message + " - " + res.status);
      } catch {
        console.error("Something went wrong");
      }
    });
  }

  function getPayment(id: string) {
    fetch(`${import.meta.env.VITE_API}/payment/${id}`).then(async (res) => {
      try {
        const payment = await res.json();
        setResponse(JSON.stringify(payment));
      } catch {
        console.error("Something went wrong");
      }
    });
  }

  return (
    <div className="h-screen w-screen flex items-center justify-center flex-col gap-4">
      <h1 className="text-3xl">Go Dashboard</h1>
      <div className="flex items-center justify-center flex-col gap-4 w-80">
        <Input
          placeholder="Optional: Enter a payment ID (Eg. 1, 2)"
          value={id}
          onChange={(e) => setId(e.currentTarget.value)}
          type="number"
        />
        <Button onClick={pingBackend}>Ping Backend</Button>
        <Card className="p-4">{response}</Card>
      </div>
    </div>
  );
}

export default App;
