import z from "zod";

export const loginSchema = z.object({
  email: z.string().email("Email khong hop le").trim(),
  password: z.string().trim().min(5, "Password phai tren 5 ki tu"),
});

export const registerSchema = loginSchema.extend({
  useName: z.string().min(5, "Tên phải trên 5 kí tự").trim(),
});

export type loginData = z.infer<typeof loginSchema>;

export type registerData = z.infer<typeof registerSchema>;
