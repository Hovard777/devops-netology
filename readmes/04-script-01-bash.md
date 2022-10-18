1. 
``` 
a=1
b=2
c=a+b
d=$a+$b
e=$(($a+$b))
```
> c = "a+b" - так как указали текст а не переменные  
d = "1+3" - команда преобразовала вывела значения переменных, но не выполнила сложение, так как по умолчанию это строки  
e = "3"   - за счет скобок мы дали команду на выполнение арифметической операции со значениями переменных 

2. 
``` 
 while (( 1 == 1 )) #add )
    do
        curl https://localhost:4757
        if (($? != 0))
        then
            date >> curl.log
        else exit #add exit condition
        fi
        sleep 10 #add timeout
    done
``` 

3. 
``` 
hosts=(192.168.0.1 173.194.222.113 87.250.250.24)
for i in {1..5}
do
date >>hosts.log
    for h in ${hosts[@]}
    do
	curl -Is $h:80 >/dev/null
        echo "    check" $h status=$? >>hosts.log
        sleep 5
    done
done
```

4. 
```
hosts=(192.168.0.1 173.194.222.113 87.250.250.24)
stat=0
While (($stat==0))
do
date >>hosts.log
    for h in ${hosts[@]}
    do
	curl -Is $h:80 >/dev/null
        echo "    check" $h status=$? >>hosts.log
        stat=$?
        
        if (($stat != 0))
        then
            $h >> error.log
        fi
        sleep 5
    done
done
```
