import Image from "next/image";

const Home = () => {
  return (
    <div className="flex-grow p-5">
      <h1 className="text-xl font-bold mb-10 ">This is my portfolio</h1>
      <p>
        <Image
          src="/images/top.webp"
          alt="Example Image"
          width={500}
          height={300}
          layout="responsive"
        />
      </p>
    </div>
  );
};

export default Home;
