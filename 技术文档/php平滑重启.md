  php平衡重启

	PHP
	查看到php-fpm的master进程号后，使用信号控制重启php-fpm：
	
	INT，TERM：立刻终止  
	QUIT：平滑终止  
	USR1：重新打开日志文件  
	USR2：平滑重载所有worker进程并重新载入配置和二进制模块    
	
	1264 主进程号
	根据信号控制重启php-fpm执行命令：kill -USR2 1264，即可重启php-fpm。
