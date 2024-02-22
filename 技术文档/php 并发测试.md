# php 并发测试

原理：通过循环 php 命令行来执行多个php脚本来实现并发  
1.php 为核心执行方法  
2.php 实现循环执行命令行，一定注意要后台执行否则成顺序执行则测试不准  
通过执行 php 2.php 执行 查看数据或者服务器负载  

```
1.php   执行代码
<?php

$username = trim($argv[1]);
$code = trim($argv[2]);

$loginUrl="http://aaa.com/api/v1/auth/login";
$loginCan=[];
$loginCan['code']=($code);
$loginCan['login_type']=1;
$loginCan['mobile']=($username);
//echo "<pre>";print_r($loginCan);exit;
$rsJson=comCurlPost($loginUrl,$loginCan);
$rs=json_decode($rsJson,true);
$token=$rs['data']['token'];
$areaId=$rs['data']['area_id'];

$bingfaArr=[
    ['url'=>'http://aaa.com/api/v1/home/areadata','can'=>['area_id'=>$areaId,'cacheget_flag'=>1],'type'=>"post"],
    ['url'=>'http://aaa.com/api/v1/home?type=1&cacheget_flag=1','can'=>[],'type'=>"get"],
    ['url'=>'http://hyjgapi.scwljg.com/api/v1/stat/map_trade_all_main?area_id='.$areaId."&is_del_cache=1",'can'=>[],'type'=>"get"],
];
//这里大概概率是并发
$multiHandle = curl_multi_init();
$curlHandles = array();
foreach ($bingfaArr as $k=>$v){
    $data=$v['can'];
   if($v['type']=="post"){
       $curl = curl_init($v['url']);
       curl_setopt($curl, CURLOPT_POST, true);
       curl_setopt($curl, CURLOPT_HTTPHEADER, [
           'Authorization: Bearer '.$token,
       ]);
       curl_setopt($curl, CURLOPT_POSTFIELDS, http_build_query($data)); // 将数据编码为 URL 字符串
       curl_setopt($curl, CURLOPT_RETURNTRANSFER, true); // 将响应保存到变量中
       curl_multi_add_handle($multiHandle, $curl);
       $curlHandles[] = $curl;

   }elseif ($v['type']=="get"){
       $ch = curl_init();
       curl_setopt($ch, CURLOPT_URL, $v['url']);
       curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
       curl_setopt($ch, CURLOPT_HTTPHEADER, [
           'Authorization: Bearer '.$token,
       ]);
       curl_multi_add_handle($multiHandle, $ch);
       $curlHandles[] = $ch;
   }
}

$running = null;
do {
    curl_multi_exec($multiHandle, $running);
    curl_multi_select($multiHandle);
} while ($running > 0);

$responses = array();
foreach ($curlHandles as $ch) {
    $responses[] = curl_multi_getcontent($ch);
    curl_multi_remove_handle($multiHandle, $ch);
    curl_close($ch);
}

curl_multi_close($multiHandle);

$text="-------------$username-------------------------"."\r\n";
$text.=date("Y-m-d H:i:s").json_encode($responses)."\r\n";

file_put_contents("/Users/zhangyanchong/work/d/l_aaa/1.txt",$text,FILE_APPEND);

 function comCurlGet($url, $data_string)
{
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);

    curl_setopt($ch, CURLOPT_HTTPHEADER, array(
        'Content-Type: application/json; charset=utf-8',
    ));
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false); // https请求 不验证证书和hosts
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
    $data = curl_exec($ch);
    curl_close($ch);
    return $data;
}

    function comCurlPost($url, $data)
    {
        // 初始化 cURL
        $curl = curl_init($url);

        curl_setopt($curl, CURLOPT_POST, true);
        curl_setopt($curl, CURLOPT_POSTFIELDS, http_build_query($data)); // 将数据编码为 URL 字符串
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true); // 将响应保存到变量中

        // 发送请求并获取响应
        $response = curl_exec($curl);
        // 检查是否有错误发生
//        if (curl_errno($curl)) {
//            echo "cURL Error: " . curl_error($curl);
//        } else {
//            // 处理响应
//            echo $response;
//        }

        curl_close($curl);
        return $response;
    }

  function comCurlTokenPost($url, $data,$token){
      // 初始化 cURL
      $curl = curl_init($url);

      curl_setopt($curl, CURLOPT_POST, true);
      curl_setopt($curl, CURLOPT_HTTPHEADER, [
          'Authorization: Bearer '.$token,
      ]);
      curl_setopt($curl, CURLOPT_POSTFIELDS, http_build_query($data)); // 将数据编码为 URL 字符串
      curl_setopt($curl, CURLOPT_RETURNTRANSFER, true); // 将响应保存到变量中

      // 发送请求并获取响应
      $response = curl_exec($curl);
      curl_close($curl);
      return $response;
  }

```  


2.php  并发执行php脚本  
```
<?php
$initUser=[
    ['name'=>'13211111111','code'=>'1111'],
    ['name'=>'12311111111','code'=>'1111'],
    ['name'=>'1331111111','code'=>'1111'],
    ['name'=>'1441111111','code'=>'1111'],
];

foreach ($initUser as $v){
    $name=$v['name'];
    $code=$v['code'];
    $text="-------------$name---开始-------------------------"."\r\n";
    file_put_contents("/Users/zhangyanchong/work/d/laradock_web/l_aaa/1.txt",$text,FILE_APPEND);
    exec("php /Users/zhangyanchong/work/d/laradock_web/l_aaa/2.php   $name $code  > /dev/null 2>&1 &");  //一定加上 > /dev/null 2>&1 & 否则不是并发
}
```