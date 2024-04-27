import type { Metadata } from "next";
import { Noto_Sans_JP } from "next/font/google";
import "./globals.css";
import SideMenu from "@/components/SideMenu";
import Footer from "@/components/Footer";

const notoSansJp = Noto_Sans_JP({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "My portfolio",
  description: "This is my portfolio",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={notoSansJp.className}>
        <div className="flex min-h-screen">
          <SideMenu />
          {children}
        </div>
        <Footer />
      </body>
    </html>
  );
}
