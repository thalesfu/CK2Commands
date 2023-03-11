package el-arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type El-arishCounty interface {
    feud.County
    BAbuaweigila阿布阿维吉拉() 	feud.Barony
    BArish阿里什() 	feud.Barony
    BBirelhamma哈迈井() 	feud.Barony
    BKharruba哈鲁巴() 	feud.Barony
    BMasyada马察达() 	feud.Barony
    BMitmatna米特马纳() 	feud.Barony
    BTukkot图科特() 	feud.Barony
    BZuwayid祖瓦德() 	feud.Barony
}

type 阿里什El-arishCounty struct {
	feud.BaseCounty
	阿布阿维吉拉Abuaweigila 	feud.Barony
	阿里什Arish 	feud.Barony
	哈迈井Birelhamma 	feud.Barony
	哈鲁巴Kharruba 	feud.Barony
	马察达Masyada 	feud.Barony
	米特马纳Mitmatna 	feud.Barony
	图科特Tukkot 	feud.Barony
	祖瓦德Zuwayid 	feud.Barony
}

func (c *阿里什El-arishCounty) BAbuaweigila阿布阿维吉拉() feud.Barony {
	return c.阿布阿维吉拉Abuaweigila
}
    
func (c *阿里什El-arishCounty) BArish阿里什() feud.Barony {
	return c.阿里什Arish
}
    
func (c *阿里什El-arishCounty) BBirelhamma哈迈井() feud.Barony {
	return c.哈迈井Birelhamma
}
    
func (c *阿里什El-arishCounty) BKharruba哈鲁巴() feud.Barony {
	return c.哈鲁巴Kharruba
}
    
func (c *阿里什El-arishCounty) BMasyada马察达() feud.Barony {
	return c.马察达Masyada
}
    
func (c *阿里什El-arishCounty) BMitmatna米特马纳() feud.Barony {
	return c.米特马纳Mitmatna
}
    
func (c *阿里什El-arishCounty) BTukkot图科特() feud.Barony {
	return c.图科特Tukkot
}
    
func (c *阿里什El-arishCounty) BZuwayid祖瓦德() feud.Barony {
	return c.祖瓦德Zuwayid
}
    
var CEl-arish阿里什 El-arishCounty = &阿里什El-arishCounty{}

func init() {
	f := CEl-arish阿里什.(*阿里什El-arishCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "787",
		Title:     "el-arish",
		TitleName: "阿里什",
		TitleCode: "c_el-arish",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布阿维吉拉Abuaweigila = BAbuaweigila阿布阿维吉拉
	f.阿布阿维吉拉Abuaweigila.SetParent(f)
	
	f.阿里什Arish = BArish阿里什
	f.阿里什Arish.SetParent(f)
	
	f.哈迈井Birelhamma = BBirelhamma哈迈井
	f.哈迈井Birelhamma.SetParent(f)
	
	f.哈鲁巴Kharruba = BKharruba哈鲁巴
	f.哈鲁巴Kharruba.SetParent(f)
	
	f.马察达Masyada = BMasyada马察达
	f.马察达Masyada.SetParent(f)
	
	f.米特马纳Mitmatna = BMitmatna米特马纳
	f.米特马纳Mitmatna.SetParent(f)
	
	f.图科特Tukkot = BTukkot图科特
	f.图科特Tukkot.SetParent(f)
	
	f.祖瓦德Zuwayid = BZuwayid祖瓦德
	f.祖瓦德Zuwayid.SetParent(f)
	
}
