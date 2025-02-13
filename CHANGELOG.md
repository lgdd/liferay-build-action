# Changelog

## [2.1.0](https://github.com/lgdd/liferay-build-action/compare/v2.0.0...v2.1.0) (2025-02-13)


### Features

* add new parameter build-extra-args ([ac493a5](https://github.com/lgdd/liferay-build-action/commit/ac493a55923df1867c57b2250b7dfd6d9ffad8f7))

## [2.0.0](https://github.com/lgdd/liferay-build-action/compare/v1.0.0...v2.0.0) (2025-02-13)


### ⚠ BREAKING CHANGES

* change default java version to 17
* update artifact name to include java version
* upgrade deprecated actions

### Bug Fixes

* artifact already exists ([a352494](https://github.com/lgdd/liferay-build-action/commit/a35249465ac1f267a7fe132f0ce71bd4874e830d))
* **ci:** missing checkout step ([a9c1ed4](https://github.com/lgdd/liferay-build-action/commit/a9c1ed44ccb4e873dbde7305dab26b446af6d4ce))
* **ci:** wrong artifact name ([23a778c](https://github.com/lgdd/liferay-build-action/commit/23a778c6ea8734d701d09dbbc36d089f67f3ed4e))


### Miscellaneous Chores

* change default java version to 17 ([f2bf1cf](https://github.com/lgdd/liferay-build-action/commit/f2bf1cfef552e22c56bc077cc97e585846334e88))
* update artifact name to include java version ([e06ce39](https://github.com/lgdd/liferay-build-action/commit/e06ce396cf35fb213b6ede5f7edc961c0bb9fb0b))
* upgrade deprecated actions ([a816a1f](https://github.com/lgdd/liferay-build-action/commit/a816a1f145f4a526e264afd077e91096753a1e25))

## 1.0.0 (2023-06-11)


### Features

* add create & upload artifacts in workflow ([1b4a267](https://github.com/lgdd/liferay-build-action/commit/1b4a2676d32402b5134aef273aef692c13c5b2ca))
* build for maven or gradle without input ([8730cb1](https://github.com/lgdd/liferay-build-action/commit/8730cb158194f5a41a90d5d51ab8f22b07dbc9f9))


### Bug Fixes

* **ci:** maven upload test ([cedf256](https://github.com/lgdd/liferay-build-action/commit/cedf256c7dca5c9f621e0c541131be44df97d19e))
* **ci:** missing step id ([8bcfc78](https://github.com/lgdd/liferay-build-action/commit/8bcfc785f27d071478e3f5cf67bd113c27c1a996))
* **ci:** syntax issue & filename with mv ([0a65a38](https://github.com/lgdd/liferay-build-action/commit/0a65a389cfa65cd68c82b98b984befae2e1f584e))
* **ci:** wrong path for checksum ([7b71f1c](https://github.com/lgdd/liferay-build-action/commit/7b71f1cf7fe1bec09b3a5d90cb71dd82edd3d5c4))
* **ci:** wrong path for new checksum ([3442c45](https://github.com/lgdd/liferay-build-action/commit/3442c4564dd0766a621f881c06ffe81b5817fce1))
* duplicate step id ([2229e8b](https://github.com/lgdd/liferay-build-action/commit/2229e8befd51d2141699bdcea4ec87f07e9a383f))
* missing dist folders ([22a207e](https://github.com/lgdd/liferay-build-action/commit/22a207e576b7ca33a38c74c602b6fbe6f8e073f3))
* remove go steps ([d66fa75](https://github.com/lgdd/liferay-build-action/commit/d66fa75182d789100c3f989184e89393e64cd962))
* requirements & bundle path for gradle ([fa59e1d](https://github.com/lgdd/liferay-build-action/commit/fa59e1d1eb500c2b86aa0a74439c76c7ce0e6d97))
* unconsistent upload name ([8cf32c6](https://github.com/lgdd/liferay-build-action/commit/8cf32c669ca58c934dd4d8ddd722ce4f2c23368d))
* unnecessary mkdir command ([0262ab6](https://github.com/lgdd/liferay-build-action/commit/0262ab69fc268996662681be29e3d2358c091cfe))
* use single quotes to test upload condition ([0e0ed02](https://github.com/lgdd/liferay-build-action/commit/0e0ed02fce96218bb10ecacdc13ad9ab7fa11e33))
* verify if dirs are empty before mv ([e40a569](https://github.com/lgdd/liferay-build-action/commit/e40a569ca8c9672ae59262097bcac01def76a40f))
* wrong directory for chmod +x ([b53ade1](https://github.com/lgdd/liferay-build-action/commit/b53ade11b4a780cca11cd94b547fae4ae1105bd4))
* wrong first step in action ([f45241f](https://github.com/lgdd/liferay-build-action/commit/f45241f198b4f5642f1b169329e0410cf8f01b5d))
* wrong location for the go binary ([cfff3d7](https://github.com/lgdd/liferay-build-action/commit/cfff3d7233d3798b69a15f0e55aa04168dd39475))
* wrong maven bundle path ([b782a62](https://github.com/lgdd/liferay-build-action/commit/b782a6222e6835c793be32f13a5b72ff12b5836b))
* wrong upload artifacts condition ([e35afcd](https://github.com/lgdd/liferay-build-action/commit/e35afcdf912d41a5aa65c77e7f91dd209a72937c))
