import {StudyCard} from "@/components/card/StudyCard.tsx";
import {StudyButtons} from "@/components/card/StudyButtons.tsx";

function StudyPage() {
    return (
        <div className="flex flex-1 flex-col justify-center gap-5">
            <div className="text-xs">Card 1 of 3</div>
            <StudyCard/>
            <StudyButtons />
        </div>
    );
}

export const Component = StudyPage
