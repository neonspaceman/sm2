import {Link, useLocation} from "react-router-dom";
import cn from "classnames"
import {ComponentType, SVGProps} from "react";

interface RouteInterface {
    route: string;
    text: string;
    Icon: ComponentType<SVGProps<SVGSVGElement>>;
}

interface DockProps {
    routes: RouteInterface[];
}

export function Dock({routes}: DockProps) {
    const {pathname} = useLocation();

    return (
        <div className="dock justify-center">
            {routes.map((route) => {
                return (
                    <Link key={route.route} to={route.route} className={cn({"dock-active": pathname === route.route})}>
                        <route.Icon className="size-[1.2em]"/>
                        <span className="dock-label">{route.text}</span>
                    </Link>
                )
            })}
        </div>
    )
}