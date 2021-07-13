create table t_article(
    id int primary key,
    title varchar(100),
    content text,
    status tinyint not null,
    userId int not null 
)
