

	Bit-map的基本思想就是用一个bit位来标记某个元素对应的Value，而Key即是该元素。由于采用了Bit为单位来存储数据，因此在存储空间方面，可以大大节省。（PS：划重点 节省存储空间）
	
	 Bitmap有什么用
	
	大量数据的快速排序、查找、去重



	<?php
	
	/*
	 *适用于64的操作系统
	 * https://blog.csdn.net/zh7314/article/details/115652185
	 * **/
	final class BitMap {
	
	    //int 位数
	    private static $phpIntSize = PHP_INT_SIZE;
	    //int最大值  Usually int(2147483647) in 32 bit systems and int(9223372036854775807) in 64 bit systems. Available since PHP 5.0.5
	    private static $phpIntMax = PHP_INT_MAX;
	    //最大位数64位 1 << 6 32位 1 << 5
	   // private static $max = (1 << 6 ) - 1;
	    private static $max =63;
	    //储存数据的变量
	    private static $data = [];
	
	    public static function addValue($n) {
	        //商   $a >> $b	Shift right（右移）	将 $a 中的位向右移动 $b 次（每一次移动都表示“除以 2”）
	        $row = $n >> 6;
	        //余数
	        $index = $n % self::$max;
	      //  echo (1 << $index); exit;
	        //或运算保证占位不被覆盖   $a << $b	Shift left（左移）	将 $a 中的位向左移动 $b 次（每一次移动都表示“乘以 2”）。
	        self::$data[$row] |= 1 << $index;
	    }
	

               /*
           * bit map 删除值
           * **/
            public  static  function delValue($n){
                $row = $n >> 6;
                //余数
                $index = $n % self::$max;
                //  echo (1 << $index); exit;
                //或运算保证占位不被覆盖   $a << $b	Shift left（左移）	将 $a 中的位向左移动 $b 次（每一次移动都表示“乘以 2”）。
                if(self::exits($n)){
                    @self::$data[$row]^= 1 << $index;
                }
            }
        
	    // 判断所在的bit为是否为1
	    public static function exits($n) {
	        $row = $n >> 6;
	        $index = $n % (self::$max);
	
	        $result = self::$data[$row] & (1 << $index);
	//        p($result);
	        return $result != 0;
	    }
	
	    public static function getData() {
	        return self::$data;
	    }
	
	    //输出数组并且排序
	  public  static  function outPut($bitmap){
	        $int_bit_size = self::$max;
	        $result =[];
	        foreach ($bitmap as $k => $item) {
	            for ($i = 0; $i < $int_bit_size; $i++) {
	                $temp = 1 << $i;
	                $flag = $temp & $item;
	                if ($flag) {
	                    $result[] = $k * $int_bit_size + $i;
	                }
	            }
	        }
	        return $result;
	    }
	}
	
	error_reporting(1);
	ini_set("display_errors",1);
	
	
	
	$arr = [0, 1, 3, 16,50, 42, 69, 18, 11, 99,12313123, 32421, 32423, 32525,500,565166962];
	echo "<pre>";print_r($arr);
	foreach ($arr as $v) {
	    BitMap::addValue($v);   //添加值
	}
	
	$tt = BitMap::getData();  //获取添加的bitmap值
	echo "<pre>";print_r($tt);
	$xin=BitMap::outPut($tt);   //获取最原始的数组
	echo  "<pre>";print_r($xin);
	
	
	$rr = BitMap::exits(501);  //判断是否存在
	
	if ($rr) {
	    echo ('ok');
	} else {
	    echo ('no');
	}
	


