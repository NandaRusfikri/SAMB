import Head from 'next/head'


import Navbar from "@/pages/navbar";


export async function getServerSideProps() {

    try {
       

        return { props: {  } };
    } catch (error) {
        console.error(error);
        return { props: {  } };
    }
}

interface HomeProps {

    error: any;
}

const Home: React.FC<HomeProps> = ({ error }) => {
    if (error) {
        return <div>An Error {error.message}</div>
    }

    return (
        <>
            <Head>
                <title>SAMB</title>
                <meta name="description" content="SAMB" />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <Navbar />
            <main>


            <h2 > <a href="https://samb-api.disewa.id/swagger/index.html">Document API Swagger</a></h2>
                


            </main>
        </>
    )
};

export default Home;
