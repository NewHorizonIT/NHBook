import apiInstance from "@/api/axios";
import { ListBook } from "@/components/common/Books";
import { Container } from "@/components/layout/Containner";
import Button from "@/components/ui/Button";
import { Hopper } from "@/components/ui/icons";
import Input from "@/components/ui/Input";
import Lable from "@/components/ui/Lable";
import Modal from "@/components/ui/Modal";
import Select from "@/components/ui/Select";
import { useQuery } from "@tanstack/react-query";
import React, { useState } from "react";

const BooksPage = () => {
  const [title, setTitle] = useState("");
  const [category, setCategory] = useState(0);
  const [maxPrice, setMaxPrice] = useState(0);
  const [minPrice, setMinPrice] = useState(0);

  const fetchCategories = async () => {
    try {
      const res = await apiInstance.get("/categories/status/0");
      console.log(res);
      return res;
    } catch (error) {
      console.log("ERR:", error);
    }
  };

  const handleFilter = () => {
    const dataFilter = {
      title,
      category,
      maxPrice,
      minPrice,
    };

    console.log(dataFilter);
  };

  const { data: categories } = useQuery({
    queryKey: ["categories"],
    queryFn: () => fetchCategories(),
    staleTime: 1000 * 60 * 5,
  });
  const [openFilter, setOpenFilter] = useState(false);

  const handleCloseFilter = () => {
    console.log(openFilter);
    setOpenFilter(false);
  };
  return (
    <Container>
      <div className="">
        <ListBook />
        <Button
          size="md"
          wide="square"
          color="primary"
          className="fixed bottom-10 right-5 "
          onClick={() => setOpenFilter(true)}
        >
          <Hopper />
        </Button>
        {openFilter && (
          <Modal onClose={handleCloseFilter}>
            <h2 className="text-xl font-semibold">Lọc sản phẩm</h2>
            <div className="my-8">
              <div>
                <Lable lable="Tên sách" className="text-lg">
                  <Input
                    size="md"
                    className="my-2 border-0"
                    placeholder="Nhập tên sách"
                    props={{ value: title }}
                    onChange={(e) => setTitle(e.target.value)}
                  ></Input>
                </Lable>
              </div>
              <div>
                <Lable lable="Thể loại" className="text-lg">
                  <Select
                    size="md"
                    options={categories}
                    className="my-2 border-0 shadow-inner"
                    onChange={(e) => setCategory(e.target.value)}
                  />
                </Lable>
              </div>
              <div>
                <Lable lable="Giá từ" className="text-lg">
                  <Input
                    type="number"
                    size="md"
                    className="border-0"
                    props={{ value: minPrice }}
                    onChange={(e) => setMinPrice(e.target.value)}
                  />
                </Lable>
              </div>
              <div>
                <Lable lable="Giá từ" className="text-lg">
                  <Input
                    type="number"
                    size="md"
                    className="border-0"
                    props={{ value: maxPrice }}
                    onChange={(e) => setMaxPrice(e.target.value)}
                  />
                </Lable>
              </div>
            </div>

            <div className="">
              <Button color="primary" onClick={handleFilter}>
                Tim kiem
              </Button>
            </div>
          </Modal>
        )}
      </div>
    </Container>
  );
};

export default BooksPage;
