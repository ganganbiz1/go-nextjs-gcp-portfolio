import { auth, googleProvider } from "./../../firebase.config";
import {
  signInWithEmailAndPassword,
  signInWithPopup,
  createUserWithEmailAndPassword,
  signInWithRedirect,
  signOut,
  UserCredential,
} from "firebase/auth";
import { Credential } from "@/types/user";

export const signInWithEmail = (
  email: string,
  password: string
): Promise<UserCredential> => {
  return signInWithEmailAndPassword(auth, email, password);
};

export const signInWithGoogle = async (): Promise<Credential> => {
  const result = await signInWithPopup(auth, googleProvider);
  const token = await result.user.getIdToken();
  const uid = result.user.uid;
  const providerId = result.user.providerData[0]?.providerId ?? "";
  const displayName = result.user.displayName;
  return {
    idToken: token,
    displayName: displayName,
    uid: uid,
    providerId: providerId,
  };
};

// もしかしたら使うかもなので、一応残しておく
export const signInWithGoogleRedirect = (): any => {
  signInWithRedirect(auth, googleProvider);
};

export const createAccount = (
  email: string,
  password: string
): Promise<UserCredential> => {
  return createUserWithEmailAndPassword(auth, email, password);
};

export const logout = (): Promise<void> => {
  return signOut(auth);
};
