# geek_camp_202303_backend
The backend system as an achievement for "Geek Camp (技育CAMP)", a hackathon organized by Supportez (サポーターズ) in March, 2023.

## Setup the Server

First, clone this repository from GitHub into your computer.

```sh
git clone https://github.com/Yuritani2000/geek_camp_202303_backend
```

Move to the downloaded repository.

```sh
cd geek_camp_202303_backend
```

Copy `.env.sample` into `.env` , in which the environmental variables are.

```sh
cp .env.sample .env
```

Run the following command to run the docker container.

```sh
docker-compose up -d
```

## Main References

This backend app derives from so many articles and other information on the Internet. I appreciate the authors of articles in which I used the code snippets.
- [Go Development Environment Creation with Docker (Japanese)](https://qiita.com/suzu12/items/177c03e8984881f1efa4#sample-%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%A0)


## Learning Memo

I wrote daily records of development of this backend system for my remainder.
I wish you could know how I have thought and done during development. [Link to Learning Memo](./learning_memo/)