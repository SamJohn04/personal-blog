import { Link, useNavigate, useParams } from "react-router";
import DefaultHeader from "../components/DefaultHeader";
import { useEffect, useState } from "react";

type Blog = {
  id: number,
  title: string,
  content: string,
  createdAt: Date,
  lastUpdatedAt: Date
}

export default function Blog() {
  const navigate = useNavigate();

  const [blog, setBlog] = useState<Blog | null>(null);
  const { id } = useParams();

  useEffect(() => {
    async function getBlog() {
      const res = await fetch(`/api/blog/${id}`);
      if (!res.ok) {
        return;
      }
      const data = await res.json();
      setBlog({
        id: data.id,
        title: data.title,
        content: data.content,
        createdAt: new Date(data.createdAt),
        lastUpdatedAt: new Date(data.lastUpdatedAt),
      })
    }
    getBlog();
  }, []);

  async function deleteBlog() {
    if (!confirm("Are you sure you want to delete the blog?")) {
      return;
    }
    const res = await fetch(`/api/blog/${id}`, {
      method: "DELETE",
      headers: {
        "Authorization": `Bearer ${localStorage.getItem("authToken")}`,
      }
    });
    if (res.ok) {
      navigate("/");
    } else {
      alert("Something went wrong");
    }
  }

  return (
    <>
      <DefaultHeader/>
      <main>
        <h1>{blog?.title}</h1>
        <p className="text-gray">{blog?.createdAt?.toDateString()}</p>
        {
          localStorage.getItem("authLevel") === "3" &&
            <div className="pad-b-10 grid cols-2 gap-4">
          <Link to={`/blog/${id}/edit`} className="grid"><button>Edit Blog</button></Link>
          <button onClick={deleteBlog}>Delete Blog</button>
          </div>
        }
        <div dangerouslySetInnerHTML={{ __html: blog?.content ?? "" }} />
      </main>
    </>
  )
}
