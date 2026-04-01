import {Search} from "lucide-react";

export function SearchInput() {
    return (
        <label className="input input-primary w-auto outline-none">
            <Search className="opacity-50" />
            <input type="text" className="grow" placeholder="Search your cards..."/>
        </label>
    )
}