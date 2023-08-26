/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    domains: ['tailwindui.com','splidejs.com',"lh3.googleusercontent.com","localhost","disewa.id","samb-api.disewa.id"]
  },
  env: {
    pathimage: "https://samb-api.disewa.id",
    apidomain: "https://samb-api.disewa.id",
  },

}

module.exports = nextConfig
