1.mongdb 版本  --4.2 版本    （bind_ip 只能绑定一个ip 我草 无情）
2. 分片服务 （副本集）  mongo3.4 以后必须的    
   配置服务器（副本集） mongo 3.4 以后必须的  
   routes 路由 mongos   （可以一台）      
   什么是副本集，自己百度一下  （就是有mongo主服务器，多个从服务器 当某个从服务器挂了 主服务器会从服务器找个运行正常的使用，如果主服务器挂了 就从从服务器找个作为主服务器） 一般为3个一个副本集   
   每个副本集请至少保证剩余2台机器，否则不能用，如果多台机器挂掉后剩余2台后，不能再挂掉，否则不能使用 （<2 台 副本集就不能用了 ）   （自己测试）
   
 3.安装  暂时没有使用默认端口 （只在一台做实验）   
 
	  routes         
	    127.0.0.1     27000  
	  配置服务器 （一个副本集）   config_set  
	    127.0.0.1     27001  
	    127.0.0.1     27002  
	    127.0.0.1     27003   
	  分片节点    （2个副本集） share1  
	    127.0.0.1     27004  
	    127.0.0.1     27005   
	    127.0.0.1     27006  

	    127.0.0.1     27007   share2  
	    127.0.0.1     27008  
	    127.0.0.1     27009  
  4.创建文件夹  
  data  数据文件下目录    

  		config1 config2 config3 mongos  share1_1 share1_2 share1_3  share2_1   share2_2  share2_3 
  log  日志  
  config 配置文件  
  pid  文件夹  
 5.先启动配置服务器  副本集

    
		config1.conf
		logpath=/zyc/mongo_share/log/config1.log
		pidfilepath=/zyc/mongo_share/pid/config1.pid
		logappend=true
		port=27001  
		fork=true
		dbpath=/zyc/mongo_share/data/config1/
		configsvr=true   # 在配置文件添加此项就行
		oplogSize=512
		replSet=config_set
		
		
		config2.conf
		logpath=/zyc/mongo_share/log/config2.log
		pidfilepath=/zyc/mongo_share/pid/config2.pid
		logappend=true
		port=27002  
		fork=true
		dbpath=/zyc/mongo_share/data/config2/
		configsvr=true   # 在配置文件添加此项就行
		oplogSize=512
		replSet=config_set
		
		config3.conf
		logpath=/zyc/mongo_share/log/config3.log
		pidfilepath=/zyc/mongo_share/pid/config3.pid
		logappend=true
		port=27003  
		fork=true
		dbpath=/zyc/mongo_share/data/config3/
		configsvr=true   # 在配置文件添加此项就行
		oplogSize=512
		replSet=config_set
		
		
		
		 ./mongod  -f   /zyc/mongo_share/config/config1.conf
		 ./mongod  -f   /zyc/mongo_share/config/config2.conf
		  ./mongod -f   /zyc/mongo_share/config/config3.conf
		
		
		
		mongo 操作
		./mongo --port 27001
		use admin
		
		config = { _id:"config_set",members:[ {_id:0,host:"127.0.0.1:27001"}, {_id:1,host:"127.0.0.1:27002"}, {_id:2,host:"127.0.0.1:27003"}] }        #定义副本集
		
		rs.initiate(config) 



 6  启动mongos 路由服务器


	    ./mongos  -f /zyc/mongo_share/config/mongos.conf
	
		mongos.conf  其中的config_set 跟上面的匹配
		logpath=/zyc/mongo_share/log/mongos.log
		pidfilepath=/zyc/mongo_share/pid/mongos.pid
		logappend=true
		port=27000
		fork=true
		configdb=config_set/127.0.0.1:27001,127.0.0.1:27002,127.0.0.1:27003


