"use client";

import React, { useState, useEffect } from "react";
import config from "@/env/config";
import Image from "next/image";
import Link from "next/link";
import { Article } from "@/types/article";

const ArticleList = () => {
  const [articles, setArticles] = useState<Article[]>([]);

  useEffect(() => {
    const fetchArticles = async () => {
      const res = await fetch(`${config.apiBaseUrl}/articles`);
      if (!res.ok) {
        console.error('Failed to fetch articles');
        return;
      }
      const json = await res.json();
      setArticles(json.data);
    };

    fetchArticles();
  }, []);

  return (
    <div className="container mx-auto px-4 py-8">
    <h1 className="text-4xl font-bold mb-6">Articles</h1>
    <div className="flex justify-end flex-grow mb-8">
      <Link href="/article/create" passHref>
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
          Create Article
        </button>
      </Link>
    </div>
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {articles.map((article) => (
        <Link href={`/article/${article.id}`} passHref>
        <div key={article.id} className="bg-white rounded-lg shadow-md p-4">
          <Image
            src={`/images/${article.imageId}.webp`}
            alt="image"
            width={50}
            height={50}
          />
          <h2 className="text-xl font-bold mt-4 mb-4"> {article.title.length > 9 ? `${article.title.substring(0, 9)}...` : article.title}</h2>
          <p className="text-gray-600 p-4"> {article.content.length > 10 ? `${article.content.substring(0, 10)}...` : article.content}</p>
        </div>
        </Link>
      ))}
    </div>
  </div>
  );
};

export default ArticleList;
