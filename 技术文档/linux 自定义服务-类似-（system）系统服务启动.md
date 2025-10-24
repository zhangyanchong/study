linux 自定义服务-类似-系统服务启动

```
1.在 /etc/systemd/system/ 路径下创建  zyc.service 文件
内容
# /etc/systemd/system/webserver.service
[Unit]
Description=data_ingestion HTTP Server
After=network.target

[Service]
Type=simple
WorkingDirectory=/home/v-zyc/web/
ExecStart=/home/v-zyc/web/main -f /home/v-zyc/web/.env.yml
Restart=always

[Install]
WantedBy=multi-user.target


===================================================
Description 服务的说明文字
After=network.target 等网络启动后再启动服务
Type=simple 主进程就是这个命令本身
WorkingDirectory 程序运行目录
ExecStart 启动命令
Restart=always 程序退出后自动重启
WantedBy=multi-user.target 随系统开机自启（正常服务器模式）

相关命令

启动服务
systemctl    start  zyc
 重启
 systemctl restart zyc

停止服务
sudo systemctl stop zyc

查看状态
sudo systemctl status zyc

查看日志
sudo journalctl -u zyc -f
```