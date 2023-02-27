package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AprutiumCounty interface {
    feud.County
    BAprutium_penne佩内() 	feud.Barony
    BAquila拉奎拉() 	feud.Barony
    BAtri阿特里() 	feud.Barony
    BAvezzano阿韦扎诺() 	feud.Barony
    BPaganica帕格尼卡() 	feud.Barony
    BPescara佩斯卡拉() 	feud.Barony
    BTeramo泰拉莫() 	feud.Barony
}

type 拉奎拉AprutiumCounty struct {
	feud.BaseCounty
	佩内Aprutium_penne 	feud.Barony
	拉奎拉Aquila 	feud.Barony
	阿特里Atri 	feud.Barony
	阿韦扎诺Avezzano 	feud.Barony
	帕格尼卡Paganica 	feud.Barony
	佩斯卡拉Pescara 	feud.Barony
	泰拉莫Teramo 	feud.Barony
}

func (c *拉奎拉AprutiumCounty) BAprutium_penne佩内() feud.Barony {
	return c.佩内Aprutium_penne
}
    
func (c *拉奎拉AprutiumCounty) BAquila拉奎拉() feud.Barony {
	return c.拉奎拉Aquila
}
    
func (c *拉奎拉AprutiumCounty) BAtri阿特里() feud.Barony {
	return c.阿特里Atri
}
    
func (c *拉奎拉AprutiumCounty) BAvezzano阿韦扎诺() feud.Barony {
	return c.阿韦扎诺Avezzano
}
    
func (c *拉奎拉AprutiumCounty) BPaganica帕格尼卡() feud.Barony {
	return c.帕格尼卡Paganica
}
    
func (c *拉奎拉AprutiumCounty) BPescara佩斯卡拉() feud.Barony {
	return c.佩斯卡拉Pescara
}
    
func (c *拉奎拉AprutiumCounty) BTeramo泰拉莫() feud.Barony {
	return c.泰拉莫Teramo
}
    
var CAprutium拉奎拉 AprutiumCounty = &拉奎拉AprutiumCounty{}

func init() {
	f := CAprutium拉奎拉.(*拉奎拉AprutiumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "929",
		Title:     "aprutium",
		TitleName: "拉奎拉",
		TitleCode: "c_aprutium",
		Baronies:  map[string]feud.Barony{},
	}

	f.佩内Aprutium_penne = BAprutium_penne佩内
	f.佩内Aprutium_penne.SetParent(f)
	
	f.拉奎拉Aquila = BAquila拉奎拉
	f.拉奎拉Aquila.SetParent(f)
	
	f.阿特里Atri = BAtri阿特里
	f.阿特里Atri.SetParent(f)
	
	f.阿韦扎诺Avezzano = BAvezzano阿韦扎诺
	f.阿韦扎诺Avezzano.SetParent(f)
	
	f.帕格尼卡Paganica = BPaganica帕格尼卡
	f.帕格尼卡Paganica.SetParent(f)
	
	f.佩斯卡拉Pescara = BPescara佩斯卡拉
	f.佩斯卡拉Pescara.SetParent(f)
	
	f.泰拉莫Teramo = BTeramo泰拉莫
	f.泰拉莫Teramo.SetParent(f)
	
}
