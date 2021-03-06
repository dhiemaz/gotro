package S

import (
	"github.com/kokizzu/gotro/B"
	"github.com/kokizzu/gotro/I"
)

// add single quote in the beginning and the end of string.
//  S.Q(`coba`) // `'coba'`
//  S.Q(`123`)  // `'123'`
func Q(str string) string {
	return `'` + str + `'`
}

// replace ` and give double quote (for table names)
//  S.ZZ(`coba"`) // `"coba&quot;"`
func ZZ(str string) string {
	str = Trim(str)
	str = Replace(str, `"`, `&quot;`)
	return `"` + str + `"`
}

// give ' to boolean value
//  S.ZB(true)  // `'true'`
//  S.ZB(false) // `'false'`
func ZB(b bool) string {
	return `'` + B.ToS(b) + `'`
}

// give ' to int64 value
//  S.ZI(23)) // '23'
//  S.ZI(03)) // '3'
func ZI(num int64) string {
	return `'` + I.ToS(num) + `'`
}

// double quote a json string
//  hai := `{'test':123,"bla":[1,2,3,4]}`
//  S.ZJ(hai) // "{'test':123,\"bla\":[1,2,3,4]}"
func ZJ(str string) string {
	str = Replace(str, `\`, `\\`)
	str = Replace(str, "\r", `\r`)
	str = Replace(str, "\n", `\n`)
	str = Replace(str, `"`, `\"`)
	return `"` + str + `"`
}

// trim, replace <, >, ', " and gives single quote
//  S.Z(`<>'"`) // `&lt;&gt;&apos;&quot;
func Z(str string) string {
	str = Trim(str)
	str = Replace(str, `<`, `&lt;`)
	str = Replace(str, `>`, `&gt;`)
	str = Replace(str, `'`, `&apos;`)
	str = Replace(str, `"`, `&quot;`)
	str = Replace(str, `\`, `\\`)
	return `'` + str + `'`
}

// replace <, >, ', " and gives single quote (without trimming)
//  S.Z(`<>'"`) // `&lt;&gt;&apos;&quot;
func ZS(str string) string {
	str = Replace(str, `<`, `&lt;`)
	str = Replace(str, `>`, `&gt;`)
	str = Replace(str, `'`, `&apos;`)
	str = Replace(str, `"`, `&quot;`)
	str = Replace(str, `\`, `\\`)
	return `'` + str + `'`
}

// replace <, >, ', ", % and gives single quote and %
//  S.ZLIKE(`coba<`))  // output '%coba&lt;%'
//  S.ZLIKE(`"coba"`)) // output '%&quot;coba&quot;%'
func ZLIKE(str string) string {
	str = Trim(str)
	str = Replace(str, `<`, `&lt;`)
	str = Replace(str, `>`, `&gt;`)
	str = Replace(str, `'`, `&apos;`)
	str = Replace(str, `"`, `&quot;`)
	str = Replace(str, `\`, `\\`)
	str = Replace(str, `%`, `\%`)
	return `'%` + str + `%'`
}

// replace <, >, ', ", % but without giving single quote
func XSS(str string) string {
	str = Trim(str)
	str = Replace(str, `<`, `&lt;`)
	str = Replace(str, `>`, `&gt;`)
	str = Replace(str, `'`, `&apos;`)
	str = Replace(str, `"`, `&quot;`)
	return str
}
