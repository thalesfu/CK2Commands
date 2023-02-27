package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FinnmarkCounty interface {
    feud.County
    BHammerfest哈默弗斯特() 	feud.Barony
    BKarsloy卡尔绥() 	feud.Barony
    BMalangen马朗恩() 	feud.Barony
    BOstervagen厄斯特沃格() 	feud.Barony
    BPiselvnes皮瑟吕尼() 	feud.Barony
    BTromso特罗姆瑟() 	feud.Barony
    BVardohus瓦尔德胡斯() 	feud.Barony
    BVarghoeya瓦尔格赫亚() 	feud.Barony
}

type 芬马克FinnmarkCounty struct {
	feud.BaseCounty
	哈默弗斯特Hammerfest 	feud.Barony
	卡尔绥Karsloy 	feud.Barony
	马朗恩Malangen 	feud.Barony
	厄斯特沃格Ostervagen 	feud.Barony
	皮瑟吕尼Piselvnes 	feud.Barony
	特罗姆瑟Tromso 	feud.Barony
	瓦尔德胡斯Vardohus 	feud.Barony
	瓦尔格赫亚Varghoeya 	feud.Barony
}

func (c *芬马克FinnmarkCounty) BHammerfest哈默弗斯特() feud.Barony {
	return c.哈默弗斯特Hammerfest
}
    
func (c *芬马克FinnmarkCounty) BKarsloy卡尔绥() feud.Barony {
	return c.卡尔绥Karsloy
}
    
func (c *芬马克FinnmarkCounty) BMalangen马朗恩() feud.Barony {
	return c.马朗恩Malangen
}
    
func (c *芬马克FinnmarkCounty) BOstervagen厄斯特沃格() feud.Barony {
	return c.厄斯特沃格Ostervagen
}
    
func (c *芬马克FinnmarkCounty) BPiselvnes皮瑟吕尼() feud.Barony {
	return c.皮瑟吕尼Piselvnes
}
    
func (c *芬马克FinnmarkCounty) BTromso特罗姆瑟() feud.Barony {
	return c.特罗姆瑟Tromso
}
    
func (c *芬马克FinnmarkCounty) BVardohus瓦尔德胡斯() feud.Barony {
	return c.瓦尔德胡斯Vardohus
}
    
func (c *芬马克FinnmarkCounty) BVarghoeya瓦尔格赫亚() feud.Barony {
	return c.瓦尔格赫亚Varghoeya
}
    
var CFinnmark芬马克 FinnmarkCounty = &芬马克FinnmarkCounty{}

func init() {
	f := CFinnmark芬马克.(*芬马克FinnmarkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "389",
		Title:     "finnmark",
		TitleName: "芬马克",
		TitleCode: "c_finnmark",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈默弗斯特Hammerfest = BHammerfest哈默弗斯特
	f.哈默弗斯特Hammerfest.SetParent(f)
	
	f.卡尔绥Karsloy = BKarsloy卡尔绥
	f.卡尔绥Karsloy.SetParent(f)
	
	f.马朗恩Malangen = BMalangen马朗恩
	f.马朗恩Malangen.SetParent(f)
	
	f.厄斯特沃格Ostervagen = BOstervagen厄斯特沃格
	f.厄斯特沃格Ostervagen.SetParent(f)
	
	f.皮瑟吕尼Piselvnes = BPiselvnes皮瑟吕尼
	f.皮瑟吕尼Piselvnes.SetParent(f)
	
	f.特罗姆瑟Tromso = BTromso特罗姆瑟
	f.特罗姆瑟Tromso.SetParent(f)
	
	f.瓦尔德胡斯Vardohus = BVardohus瓦尔德胡斯
	f.瓦尔德胡斯Vardohus.SetParent(f)
	
	f.瓦尔格赫亚Varghoeya = BVarghoeya瓦尔格赫亚
	f.瓦尔格赫亚Varghoeya.SetParent(f)
	
}
