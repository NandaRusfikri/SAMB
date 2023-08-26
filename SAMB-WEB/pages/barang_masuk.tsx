import Head from 'next/head'

import styles from '@/styles/Home.module.css'
import React, { useState } from "react";
import { useRouter } from 'next/router'
import axios from 'axios';
import Image from 'next/image'
import Link from "next/link";
import Navbar from "@/pages/navbar";


export async function getServerSideProps() {

    try {
        // const ListVenueResponse = await axios.post(process.env.apidomain + '/api/v1/barang_masuk/list', {});
        // const ListVenue = ListVenueResponse.data.data;

        const BarangMasukResponse = await axios.get(process.env.apidomain + '/api/v1/barang_masuk/header', {});
        const ListBarangMasuk = BarangMasukResponse.data.data;

        return { props: { ListBarangMasuk } };
    } catch (error) {
        console.error(error);
        return { props: { ListBarangMasuk: [] } };
    }
}

interface HomeProps {
    ListBarangMasuk: any[];
    // ListVenue: any[];
    error: any;
}

const Home: React.FC<HomeProps> = ({ ListBarangMasuk, error }) => {
    if (error) {
        return <div>An Error {error.message}</div>
    }
    const [selectedCategory, setCategory] = useState<string | number>();


    return (
        <>
            <Head>
                <title>Add Barang Masuk </title>
                <meta name="description" content="Sewa Lapangan Olahraga dengan jadwal" />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <Navbar />
            <main>


                <div className={styles.center}>

                    <h1 className={styles.title}>Barang Masuk</h1>
                    <div className="mx-auto max-w-2xl py-4 px-5 sm:py-8 sm:px-6 lg:max-w-7xl lg:px-8">
                        <Link href={`/add_barang_masuk_header`}>
                            <button
                                className=" bg-secondary-500 text-white active:bg-pink-600 font-bold uppercase text-sm px-6 py-2 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-5 ease-linear transition-all duration-150"
                                type="button"
                            >
                                Tambah

                            </button>
                        </Link>

                        <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
                            <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                                <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                                    <tr>
                                        <th scope="col" className="px-6 py-3">
                                            TrxInNo
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            WareHouse
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Date
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Supplier
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Notes
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                          Action
                                        </th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {ListBarangMasuk.map((item) => (
                                        <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                                            <th scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                                {item.TrxInNo}
                                            </th>
                                            <td className="px-6 py-4">
                                                {item.MasterWareHouse.WhsName}
                                            </td>
                                            <td className="px-6 py-4">
                                                {item.TrxInDate.substring(0, 10)}
                                            </td>
                                            <td className="px-6 py-4">
                                                {item.MasterSupplier.SupplierName}
                                            </td>
                                            <td className="px-6 py-4">
                                                {item.TrxInNotes}
                                            </td>
                                            <td className="px-6 py-4 text-left">
                                            <Link href={`barang_masuk_header/${String(item.TrxInPK)}`}>
                                            <button type="button" className="mt-100  text-white bg-green-500 hover:bg-green-600 focus:ring-4 focus:outline-none focus:ring-green-400 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:hover:bg-green-700 dark:focus:ring-green-400 mr-2 mb-2">
                                                Detail</button>
                                                    
                                                    </Link>
                                            
                                               
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
};

export default Home;
