package S

// String support package
import (
	"encoding/json"
	"strconv"
	"strings"

	"crypto/sha256"
	"encoding/base64"
	"regexp"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
)

// check whether the input string (first arg) starts with a certain character (second arg) or not.
//  S.StartsWith(`adakah`,`ad`) // bool(true)
//  S.StartsWith(`adakah`,`bad`) // bool(false)
func StartsWith(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// check whether the input string (first arg) ends with a certain character (second arg) or not.
//  S.EndsWith(`adakah`,`ah`)) // bool(true)
//  S.EndsWith(`adakah`,`aka`)) // bool(false)
func EndsWith(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// check whether the input string (first arg) contains a certain sub string (second arg) or not.
//  S.Contains(`komputer`,`om`)) // bool(true)
//  S.Contains(`komputer`,`opu`)) // bool(false)
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// compare two input string (first arg) equal with another input string (second arg).
//  S.Equals(`komputer`,`komputer`)) // bool(true)
//  S.Equals(`komputer`,`Komputer`)) // bool(false)
func Equals(strFirst, strSecond string) bool {
	return strFirst == strSecond
}

// compare two input string (first arg) equal with ignoring case another input string (second arg).
//  S.EqualsIgnoreCase(`komputer`,`komputer`)) // bool(true)
//  S.EqualsIgnoreCase(`komputer`,`Komputer`)) // bool(true)
func EqualsIgnoreCase(strFirst, strSecond string) bool {
	strFirst = ToLower(strFirst)
	strSecond = ToLower(strSecond)
	return strFirst == strSecond
}

// count how many specific character (first arg) that the string (second arg) contains
//  S.Count(`komputeer`,`e`))// output int(2)
func Count(str, substr string) int {
	return strings.Count(str, substr)
}

// erase spaces from left and right
//  S.Trim(` withtrim:  `) // `withtrim:`
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// replace all substring with another substring
//  S.Replace(`bisa`,`is`,`us`) // `busa`
func Replace(haystack, needle, gold string) string {
	return strings.Replace(haystack, needle, gold, -1)
}

// change the characters in string to lowercase
//  S.ToLower(`BIsa`) // "bisa"
func ToLower(str string) string {
	return strings.ToLower(str)
}

// change the characters in string to uppercase
// S.ToUpper(`bisa`) // "BISA"
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// get character at specific index, utf-8 safe
//  S.CharAt(`Halo 世界`, 5) // `世` // utf-8 example, if characters not shown, it's probably your font/editor/plugin
//  S.CharAt(`Halo`, 3) // `o`
func CharAt(str string, index int) string {
	for in, ch := range str {
		if in == index {
			return string(ch)
		}
	}
	return ``
}

// remove character at specific index, utf-8 safe
//  S.RemoveCharAt(`Halo 世界`, 5) // `Halo 界` --> utf-8 example, if characters not shown, it's probably your font/editor/plugin
//  S.RemoveCharAt(`Halo`, 3) // `Hal`
func RemoveCharAt(str string, index int) string {
	var chars []byte
	for in, ch := range str {
		if index != in {
			chars = append(chars, string(ch)...)
		}
	}
	return string(chars)
}

// Change first letter for every word to uppercase
//  S.ToTitle(`Disa dasi`)) // output "Disa Dasi"
func ToTitle(str string) string {
	return strings.Title(str)
}

// simplified ternary operator (bool ? val : 0), returns second argument, if the condition (first arg) is true, returns empty string if not
//  S.If(true,`a`) // `a`
//  S.If(false,`a`) // ``
func If(b bool, yes string) string {
	if b {
		return yes
	}
	return ``
}

// ternary operator (bool ? val1 : val2), returns second argument if the condition (first arg) is true, third argument if not
//  S.IfElse(true,`a`,`b`) // `a`
//  S.IfElse(false,`a`,`b`) // `b`
func IfElse(b bool, yes, no string) string {
	if b {
		return yes
	}
	return no
}

// coalesce, return first non-empty string
//  S.IfEmpty(``,`2`) // `2`
//  S.IfEmpty(`1`,`2`) // `1`
func IfEmpty(str1, str2 string) string {
	if str1 != `` {
		return str1
	}
	return str2
}

// coalesce, return first non-empty string
//  S.Coalesce(`1`,`2`) // `1`
//  S.Coalesce(``,`2`) // `2`
//  S.Coalesce(``,``,`3`) // `3`
func Coalesce(strs ...string) string {
	for _, str := range strs {
		if str != `` {
			return str
		}
	}
	return ``
}

// convert string to int64, returns 0 and silently print error if not valid
//  S.ToI(`1234`) // 1234
//  S.ToI(`1a`) // 0
func ToI(str string) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	L.IsError(err, str)
	return val
}

// convert to int64 with check
//  S.AsI(`1234`) // 1234, true
//  S.AsI(`1abc`) // 0, false
func AsI(str string) (int64, bool) {
	res, err := strconv.ParseInt(str, 10, 64)
	return res, err == nil
}

// convert string to float64, returns 0 and silently print error if not valid
//  S.ToF(`1234.5`) // 1234.5
//  S.ToF(`1a`) // 0.0
func ToF(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	L.IsError(err, str)
	return val
}

// convert to float64 with check
//  S.AsF(`1234.5`) // 1234.5, true
//  S.AsF(`1abc`) // 0.0, false
func AsF(str string) (float64, bool) {
	res, err := strconv.ParseFloat(str, 64)
	return res, err == nil
}

// convert JSON object to map[string]interface{}, silently print and return empty map if failed
//  json_str := `{"test":123,"bla":[1,2,3,4]}`
//  map1 := S.JsonToMap(json_str)
func JsonToMap(str string) (res map[string]interface{}) {
	res = map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &res)
	L.IsError(err, str)
	return
}

