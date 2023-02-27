package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HajarCounty interface {
    feud.County
    BAlhamra哈姆拉() 	feud.Barony
    BHaran哈兰() 	feud.Barony
    BIbri伊卜里() 	feud.Barony
    BJabrin贾布林() 	feud.Barony
    BMasruq迈斯鲁格() 	feud.Barony
    BSuhar苏哈尔() 	feud.Barony
    BYanqui延古勒() 	feud.Barony
}

type 苏哈尔HajarCounty struct {
	feud.BaseCounty
	哈姆拉Alhamra 	feud.Barony
	哈兰Haran 	feud.Barony
	伊卜里Ibri 	feud.Barony
	贾布林Jabrin 	feud.Barony
	迈斯鲁格Masruq 	feud.Barony
	苏哈尔Suhar 	feud.Barony
	延古勒Yanqui 	feud.Barony
}

func (c *苏哈尔HajarCounty) BAlhamra哈姆拉() feud.Barony {
	return c.哈姆拉Alhamra
}
    
func (c *苏哈尔HajarCounty) BHaran哈兰() feud.Barony {
	return c.哈兰Haran
}
    
func (c *苏哈尔HajarCounty) BIbri伊卜里() feud.Barony {
	return c.伊卜里Ibri
}
    
func (c *苏哈尔HajarCounty) BJabrin贾布林() feud.Barony {
	return c.贾布林Jabrin
}
    
func (c *苏哈尔HajarCounty) BMasruq迈斯鲁格() feud.Barony {
	return c.迈斯鲁格Masruq
}
    
func (c *苏哈尔HajarCounty) BSuhar苏哈尔() feud.Barony {
	return c.苏哈尔Suhar
}
    
func (c *苏哈尔HajarCounty) BYanqui延古勒() feud.Barony {
	return c.延古勒Yanqui
}
    
var CHajar苏哈尔 HajarCounty = &苏哈尔HajarCounty{}

func init() {
	f := CHajar苏哈尔.(*苏哈尔HajarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "869",
		Title:     "hajar",
		TitleName: "苏哈尔",
		TitleCode: "c_hajar",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈姆拉Alhamra = BAlhamra哈姆拉
	f.哈姆拉Alhamra.SetParent(f)
	
	f.哈兰Haran = BHaran哈兰
	f.哈兰Haran.SetParent(f)
	
	f.伊卜里Ibri = BIbri伊卜里
	f.伊卜里Ibri.SetParent(f)
	
	f.贾布林Jabrin = BJabrin贾布林
	f.贾布林Jabrin.SetParent(f)
	
	f.迈斯鲁格Masruq = BMasruq迈斯鲁格
	f.迈斯鲁格Masruq.SetParent(f)
	
	f.苏哈尔Suhar = BSuhar苏哈尔
	f.苏哈尔Suhar.SetParent(f)
	
	f.延古勒Yanqui = BYanqui延古勒
	f.延古勒Yanqui.SetParent(f)
	
}
