import apiInstance from "@/api/axios";
import { useQuery } from "@tanstack/react-query";
import React from "react";

const Sidebar = () => {
  const fetchCategory = async () => {
    const res = await apiInstance.get("/categories/");
    console.log(res.data);
    return res.data;
  };

  const { data } = useQuery({
    queryKey: ["categories"],
    queryFn: () => fetchCategory(),
  });
  return (
    <div className="p-4 bg-base-200 shadow rounded-lg sm:block hidden w-[200px] sticky top-[65px] ">
      <h3>Lọc sách</h3>
    </div>
  );
};

export default Sidebar;
