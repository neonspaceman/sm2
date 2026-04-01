import {fetchCards} from "@/api/card/requests.ts";
import {useQuery} from "@tanstack/react-query";

export function useCardList() {
    return useQuery({
        queryFn: () => fetchCards(),
        queryKey: ["cards"]
    })
}