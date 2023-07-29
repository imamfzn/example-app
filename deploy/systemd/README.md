# Systemd Service

For the first time:

```sh
sudo systemctl daemon-reload
sudo systemctl restart telefun
sudo systemctl enable telefun
```

To see logs:

```sh
sudo journalctl -u telefun.service
```

To stop:
```sh
sudo systemctl stop telefun
```

To start:
```sh
sudo systemctl start telefun
```
