import {
  Menubar,
  MenubarContent,
  MenubarItem,
  MenubarMenu,
  MenubarTrigger,
} from "@/components/ui/menubar";
import Link from "next/link";

export default function Menu() {
  return (
    <Menubar className="border-0 gap-4 text-xl">
      <MenubarMenu>
        <MenubarTrigger>
          <Link href={"/"}>Trang chủ</Link>
        </MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger>Danh mục</MenubarTrigger>
        <MenubarContent>
          <MenubarItem>
            <Link href={"/category/programing"}>Programing</Link>
          </MenubarItem>
          <MenubarItem>
            <Link href={"/category/education"}>Education</Link>
          </MenubarItem>
          <MenubarItem>
            <Link href={"/category/novel"}>Novel</Link>
          </MenubarItem>
        </MenubarContent>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger>
          <Link href={"/new-books"}>Sách mới</Link>
        </MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger>
          <Link href={"/sale"}>Khuyến mãi</Link>
        </MenubarTrigger>
      </MenubarMenu>
      <MenubarMenu>
        <MenubarTrigger>
          <Link href={"/contact"}>Liên hệ</Link>
        </MenubarTrigger>
      </MenubarMenu>
    </Menubar>
  );
}
