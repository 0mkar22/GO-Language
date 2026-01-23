create table car (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    make VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    price NUMERIC(19,2) NOT NULL
);

create table person (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	email VARCHAR(150),
	gender VARCHAR(50) NOT NULL,
	date_of_birth DATE NOT NULL,
	country_of_birth VARCHAR(50) NOT NULL,
    car_id BIGINT REFERENCES car(id),
    UNIQUE(car_id)
);



insert into person (first_name, last_name, email, gender, date_of_birth, country_of_birth) values
('John', 'Doe', 'john.doe@example.com', 'Male', '1990-01-01', 'USA');
insert into person (first_name, last_name, email, gender, date_of_birth, country_of_birth) values
('Jane', 'Smith', 'jane.smith@example.com', 'Female', '1985-05-15', 'Canada');  
insert into person (first_name, last_name, email, gender, date_of_birth, country_of_birth) values
('Alice', 'Johnson', 'alice.johnson@example.com', 'Female', '1992-07-20', 'UK');

insert into car (make, model, price) values
('Toyota', 'Camry', 24000.00);
insert into car (make, model, price) values
('Honda', 'Civic', 22000.00);
