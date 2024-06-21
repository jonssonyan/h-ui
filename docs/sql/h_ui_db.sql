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
VALUES ('sysadmin', '02f382b76ca1ab7aa06ab03345c7712fd5b971fb0c0f2aef98bac9cd',
        'ee85cbdbb387a61de3d1d52e6773cc31e28c7913b32f4b7aa44ec61b', -1, 0, 0, 253370736000000, 6, 'admin');
create table config
(
    id          INTEGER              not null
        primary key autoincrement,
    `key`       TEXT      default '' not null
        unique,
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