import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import {RouterProvider} from "react-router";
import {router} from "@/navigation/router.ts";
import "./daisyui.css"
import './index.scss'
import {QueryClientProvider} from "@tanstack/react-query";
import {queryClient} from "@/api/client/query-client.ts";
import {retrieveLaunchParams} from "@tma.js/sdk-react";
import {EnvUnsupported} from "@/components/app/EnvUnsupported.tsx";
import {init} from "@/init.ts";
// Mock the environment in case, we are outside Telegram.
import './miniapp-mock-env.ts';
// import {ReactQueryDevtoolsPanel} from "@tanstack/react-query-devtools";

const root = createRoot(document.getElementById('root')!);

try {
    const launchParams = retrieveLaunchParams();
    const {tgWebAppPlatform: platform} = launchParams;
    const debug = (launchParams.tgWebAppStartParam || '').includes('debug') || import.meta.env.DEV;

    // Configure all application dependencies.
    await init({
        debug,
        eruda: debug && ['ios', 'android'].includes(platform),
        mockForMacOS: platform === 'macos',
    })
        .then(() => {
            root.render(
                <StrictMode>
                    <QueryClientProvider client={queryClient}>
                        <RouterProvider router={router}/>
                        {/*<ReactQueryDevtoolsPanel />*/}
                    </QueryClientProvider>
                </StrictMode>
            );
        });
} catch (e) {
    root.render(<EnvUnsupported />);
}
