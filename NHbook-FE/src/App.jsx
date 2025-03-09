import { RouterProvider } from "react-router-dom";
import useAuthStore from "./store/AuthStore";
import router from "./router/router";

function App() {
  const isAuth = useAuthStore();
  console.log(isAuth);
  return <RouterProvider router={router}></RouterProvider>;
}

export default App;
