# Identavatar

![GitHub](https://img.shields.io/github/license/steelWinds/identavatar)
![GitHub top language](https://img.shields.io/github/languages/top/steelWinds/identavatar)

</br>
<div align="center">
  <img width="300" height="230" src="./public/logo.svg">
  <h3 align="center">IdentAvatar</h3>
</div>

# About The Project

### Built with
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Getting Started

### Installation
1. Clone the repo
   ```bash
   git clone https://github.com/steelWinds/identavatar
   ```
2. Install modules
   ```bash
   go mod tidy
   ```
3. Build docker image and run it
   ```bash
   make build
   ```
3. Or run docker-compose with Air
   ```bash
   make up
   ```
### Usage

#### Query params:
- ```squares: Int (required)``` - amount of squares
- ```size: Int (required)``` - size of square 
- ```word: String (required)``` - your word 

#### Example
```bash
http://localhost:3180/?squares=6&size=30&word=mycoolword
```
## License

Distributed under the MPL v2 License. See LICENSE.txt for more information.

## Contact

[@steelWinds](https://github.com/steelWinds) | kirillsurov0@gmail.com | [t.me/bladeVrtx](https://t.me/bladeVrtx)
