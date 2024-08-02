# lavarel   自己创建的简单异步 第一版

```
 异步实现东西
   例如 队列 TmpTaskA  TmpTaskB   2个队列
   1. 能同时执行队列  TmpTaskA  TmpTaskB 里面的任务 （根据创建时间倒排的）
   2. TmpTaskA 任务报错 不影响其 TmpTaskB 里面的任务
   3. TmpTaskA TmpTaskB 任务修改代码会 几秒后生效 不用启动相关服务
   4. 每种类型的任务，能在（几秒内）同时并发（10个可配置）任务
   5. 如果任务报错会记录在错误日志（tb_queue_job_fail）表里面 
---------------------------------------------------------------------------
 1.扩展 php 需要安装swoole 扩展  
 2. lavarel  composer.json 添加    
    "swooletw/laravel-swoole": "^2.13",         
      "ext-swoole": "*"
 3.守护相关进程
  php swoole_server.php >> swoole_server.log 2>&1 &   错误输出到文件用于查看    
  
4.文件介绍 在 lavarel 核心app目录下面创建 Task目录
    1.其中目录下面Core  TaskBase.php 为核心文件 
    2.其他的为使用文件  
        TmpTaskB.php
        TmpTaskA.php
5. 使用 相关代码 
   5.1在Task 目录下创建 文件 TmpTaskB.php   handle($args) 方法为执行方法 $args 参数数组
   5.2 其他文件里面写入 
          $base=new  TaskBase();
          $tmpData=[];
          $tmpData['id']=$check_id;
          $base->createTask("TmpTaskB",$tmpData);   
          //TmpTaskB 为文件 及其队列名称  $tmpData 为参数
          
6.相关代码介绍 
  1.入口文件  swoole_server.php
 ```
    <?php
    
    /*
     * 需要安装一下swoole  php 建议7.4 或者更高  7.3 sleep 有问题
     * 需要在程序执行 php swoole_server.php   守护
     * php swoole_server.php >> swoole_server.log 2>&1 &   错误输出到文件用于查看
     * 使用： 在App\Task  建立自己的类  一个类相当于一种类型
     * **/
    
    use Swoole\Process;
    
    set_time_limit(0);
    ini_set("memory_limit", '3080m');
    //因为有多进程  所以协成必须关闭掉 Timer 默认是协成 估计是版本问题
    ini_set('swoole.enable_coroutine', 'Off');
    
    require __DIR__ . '/vendor/autoload.php';
    
    
    $app = require_once __DIR__ . '/bootstrap/app.php';
    $kernel = $app->make(Illuminate\Contracts\Console\Kernel::class);
    $kernel->bootstrap();
    
    // 定时器处理任务
    Swoole\Timer::tick(2000, function () {
       // echo "start"; echo  "\r\n";
        require_once __DIR__ . '/app/Task/Core/TaskBase.php';
        $scheduler = new \App\Task\Core\TaskBase();
        $scheduler->dealTask();
    });
    
    // 等待所有子进程结束
    //while (($status = Process::wait()) !== false) {
    //    // 处理子进程状态
    //}
    
    // 添加信号处理以回收子进程
    Process::signal(SIGCHLD, function () {
        while ($ret = Process::wait(false)) {
            echo "Process with PID {$ret['pid']} exited with code {$ret['code']}\n";
        }
    });
    
    //不加这个会一直报错
    Swoole\Event::wait();

 ```  
 核心文件TaskBase.php
 ```
     <?php
    
    namespace App\Task\Core;
    
    use App\Models\QueueJob;
    use App\Models\QueueJobFail;
    use Swoole\Coroutine;
    use Swoole\Process;
    use Illuminate\Support\Facades\DB;
    use PDO;
    class  TaskBase
    {
        /*
         * $fileName 即为队列名称
         * $param  参数
         * 一个文件 一个队列  队列名称跟文件名的一样的 不需要重新写
         * 默认执行 文件的handle 方法
         * ***/
        /*
         *生成任务
         * **/
        function createTask($fileName, $param = [])
        {
            $insert = [];
            $insert['queue_file'] = $fileName;
            $insert['param_json'] = json_encode($param);
            $insert['created_at'] = $insert['updated_at'] = date("Y-m-d H:i:s");
            $id = QueueJob::insertGetId($insert);
            return $id;
        }
    
        /*
         * 核心方法处理任务
         * **/
        function dealTask()
        {
            //要处理的任务
            $taskList = $this->getTask();
    //        echo  "11";echo "33";
    //        echo "<pre>";print_r($taskList);
            if (empty($taskList) || count($taskList)<=0) {
                return true;
            }
    //        echo  json_encode($taskList); echo "\r\n";
    
    //        $taskData=[];  //进程
            foreach ($taskList as $k => $v) {
                $process = new Process(function (Process $process) use ($v) {
                    // 注册一个关闭函数来处理致命错误  必须写在最前面
                    register_shutdown_function(function () use ($v) {
                        $error = error_get_last();
                        if ($error) {
                            $taskId = $v['id'];
                            $queueFile = $v['queue_file'];
                            //错误信息
                            $this->errorInfoToDb($taskId,$queueFile,$error['message']);
                        }
                    });
                    set_time_limit(0);
                    // 在每个子进程中创建新的数据库连接 (否则会报错)
                    DB::reconnect();
                    DB::connection()->getPdo()->setAttribute(PDO::ATTR_TIMEOUT,7200); // 设置数据库连接的超时时间
                    $parameters = json_decode($v['param_json'], true);
                    //方法和类
                    $methodName = "handle";
                    $namespace = "App\Task\\";
                    $queueFile = $v['queue_file'];
                    $tmpObj = $namespace . $queueFile;
                    $taskId = $v['id'];
    
                    //文件不存在或者 类不对
                    if (class_exists($tmpObj)==false){
                        $this->errorInfoToDb($taskId,$queueFile,"类文件不正确,请看一下队列是否和文件名字一样");
                        return;
                    }
    
                    $handler = new  $tmpObj;
    
                    //设置超时时间 默认就是无限期执行
    //                $timeout = 3600*2; // 设定超时时间为10秒
    //                // 设置一个定时器，超时后退出进程
    //                Swoole\Timer::after($timeout * 1000, function() use ($process) {
    //                    echo "Process {$process->pid} timed out\n";
    //                    $process->exit(1); // 退出子进程，返回状态码 1
    //                });
    
                    //设置内存大小 不受限制
                    ini_set("memory_limit", '-1');
    
    
                    if (method_exists($handler, $methodName)) {
                        // 更新任务状态为正在处理
                        $this->taskStatusChange($taskId, 2);
                        try {
                            // call_user_func_array([$handler, $methodName], $parameters);
                            // 将参数作为一个数组传递
                            $handler->$methodName($parameters);
    
                            // 更新任务状态为已完成
                            $this->taskStatusChange($taskId, 3);
                        } catch (\Exception $e) {
                            echo "Caught exception in task for task ID $taskId: " . $e->getMessage() . "\n";
                            //错误信息
                            $this->errorInfoToDb($taskId,$queueFile,$e->getMessage());
                        }
    
                        //  $this->runningTasks[$taskId] = true;
                    } else {   //存入错入
                        // echo "Method $methodName does not exist for task type $taskType\n";
                    }
                });
                $process->start();
            }
        }
    
        function  errorInfoToDb($taskId,$queueFile,$content){
            // 更新任务状态为失败
            $this->taskStatusChange($taskId, 4);
            $insert = [];
            $insert['queue_file'] = $queueFile;
            $insert['job_id'] = $taskId;
            $insert['content_data'] = $content;
            $insert['created_at'] = $insert['updated_at'] = date("Y-m-d H:i:s");
            $id=QueueJobFail::insertGetId($insert);
            return $id;
        }
    
        /*
         * 更新任务状态
         * **/
        function taskStatusChange($taskId, $status)
        {
            $update = [];
            $update['updated_at'] = date("Y-m-d H:i:s");
            $update['status'] = $status;
            QueueJob::where("id", $taskId)->update($update);
            return true;
        }
    
        /*
         * 获取除了 结束 和进行中
         * 获取当前任务中没种类型中的一个
         * 每个任务最多执行2个
         * 重试3次数（--）
         * ***/
        function getTask()
        {
            $tryNum=3;
            $taskTypeLimit=10;
            $data = [];
            /*获取当前有多少种类型的任务 (不包括正在进行中的)**/
            $queueNameIds = QueueJob::whereIn("status", [1])->where('try_num',"<=",$tryNum)->pluck("queue_file")->toArray();
            $queueNameIds=array_unique($queueNameIds);
            if (!empty($queueNameIds)) {
                foreach ($queueNameIds as $queueFile) {
                    //超过2个正在进行  直接跳过
                    $tmpCount=QueueJob::whereIn("status", [2])->where("queue_file", $queueFile)->count();
                    if($tmpCount>=$taskTypeLimit){
                        continue;
                    }
                    $queueJobData = QueueJob::whereIn("status", [1])->where("queue_file", $queueFile)->where('try_num',"<=",$tryNum)->orderBy("created_at", "desc")->first();
                    if (!empty($queueJobData)){
                        $data[] = $queueJobData;
                    }
                }
            }
            return $data;
        }
    
    }
 ```
 TmpTaskA.php 文件
```
    <?php
    namespace App\Task;
    
    use App\Http\Utils\FileLog;
    use App\Models\Company;
    
    class TmpTaskA{
    
        function  handle($args){
           $data=[];
           $data['company_address']="公司地址".$args['id'];
            Company::where("id",$args['id'])->update($data);
        }
        function bbb(){
    
        }
    } 
``` 
 

   
       
 核心表 
 核心表 tb_queue_job   任务表
 CREATE TABLE `tb_queue_job` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `try_num` tinyint(1) DEFAULT '0' COMMENT '尝试次数',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `queue_file` varchar(255) DEFAULT '' COMMENT '队列名字也是文件名字',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态 1 默认 2进行中   3已结束  4.报错重试种',
  `param_json` varchar(255) DEFAULT '' COMMENT '参数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8; 

 队列报错任务表
CREATE TABLE `tb_queue_job_fail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `job_id` int(11) DEFAULT NULL COMMENT '任务id',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `content_data` longtext COMMENT '异常信息',
  `queue_file` varchar(255) DEFAULT '' COMMENT '文件名字',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='队列失败表';   




```  
