<?php
/**
 */
namespace Lib\Components;


use Lib\Components\Fastdfs\Storage;
use Lib\Components\Fastdfs\Tracker;




class UploadFile
{
    private $tracker = null;
    private $storage = null;
    private $fdfsClient = null;
    private $storageInfo = null;

    public function __construct($config)
    {

        $this->tracker = new Tracker($config['url'], $config['port']);
        $this->storageInfo = $this->tracker->applyStorage($config['group']);
        $this->storage = new Storage($config['url'], $this->storageInfo['storage_port']);
        //var_dump($this->storage,$this->storageInfo);die();
    }

    public function uploadFile($file)
    {
       //var_dump($this->storageInfo['storage_index'],$file);die();
        return $this->storage->uploadFile($this->storageInfo['storage_index'],$file);
    }
}