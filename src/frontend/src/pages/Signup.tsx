import { useState, type FormEvent } from "react";
import DefaultHeader from "../components/DefaultHeader";

export default function Signup() {
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  async function register(event: FormEvent) {
    event.preventDefault();
    const res = await fetch("/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username,
        email,
        password,
      })
    });

    if (res.ok) {
      alert("Success!");
    } else {
      alert("Something went wrong...");
    }
  }

  return (
    <>
      <DefaultHeader />
      <main>
        <form
          className="signup-signin-form"
          onSubmit={register}>
          <label htmlFor="username">Username</label>
          <input id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            type="text"
            autoComplete="username" required />

          <label htmlFor="email">Email</label>
          <input id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            type="email"
            autoComplete="email" required />

          <label htmlFor="password">Password</label>
          <input id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            type="password" required />

          <button
            className="col-span-full"
            type="submit">Sign Up</button>
        </form>
      </main>
    </>
  )
}
