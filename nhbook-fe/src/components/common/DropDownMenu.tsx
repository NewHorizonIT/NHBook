import React from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import Link from "next/link";
const DropDownMenuAccount = () => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild className="border-0 cursor-pointer">
        <Avatar>
          <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        <DropdownMenuItem>
          <Link href={"/account/profile"}>Thông tin tài khoản</Link>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Link href={"/account/profile"}>Thông tin tài khoản</Link>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default DropDownMenuAccount;
