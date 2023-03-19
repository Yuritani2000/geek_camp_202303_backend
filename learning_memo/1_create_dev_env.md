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

At first, I searched for articles on Goole to know how to write a `docker-compose.yml` file or something to create my environment. 

The following article would be help.

[Link to a Qiita article (Japanese)](https://qiita.com/suzu12/items/177c03e8984881f1efa4)
[Link to a Zenn article (Japanese)](https://zenn.dev/yusuke49/articles/9ed37838861b1d)

