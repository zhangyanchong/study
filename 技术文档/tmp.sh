#!/bin/bash

function rand(){
    min=$1
    max=$(($2-$min+1))
    num=$(cat /proc/sys/kernel/random/uuid | cksum | awk -F ' ' '{print $1}')
    echo $(($num%$max+$min))
}


  for file in `ls  /zyc/tmp/*`
    do
        if test -f $file
        then
            rnd=$(rand 11 20)
           #echo "file $rnd:  $file"
          cmd="sed -i '${rnd}d' $file"
           #echo $cmd
          eval $cmd
        fi
    done


