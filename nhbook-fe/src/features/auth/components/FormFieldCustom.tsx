import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { HTMLInputTypeAttribute } from "react";
import { Control } from "react-hook-form";

interface FormFieldCustomProps {
  control: Control<any>;
  name: string;
  label?: string;
  typeInput: HTMLInputTypeAttribute;
}

const FormFieldCustom = ({
  control,
  name,
  label,
  typeInput,
}: FormFieldCustomProps) => {
  return (
    <div className="my-5">
      <FormField
        control={control}
        name={name}
        render={({ field }) => (
          <FormItem>
            <FormLabel>{label}</FormLabel>
            <FormControl>
              <Input placeholder="shadcn" {...field} type={typeInput} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
    </div>
  );
};

export default FormFieldCustom;
