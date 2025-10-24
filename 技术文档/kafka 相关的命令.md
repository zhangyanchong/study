# kafka 相关的命令

``` 
首先需要安装一套kafka 需要的命令 查询百度
 /*
 *10.1.1.1  ip 地址 39092 端口号
 * aaa_mid topic 名字
 * aaa_group  消费组名 
 */
 1.查看某个topic 是否存在  

     kafka-topics.sh --bootstrap-server 10.1.1.1:39092 --list | grep aaa_mid
     	•	有输出 → 存在
    	•	没输出 → 不存在
2.查看topic 是否有数据写入
     watch -n 5 -t "date '+%H:%M:%S'; ./kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list 10.1.1.1:39092 --topic aaa_mid --time -1"	
     查看数值是否增长 有增长就是有数据写入
3.查看topic 是否有数据消费
     watch -n 5 "./kafka-consumer-groups.sh --bootstrap-server 10.1.1.1:39092 --describe --group aaa_group"
     CURRENT-OFFSET   总过多少个
     LOG-END-OFFSET   消费到哪一个
     LAG              还剩多少没消费
 
4.查看topic最近的的一条内容  只是查看 
    首先 看看  watch -n 5 -t "date '+%H:%M:%S'; ./kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list 10.1.1.1:39092 --topic aaa_mid --time -1"	
    看看到哪里了 假如 0 分区到 5919234
    可以去 5919234-1 为 offset
      /*
      *partition 哪个分区
      *offset 从哪个开始  
      * max-messages
      **/
      ./kafka-console-consumer.sh  --bootstrap-server 10.1.1.1:39092 --topic aaa_mid --partition 0 --offset 591923 --max-messages 1
5. 在某个topic 添加一条数据  （暂时无法指定分区）
    kafka.tools.GetOffsetShell --broker-list 10.1.1.1:39092 --topic aaa_mid --time -1
    查看一下 每个分区多少条  记录一下
    ./kafka-console-producer.sh   --broker-list 10.1.1.1:39092   --topic aaa_mid
    然后回车 输入个json
    {"id":102,"name":"xin"}
    
    //再看一下 哪个分区 多了记录
    kafka.tools.GetOffsetShell --broker-list 10.1.1.1:39092 --topic aaa_mid --time -1
    
    //在根据 上面的4  查看数据 
    ./kafka-console-consumer.sh  --bootstrap-server 10.1.1.1:39092 --topic aaa_mid --partition 0 --offset 591923 --max-messages 1
    
6.在某个topic 消费一条数据
  6.1 查看 这个组 有没有消费  如果组不存在 可以先执行6.3
./kafka-consumer-groups.sh --bootstrap-server 10.1.1.1:39092  --describe --group tmp_group
  6.2 添加一条数据 
   ./kafka-console-producer.sh   --broker-list 10.1.1.1:39092  --topic aaa_mid
   {"id":110,"name":"哈多"}
 6.3 创建一个临时的 topic 组 消费 aaa_mid 消费一条记录退出
    ./kafka-console-consumer.sh --bootstrap-server 10.1.1.1:39092 --topic aaa_mid --group tmp_group --max-messages 1  
       
```