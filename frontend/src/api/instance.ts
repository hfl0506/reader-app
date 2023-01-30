import axios from "axios";
import { getSession, signOut } from "next-auth/react";

const BASE_URL = process.env.NEXT_PUBLIC_API_HOST!;

const instance = axios.create({
  baseURL: BASE_URL,
  timeout: 20000,
  headers: {
    "Content-Type": "application/json",
  },
});

instance.interceptors.request.use(async (request) => {
  const session = await getSession();
  if (session) {
    request.headers = request?.headers ?? {};
    request.headers.Authorization = `Bearer ${session?.accessToken}`;
  }

  return request;
});

instance.interceptors.response.use(
  async (response) => {
    return response;
  },
  function (error) {
    if (
      error?.response?.status === 401 &&
      error?.response?.config?.headers?.Authorization
    ) {
      signOut();
    }
    return Promise.reject(error);
  }
);

export default instance;
