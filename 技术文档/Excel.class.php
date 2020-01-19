<?php 

/*****************************************/
/*  excel 导入导出Excel操作类(PHPExcel类) *
    @author xfan
    @date 2018/04/28				      *
/*****************************************/

class Excel {
	/*
     * 写入excel 并下载
     * @datas 导出的数据
     * 例如：array(
     *     0=>array('aaa','男','152'),
     *     1=>array('aaa','男','152'),
     * )
     * @filename 文件名称(只是名称)
     * $head  导出excel 头部标题  例如： array('姓名','性别','年龄')
     * 
     */
    public static function export_excel($datas, $head, $filename='data') {
        require 'PHPExcel/PHPExcel.php';
        require 'PHPExcel/PHPExcel/Writer/Excel2007.php';

        $objPHPExcel = new \PHPExcel();
        $objPHPExcel->getProperties()->setCreator("Phpmarker")->setLastModifiedBy("Phpmarker")->setTitle("Phpmarker")->setSubject("Phpmarker")->setDescription("Phpmarker")->setKeywords("Phpmarker")->setCategory("Phpmarker");
        // Set active sheet index to the first sheet, so Excel opens this as the first sheet
        $objPHPExcel->setActiveSheetIndex(0);  
        $objActSheet = $objPHPExcel->getActiveSheet();

        //默认列
        $column = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z', 'AA', 'AB', 'AC', 'AD', 'AE', 'AF', 'AG', 'AH'];
        //设置head
        $count = count($head);
        for ($i=0; $i < $count; $i++) { 
            $key = $column[$i].'1'; 
            $objActSheet->setCellValue($key, $head[$i]);
        }  
        // Rename worksheet
        $objActSheet->setTitle('Phpmarker-' . date('Y-m-d')); 
        
        $objActSheet->getDefaultRowDimension()->setRowHeight(15);
        $objActSheet->freezePane('A2');	
        $i = 2;
        foreach($datas as $data){ 
            for ($j=0; $j < $count; $j++) { 
                $objActSheet->setCellValue($column[$j] . $i, $data[$j]);
                $objActSheet->setCellValueExplicit($column[$j]. $i, $data[$j],\PHPExcel_Cell_DataType::TYPE_STRING);
                $objActSheet->getStyle($column[$j] . $i)->getAlignment()->setHorizontal(\PHPExcel_Style_Alignment::VERTICAL_CENTER);
                $objActSheet->getStyle($column[$j] . $i)->getNumberFormat()->setFormatCode("@");
                $objActSheet->getStyle($column[$j] . $i)->getAlignment()->setWrapText(true); 
            } 
	        $i++;
        } 
        
        ob_end_clean();//清除缓冲区,避免乱码 
        header('Content-Type: application/vnd.ms-excel');
        header('Content-Disposition: attachment;filename="'.$filename.'.xlsx"');
        header('Cache-Control: max-age=0');
        // If you're serving to IE 9, then the following may be needed
        header('Cache-Control: max-age=1'); 
        // If you're serving to IE over SSL, then the following may be needed
        header('Expires: Mon, 26 Jul 1997 05:00:00 GMT'); // Date in the past
        header('Last-Modified: ' . gmdate('D, d M Y H:i:s') . ' GMT'); // always modified
        header('Cache-Control: cache, must-revalidate'); // HTTP/1.1
        header('Pragma: public'); // HTTP/1.0

        $objWriter = new \PHPExcel_Writer_Excel5($objPHPExcel);
        $objWriter->save('php://output'); 
    }

    /*
     * 读取excel数据并返回
     * @file  文件路径
     * @endColumn  最后列字母标识(读到该列停止读取)
     */
    public static function read_excel($fileDir, $startRow=1, $endColumn=24) {
        require 'PHPExcel/PHPExcel.php';
        require 'PHPExcel/PHPExcel/IOFactory.php';
        
        $inputFileType = \PHPExcel_IOFactory::identify($fileDir);
        $objReader = \PHPExcel_IOFactory::createReader($inputFileType);  
        $objReader->setReadDataOnly(true);//只读去数据，忽略里面各种格式等
        $objPHPExcel = $objReader->load($fileDir);
        $objActiveSheet = $objPHPExcel->getActiveSheet();//获取当前活动sheet
        $totalRow = $objActiveSheet->getHighestRow();//总行数
        $totalColumn = $objActiveSheet->getHighestColumn();
        //默认列
        $arr = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z', 'AA', 'AB', 'AC', 'AD', 'AE', 'AF', 'AG', 'AH'];
         
        $data = array();

        $endColumn = strtoupper(trim($endColumn));
        for ($row=$startRow; $row <= $totalRow; $row++) { 
        	for ($column=0; $column <= $endColumn; $column++) { 
        		$data[$row-1][$column] = $objActiveSheet->getCellByColumnAndRow($column, $row)->getValue();
        	}
        }

        return $data;
    } 
}

