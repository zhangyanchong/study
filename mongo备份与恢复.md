 1.mongo 分片 备份  和  恢复

		 备份指定库   
		
		 mongodump --host 127.0.0.1 --port 27000  -d zyc -o /zyc/mongo_share/tmp  
		 它会在/tmp/目录下面生成一个mydb的目录   
		 
		 备份所有库   
		 mongodump --host 127.0.0.1 --port 20000 -o /tmp/mongobak/alldatabase   
		 
		 指定备份集合   
		 mongodump --host 127.0.0.1 --port 20000 -d mydb -c c1 -o /tmp/mongobak/   
		 它依然会生成mydb目录，再在这目录下面生成两个文件   
		
		导出集合为json文件   
		mongoexport --host 127.0.0.1 --port 20000 -d mydb -c c1 -o /tmp/mydb2/1.json 


		
		恢复所有库   
		mongorestore -h 127.0.0.1 --port 20000 --drop dir/ //其中dir是备份所有库的目录名字，其中--drop可选，意思是当恢复之前先把之前的数据删除，不建议使用   
		
		恢复指定库   mongorestore -d zyc /zyc/mongo_share/tmp/zyc  //-d跟要恢复的库名字，dir就是该库备份时所在的目录  
		
		恢复集合  mongorestore -d mydb -c testc dir/mydb/testc.bson // -c后面跟要恢复的集合名字，dir是备份mydb库时生成文件所在路径，这里是一个bson文件的路径   
		
		导入集合   mongoimport -d mydb -c testc --file /tmp/testc.json  

      恢复带密码的指定库 
      /mnt/mongo_share/mongodb42/bin/mongorestore   -h 127.0.0.1:27017  -d calorie_log /mnt/bak_mongo_data/20210127/calorie_log   -u 用户名 -p 密码  --authenticationDatabase admin
