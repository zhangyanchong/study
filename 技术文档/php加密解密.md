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
?>
```
