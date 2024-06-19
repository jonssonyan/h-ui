# FAQ

## Command Line

Give executable permissions

```bash
chmod +x ./h-ui
```

- Check the system version

  ```bash
  ./h-ui -v
  ```

- Start with a custom web port

  ```bash
  ./h-ui -p [port]
  ```

# Deployment Issues

## Meaning of folders in a project

- bin: Hysteria2 executable and configuration files

- data: database files
- export: all exported files
- logs: system logs and Hysteria2 operation logs

## Startup failed

View log files

1. View error description in the command line

2. View logs/h-ui.log

# Usage Issues

## How to manage certificates?

ACME method (recommended) The system automatically manages. For self-owned certificates, you need to maintain them
manually and replace them manually after expiration.

## Own certificate, set Hysteria2 certificate path

Upload the certificate file to the server, and fill in the absolute path of the certificate when configuring Hysteria2

## Why only Hysteria2 version >= v2.4.4 is supported?

Lower versions do not support the latest API

## Log export failed

No log file