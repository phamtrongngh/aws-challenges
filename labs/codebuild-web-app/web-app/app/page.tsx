import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <div className="h-screen flex flex-col items-center justify-center">
      <Image src="/aws.png" alt="Logo" width={200} height={200} />
      <p
        className="
          text-center
          text-2xl
          text-gray-600
          font-semibold
          p-4
        "
      >
        Hello! This is a simple NextJS web app for the {" "}
        <Link
          className="text-blue-500 hover:underline"
          href="https://github.com/phamtrongngh/aws-labs-for-kids"
        >
          AWS Labs for Kids series. Changed!
        </Link>
      </p>
    </div>
  );
}
