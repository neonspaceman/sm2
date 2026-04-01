import {forwardRef} from "react";

interface RemoveCardDialogProps {
    question: string;
    answer: string;
    onConfirm: () => void;
}

export const RemoveCardDialog = forwardRef<HTMLDialogElement, RemoveCardDialogProps>(({
    question,
    answer,
    onConfirm,
}, ref) => {
    return (
        <dialog ref={ref} className="modal">
            <div className="modal-box">
                <h3 className="text-lg font-bold">Hello!</h3>
                <div className="py-4 flex flex-col gap-3">
                    <p>Are you sure to remove the card?</p>
                    <p>Q: {question}</p>
                    <p>Q: {answer}</p>
                </div>
                <div className="modal-action">
                    <button className="btn btn-soft btn-error" onClick={onConfirm}>Yes, remove it!</button>
                    <form method="dialog">
                        <button className="btn btn-primary">Cancel</button>
                    </form>
                </div>
            </div>
        </dialog>
    )
})