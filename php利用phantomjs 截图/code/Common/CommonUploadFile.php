<?php

namespace Common;

use Lib\Components\UploadFile;

class CommonUploadFile
{

    //上传文件到远程服务器
    public static function uploadImageFileToRemote($upload_file_name, &$msg='')
    {
        $imageUpload = config("imageUpload");
        $upload = new UploadFile($imageUpload);
        //将临时图片上传到图片服务器
        $file_info = $upload->uploadFile($upload_file_name);//上传主文件
        if (empty($file_info)) {
            $msg = "文件上传失败";
            return false;
        }
        unlink($upload_file_name);//删除临时文件
        return $file_info;
    }

}