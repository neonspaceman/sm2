function ProfilePage() {
    return (
        <div className="flex flex-1 flex-col justify-center gap-5">
            <div className="avatar mx-auto">
                <div className="ring-primary ring-offset-base-100 w-24 rounded-full ring-3 ring-offset-5">
                    <img src="https://img.daisyui.com/images/profile/demo/spiderperson@192.webp"/>
                </div>
            </div>
            <div className="text-center text-xl font-extrabold">Dmitry</div>
        </div>
    );
}

export const Component = ProfilePage
