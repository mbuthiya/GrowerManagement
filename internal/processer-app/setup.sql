drop table Users cascade;
drop table Processors cascade;
drop table Workers cascade;
drop table Farms cascade;
drop table Notifications cascade;
drop table Appointments cascade;
drop table Crops cascade;
drop table Crop_Farm cascade;
drop table Practices cascade;
drop table Animals cascade;
drop table Farm_Animals cascade;
drop table Reports cascade;
drop table Crops_Management cascade;
drop table Management_Conditions cascade;
drop table Crop_Inputs cascade;



create table Users(
    id serial primary key,
    email varchar(255),
    password text
);

create table Processors(
    id serial primary key,
    user_id integer references Users(id),
    company_name varchar(255),
    country varchar(100)
);

create table Workers(
    id serial primary key,
    user_id integer references Users(id),
    processor_id integer references Processors(id),
    first_name varchar(100),
    last_name varchar(100),
    phone_number integer
);

create table Farms(
    id serial primary key,
    processor_id integer references Processors(id),
    date_added Date,
    lattitude varchar,
    longitude varchar,
    farm_code varchar(50)
);

create table Notifications(
    id serial primary key,
    user_id integer references Users(id),
    notification_time timestamp,
    notification_message text
);

create table Appointments(
    id serial primary key,
    worker_id integer references Workers(id),
    farm_id integer references Farms(id),
    appointment_type varchar(50)
);

create table Crops(
    id serial primary key,
    crop_name varchar(100)
    
);

create table Crop_Farm(
    id serial primary key,
   crop_id integer references Crops(id),
   farm_id integer references Farms(id),
   frequency varchar(30),
   crop_year varchar,
   yield integer,
   planting_date Date
);

create table Practices(
    id serial primary key,
    farm_id integer references Farms(id),
    flood_irrigation varchar,
    sprinkler_irrigation varchar,
    drip_irrigation varchar,
    natural_enemies varchar,
    animal_manure varchar,
    green_manure varchar,
    compost_use varchar,
    manual_weeding varchar
);

create table Animals(
    id serial primary key,
    animal_name varchar
);

create table Farm_Animals(
    id serial primary key,
    farm_id integer references Farms(id),
    animal_id integer references Animals(id),
    number_of_animals integer,
    percent_as_manure integer

);

create table Reports(
    id serial primary key,
    farm_id integer references Farms(id),
    worker_id integer references Workers(id),
    report_date Date,
    report_approved varchar,
    report_comments text
);

create table Crops_Management(
     id serial primary key,
     report_id integer references Reports(id),
     crop_id integer references Crop_Farm(id)

);

create table Management_Conditions(
    id serial primary key,
    crop_management_id integer references Crops_Management(id),
    activity varchar,
    condition_status varchar,
    condition_justification text
);

create table Crop_Inputs(
     id serial primary key,
     crop_management_id integer references Crops_Management(id),
     product varchar,
     quantity integer,
     input_date Date

);


