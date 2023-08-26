
import styles from '@/styles/Home.module.css'
import Head from 'next/head'
import Navbar from "@/pages/navbar";
import axios from "axios";
import Link from "next/link";


export async function getServerSideProps({ params, req  }) {


    try {




        const ReqVenue = await axios.get(process.env.apidomain + '/api/v1/barang_masuk/header/'+params.id, );
        const DataBarangMasukHeader = ReqVenue.data.data;

        const ReqDetail = await axios.get(process.env.apidomain + '/api/v1/barang_masuk/detail/'+params.id, );
        const DataBarangMasuDetail = ReqDetail.data.data;

        return {props: { DataBarangMasukHeader,DataBarangMasuDetail,params}};
    } catch (error) {
        console.error(error);
        return {props: {DataBarangMasukHeader: {},DataBarangMasuDetail:[],params:0}};
    }




}

const DetailVenue = ({DataBarangMasukHeader,DataBarangMasuDetail,params}) => {


  
    return (
        <>
            <Head>
                {/* <title>{title}</title> */}
                {/* <meta name="description" content={"Sewa Lapangan "+DataVenue.category?.name+" "+DataVenue.org?.city
                    +" "+DataVenue.org?.name}/>
                <meta name="viewport" content="width=device-width, initial-scale=1"/>
                <link rel="icon" href="/favicon.ico"/> */}
            </Head>
            <Navbar/>
            <main>


            <div className={styles.center}>
                    <div className="mx-auto max-w-2xl py-4 px-5 sm:py-8 sm:px-6 lg:max-w-7xl lg:px-8">

                    <div>

<div className=" border-t border-gray-100">
    <dl className="divide-y divide-gray-100">
        <div className="px-4 py-1 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt className="text-sm font-bold  leading-6 text-gray-900">Trx Masuk No</dt>
            <dd className="mt-1  text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{DataBarangMasukHeader.TrxInNo}</dd>
        </div>
        <div className="px-4 py-1 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt className="text-sm font-bold leading-6 text-gray-900">WareHouse Name</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">WA {DataBarangMasukHeader.MasterWareHouse.WhsName}</dd>
        </div>
        <div className="px-4 py-1 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt className="text-sm font-bold leading-6 text-gray-900">Date</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{DataBarangMasukHeader.TrxInDate}</dd>
        </div>
        <div className="px-4 py-1 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt className="text-sm font-bold leading-6 text-gray-900">Supplier Name</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                {DataBarangMasukHeader.MasterSupplier.SupplierName}
            </dd>

          


        </div>
     

    </dl>
</div>
</div>


                        <Link href={`/add_barang_masuk_detail/${String(params.id)}`}>
                            <button
                                className=" bg-secondary-500 text-white active:bg-pink-600 font-bold uppercase text-sm px-6 py-2 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-5 ease-linear transition-all duration-150"
                                type="button"
                            >
                                Tambah Detail

                            </button>
                        </Link>

                        <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
                            <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                                <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                                    <tr>
                                        <th scope="col" className="px-6 py-3">
                                            ProductName
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Quantity Dus
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Quantity Pcs
                                        </th>
                                     
                                    </tr>
                                </thead>
                                <tbody>
                                    {DataBarangMasuDetail.map((item) => (
                                        <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                                            <th scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                                {item.TrxInDProduct.ProductName}
                                            </th>
                                            <td className="px-6 py-4">
                                                {item.TrxInDQtyDus}
                                            </td>
                                            <td className="px-6 py-4">
                                                {item.TrxInDQtyPcs}
                                            </td>
                                           
                                        </tr>))
                                    }


                                </tbody>
                            </table>
                        </div>



                    </div>


                </div>


            </main>


        </>
    )
}

export default DetailVenue;
