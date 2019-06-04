<?php
return [
    'rootPath'=>dirname(__FILE__),
    'imageUpload' => [      //fastdfs  服务器
        'port' => '22122',
        'group' => 'group1',
        'url' => '192.168.1.1',
    ],
    'gearmanService'=>"10.10.10.92:4730,10.10.10.90:4730",  //gearmand Job 服务器名称
    'phantomJsPath'=>'/home/zhangyanchong/web/work/phantomjs/',   //根据服务写正确路径
];
