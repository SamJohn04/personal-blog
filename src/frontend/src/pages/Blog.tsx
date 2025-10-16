import { useParams } from "react-router";
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
  return (
    <>
      <DefaultHeader/>
      <main>
        <h1>{blog?.title}</h1>
        <h3>{blog?.createdAt?.toDateString()}</h3>
        {blog?.content}
      </main>
    </>
  )
}
