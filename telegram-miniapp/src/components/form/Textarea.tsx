import cn from "classnames";
import {useFormContext} from "react-hook-form";

interface TextareaProps {
    label: string;
    name: string;
    placeholder?: string;
}

export function Textarea({label, name, placeholder}: TextareaProps) {
    const {register, formState: {errors}} = useFormContext();

    return (
        <fieldset className="fieldset">
            <label className="label">{label}</label>
            <textarea
                placeholder={placeholder}
                className={
                    cn("textarea w-auto outline-none", {
                        "input-primary": !errors.answer,
                        "input-error": !!errors.answer
                    })
                }
                {...register(name)}
            ></textarea>
            {errors.answer && (
                <p className="text-error">
                    {errors.answer.message?.toString()}
                </p>
            )}
        </fieldset>
    )
}