import NextAuth from "next-auth";
import { JWT } from "next-auth/jwt";

declare module "next/app" {
  type AppProps<P = Record<string, unknown>> = {
    Component: NextComponentType<NextPageContext, any, P>;
    router: Router;
    __N_SSG?: boolean;
    __N_SSP?: boolean;
    pageProps: P & {
      /** Initial session passed in from `getServerSideProps` or `getInitialProps` */
      session?: Session;
    };
  };
}

declare module "next-auth" {
  interface User {
    accessToken: string;
    accessTokenExpiresAt: number;
    refreshToken: string;
    refreshTokenExpiresAt: number;
    userId: string;
    roles: string[];
  }

  interface Session extends User {
    accessToken: string;
    accessTokenExpiresAt: number;
    userId: string;
    roles: string[];
    isDeveloper: boolean;
    isAuthor: boolean;
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    userId: string;
    accessToken: string;
    accessTokenExpiresAt: number;
    refreshToken: string;
    refreshTokenExpiresAt: number;
    roles: string[];
  }
}
