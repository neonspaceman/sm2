import {createHashRouter, redirect} from "react-router-dom";
import {ROUTES} from "@/navigation/routes.ts";
import {App} from "@/components/app/App.tsx";
import {RootErrorBoundary} from "@/components/error-boundary/ErrorBoundary.tsx";
import {LoadingScreen} from "@/components/loading-screen/LoadingScreen.tsx";

export const router = createHashRouter([
    {
        path: '/',
        Component: App,
        ErrorBoundary: RootErrorBoundary,
        HydrateFallback: LoadingScreen,
        children: [
            {
                path: ROUTES.HOME,
                lazy: () => import("@/pages/IndexPage.tsx"),
            },
            {
                path: ROUTES.STUDY,
                lazy: () => import('@/pages/StudyPage.tsx'),
            },
            {
                path: ROUTES.ADD_CARD,
                lazy: () => import('@/pages/AddPage.tsx'),
            },
            {
                path: ROUTES.PROFILE,
                lazy: () => import('@/pages/ProfilePage.tsx'),
            },
            {
                path: '*',
                loader: () => redirect(ROUTES.HOME),
            },
        ],
    },
]);
