CREATE TABLE go_employees.employees (
	id BIGINT UNSIGNED auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	salary BIGINT UNSIGNED NULL,
	age INT NULL,
	profile_image varchar(100) NULL,
	CONSTRAINT employees_pk PRIMARY KEY (id)
)