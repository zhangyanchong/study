1 php 导出高级的excel 可导出比较复杂的excel

![1](https://github.com/zhangyanchong/php/blob/master/img/1.png)

   可导出合并单元格，颜色，大小 具体用法是 （什么样的都能实现）

2  php 写法
<pre><code>    
  /**
 * 导出文件
 * @return string
 */
public  function  export()
{
$file_name = "成绩单-".date("Y-m-d H:i:s",time());
$file_suffix = "xls";
$data['a']=1;
$data['b']="zhangsan";
header("Content-Type: application/vnd.ms-excel");
header("Content-Disposition: attachment; filename=$file_name.$file_suffix");
//根据业务，自己进行模板赋值。
$this->display("1.php",$data);
}
</code></pre>

3  模板写法 1.php  （只要页面能写出来的excel 样式） 就能导出什么样格式的的excel
	
	<html xmlns:o="urn:schemas-microsoft-com:office:office"
	xmlns:x="urn:schemas-microsoft-com:office:excel"
	xmlns="http://www.w3.org/TR/REC-html40">
	<head>
	<meta http-equiv=Content-Type content="text/html; charset=utf-8">
	<meta name=ProgId content=Excel.Sheet>
	<meta name=Generator content="Microsoft Excel 11">
	</head>
	<body>
	<table border=1 cellpadding=0 cellspacing=0 width="100%" >
	     <tr>
	         <td colspan="5" align="center">
	             <h2>成绩单</h2>
	         </td>
	     </tr>
	     <tr>
	         <td style='width:54pt' align="center">编号</td>
	         <td style='width:54pt' align="center">姓名</td>
	         <td style='width:54pt' align="center">语文</td>
	         <td style='width:54pt' align="center">数学</td>
	         <td style='width:54pt' align="center">英语</td>
	     </tr>
	     <tr>
	        <td align="center">1</td>
	        <td style="background-color: #00CC00;" align="center"><?php $data['a'] ?></td>
	        <td style="background-color: #00adee;" align="center">90</td>
	        <td style="background-color: #00CC00;" align="center">85</td>
	        <td style="background-color: #00adee;" align="center">100</td>
	     </tr>
	    <tr>
	        <td align="center">2</td>
	        <td style="background-color: #00CC00;" align="center">Tom</td>
	        <td style="background-color: #00adee;" align="center">99</td>
	        <td style="background-color: #00CC00;" align="center">85</td>
	        <td style="background-color: #00adee;" align="center">80</td>
	    </tr>
	</table>
	</body>
	</html>







