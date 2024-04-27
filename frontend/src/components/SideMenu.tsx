import React from "react";
import Image from "next/image";
import { Button } from "./ui/button";
import Link from "next/link";

function SideMenu() {
  return (
    <div className="hidden lg:flex flex-col w-64 min-h-screen bg-blue-300 p-5 border-r border-gray-300">
      <h1 className="font-bold text-lg pb-4 mb-4">
        <Image
          src="/images/logo.jpg"
          alt="Example Image"
          width={50}
          height={50}
        />
      </h1>
      <ul>
      <li>
          <Link href="/" passHref>
            <Button variant="link">
              <p className="hover:bg-blue-300">Home</p>
            </Button>
          </Link>
        </li>
        <li>
          <Link href="/article" passHref>
            <Button variant="link">
              <p className="hover:bg-blue-300">Article</p>
            </Button>
          </Link>
        </li>
        <li>
          <Button variant="link">
            <p className="hover:bg-blue-300">Dummy</p>
          </Button>
        </li>
        <li>
          <Button variant="link">
            <p className="hover:bg-blue-300">Dummy</p>
          </Button>
        </li>
      </ul>
    </div>
  );
}

export default SideMenu;
