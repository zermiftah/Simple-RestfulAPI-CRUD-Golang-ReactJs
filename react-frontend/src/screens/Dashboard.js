import NavbarSidebar from "../components/NavbarSidebar"

export default function Dashboard() {


    return (
        <>
            <div className="min-h-full">
                <NavbarSidebar />
                <main className="flex-1 pb-8">
                    {/* Page header */}
                    <div className="bg-white shadow">
                        <div className="px-4 sm:px-6 lg:max-w-6xl lg:mx-auto lg:px-8">
                            <div className="py-6 md:flex md:items-center md:justify-between lg:border-t lg:border-gray-200">
                                <div className="flex-1 min-w-0">



                                </div>
                                <div className="mt-6 flex space-x-3 md:mt-0 md:ml-4">
                                </div>
                            </div>
                        </div>
                    </div>


                </main>
            </div>

        </>
    )
}
