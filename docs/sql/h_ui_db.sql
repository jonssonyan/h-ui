create table account
(
    id             INTEGER                  not null
        primary key autoincrement,
    username       TEXT      default ''     not null
        unique,
    `pass`         TEXT      default ''     not null,
    con_pass       TEXT      default ''     not null,
    quota          INTEGER   default 0      not null,
    download       INTEGER   default 0      not null,
    upload         INTEGER   default 0      not null,
    expire_time    INTEGER   default 0      not null,
    kick_util_time INTEGER   default 0      not null,
    device_no      INTEGER   default 3      not null,
    `role`         TEXT      default 'user' not null,
    deleted        INTEGER   default 0      not null,
    create_time    TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime')),
    update_time    TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime'))
);
create index account_deleted_index
    on account (deleted);
create index account_con_pass_index
    on account (con_pass);
create index account_pass_index
    on account (`pass`);
create index account_username_index
    on account (username);
INSERT INTO account (username, `pass`, con_pass, quota, download, upload, expire_time, device_no, role)
VALUES ('sysadmin', 'f8cdb04495ded47615258f9dc6a3f4707fd2405434fefc3cbf4ef4e6',
        'c7591c31adf8af0b6b8ae8cbbccd8d1aaa0c7bb068f576bddb6378d5', -1, 0, 0, 253370736000000, 6, 'admin');
create table config
(
    id          INTEGER              not null
        primary key autoincrement,
    `key`       TEXT      default '' not null,
    `value`     TEXT      default '' not null,
    remark      TEXT      default '' not null,
    create_time TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime')),
    update_time TIMESTAMP default (datetime(CURRENT_TIMESTAMP, 'localtime'))
);
create index account_key_index
    on config (`key`);
INSERT INTO config (key, value, remark)
VALUES ('H_UI_WEB_PORT', '8081', 'H UI Web 端口');
INSERT INTO config (key, value, remark)
VALUES ('JWT_SECRET', hex(randomblob(10)), 'JWT 密钥');
INSERT INTO config (key, value, remark)
VALUES ('HYSTERIA2_ENABLE', '0', 'Hysteria2 开关');
INSERT INTO config (key, value, remark)
VALUES ('HYSTERIA2_CONFIG', '', 'Hysteria2 配置');
INSERT INTO config (key, value, remark)
VALUES ('HYSTERIA2_TRAFFIC_TIME', '1', 'Hysteria2 流量倍数');