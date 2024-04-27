"use client"

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import React, { useState } from 'react';
import { signInWithEmail, signInWithGoogle, createAccount, logout,signInWithGoogleRedirect } from '@/lib/auth';

const Login = () => {

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async () => {
    try {
      await signInWithEmail(email, password);
      console.log("Logged in successfully!");
    } catch (error) {
      console.error(error);
    }
  };

  const handleGoogleLogin = async () => {
    try {
      const result = await signInWithGoogle();
      console.log("Logged in with Google!");
      console.log(result);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="max-w-lg w-full p-8 bg-white shadow-xl rounded-xl">
        <h2 className="text-3xl font-bold mb-8 text-center">ログイン</h2>
        <div className="space-y-6">
          <Input placeholder="メールアドレス" value={email} onChange={(e) => setEmail(e.target.value)}/>
          <Input type="password" placeholder="パスワード" value={password} onChange={(e) => setPassword(e.target.value)}/>
          <Button  onClick={handleLogin} className="w-full">ログイン</Button>
          <Button onClick={handleGoogleLogin} className="w-full bg-blue-500 hover:bg-blue-600 text-white">
            Googleでログイン
          </Button>
        </div>
      </div>
    </div>
  );
};

export default Login;
