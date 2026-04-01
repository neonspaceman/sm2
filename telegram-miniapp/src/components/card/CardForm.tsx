import {FormProvider, useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {cardScheme} from "@/components/card/schemas.ts";
import {LayersPlus} from "lucide-react";
import {Textarea} from "@/components/form/Textarea.tsx";

export function CardForm() {
    const form = useForm({
        resolver: zodResolver(cardScheme),
    });

    const onSubmit = form.handleSubmit((data) => {
        console.log(data);
    })

    return (
        <FormProvider {...form}>
            <form onSubmit={onSubmit} className="w-auto">
                <Textarea label="Question" name="question" placeholder="What is the concept you want to memorize?" />
                <Textarea label="Answer" name="answer" placeholder="The solution or definition..." />
                <button className="btn btn-primary mt-4" type="submit"><LayersPlus/> Add to deck</button>
            </form>
        </FormProvider>
    )
}