import {Loading} from "@/components/loading/Loading.tsx";

export function LoadingScreen() {
    return (
        <div className="flex flex-col items-center min-h-screen justify-center">
            <Loading />
        </div>
    )
}