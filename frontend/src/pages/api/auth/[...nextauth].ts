import { NextApiRequest, NextApiResponse } from "next";
import NextAuth, { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

export const authOptions: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        email: { label: "Email", type: "email" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        try {
          const response = await signIn();
        } catch (error) {
          return null;
        }
      },
    }),
  ],
  pages: {
    signIn: "/login",
    error: "/login",
  },
  callbacks: {
    async signIn({ user }) {
      if (user?.roles)
        return (
          user?.roles.find(
            (role) => role == "admin" || role == "developer" || role == "author"
          ) === undefined
        );
      return false;
    },
    async redirect({ baseUrl }) {
      return baseUrl;
    },
    async jwt({ token, user }) {
      if (user) {
        return {
          ...token,
          accessToken: user.accessToken,
          accessTokenExpiresAt: user.accessTokenExpiresAt,
          refreshToken: user.refreshToken,
          refreshTokenExpiresAt: user.refreshTokenExpiresAt,
          userId: user.userId,
          roles: user.roles,
        };
      }

      if (token.accessTokenExpiresAt < Date.now() / 1000) {
        return withLock(token, refreshAccessToken);
      }

      return token;
    },
    async session({ session, token }) {
      session.userId = token.userId;
      session.accessToken = token.accessToken;
      session.accessTokenExpiresAt = token.accessTokenExpiresAt;
      session.roles = token.roles;

      session.isDeveloper =
        session.roles &&
        session.roles.find(
          (value) => value == "developer" || value == "admin"
        ) != undefined;
      session.isAuthor =
        session.roles &&
        session.roles.find((value) => value == "author") != undefined;

      return session;
    },
  },
  session: {
    maxAge: 30 * 24 * 60 * 60,
  },
  debug: process.env.NODE_ENV !== "production",
};

export default (req: NextApiRequest, res: NextApiResponse) =>
  NextAuth(req, res, authOptions);
