import {useReward} from "react-rewards";

export function StudyButtons() {
    const {reward} = useReward('rewardId', 'confetti');

    return (
        <div className="flex flex-col gap-2">
            <button className="btn btn-soft btn-primary">
                <span className="font-extrabold">Again</span>
                <span className="text-xs font-thin">(1m)</span>
            </button>
            <button className="btn btn-soft btn-primary">
                <span className="font-extrabold">Hard</span>
                <span className="text-xs font-thin">(2d)</span>
            </button>
            <button className="btn btn-soft btn-primary">
                <span className="font-extrabold">Good</span>
                <span className="text-xs font-thin">(4d)</span>
            </button>
            <button className="btn btn-soft btn-primary" onClick={reward}>
                <span id="rewardId"/>
                <span className="font-extrabold">Easy</span>
                <span className="text-xs font-thin">(7d)</span>
            </button>
        </div>
    )
}