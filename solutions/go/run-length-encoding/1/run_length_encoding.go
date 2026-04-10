package runlengthencoding

import (
    "strconv"
    "unicode"
)

func RunLengthEncode(input string) string {
	decodeStr := ""
	count := 1
  
    if len(input) == 0 { return input }
    
    for i, _ := range input {
        if len(input) !=  i+1 && input[i] == input[i+1] {
            count += 1
        }else if count == 1  {
            decodeStr = decodeStr + string(input[i])
        }else {            
            decodeStr = decodeStr + strconv.Itoa(count) + string(input[i])            
            count = 1
        }
    }

    return decodeStr
}

func RunLengthDecode(input string) string {
    encodedStr := ""
    digit := ""
	for i, ch := range input {         
        if unicode.IsDigit(ch) {
        	digit = digit + string(ch)        	
        }else if i!=0 && (input[i-1] >= '0' && input[i-1] <= '9'){
            num, _ := strconv.Atoi(digit)
            	for j:=0; j < num; j++ {
                	encodedStr = encodedStr + string(ch)
            	}
        		digit = ""
        }else {
            encodedStr += string(ch)
        }
    }
    return encodedStr
}
