import React from "react";
import * as z from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { NavLink } from "react-router-dom";
import Lable from "@/components/ui/Lable";
import Input from "@/components/ui/Input";
import Button from "@/components/ui/Button";

const loginSchema = z.object({
  email: z.string().email("Email khong hop le"),
  password: z
    .string()
    .min(4, "password toi thieu 4 ki tu")
    .max(20, "password khong duco qua 20 ki tu"),
});

const registerSchema = loginSchema.extend({
  userName: z.string().min(5, "ten toi thieu 5 ki tu"),
});

const AuthForm = ({ type, eventSubmit }) => {
  const schema = type === "login" ? loginSchema : registerSchema;

  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm({ resolver: zodResolver(schema) });
  return (
    <form
      action=""
      onSubmit={handleSubmit(eventSubmit)}
      className="p-10 bg-base-300 shadow-xl rounded w-full min-w-[400px]"
    >
      {/* Heading */}
      {type === "login" ? (
        <h2 className="text-3xl text-center mb-10">Welcome back NHBook</h2>
      ) : (
        <h2 className="text-3xl text-center mb-10">Welcome to NHBook</h2>
      )}
      {/* Body */}
      <div className="flex gap-4 flex-col">
        {type === "login" ? (
          <></>
        ) : (
          <div>
            <div className="flex gap-4 justify-between items-center">
              <Lable lable="UserName" className="min-w-[80px] text-md" />
              <Input
                placeholder="EX: aquan"
                size="md"
                props={{ ...register("userName") }}
              />
            </div>
            <p className="text-error text-sm">{errors?.userName?.message}</p>
          </div>
        )}
        <div>
          <div className="flex gap-4 justify-between items-center">
            <Lable lable="Email" className="min-w-[80px] text-md" />
            <Input
              type="email"
              placeholder="EX: aquan@gmail.com"
              size="md"
              props={{ ...register("email") }}
            />
          </div>
          <p className="text-error text-sm">{errors?.email?.message}</p>
        </div>
        <div>
          <div className="flex gap-4 justify-between items-center">
            <Lable lable="Password" className="min-w-[80px] text-md" />
            <Input
              placeholder="Nhap mat khau"
              props={{ ...register("password") }}
              size="md"
            />
          </div>
          <p className="text-error text-sm">{errors?.password?.message}</p>
        </div>
      </div>
      {/* Action */}
      <div className="flex flex-col justify-center items-center mt-10">
        <Button size="lg" wide="wide" color="primary">
          {type === "login" ? "Dang nhap" : "Dang ki"}
        </Button>
        <p className="text-sm mt-3">
          {type === "login" ? (
            <>
              Ban chua co tai khoan?
              <NavLink to="/auth/register" className="text-primary">
                Dang ki
              </NavLink>
            </>
          ) : (
            <>
              Ban da co tai khoan?
              <NavLink to="/auth/login" className="text-primary">
                Dang nhap
              </NavLink>
            </>
          )}
        </p>
      </div>
    </form>
  );
};

export default AuthForm;
