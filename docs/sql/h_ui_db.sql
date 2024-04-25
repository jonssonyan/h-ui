create table account
(
    id          INTEGER              not null
        primary key autoincrement,
    username    TEXT      default '' not null,
    `pass`      TEXT      default '' not null,
    quota       INTEGER   default 0  not null,
    download    INTEGER   default 0  not null,
    upload      INTEGER   default 0  not null,
    expire_time INTEGER   default 0  not null,
    is_admin    INTEGER   default 0  not null,
    deleted     INTEGER   default 0  not null,
    create_time TIMESTAMP default CURRENT_TIMESTAMP,
    update_time TIMESTAMP default CURRENT_TIMESTAMP
);

create index account_deleted_index
    on account (deleted);

create index account_pass_index
    on account (`pass`);

create index account_username_index
    on account (username);

INSERT INTO account (username, `pass`, quota, download, upload, expire_time, is_admin)
VALUES ('admin', 'f8cdb04495ded47615258f9dc6a3f4707fd2405434fefc3cbf4ef4e6', -1, 0, 0, 253370736000000, 1);


create table config
(
    id          INTEGER              not null
        primary key autoincrement,
    `key`       TEXT      default '' not null,
    `value`     TEXT      default '' not null,
    remark      TEXT      default '' not null,
    create_time TIMESTAMP default CURRENT_TIMESTAMP,
    update_time TIMESTAMP default CURRENT_TIMESTAMP
);

create index account_key_index
    on config (`key`);

INSERT INTO config (key, value, remark)
VALUES ('H_UI_WEB_PORT', '8081', 'H UI Web 端口');
INSERT INTO config (key, value, remark)
VALUES ('HYSTERIA2_CONFIG', '{}', 'Hysteria2 配置');
INSERT INTO config (key, value, remark)
VALUES ('JWT_SECRET', randomblob(10), 'JWT 密钥');