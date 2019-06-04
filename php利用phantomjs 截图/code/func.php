<?php
/*
 * 通用function
 * **/
function config($name=""){
     $data=require(dirname(__FILE__)."/config.php");
    if (empty($name)) {
        return $data;
    }
    return $data[$name];
}

/*数据返回**/
function jsonReturn($status=200,$msg="",$data=[]){
     $tmp['status']=$status;
     $tmp['message']=$msg;
     $tmp['data']=$data;
     return  json_encode($tmp);
}

/*
 *只获取页面状态码
 * **/
function  curlCode($url){
    $ch = curl_init($url);
    curl_setopt($ch, CURLOPT_HEADER, true);    // we want headers
    curl_setopt($ch, CURLOPT_NOBODY, true);    // we don't need body
    curl_setopt($ch, CURLOPT_RETURNTRANSFER,1);
    curl_setopt($ch, CURLOPT_TIMEOUT,10);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
    curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/4.0 (compatible; MSIE 5.01; Windows NT    5.0');
    $output = curl_exec($ch);
    $httpcode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);
    return $httpcode;
}

