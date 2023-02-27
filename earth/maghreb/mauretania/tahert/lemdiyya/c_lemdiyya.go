package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LemdiyyaCounty interface {
    feud.County
    BAindefla艾因德夫拉() 	feud.Barony
    BAintorki艾因多吉() 	feud.Barony
    BBlida布利达() 	feud.Barony
    BElhachem哈希姆() 	feud.Barony
    BElrhenb艾里亨布() 	feud.Barony
    BMedea梅德阿() 	feud.Barony
    BRelizane赫利赞() 	feud.Barony
    BSidimohammed西迪穆罕默德() 	feud.Barony
    BTiaret提亚雷特() 	feud.Barony
}

type 提亚雷特LemdiyyaCounty struct {
	feud.BaseCounty
	艾因德夫拉Aindefla 	feud.Barony
	艾因多吉Aintorki 	feud.Barony
	布利达Blida 	feud.Barony
	哈希姆Elhachem 	feud.Barony
	艾里亨布Elrhenb 	feud.Barony
	梅德阿Medea 	feud.Barony
	赫利赞Relizane 	feud.Barony
	西迪穆罕默德Sidimohammed 	feud.Barony
	提亚雷特Tiaret 	feud.Barony
}

func (c *提亚雷特LemdiyyaCounty) BAindefla艾因德夫拉() feud.Barony {
	return c.艾因德夫拉Aindefla
}
    
func (c *提亚雷特LemdiyyaCounty) BAintorki艾因多吉() feud.Barony {
	return c.艾因多吉Aintorki
}
    
func (c *提亚雷特LemdiyyaCounty) BBlida布利达() feud.Barony {
	return c.布利达Blida
}
    
func (c *提亚雷特LemdiyyaCounty) BElhachem哈希姆() feud.Barony {
	return c.哈希姆Elhachem
}
    
func (c *提亚雷特LemdiyyaCounty) BElrhenb艾里亨布() feud.Barony {
	return c.艾里亨布Elrhenb
}
    
func (c *提亚雷特LemdiyyaCounty) BMedea梅德阿() feud.Barony {
	return c.梅德阿Medea
}
    
func (c *提亚雷特LemdiyyaCounty) BRelizane赫利赞() feud.Barony {
	return c.赫利赞Relizane
}
    
func (c *提亚雷特LemdiyyaCounty) BSidimohammed西迪穆罕默德() feud.Barony {
	return c.西迪穆罕默德Sidimohammed
}
    
func (c *提亚雷特LemdiyyaCounty) BTiaret提亚雷特() feud.Barony {
	return c.提亚雷特Tiaret
}
    
var CLemdiyya提亚雷特 LemdiyyaCounty = &提亚雷特LemdiyyaCounty{}

func init() {
	f := CLemdiyya提亚雷特.(*提亚雷特LemdiyyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "830",
		Title:     "lemdiyya",
		TitleName: "提亚雷特",
		TitleCode: "c_lemdiyya",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因德夫拉Aindefla = BAindefla艾因德夫拉
	f.艾因德夫拉Aindefla.SetParent(f)
	
	f.艾因多吉Aintorki = BAintorki艾因多吉
	f.艾因多吉Aintorki.SetParent(f)
	
	f.布利达Blida = BBlida布利达
	f.布利达Blida.SetParent(f)
	
	f.哈希姆Elhachem = BElhachem哈希姆
	f.哈希姆Elhachem.SetParent(f)
	
	f.艾里亨布Elrhenb = BElrhenb艾里亨布
	f.艾里亨布Elrhenb.SetParent(f)
	
	f.梅德阿Medea = BMedea梅德阿
	f.梅德阿Medea.SetParent(f)
	
	f.赫利赞Relizane = BRelizane赫利赞
	f.赫利赞Relizane.SetParent(f)
	
	f.西迪穆罕默德Sidimohammed = BSidimohammed西迪穆罕默德
	f.西迪穆罕默德Sidimohammed.SetParent(f)
	
	f.提亚雷特Tiaret = BTiaret提亚雷特
	f.提亚雷特Tiaret.SetParent(f)
	
}
