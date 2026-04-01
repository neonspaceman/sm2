import {useState} from "react";
import cn from "classnames"

export function StudyCard() {
    const [isFlipped, setFlipped] = useState(false)

    return (
        <div className="perspective-1000">
            <div className={cn("card cursor-pointer card-border bg-base-100 card-xl aspect-[4/3] card-3d", {"card-3d-flipped": isFlipped})} onClick={() => setFlipped(v => !v)}>
                <div className="absolute h-full w-full card-body gap-2 front">
                    <p className="place-content-center text-center">What is the capital of France?</p>
                    <p className="flex-none text-center text-secondary text-sm">Tap to reveal answer</p>
                </div>
                <div className="absolute h-full w-full rotate-y-180 card-body gap-2 back">
                    <p className="place-content-center text-sm text-center">What is the capital of France?</p>
                    <div className="divider"></div>
                    <p className="font-bold place-content-center text-center">Paris</p>
                </div>
            </div>
        </div>
    )
}