import {
    Web3Storage
} from './bundle.esm.min.js'

let API_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweEMzQzUyYzk1Y0NBMTg5RGFjQ0FmMjM0OURFQTExQWQxQjk2NDdlYjgiLCJpc3MiOiJ3ZWIzLXN0b3JhZ2UiLCJpYXQiOjE2NTM4MDU0ODYwNzYsIm5hbWUiOiJseWptcnkifQ.taP2rbdPWPgHE7g7eQV7BM9IRdu3pCqaruI7jCEPZXg";
// Construct with token and endpoint
const client = new Web3Storage({
    token: API_TOKEN
})

const fileInp = document.getElementById('fileInp')
console.log("fileInp = ", fileInp)
cube.onclick = function () {
    fileInp.click()
}
fileInp.onchange = async function () {
    // Pack files into a CAR and send to web3.storage
    const rootCid = await client.put(fileInp.files) // Promise<CIDString>

    // Get info on the Filecoin deals that the CID is stored in
    const info = await client.status(rootCid) // Promise<Status | undefined>

    console.log("info = ", info)
}