<?php
$url="http://www.baidu.com"; //要获取的内容的html页面地址
$phantomJs="";//phantomJs 文件夹的位置
$phanjs=$phantomJs."bin/phantomjs";
$zycjs=phantomJs."examples/zyc.js";
$command="$phanjs   $zycjs     $url ";
$rs=shell_exec($command);