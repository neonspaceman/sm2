import {SearchInput} from "@/components/search-input/SearchInput.tsx";
import {Card} from "@/components/card/Card.tsx";
import {useCardList} from "@/api/card/hooks.ts";
import {Loading} from "@/components/loading/Loading.tsx";

function IndexPage() {
    const {data, isPending} = useCardList()

    return (
        <>
            <SearchInput/>
            {isPending && (
                <div className="pt-5 mx-auto">
                    <Loading/>
                </div>
            )}
            {!isPending && (
                data?.map((card) => (
                    <Card
                        key={card.id}
                        id={card.id}
                        answer={card.answer}
                        question={card.question}
                        createdAt={new Date()}
                        easy={2.5}
                        interval={2}
                    />
                ))
            )}
        </>
    );
}

export const Component = IndexPage
