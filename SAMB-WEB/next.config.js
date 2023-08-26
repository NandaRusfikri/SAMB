/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    domains: ['tailwindui.com','splidejs.com',"lh3.googleusercontent.com","localhost","disewa.id","samb-api.disewa.id"]
  },
  env: {
    pathimage: "http://disewa.id:2225",
    apidomain: "http://disewa.id:2225",
  },
  // env: {
  //   pathimage: "http://localhost:2224",
  //   apidomain: "http://localhost:2224",
  // },
}

module.exports = nextConfig
