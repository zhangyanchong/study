<?php

spl_autoload_register('custom_autoload');
function custom_autoload($classname)
{

	$_file = str_replace('\\', '/', $classname);
	$pathinfo = pathinfo($_file);

	$dir = $pathinfo['dirname'];
	if ($dir[0] == '/') {
		$dir = substr($dir, 1);
	}
 	$include_file = dirname(__FILE__).DIRECTORY_SEPARATOR . $dir . DIRECTORY_SEPARATOR . $pathinfo['basename'] . '.php';
	if(is_file($include_file)){
		include $include_file;
	}

}
