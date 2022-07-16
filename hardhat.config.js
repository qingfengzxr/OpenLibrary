require("@nomiclabs/hardhat-waffle");
// require("@nomiclabs/hardhat-ganache");
require('hardhat-contract-sizer');

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.12",
  contractSizer: {
    alphaSort: true,
    runOnCompile: true,
    disambiguatePaths: false,
  },
  networks: {
    gw_devnet_v1: {
      url: `http://127.0.0.1:8024`,
      accounts: [`0x9d5bc55413c14cf4ce360a6051eacdc0e580100a0d3f7f2f48f63623f6b05361`],
    },
    gw_testnet_l2: {
      url: `https://godwoken-testnet-v1.ckbapp.dev`,
      accounts: [`0x29cc0c552c6b610cd24699628606d45aa80ea7cff50c043f38d9178fa1d579c5`],
    }
  },
};


