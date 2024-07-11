CREATE TABLE IF NOT EXISTS account
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
CREATE INDEX IF NOT EXISTS account_deleted_index ON account (deleted);
CREATE INDEX IF NOT EXISTS account_username_index ON account (username);
CREATE INDEX IF NOT EXISTS account_con_pass_index ON account (con_pass);
CREATE INDEX IF NOT EXISTS account_pass_index ON account (pass);
INSERT INTO account (id, username, pass, con_pass, quota, download, upload, expire_time, device_no, role)
SELECT 1 ,'sysadmin', '02f382b76ca1ab7aa06ab03345c7712fd5b971fb0c0f2aef98bac9cd', 'sysadmin.sysadmin', -1, 0, 0, 253370736000000, 6, 'admin'
    WHERE NOT EXISTS (SELECT 1 FROM account WHERE id = 1);
CREATE TABLE IF NOT EXISTS config
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    key         TEXT NOT NULL UNIQUE DEFAULT '',
    value       TEXT NOT NULL        DEFAULT '',
    remark      TEXT NOT NULL        DEFAULT '',
    create_time TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS config_key_index ON config (key);
INSERT INTO config (key, value, remark)
SELECT 'H_UI_WEB_PORT', '8081', 'H UI Web Port'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'H_UI_WEB_PORT');
INSERT INTO config (key, value, remark)
SELECT 'H_UI_CRT_PATH', '', 'H UI Crt File Path'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'H_UI_CRT_PATH');
INSERT INTO config (key, value, remark)
SELECT 'H_UI_KEY_PATH', '', 'H UI Key File Path'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'H_UI_KEY_PATH');
INSERT INTO config (key, value, remark)
SELECT 'JWT_SECRET', hex(randomblob(10)), 'JWT Secret'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'JWT_SECRET');
INSERT INTO config (key, value, remark)
SELECT 'HYSTERIA2_ENABLE', '0', 'Hysteria2 Switch'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'HYSTERIA2_ENABLE');
INSERT INTO config (key, value, remark)
SELECT 'HYSTERIA2_CONFIG', '', 'Hysteria2 Config'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'HYSTERIA2_CONFIG');
INSERT INTO config (key, value, remark)
SELECT 'HYSTERIA2_TRAFFIC_TIME', '1', 'Hysteria2 Traffic Time'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'HYSTERIA2_TRAFFIC_TIME');
INSERT INTO config (key, value, remark)
SELECT 'HYSTERIA2_CONFIG_REMARK', '', 'Hysteria2 Config Remark'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'HYSTERIA2_CONFIG_REMARK');
INSERT INTO config (key, value, remark)
SELECT 'HYSTERIA2_CONFIG_PORT_HOPPING', '', 'Hysteria2 Config Port Hopping'
    WHERE NOT EXISTS (SELECT 1 FROM config WHERE key = 'HYSTERIA2_CONFIG_PORT_HOPPING');