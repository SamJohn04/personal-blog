import { useState, type FormEvent } from "react";
import DefaultHeader from "../components/DefaultHeader";
import { useNavigate } from "react-router";

export function CreateBlog() {
  const navigate = useNavigate();

  const [title, setTitle] = useState<string>("");
  const [content, setContent] = useState<string>("");

  const [loading, setLoading] = useState<boolean>(false);

  async function create(event: FormEvent) {
    event.preventDefault();
    setLoading(true);

    const res = await fetch("/api/blog", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${localStorage.getItem("authToken")}`,
      },
      body: JSON.stringify({
        title,
        content,
      }),
    });

    if (res.ok) {
      navigate("/");
    } else if (res.status === 401) {
      alert("Unauthorized; logging user out.");
      localStorage.removeItem("authToken");
      localStorage.removeItem("authLevel");
      localStorage.removeItem("email");
      navigate("/");
    } else {
      console.error(await res.text());
      alert("Something went wrong!");
      setLoading(false);
    }
  }

  return (
    <>
      <DefaultHeader />
      <main>
        <h1>New Blog</h1>
        <form onSubmit={create}>
          <label htmlFor="title">Title</label>
          <input id="title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required />
          <label htmlFor="content" className="col-span-full">Content</label>
          <textarea id="content"
            className="col-span-full"
            rows={16}
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required />
          <button className="col-span-full" disabled={loading} type="submit">Create Blog Post</button>
        </form>
      </main>
    </>
  );
}
