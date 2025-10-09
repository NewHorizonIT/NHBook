import React from "react";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "../ui/pagination";

interface PagnitionProps {
  currentPage: number;
  maxPage: number;
}

const PaginationCustom = ({ currentPage, maxPage }: PagnitionProps) => {
  return (
    <Pagination className="my-6">
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious href="#" />
        </PaginationItem>
        {currentPage - 1 > 0 && (
          <PaginationItem>
            <PaginationLink href="#">{currentPage - 1}</PaginationLink>
          </PaginationItem>
        )}
        <PaginationItem>
          <PaginationLink href="#" isActive>
            {currentPage}
          </PaginationLink>
        </PaginationItem>
        <div className="items-center md:flex hidden">
          {currentPage + 1 < maxPage && (
            <PaginationItem>
              <PaginationLink href="#">{currentPage + 1}</PaginationLink>
            </PaginationItem>
          )}
        </div>
        <PaginationItem>
          <PaginationNext href="#" />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
};

export default PaginationCustom;
