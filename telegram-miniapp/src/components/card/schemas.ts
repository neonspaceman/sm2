import {z} from "zod"

export const cardScheme = z.object({
    question: z.string().trim().min(1),
    answer: z.string().trim().min(1)
})