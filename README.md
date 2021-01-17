# zplat

[![](https://github.com/saltbo/zplat/workflows/build/badge.svg)](https://github.com/saltbo/zplat/actions?query=workflow%3Abuild)
[![](https://codecov.io/gh/saltbo/zplat/branch/master/graph/badge.svg)](https://codecov.io/gh/saltbo/zplat)
[![](https://wakatime.com/badge/github/saltbo/zplat.svg)](https://wakatime.com/badge/github/saltbo/zplat)
[![](https://api.codacy.com/project/badge/Grade/88817db9b3b04c0293c9d001d574a5ef)](https://app.codacy.com/manual/saltbo/zplat?utm_source=github.com&utm_medium=referral&utm_content=saltbo/zplat&utm_campaign=Badge_Grade_Dashboard)
[![](https://img.shields.io/github/v/release/saltbo/zplat.svg)](https://github.com/saltbo/zplat/releases)
[![](https://img.shields.io/github/license/saltbo/zplat.svg)](https://github.com/saltbo/zplat/blob/master/LICENSE)

English | [🇨🇳中文](https://saltbo.cn/zplat)

## Features

- 安装系统：引导用户进行安装
- 配置系统：帮助开发者进行配置管理
- 用户系统：提供注册登录，个人配置等功能

## Endpoints

<!-- markdown-swagger -->

Endpoint            | Method | Auth? | Description
 ------------------- | ------ | ----- | --------------------------
`/tokens`           | POST   | No    | 用于账户登录和申请密码重置
`/users`            | GET    | No    | 获取一个用户信息
`/users`            | POST   | No    | 注册一个用户
`/users/{email}`    | PATCH  | No    | 用于账户激活和密码重置
`/users/{username}` | GET    | No    | 获取一个用户信息

<!-- /markdown-swagger -->

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.

## Contact us

- [Author Blog](https://saltbo.cn).

## Author

- [saltbo](https://github.com/saltbo)

## License

- [MIT](https://github.com/saltbo/zplat/blob/master/LICENSE)