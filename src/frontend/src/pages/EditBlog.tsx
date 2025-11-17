import { useEffect, useState, type FormEvent } from "react";
import DefaultHeader from "../components/DefaultHeader";
import { useNavigate, useParams } from "react-router";

export function EditBlog() {
  const navigate = useNavigate();

  const { id } = useParams();

  const [title, setTitle] = useState<string>("");
  const [content, setContent] = useState<string>("");

  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    async function getBlog() {
      const res = await fetch(`/api/blog/${id}?edit=true`);
      if (!res.ok) {
        return;
      }
      const data = await res.json();
      setTitle(data.title);
      setContent(data.content);
      setLoading(false);
    };
    getBlog();
  }, []);

  async function update(event: FormEvent) {
    event.preventDefault();
    setLoading(true);

    const res = await fetch(`/api/blog/${id}`, {
      method: "PUT",
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
        <h1>Edit Blog</h1>
        <form onSubmit={update}>
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
          <button className="col-span-full" disabled={loading} type="submit">Edit Blog Post</button>
        </form>
      </main>
    </>
  );
}
