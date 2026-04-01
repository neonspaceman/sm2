import {client} from "@/api/client/client.ts";
import {CardsResponse} from "@/api/card/types.ts";

export const fetchCards = async () => {
    const response = await client.get<CardsResponse>('/cards');
    return response.data.data
}