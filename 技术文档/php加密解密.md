```php
<?php
echo $str="123123";
$key="!@#abc";
echo "<br>";
echo  $val=jiami($str,$key);
echo "<br>";
echo jiemi($val,$key);

function jiami($val,$str){
    $iv = "123456781234567812345678";  //跟解密一样  可以写到缓存中 这个是变化的才会生成变化的加密字符串
    $result = openssl_encrypt($val, "aes-128-cbc", $str, true,$iv);
    return base64_encode($result);
}

function jiemi($val,$str){
    $iv = "123456781234567812345678";  //跟加密一样 可以写到缓存中
    $val=base64_decode($val);
     $result = openssl_decrypt($val, "aes-128-cbc", $str, true,$iv);
    var_dump($result);
    return $result;
}




class Encrypt {
    protect  $pass= md5(base64_encode("123abc456"));
	public static function encryptAes($data = '') {
	    

		if(function_exists('openssl_encrypt')) {
			return openssl_encrypt($data, "aes128", $this->pass);
		} else {
			//TODO
		}
		
	}

	public static function decryptAes($data = '') {
	
		if(function_exists('openssl_decrypt')) {
			return openssl_decrypt($data, "aes128", $this->pass);
		}
		
	}
}

?>
```