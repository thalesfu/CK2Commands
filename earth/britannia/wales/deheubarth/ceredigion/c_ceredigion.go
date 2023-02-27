package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CeredigionCounty interface {
    feud.County
    BAberaeron阿伯赖伦() 	feud.Barony
    BAberystwyth阿伯里斯特威斯() 	feud.Barony
    BCardigan卡迪根() 	feud.Barony
    BLampeter兰彼得() 	feud.Barony
    BLlanbadarn兰巴达恩() 	feud.Barony
    BLlanddewi_brefi兰瑟威布雷维() 	feud.Barony
    BTregaron特雷加伦() 	feud.Barony
}

type 凯雷迪吉昂CeredigionCounty struct {
	feud.BaseCounty
	阿伯赖伦Aberaeron 	feud.Barony
	阿伯里斯特威斯Aberystwyth 	feud.Barony
	卡迪根Cardigan 	feud.Barony
	兰彼得Lampeter 	feud.Barony
	兰巴达恩Llanbadarn 	feud.Barony
	兰瑟威布雷维Llanddewi_brefi 	feud.Barony
	特雷加伦Tregaron 	feud.Barony
}

func (c *凯雷迪吉昂CeredigionCounty) BAberaeron阿伯赖伦() feud.Barony {
	return c.阿伯赖伦Aberaeron
}
    
func (c *凯雷迪吉昂CeredigionCounty) BAberystwyth阿伯里斯特威斯() feud.Barony {
	return c.阿伯里斯特威斯Aberystwyth
}
    
func (c *凯雷迪吉昂CeredigionCounty) BCardigan卡迪根() feud.Barony {
	return c.卡迪根Cardigan
}
    
func (c *凯雷迪吉昂CeredigionCounty) BLampeter兰彼得() feud.Barony {
	return c.兰彼得Lampeter
}
    
func (c *凯雷迪吉昂CeredigionCounty) BLlanbadarn兰巴达恩() feud.Barony {
	return c.兰巴达恩Llanbadarn
}
    
func (c *凯雷迪吉昂CeredigionCounty) BLlanddewi_brefi兰瑟威布雷维() feud.Barony {
	return c.兰瑟威布雷维Llanddewi_brefi
}
    
func (c *凯雷迪吉昂CeredigionCounty) BTregaron特雷加伦() feud.Barony {
	return c.特雷加伦Tregaron
}
    
var CCeredigion凯雷迪吉昂 CeredigionCounty = &凯雷迪吉昂CeredigionCounty{}

func init() {
	f := CCeredigion凯雷迪吉昂.(*凯雷迪吉昂CeredigionCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1948",
		Title:     "ceredigion",
		TitleName: "凯雷迪吉昂",
		TitleCode: "c_ceredigion",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯赖伦Aberaeron = BAberaeron阿伯赖伦
	f.阿伯赖伦Aberaeron.SetParent(f)
	
	f.阿伯里斯特威斯Aberystwyth = BAberystwyth阿伯里斯特威斯
	f.阿伯里斯特威斯Aberystwyth.SetParent(f)
	
	f.卡迪根Cardigan = BCardigan卡迪根
	f.卡迪根Cardigan.SetParent(f)
	
	f.兰彼得Lampeter = BLampeter兰彼得
	f.兰彼得Lampeter.SetParent(f)
	
	f.兰巴达恩Llanbadarn = BLlanbadarn兰巴达恩
	f.兰巴达恩Llanbadarn.SetParent(f)
	
	f.兰瑟威布雷维Llanddewi_brefi = BLlanddewi_brefi兰瑟威布雷维
	f.兰瑟威布雷维Llanddewi_brefi.SetParent(f)
	
	f.特雷加伦Tregaron = BTregaron特雷加伦
	f.特雷加伦Tregaron.SetParent(f)
	
}
