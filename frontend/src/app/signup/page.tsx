import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import React from "react";

const Signup = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="max-w-lg w-full p-8 bg-white shadow-xl rounded-xl">
        <h2 className="text-3xl font-bold mb-8 text-center">サインアップ</h2>
        <div className="space-y-6">
          <Input placeholder="メールアドレス" />
          <Input type="password" placeholder="パスワード" />
          <Button className="w-full">サインアップ</Button>
          <Button className="w-full bg-blue-500 hover:bg-blue-600 text-white">
            Googleでサインアップ
          </Button>
        </div>
      </div>
    </div>
  );
};

export default Signup;
