import { lazy } from "react";
import { createBrowserRouter } from "react-router-dom";
import DefaultLayout from "../layouts/DefaultLayout";
import { AuthLayout } from "@/layouts";
import ErrorPage from "@/pages/ErrorPage";

const HomePage = lazy(() => import("@/pages/HomePage"));
const AboutPage = lazy(() => import("@/pages/AboutPage"));
const BooksPage = lazy(() => import("@/pages/BooksPage"));
const HistoryOrderPage = lazy(() => import("@/pages/HistoryOrderPage"));
const RegisterPage = lazy(() => import("@/pages/RegisterPage"));
const LoginPage = lazy(() => import("@/pages/LoginPage"));

const router = createBrowserRouter([
  {
    path: "/",
    element: <DefaultLayout />,
    children: [
      {
        path: "",
        element: <HomePage />,
      },
      {
        path: "/about",
        element: <AboutPage />,
      },
      {
        path: "/books",
        element: <BooksPage />,
      },
      {
        path: "/history",
        element: <HistoryOrderPage />,
      },
    ],
  },
  {
    path: "/auth",
    element: <AuthLayout />,
    children: [
      {
        path: "register",
        element: <RegisterPage />,
      },
      {
        path: "login",
        element: <LoginPage />,
      },
    ],
  },
  {
    path: "*",
    element: <ErrorPage />,
  },
]);

export default router;
