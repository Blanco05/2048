package main

import(
	"fmt"
)

func cancella(str string) string{
	var ris string
	for _, c:=range str{
		if c!='*'{
			ris += string(c)
		}
	}
	return ris
}

/*
func conta_parole(str string) int{
	var count, j int
	for i:=0; i<len(str)-1; i++{
		j = i + 1
		if str[i] != ' ' && str[j] == ' '{
			count++
		}
	}
	return count + 1
}
*/
func main(){
	var str string
	str = "ca**ne**n*e**ro"
	fmt.Println(cancella(str))
}