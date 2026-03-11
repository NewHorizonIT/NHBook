"use client";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Form } from "@/components/ui/form";
import { RegisterSchema, RegisterSchemaType } from "@/features/auth/validate";
import FormFieldCustom from "@/features/auth/components/FormFieldCustom";

const RegisterForm = () => {
  const form = useForm<RegisterSchemaType>({
    resolver: zodResolver(RegisterSchema),
    defaultValues: {
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  // Function to handle form submission
  const onSubmit = (data: RegisterSchemaType) => {
    console.log("Form Data:", data);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <h1 className="text-2xl font-bold mb-5 text-center">
          Đăng ký tài khoản NHBook
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
        <FormFieldCustom
          control={form.control}
          name="confirmPassword"
          label="Xác nhận mật khẩu"
          typeInput="password"
        />
        <Button
          type="submit"
          className="w-full cursor-pointer"
          variant={"secondary"}
        >
          Đăng ký
        </Button>
        <Button
          type="button"
          className="w-full mt-3 cursor-pointer"
          variant={"secondary"}
        >
          Đăng ký với Google
        </Button>
        <p className="mt-4 text-sm text-center">
          Bạn đã có tài khoản?{" "}
          <Link href="/login" className="text-blue-500">
            Đăng nhập
          </Link>
        </p>
      </form>
    </Form>
  );
};

export default RegisterForm;
