import { lazy } from "react";
import { createBrowserRouter } from "react-router-dom";
import DefaultLayout from "../layouts/DefaultLayout";

const HomePage = lazy(() => import("@/pages/HomePage"));
const AboutPage = lazy(() => import("@/pages/AboutPage"));
const CategoryPage = lazy(() => import("@/pages/CategoryPage"));
const HistoryOrderPage = lazy(() => import("@/pages/HistoryOrderPage"));

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
        path: "/category",
        element: <CategoryPage />,
      },
      {
        path: "/history",
        element: <HistoryOrderPage />,
      },
    ],
  },
]);

export default router;
