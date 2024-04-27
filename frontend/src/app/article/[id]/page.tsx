"use client";

import React from "react";
import Image from "next/image";
import { Article } from "@/types/article";
import config from "@/env/config";

const ArticleDetail = async ({ params }: { params: { id: number } }) => {
  const article = await getDetailArtcle(params.id);

  return (
    <div className="max-w-3xl mx-auto p-5">
      <Image
        src={`/images/${article.imageId}.webp`}
        alt="image"
        width={200}
        height={100}
      />
      <h1 className="text-4xl text-center mb-10 mt-10 border-2 border-gray-300 rounded-lg shadow p-3">
        {article.title}
      </h1>
      <p>{article.content}</p>
    </div>
  );
};

export default ArticleDetail;

const getDetailArtcle = async (id: number): Promise<Article> => {
  const res = await fetch(`${config.apiBaseUrl}/articles/${id}`, {
    next: { revalidate: 60 },
  });

  if (!res.ok) {
    throw new Error("エラーが発生しました。");
  }

  await new Promise((resolve) => setTimeout(resolve, 1500));

  const json = await res.json();

  return json.data;
};
