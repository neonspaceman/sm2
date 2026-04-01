import {CardForm} from "@/components/card/CardForm.tsx";

function AddPage() {
    return (
        <div className="max-w-lg mx-auto flex flex-1 flex-col justify-center gap-3">
            <div className="text-xl font-bold">Add new card</div>
            <div className="text-sm">Create a flashcard to expand your knowledge base.</div>
            <CardForm />
        </div>
    );
}

export const Component = AddPage
