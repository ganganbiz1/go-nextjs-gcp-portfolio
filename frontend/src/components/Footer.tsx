import React from "react";
import Image from "next/image";
import { Button } from "./ui/button";
import Link from "next/link";

const Footer = () => {
  return (
    <div className="lg:hidden flex justify-between items-center w-full p-5 bg-gray-100 text-center fixed bottom-0">
    <div className="flex-grow text-left ml-10">
      <Link href="/" passHref>
        <Button variant="link">
          <p className="hover:bg-blue-300">Home</p>
        </Button>
      </Link>
    </div>
    <div className="flex-grow text-right mr-10">
      <Link href="/article" passHref>
        <Button variant="link">
          <p className="hover:bg-blue-300">Article</p>
        </Button>
      </Link>
    </div>
  </div>
  );
};

export default Footer;
