<?php
header("X-Accel-Buffering: no");
for ($i=0; $i<4; $i++)
{
echo $i.'<br />';
ob_flush(); //推出用户缓存
flush(); //推出系统缓存
sleep(1);
}
echo "daa";

exit;
