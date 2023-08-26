import Link from "next/link";
import React, { useRef, useState } from 'react';
import axios from 'axios';
import axiosInstance from '../axiosConfig';
import { useRouter } from 'next/router';
import Head from "next/head";
import Navbar from "./navbar";

export async function getStaticProps() {
    try {



        const ListWareHouseResponse = await axios.get(process.env.apidomain + '/api/v1/warehouses', {});
        const ListWareHouse = ListWareHouseResponse.data.data;

        const ListSuppliersResponse = await axios.get(process.env.apidomain + '/api/v1/suppliers', {});
        const ListSupplier = ListSuppliersResponse.data.data;

        return { props: {  ListWareHouse, ListSupplier } };
    } catch (error) {
        console.error(error);
        return { props: {  ListWareHouse: [], ListSupplier: [] } };
    }
}

const AddBarangMasuk = ({  ListWareHouse, ListSupplier }) => {

    const router = useRouter();

    const [selectedWareHouse, setWareHouse] = useState(1)
    const [selectedSupplier, setSupplier] = useState(1)





    const [formValue, setformValue] = React.useState({
        TrxInNo: '',
        TrxInNotes: '',

    });


    const handleSubmit = async (event) => {
        event.preventDefault();


        var Request = {
            TrxInNo: formValue.TrxInNo,
            TrxInNotes: formValue.TrxInNotes,
            TrxInSuppIdf: parseInt(selectedSupplier),
            WhsIdf: parseInt(selectedWareHouse),
        }

        // console.log("Request ==> "  , Request)

        var Json = JSON.stringify(Request);

        console.log("Request ==> "  , Json)

        try {
            // make axios post request
            const response = await axiosInstance.post('/api/v1/barang_masuk/header', Json,{
                headers: {
                    'Content-Type': 'application/json',
                }
            });
            router.push('/barang_masuk');
            console.log(response)
        } catch (error) {
            console.log(error)
        }
    }


    const handleChange = (event) => {
        setformValue({
            ...formValue,
            [event.target.name]: event.target.value
        });
    }
    return (
        <>


            <Head>
                <title>Disewa.id</title>
                <meta name="description" content="sewa lapangan gor venue" />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <Navbar />

            <main>


                <div className="grid grid-cols-8 gap-2">
                    <div className="lg:col-start-3 lg:col-end-7 col-span-8  border-2 rounded-md border-neutral-900">
                        <div className="relative p-4">
                            <form onSubmit={handleSubmit}>
                                <div className="space-y-12">

                                    <div className="border-b border-gray-900/10 pb-12">
                                        <h2 className="text-base font-semibold leading-7 text-gray-900">Add Barang Masuk Header</h2>





                                        <div className="mt-10 grid grid-cols-1 gap-x-6  sm:grid-cols-6">
                                            <div className="sm:col-span-3">
                                                <label htmlFor="name"
                                                    className="block text-sm font-medium leading-6 text-gray-900">
                                                    Trx No
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        value={formValue.TrxInNo}
                                                        onChange={handleChange}
                                                        type="text"
                                                        name="TrxInNo"
                                                        id="TrxInNo"
                                                        placeholder={'TRX-001'}
                                                        autoComplete="given-name"
                                                        className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    />
                                                </div>
                                            </div>
                                          

                                            <div className="sm:col-span-3">
                                                <label htmlFor="country"
                                                    className="block text-sm font-medium leading-6 text-gray-900">
                                                    WareHouse
                                                </label>
                                                <div className="mt-2">


                                                    <select value={selectedWareHouse} onChange={(e) => {
                                                        setWareHouse(e.target.value);
                                                    }}
                                                        id="warehouse"
                                                        name="warehouse"
                                                        autoComplete="country-name"
                                                        className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    >
                                                        {/* <option selected>Choose a Product</option> */}

                                                        {ListWareHouse.map((item) => (
                                                            <option value={item.WhsPK}>{item.WhsName}</option>
                                                        ))}
                                                    </select>
                                                </div>
                                            </div>

                                          

                                            <div className="sm:col-span-3">
                                                <label htmlFor="country"
                                                    className="block text-sm font-medium leading-6 text-gray-900">
                                                    Supplier
                                                </label>
                                                <div className="mt-2">


                                                    <select value={selectedSupplier} onChange={(e) => {
                                                        setSupplier(e.target.value);
                                                    }}
                                                        id="supplier"
                                                        name="supplier"
                                                        autoComplete="country-name"
                                                        className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    >
                                                        {/* <option selected>Choose a Product</option> */}

                                                        {ListSupplier.map((item) => (
                                                            <option value={item.SupplierPK}>{item.SupplierName}</option>
                                                        ))}
                                                    </select>
                                                </div>
                                            </div>

                                          

                                            <div className="sm:col-span-3">
                                                <label htmlFor="address"
                                                    className="block text-sm font-medium leading-6 text-gray-900">
                                                    Notes
                                                </label>
                                                <div className="mt-2">
                                                    <textarea
                                                        value={formValue.TrxInNotes}
                                                        onChange={handleChange}
                                                        id="TrxInNotes"
                                                        name="TrxInNotes"
                                                        placeholder="sambalnya dipisah"
                                                        rows={2}
                                                        className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                        defaultValue={''}
                                                    />
                                                </div>
                                            </div>
                                        </div>


                                    </div>


                                </div>


                                <div className="mt-6 flex items-center justify-end gap-x-6">
                                    <button type="button" className="text-sm font-semibold leading-6 text-gray-900">
                                        <Link href={`/barang_masuk`}>Cancel</Link>
                                    </button>
                                    <button
                                        type="submit"
                                        className="rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                    >
                                        Save
                                    </button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </main>
        </>
    )
}
export default AddBarangMasuk;
