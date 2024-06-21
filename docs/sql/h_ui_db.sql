CREATE TABLE account
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    username       TEXT    NOT NULL UNIQUE DEFAULT '',
    pass           TEXT    NOT NULL        DEFAULT '',
    con_pass       TEXT    NOT NULL        DEFAULT '',
    quota          INTEGER NOT NULL        DEFAULT 0,
    download       INTEGER NOT NULL        DEFAULT 0,
    upload         INTEGER NOT NULL        DEFAULT 0,
    expire_time    INTEGER NOT NULL        DEFAULT 0,
    kick_util_time INTEGER NOT NULL        DEFAULT 0,
    device_no      INTEGER NOT NULL        DEFAULT 3,
    role           TEXT    NOT NULL        DEFAULT 'user',
    deleted        INTEGER NOT NULL        DEFAULT 0,
    create_time    TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    update_time    TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX account_deleted_index ON account (deleted);
CREATE INDEX account_username_index ON account (username);
CREATE INDEX account_con_pass_index ON account (con_pass);
CREATE INDEX account_pass_index ON account (pass);
INSERT INTO account (username, pass, con_pass, quota, download, upload, expire_time, device_no, role)
VALUES ('sysadmin', '02f382b76ca1ab7aa06ab03345c7712fd5b971fb0c0f2aef98bac9cd',
        'ee85cbdbb387a61de3d1d52e6773cc31e28c7913b32f4b7aa44ec61b', -1, 0, 0, 253370736000000, 6, 'admin');
CREATE TABLE config
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    key         TEXT NOT NULL UNIQUE DEFAULT '',
    value       TEXT NOT NULL        DEFAULT '',
    remark      TEXT NOT NULL        DEFAULT '',
    create_time TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX config_key_index ON config (key);
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