// convert JSON object to []interface{}, silently print and return empty slice of interface if failed
//  json_str := `[1,2,['test'],'a']`
//  arr := S.JsonToArr(json_str)
func JsonToArr(str string) (res []interface{}) {
	res = []interface{}{}
	err := json.Unmarshal([]byte(str), &res)
	L.IsError(err, str)
	return
}

// convert JSON object to map[string]interface{} with check
//  json_str := `{"test":123,"bla":[1,2,3,4]}`
//  map1, ok := S.JsonAsMap(json_str)
func JsonAsMap(str string) (res map[string]interface{}, ok bool) {
	res = map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &res)
	ok = err != nil
	return
}

// convert JSON object to []interface{} with check
//  json_str := `[1,2,['test'],'a']`
//  arr, ok := S.JsonAsArr(json_str)
func JsonAsArr(str string) (res []interface{}, ok bool) {
	res = []interface{}{}
	err := json.Unmarshal([]byte(str), &res)
	ok = err != nil
	return
}

// split a string (first arg) by characters (second arg) into array of strings (output).
//  S.Split(`biiiissssa`,`i`) // output []string{"b", "", "", "", "ssssa"}
func Split(str, sep string) []string {
	return strings.Split(str, sep)
}

// append padStr to left until length is lenStr
func PadLeft(s string, padStr string, lenStr int) string {
	var padCount int
	padCount = I.MaxOf(lenStr-len(s), 0)
	return strings.Repeat(padStr, padCount) + s
}

// append padStr to right until length is lenStr
func PadRight(s string, padStr string, lenStr int) string {
	var padCount int
	padCount = I.MaxOf(lenStr-len(s), 0)
	return s + strings.Repeat(padStr, padCount)
}

// hash password with sha256 (without salt)
func HashPassword(pass string) string {
	res1 := []byte(pass)
	res2 := sha256.Sum256(res1)
	res3 := res2[:]
	return base64.StdEncoding.EncodeToString(res3)
}

// return empty string if str is not a valid email
func ValidateEmail(str string) string {
	res := strings.Split(str, `@`)
	if len(res) != 2 {
		return ``
	}
	if (strings.Trim(res[0], `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&'*+-/=?^_{|}~.`)) == `` {
		if strings.Trim(res[1], `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-.`) == `` {
			return str
		}
	}
	return ``
}

var re_non_phone = regexp.MustCompile(`[^-+\d,]+`)

// remove invalid characters of a phone number
func ValidatePhone(str string) string {
	return re_non_phone.ReplaceAllString(str, ``)
}
