import Link from "next/link";
import React, { useRef, useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';
import Head from "next/head";
import axiosInstance from '../../axiosConfig';
import Navbar from "@/pages/navbar";


export async function getServerSideProps({ params, req  }) {
    try {



        const ListProductResponse = await axios.get(process.env.apidomain + '/api/v1/products', {});
        const ListProduct = ListProductResponse.data.data;

        const ListBarangMasukHeaderResponse = await axios.get(process.env.apidomain + '/api/v1/barang_masuk/header', {});
        const ListBarangMasukHeader = ListBarangMasukHeaderResponse.data.data;


        console.log("ListBarangMasukHeader ==> " , ListBarangMasukHeader)
        console.log("ListProduct ==> " , ListProduct)
        return { props: {  ListProduct, ListBarangMasukHeader,params } };
    } catch (error) {
        console.error(error);
        return { props: {  ListProduct: [], ListBarangMasukHeader: [] ,params:0} };
    }
}

const AddBarangMasuk = ({  ListProduct, ListBarangMasukHeader,params }) => {

    const router = useRouter();

    const [selectedHeader, setHeader] = useState(params.id)
    const [selectedProduct, setProduct] = useState(1)





    const [formValue, setformValue] = React.useState({
        TrxInDQtyDus: 1,
        TrxInDQtyPcs: 1,

    });


    const handleSubmit = async (event) => {
        event.preventDefault();


        var Request = {
            TrxInDQtyDus: parseInt(formValue.TrxInDQtyDus),
            TrxInDQtyPcs: parseInt(formValue.TrxInDQtyPcs),
            TrxInDProductIdf: parseInt(selectedProduct),
            TrxInIDF: parseInt(params.id),
        }

        // console.log("Request ==> "  , Request)

        var Json = JSON.stringify(Request);

        console.log("Request ==> "  , Json)

        try {
            // make axios post request
            const response = await axiosInstance.post('/api/v1/barang_masuk/detail', Json,{
                headers: {
                    'Content-Type': 'application/json',
                }
            });
            router.push(' /barang_masuk_header/'+selectedHeader);
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
                                        <h2 className="text-base font-semibold leading-7 text-gray-900">Add Barang Masuk Detail</h2>





                                        <div className="mt-10 grid grid-cols-1 gap-x-6  sm:grid-cols-6">
                                          

                                            <div className="sm:col-span-3">
                                                <label htmlFor="country"
                                                    className="block text-sm font-medium leading-6 text-gray-900">
                                                    Barang Masuk Header
                                                </label>
                                                <div className="mt-2">


                                                    <select value={selectedHeader} onChange={(e) => {
                                                        setHeader(e.target.value);
                                                    }}
                                                        id="warehouse"
                                                        name="warehouse"
                                                        disabled="true"
                                                        autoComplete="country-name"
                                                        className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    >
                                                        {/* <option selected>Choose a Product</option> */}

                                                        {ListBarangMasukHeader.map((item) => (
                                                            <option value={item.TrxInPK}>{item.TrxInNo}</option>
                                                        ))}
                                                    </select>
                                                </div>
                                            </div>

                                          

                                            <div className="sm:col-span-3">
                                                <label htmlFor="country"
                                                    className="block text-sm font-medium leading-6 text-gray-900">
                                                    Product
                                                </label>
                                                <div className="mt-2">


                                                    <select value={selectedProduct} onChange={(e) => {
                                                        setProduct(e.target.value);
                                                    }}
                                                        id="supplier"
                                                        name="supplier"
                                                        autoComplete="country-name"
                                                        className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    >

                                                        {ListProduct.map((item) => (
                                                            <option value={item.ProductPK}>{item.ProductName}</option>
                                                        ))}
                                                    </select>
                                                </div>
                                            </div>

                                          

                                          
                                            <div className="sm:col-span-3">
                                                <label htmlFor="price"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                   Quantity Dus
                                                </label>


                                                <div className="mt-2 flex">
                                                          <span
                                                              className="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200 border border-r-0 border-gray-300 rounded-l-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                                                              Dus.
                                                          </span>

                                                    <input value={formValue.TrxInDQtyDus}
                                                           onChange={handleChange}
                                                           type="number"
                                                           name="TrxInDQtyDus"
                                                           id="TrxInDQtyDus"
                                                           className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                           placeholder="60000"/>
                                                </div>


                                            </div>


                                            <div className="sm:col-span-3">
                                                <label htmlFor="price"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                   Quantity Pcs
                                                </label>


                                                <div className="mt-2 flex">
                                                          <span
                                                              className="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200 border border-r-0 border-gray-300 rounded-l-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                                                              Pcs.
                                                          </span>

                                                    <input value={formValue.TrxInDQtyPcs}
                                                           onChange={handleChange}
                                                           type="number"
                                                           name="TrxInDQtyPcs"
                                                           id="TrxInDQtyPcs"
                                                           className="rounded-none rounded-r-lg bg-gray-50 border text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm border-gray-300 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                           placeholder="60000"/>
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
