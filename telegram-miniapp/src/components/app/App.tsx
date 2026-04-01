import {Outlet} from "react-router-dom";
import {Dock} from "@/components/dock/Dock.tsx";
import {ROUTES} from "@/navigation/routes.ts";
import {BadgePlus, Brain, House, User} from "lucide-react";

export function App() {
    return (
        <>
            <div className="min-h-screen p-5 pb-[4rem] flex flex-col max-w-2xl mx-auto">
                <Outlet/>
            </div>
            <Dock
                routes={[
                    {route: ROUTES.HOME, text: "Cards", Icon: House},
                    {route: ROUTES.STUDY, text: "Study", Icon: Brain},
                    {route: ROUTES.ADD_CARD, text: "New", Icon: BadgePlus},
                    {route: ROUTES.PROFILE, text: "Profile", Icon: User},
                ]}
            />
        </>
    )
}
