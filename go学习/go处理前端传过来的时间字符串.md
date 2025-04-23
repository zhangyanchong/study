
```
// 数据库连接信息
	dsn := "root:7$X5OO%guEzt5f@tcp(localhost:3306)/huo?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:7$X5OO%guEzt5f@tcp(localhost:3306)/huo?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db.Exec("SET time_zone = '+08:00'")
	// 假设你要更新的 device_code 的 id 是 1
	// 假设前端传来的 expireTime 字符串
	frontendExpireTime := "2011-11-11 11:11:11"

	// 解析时间时显式指定时区为本地时区
	loc, err := time.LoadLocation("Asia/Shanghai") // 指定时区
	if err != nil {
		fmt.Println("时区加载失败:", err)
		return
	}

	// 将字符串解析为 time.Time
	expireTime, err := time.ParseInLocation("2006-01-02 15:04:05", frontendExpireTime, loc)
	if err != nil {
		fmt.Println("时间解析错误:", err)
		return
	}

	// 假设你要更新的 device_code 的 id 是 76
	deviceCode := DeviceCode{
		Code:       "new_code_123",
		ExpireTime: expireTime, // 使用解析后的 expireTime
		Num:        10,         // 使用时长 10 天
		UpdatedAt:  time.Now(),
	}

	// 更新操作
	if err := db.Model(&DeviceCode{}).Where("id = ?", 76).Updates(deviceCode).Error; err != nil {
		panic("failed to update record")
	}

	fmt.Println("成功")
	
	
	*******
	   平常我们gorm  链接数据库都会加入 写 loc=Local 
	   但是我们前端传过来的是字符串的时间 例如  2011-11-11 11:11:11 入库会自动加8小时  （sql 也对 就是入库的时候不对）
	   这个因为go的time包导致的 需要处理一下
	   // 解析时间时显式指定时区为本地时区
    	loc, err := time.LoadLocation("Asia/Shanghai") // 指定时区
    	if err != nil {
    		fmt.Println("时区加载失败:", err)
    		return
    	}
    
    	// 将字符串解析为 time.Time
    	expireTime, err := time.ParseInLocation("2006-01-02 15:04:05", frontendExpireTime, loc)
    	if err != nil {
    		fmt.Println("时间解析错误:", err)
    		return
    	}

	******* 
 
```
