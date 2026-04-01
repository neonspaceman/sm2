import {Page} from "@/components/page/Page.tsx";

export function EnvUnsupported() {
    return (
        <div className="hero min-h-screen">
            <div className="hero-content text-center">
                <div className="max-w-md">
                    <h1 className="text-5xl font-bold">Oops!</h1>
                    <p className="py-6">You are using too old Telegram client to run this application</p>
                    <img className="mx-auto max-w-50" alt="Telegram sticker" src="https://xelene.me/telegram.gif"/>
                </div>
            </div>
        </div>
    );
}