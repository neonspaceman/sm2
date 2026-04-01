import {Clock8, Trash2, Zap} from "lucide-react";
import {RemoveCardDialog} from "@/components/card/RemoveCardDialog.tsx";
import {useRef} from "react";

interface CardProps {
    id: string;
    question: string;
    answer: string;
    createdAt: Date;
    easy: Number;
    interval: Number;
}

export function Card({id, question, answer, createdAt, easy, interval}: CardProps) {
    const ref = useRef<HTMLDialogElement>(null);

    const onConfirm = () => {
        ref?.current?.close();
        alert(`Delete card with id: ${id}`);
    }

    return (
        <>
            <RemoveCardDialog ref={ref} question={question} answer={answer} onConfirm={onConfirm} />
            <div className="card w-auto card-border bg-base-100 mt-5">
                <div className="card-body">
                    <h2 className="card-title flex justify-between content-center gap-4 mb-2">
                        <div>{question}</div>
                        <button onClick={() => ref?.current?.show()} className="btn btn-circle btn-soft btn-error">
                            <Trash2 className="w-4"/>
                        </button>
                    </h2>
                    <p>{answer}</p>
                    <div className="mt-4 flex items-center justify-between text-xs text-gray-400 dark:text-gray-500">
                        <div className="flex items-center gap-4">
                            <span className="flex items-center gap-1"><Clock8
                                className="w-4"/> {createdAt.toLocaleString()}</span>
                            <span className="flex items-center gap-1"><Zap
                                className="w-4"/> Ease: {easy.toString()}</span>
                        </div>
                        <span>Interval: {interval.toString()}d</span></div>
                </div>
            </div>
        </>
    )
}