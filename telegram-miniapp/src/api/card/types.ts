export type CardsResponse = ApiResponse<Array<{
    id: string;
    question: string;
    answer: string;
}>>;