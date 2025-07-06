const ContentStorage = artifacts.require("ContentStorage");

module.exports = function (deployer) {
  // Triển khai smart contract ContentStorage
  deployer.deploy(ContentStorage);
};
