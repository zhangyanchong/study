<?php

namespace   work\screen;

define('ROOT_DIR', realpath(dirname(__FILE__) . '/../') . DIRECTORY_SEPARATOR);//doument root

/*基本文件包含**/
require(ROOT_DIR."config.php");
require(ROOT_DIR."func.php");
require(ROOT_DIR."/autoload.php");
/*基本文件包含**/

/**用的的类*/
use Common\CommonUploadFile;


$gearmanService=config("gearmanService");
$worker=new \GearmanWorker();
//$worker->addServer($gearmanService['ip'],$gearmanService['port']);  //连接到Job server 上

$worker->addServers($gearmanService);  //连接到Job server 上


/**核心方法**/
$worker->addFunction("screenshot",function($job){
    $b_time=microtime(true);  //开始时间

    $filePath=config("rootPath")."/upload/tmp/";
    $msg="";
    $nodeJsPath=config("phantomJsPath");   //phantomJs 路径
    $can=json_decode($job->workload(),true);
    //   echo "<pre>";print_r($can);exit;
    if(!isset($can['url']) || empty($can['url'])){
        $msg="url 不能为空";
        return jsonReturn("201",$msg);
    }
    $url=$can['url'];
//    $httpCode=curlCode($url);
//    if($httpCode!=200){
//        $msg="httpcode:$httpCode 不正确，联系管理员";
//        return jsonReturn("201",$msg);
//    }
    if(!filter_var($url,FILTER_VALIDATE_URL)){
        $msg="url 不正确";
        return jsonReturn("201",$msg);
    }
    if(!isset($can['webWidth']) || empty($can['webWidth'])){
        $msg="webWidth 不能为空";
        return jsonReturn("201",$msg);
    }
    $webWidth=intval($can['webWidth']);
    if($webWidth<100 || $webWidth>3000){
        $msg="webWidth 只能在100 到3000 之间";
        return jsonReturn("201",$msg);
    }
    if(!isset($can['webDomain']) || empty($can['webDomain'])){
        $domain="";
    }else{
        $domain=$can['webDomain'];
    }

    $file=$filePath."zyc_".time().rand(10000,90000).".png";
    if(isset($can['cookie']) || !empty($can['cookie'])){
        $cookie=urlencode($can['cookie']);  //传输的是json
        $command=$nodeJsPath."bin/phantomjs  ".$nodeJsPath."examples/rasterize.js  $url    $file  $webWidth $cookie $domain";
    }else{
        $command=$nodeJsPath."bin/phantomjs  ".$nodeJsPath."examples/rasterize.js  $url    $file  $webWidth";
    }
    $command=escapeshellcmd($command);
    $rs=system ($command,$rs);
    $commonUploadFile=new CommonUploadFile();
    $uploadRs=$commonUploadFile->uploadImageFileToRemote($file,$msg);
    if(!$uploadRs){
        return jsonReturn("201",$msg);
    }
    $uploadRs['fullUrl']="http://".config("imageUpload")['url']."/".config("imageUpload")['group']."/".$uploadRs['path'];

    $e_time=microtime(true);  //技术时间
    $totalTime=number_format($e_time-$b_time, 3);
    $uploadRs['time']=$totalTime;
     file_put_contents(ROOT_DIR."upload/time/".date("Ymd")."time.txt","time:".$totalTime."秒，url:$url".PHP_EOL,FILE_APPEND);
    return jsonReturn("200","",$uploadRs);

});

/**核心方法**/

while ($worker->work()){
    if ($worker->returnCode() !== GEARMAN_SUCCESS)
    {
        echo "Something Wrong" . PHP_EOL;
    }
};









