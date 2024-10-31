--- my_first_golang
create table customer (
	id int auto_increment,
	customer_name varchar(100) not null,
	customer_gender char(1) not null,
	customer_identity_number varchar(50) not null,
	primary key (id)
);

select *
from customer;