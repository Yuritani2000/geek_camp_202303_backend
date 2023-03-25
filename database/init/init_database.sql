CREATE DATABASE IF NOT EXISTS geek;

CREATE TABLE trip (
    id INEGER NOT NULL,
    host_id CHAR(10),
    host_name CHAR(10),
    car_license CHAR(25),
    car_name CHAR(25),
    passenger_limit INTEGER,
    location_from CHAR(30),
    location_to CHAR(30),
    PRIMARY KEY id
);

CREATE TABLE request (
    id INEGER NOT NULL,
    passenger_id CHAR(10),
    passenger_name CHAR(10),
    trip_id INTEGER,
    PRIMARY KEY id,
    FOREIGN KEY trip_id REFERENCES trip
)