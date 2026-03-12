"use client";

import { Form } from "@/components/ui/form";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import FormFieldCustom from "@/features/auth/components/FormFieldCustom";
import { LoginSchema, LoginSchemaType } from "@/features/auth/validate";

const LoginForm = () => {
  // use react hook form with zod resolver and LoginSchema
  const form = useForm<LoginSchemaType>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  // Function to handle form submission
  const onSubmit = (data: LoginSchemaType) => {
    console.log("Form Data:", data);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <h1 className="text-2xl font-bold mb-5 text-center">
          Đăng nhập vào NHBook
        </h1>
        <FormFieldCustom
          control={form.control}
          name="email"
          label="Email"
          typeInput="email"
        />
        <FormFieldCustom
          control={form.control}
          name="password"
          label="Mật khẩu"
          typeInput="password"
        />
        <Button
          type="submit"
          className="w-full cursor-pointer"
          variant={"secondary"}
        >
          Đăng nhập
        </Button>
        <Button
          type="button"
          className="w-full mt-3 cursor-pointer"
          variant={"secondary"}
        >
          Đăng nhập với Google
        </Button>
        <p className="mt-4 text-sm text-center">
          Bạn chưa có tài khoản?{" "}
          <Link href="/register" className="text-blue-500 underline">
            Đăng ký
          </Link>
        </p>
      </form>
    </Form>
  );
};

export default LoginForm;
