package main

import (
	"fmt"
	"strconv"
	"strings"
)


func RemoveIndex(s []string, index int)[]string {
	return append(s[:index], s[index+1:]...)
}



func RemoveDuplicate(arr []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range arr {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func ManagePunc(s []string, index int) []string {
	if s[index] == "," || s[index] == "." || s[index] == ";" || s[index] == "!" || s[index] == "?" || s[index] == ":" || s[index] =="..." || s[index] == "!?" {
				ponctuation := s[index]
					s[index-1] = s[index-1]+ponctuation
			RemoveIndex(s, index)
}
return s
}


func GetUpNumFunc(str string, s []string, index int) []string  {
	//"(up,2)"
	getStrElt := (string(str[4]))
	getNumber, err := strconv.Atoi(getStrElt)

	if err != nil{
		panic(err)
	}
	
 if s[index] == str {
	 for i := 1; i <= getNumber; i++{
		lowerCase := s[index - i]
			upperCase := strings.ToUpper(lowerCase)
			s[index - i] = upperCase
	 }
 }
  RemoveIndex(s, index)
	return s
}

func GetCapNumFunc(str string, s []string, index int) []string  {
//"(cap,5)"
	getStrElt := (string(str[5]))
	getNumber, err := strconv.Atoi(getStrElt)

	if err != nil{
		panic(err)
	}
	
 if s[index] == str {
	 for i := 1; i <= getNumber; i++{
		st := s[index-i]
			cap := strings.ToUpper(st[:1]) + st[1:]
			s[index-i] = cap
	 }
 }
	RemoveIndex(s, index)
	return s
}


func GetLowNumFunc(str string, s []string, index int) []string  {
//"(low,5)" ,
	getStrElt := (string(str[5]))
	getNumber, err := strconv.Atoi(getStrElt)

	if err != nil{
		panic(err)
	}
	
 if s[index] == str {
	 for i := 1; i <= getNumber; i++{
		upperCase := s[index - i]
		lowerCase := strings.ToLower(upperCase)
		s[index - i] = lowerCase
	 }
 }
	RemoveIndex(s, index)
	ManagePunc(s, index)
	return s
}

func HexToDeci(s []string, index int)[]string {

	hexadecimal_num := s[index-1]
	decimal_num , err := strconv.ParseInt(hexadecimal_num, 16, 64)

	if err != nil {
		panic(err)
	}

	s[index-1] = strconv.Itoa(int(decimal_num))
	RemoveIndex(s, index)
	ManagePunc(s, index)
	return s
}



func BinFunc(s []string, index int) []string {

	deximal_num_Str := s[index-1]
	deximal_num, err := strconv.Atoi(deximal_num_Str)

	if err != nil {
		panic(err)
	}

	if deximal_num == 10 {
		s[index-1] = strconv.Itoa(2)
	}
	RemoveIndex(s, index)
	ManagePunc(s, index)
	return s
}

func UpFunc(s []string, index int)[]string{

	lowerCase := s[index - 1]
	upperCase := strings.ToUpper(lowerCase)
	s[index - 1] = upperCase
	RemoveIndex(s, index)
	ManagePunc(s, index)
	return s
}

func LowFunc(s []string, index int)[]string{

	upperCase := s[index - 1]
	lowerCase := strings.ToLower(upperCase)
	s[index - 1] = lowerCase
  RemoveIndex(s, index)
	ManagePunc(s, index)
	return s
}

func CapFunc(s []string, index int)[]string{

	str := s[index-1]
	cap := strings.ToUpper(str[:1]) + str[1:]
	s[index-1] = cap
	RemoveIndex(s, index)
	ManagePunc(s, index)
	return s
}

func AddNFunc(s []string, index int)[]string{
	if (s[index] == "a" || s[index] == "A") && (strings.HasPrefix(s[index+1], "a") == true || strings.HasPrefix(s[index+1], "e") == true || strings.HasPrefix(s[index+1], "i") == true || strings.HasPrefix(s[index+1], "o") == true || strings.HasPrefix(s[index+1], "u") == true || strings.HasPrefix(s[index+1], "h") == true) {
				s[index] = s[index]+"n"
			}

			return s
}

func PuncFunc(s []string, index int) []string{
	if s[index] == "," || s[index] == "." || s[index] == ";" || s[index] == "!" || s[index] == "?" || s[index] == ":" || s[index] =="..." || s[index] == "!?" {
				ponctuation := s[index] // "(low)" ,
				if strings.HasPrefix(s[index-1], "(") == true{
          s[index-2] = s[index-2]+ponctuation
					RemoveIndex(s, index)
				} else {
					s[index-1] = s[index-1]+ponctuation
				}
			RemoveIndex(s,index)
		}  else if strings.HasPrefix(s[index], ".") == true || strings.HasPrefix(s[index], ",") == true || strings.HasPrefix(s[index], ";") == true || strings.HasPrefix(s[index], "!") == true || strings.HasPrefix(s[index], "?") == true || strings.HasPrefix(s[index], ":") == true || strings.HasPrefix(s[index], "!?") == true || strings.HasPrefix(s[index], "...") == true {
			if strings.HasPrefix(s[index-1], "(") == true {
      s[index - 2] = s[index - 2]+s[index][:1]
			} else {
				s[index - 1] = s[index - 1]+s[index][:1]
			}
			s[index] = s[index][1:]
		}
		return s
}

func ManageAp(s []string, index int)[]string{
	if s[index] =="'" && s[index+2] =="'" {
			s[index+1] = s[index]+s[index+1]+s[index+2]
			RemoveIndex(s, index)
			RemoveIndex(s, index+1)
		} else if s[index] == "'" && s[index+2] != "'" {
			for j := index+3; j <= int(len(s) -1); j++ {
				if s[j] == "'" {
					s[index+1] = s[index]+s[index+1]
					s[j-1] = s[j-1]+s[j]
					RemoveIndex(s,index)
					RemoveIndex(s, j-1)
				}
			}
		}
		return s
}

func main() {
	str := "1E (hex)  files were added .It has been 10 (bin) years .Ready , set , go (up) ! I should STOP SHOUTING (low) ... Welcome to the Brooklyn bridge (cap) . This is so exciting (up) . Bonjour will be exactly what  you are saying . Hi (up,3) i am new ' in ' this school. It just a test to check (cap) something. Punctuation tests are ... kinda boring ,don't you think !? Là ' je suis entrain ' de test !? A amazing rock! "
	s := strings.Split(str, " ")
	fmt.Println(s)

	for i :=0 ; i <= int(len(s) -1); i++{
		
		if s[i] == "(hex)" {

			HexToDeci(s, i)

		} else if s[i]=="(bin)" {

			BinFunc(s, i)

		} else if s[i] == "(up)" {

			UpFunc(s, i)

		} else if s[i] == "(low)" {

			LowFunc(s, i)

		} else if s[i] == "(cap)" {

			CapFunc(s, i)
     //"(up,2)"
		} else if int(len(s[i])) == 6 && string(s[i][1]) == "u" && string(s[i][2]) == "p" {

			GetUpNumFunc(s[i], s, i)

		} else if int(len(s[i])) == 7 && string(s[i][1]) == "c" && string(s[i][2]) == "a" && string(s[i][3]) =="p" {

			GetCapNumFunc(s[i], s, i)

		} else if int(len(s[i])) == 7 && string(s[i][1]) == "l" && string(s[i][2]) == "o" && string(s[i][3]) =="w" {

			GetLowNumFunc(s[i], s, i)

		} 
		
		PuncFunc(s, i)
		
		AddNFunc(s, i)

		// ' tet ' -> 'tet'  ' trt ezrz troro ' -> 'trt ezrz troro'
		ManageAp(s, i)
	}
	//CheckLastElt(s)

	 Result := strings.Join(RemoveDuplicate(s), " ")
	fmt.Println(Result)
}