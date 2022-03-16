create table users
(
    id bigint unsigned auto_increment primary key
);

create table multiple_many
(
    user_1 bigint unsigned not null,
    user_2 bigint unsigned not null,

    constraint unique_pair
        unique (user_1, user_2),
    constraint multiple_many_user_1_foreign
        foreign key (user_1) references users (id)
            on delete cascade,
    constraint multiple_many_user_2_foreign
        foreign key (user_2) references users (id)
            on delete cascade
);

INSERT INTO users (id) VALUES (1);
INSERT INTO users (id) VALUES (2);
INSERT INTO multiple_many (user_1, user_2) VALUES (1, 2);