7.分片副本集配置   2个副本集

	  第一组分片
	   share1_1.conf
	logpath=/zyc/mongo_share/log/share1_1.log
	pidfilepath=/zyc/mongo_share/pid/share1.pid
	directoryperdb=true
	logappend=true
	port=27004
	fork=true
	dbpath=/zyc/mongo_share/data/share1_1/
	oplogSize=512
	replSet=share1
	shardsvr=true
	   
	   
	  share1_2.conf
	logpath=/zyc/mongo_share/log/share1_2.log
	pidfilepath=/zyc/mongo_share/pid/share2.pid
	directoryperdb=true
	logappend=true
	port=27005
	fork=true
	dbpath=/zyc/mongo_share/data/share1_2/
	oplogSize=512
	replSet=share1
	shardsvr=true
	
	
	  share1_3.conf
	logpath=/zyc/mongo_share/log/share1_3.log
	pidfilepath=/zyc/mongo_share/pid/share1_3.pid
	directoryperdb=true
	logappend=true
	port=27006
	fork=true
	dbpath=/zyc/mongo_share/data/share1_3/
	oplogSize=512
	replSet=share1
	shardsvr=true
	
	
	
	
	 ./mongod  -f   /zyc/mongo_share/config/share1_1.conf
	 ./mongod  -f   /zyc/mongo_share/config/share1_2.conf
	 ./mongod -f   /zyc/mongo_share/config/share1_3.conf   
	
	
	 ./mongo --port 27004
	 
	 use admin
	
	 shard = {_id:"share1",members:[{_id:0,host:"127.0.0.1:27004"},{_id:1,host:"127.0.0.1:27005"}, {_id:2,host:"127.0.0.1:27006"}]}
	
	 rs.initiate(shard)



	 第二组分片
	
	   share2_1.conf
	logpath=/zyc/mongo_share/log/share2_1.log
	pidfilepath=/zyc/mongo_share/pid/share2_1.pid
	directoryperdb=true
	logappend=true
	port=27007
	fork=true
	dbpath=/zyc/mongo_share/data/share2_1/
	oplogSize=512
	replSet=share2
	shardsvr=true
	   
	   
	  share2.conf
	logpath=/zyc/mongo_share/log/share2_2.log
	pidfilepath=/zyc/mongo_share/pid/share2_2.pid
	directoryperdb=true
	logappend=true
	port=27008
	fork=true
	dbpath=/zyc/mongo_share/data/share2_2/
	oplogSize=512
	replSet=share2
	shardsvr=true
	
	
	  share3.conf
	logpath=/zyc/mongo_share/log/share2_3.log
	pidfilepath=/zyc/mongo_share/pid/share2_3.pid
	directoryperdb=true
	logappend=true
	port=27009
	fork=true
	dbpath=/zyc/mongo_share/data/share2_3/
	oplogSize=512
	replSet=share2
	shardsvr=true
	
	
	
	
	 ./mongod  -f   /zyc/mongo_share/config/share2_1.conf
	 ./mongod  -f   /zyc/mongo_share/config/share2_2.conf
	 ./mongod -f   /zyc/mongo_share/config/share2_3.conf   
	
	
	 ./mongo --port 27007
	 
	 use admin
	
	 shard = {_id:"share2",members:[{_id:0,host:"127.0.0.1:27007"},{_id:1,host:"127.0.0.1:27008"}, {_id:2,host:"127.0.0.1:27009"}]}
	
	 rs.initiate(shard)



8.启动mongos 路由数据库 添加分片集
	 ./mongo --port 27000
	 
	 sh.addShard("share1/127.0.0.1:27004,127.0.0.1:27005,127.0.0.1:27006");
	 sh.addShard("share2/127.0.0.1:27007,127.0.0.1:27008,127.0.0.1:27009");
	
	
	 sh.status()  
	   查看shards  判断是否成功
	
	
	
	-个集合只有且只能有一个分片键，一旦分片键确定好之后就不能更改  （重要呀）
	  8.1 hash分片 指定zyc数据库开始分片  
	   sh.enableSharding("zyc")  需要分片的数据库       
	   sh.shardCollection("zyc.nihao",{"id":"hashed"})     //已hash的方式分片  最主要的用这个 
	
		for(let i=10000;i<90000;i++){
		    db.nihao.insert({"name":'name'+i,"id":i});
		}
	
		//查看分片的情况
		db.nihao.getShardDistribution()
	
	
	 8.2 range 分片  有个默认块 一般大于64M 才会均衡数据
	
	 测试  改变一下块大小否则数据太小无法看出来
	  use config
	  db.settings.save({_id: "chunksize", value: 2})
	  db.settings.find({_id: "chunksize"})
	
	
	  sh.enableSharding("zyc")  需要分片的数据库  
	  sh.shardCollection("zyc.shenghuo",{"userid":1})     
	
	  for (i= 10000; i <=200000; i++){
	   db.shenghuo.insert({age:(i%100), name:"user"+i,userid:i})
	  } 
	
	  //查看分片的情况
		db.nihao.getShardDistribution()




