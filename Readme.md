# ssh-host-manager 

.ssh/configをパースしてIPを一覧表示してくれるようにする


# 将来的な機能
- [x] 一覧表示
- [ ] 追加・削除


# install
```
git clone https://github.com/h16110s/ssh-host-manager.git
cd ssh-host-manager
sudo make install
```

# command

## list
```zsh
ssh-hosts
```

```zsh
|      HOST      |    HOSTNAME     |
|----------------|-----------------|
| *              |                 |
| mac            | 192.168.10.10   |
| raspi          | 192.168.0.20    |
| desktop        | 192.168.0.30    |
| hlab           | 192.168.0.30    |
| github.com     | github.com      |
| github.com.sub | github.com      |
| dns            | 192.168.12.10   |
| mg01           | 192.168.12.11   |
| host01         | 192.168.12.12   |
| host02         | 192.168.12.13   |
| memo-server    | 192.168.11.19   |
| gcp            | 192.168.100.100 |
| gcp2           | 192.168.100.101 |
```

