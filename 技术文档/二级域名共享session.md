配置方法一：在php.ini配置文件中设置   session.cookie_domain   的域名值，重启php-fpm  
php
session.cookie_domain = .zixuephp.net  
配置方法二：在php公共代码前设置
php
ini_set('session.cookie_domain','.zixuephp.net');

代码如下：

1.php

	<?php
	ini_set("session.cookie_domain",'.nbsay.cn');//注：此句必须放在session_start()之前
	session_start();
	$_SESSION['abc']=4556;
	echo "<a href='http://my.nbsay.cn/2.php'>tiaozhuan<a>";            

2.php

	<?php
	ini_set("session.cookie_domain",'.nbsay.cn');//注：此句必须放在session_start()之前  
	session_start();
	echo "<pre>";print_r($_SESSION);
	~                              
                                                                                                                                                          
