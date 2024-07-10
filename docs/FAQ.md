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

## Meaning of folders in project

- bin: Hysteria2 executable and configuration files
- data: database files
- export: all exported files
- logs: system logs and Hysteria2 operation logs

# Deployment Issues

## Unable to access the panel

- Check if h-ui is running normally
- Check if the firewall allows ports
- Check if the protocol is correct, http:// or https://

## h-ui startup failed

View h-ui logs through the log menu to eliminate the cause of the error

## Hysteria2 startup failed

- Hysteria2 configuration error, view Hysteria2 logs through the log menu to eliminate the cause of the error
- It takes some time to apply for a certificate using ACME for the first time. Please wait patiently for Hysteria2 to start and refresh the page

# Usage Issues

## How to manage certificates?

ACME method (recommended) The system automatically manages. The self-owned certificate method requires manual maintenance and manual replacement after expiration

## How to set the Hysteria2 certificate path for your own certificate?

Upload the certificate file to the server. When configuring Hysteria2, fill in the absolute path of the certificate. When deploying through Docker, you need to upload the certificate file to the volume mapping folder. Recommended: /h-ui/bin

## Why only Hysteria2 version >= v2.4.4 is supported?

Hysteria2 with lower versions does not support the latest API

## Log export failed

No log file