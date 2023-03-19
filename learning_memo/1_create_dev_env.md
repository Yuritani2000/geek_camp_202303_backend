# 1. Creation of Development Environment

## Try development with Docker

I has been intended to develop this system with Docker since I decided to create backend.

This is because some reasons:
- Docker never conaminate my local environment and would never have any effects on my future projects with Go
- By using `docker-compose.yml`, I can easily combinate and save whole the environment of this backend system, or Go and other backend components like RDB .
- I have been using pre-constructed `docker-compose.yml` files in other projects, but not understood the practical usage, so this is also my challenge for Docker and docker-compose.

## Search for a Way to Create an Environment

I am intend to use Go and PostgreSQL for backend systems.
I choose PostgreSQL because I was using it in a class in my college and have a little experience about it.

At first, I searched for articles on Google to know how to write a `docker-compose.yml` file or something to create my environment. 

The following article would be help.

[Link to a Qiita article (Japanese)](https://qiita.com/suzu12/items/177c03e8984881f1efa4)
[Link to a Zenn article (Japanese)](https://zenn.dev/yusuke49/articles/9ed37838861b1d)

I can easily copy and pase the content of `docker-compose.yml` on the Qiita atricle, but it is important to evaluate multiple information resources to understand the meaning of each lines of the file. That's I chose these 2 articles.

I found some words that I don't know. I have to complete this backend as soon as possible because it is a hackathon, but I searched for the meaning of words roughly.

- ["Boilerplate" from Qiita Article (Japanese)](https://qiita.com/e99h2121/items/995671a83be1fa7b74bf)
  - It seems like a template to use by copy and pasete
- [Base Image description by IBM](https://www.ibm.com/docs/en/order-management-sw/9.5.0?topic=docker-base-image)
  - A kind of docker image that is used to create other docker images.
  - I found that the [Zenn Article](https://zenn.dev/yusuke49/articles/9ed37838861b1d) used this word for a sample code snippets on it. 
- [An article about Alpine Linux (Japanese)](https://charlie1012.hatenablog.jp/entry/2021/01/14/090000)
  - A kind of linux disribution that is minimal and fast.
  - It is popular due to the speed, but [some people warns about compatibility](https://blog.inductor.me/entry/alpine-not-recommended) of [libc](https://linuxjf.osdn.jp/JFdocs/libc-intro.html), that is C language libraries that is used on Linux.

This time, I decided to use Alpine as the Qiita and Zenn article used, in order to know how it is. I will decide whether i continue to use Alpine in the future by considering the benefits and problems.

Finally, I created the following `docker-compose.yml` file. I mainly quoted the [Qiita Atcile](https://qiita.com/suzu12/items/177c03e8984881f1efa4) reading the [Zenn Article](https://zenn.dev/yusuke49/articles/9ed37838861b1d) as well.

Docker container has been successfully created, and the test code on the Qiita Article has successfully worked.
It is obvious that this works successfully because I just simply mimitated the code of the Qiita Article...

Anyway, I could create Go development environment with Docker for the first time!
I will gather more information to improve and create my own backend during the remaining time!


<!-- The articles I quoted were a little bit old, so I modified [Docker Compose File Version](https://docs.docker.jp/compose/compose-file/compose-versioning.html#compose-file-version-3) to use as newer Docker as possible. -->

