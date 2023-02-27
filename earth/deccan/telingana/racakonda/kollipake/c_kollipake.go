package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KollipakeCounty interface {
    feud.County
    BChilkur_balaji赤库尔巴拉吉() 	feud.Barony
    BDevaryamjal德瓦尔亚加利() 	feud.Barony
    BGolkonda瞿罗军荼() 	feud.Barony
    BKarmanghat卡尔曼哈特() 	feud.Barony
    BKollipake拘梨波稽() 	feud.Barony
    BPahla波诃罗() 	feud.Barony
    BRatia罗底耶() 	feud.Barony
}

type 拘梨波稽KollipakeCounty struct {
	feud.BaseCounty
	赤库尔巴拉吉Chilkur_balaji 	feud.Barony
	德瓦尔亚加利Devaryamjal 	feud.Barony
	瞿罗军荼Golkonda 	feud.Barony
	卡尔曼哈特Karmanghat 	feud.Barony
	拘梨波稽Kollipake 	feud.Barony
	波诃罗Pahla 	feud.Barony
	罗底耶Ratia 	feud.Barony
}

func (c *拘梨波稽KollipakeCounty) BChilkur_balaji赤库尔巴拉吉() feud.Barony {
	return c.赤库尔巴拉吉Chilkur_balaji
}
    
func (c *拘梨波稽KollipakeCounty) BDevaryamjal德瓦尔亚加利() feud.Barony {
	return c.德瓦尔亚加利Devaryamjal
}
    
func (c *拘梨波稽KollipakeCounty) BGolkonda瞿罗军荼() feud.Barony {
	return c.瞿罗军荼Golkonda
}
    
func (c *拘梨波稽KollipakeCounty) BKarmanghat卡尔曼哈特() feud.Barony {
	return c.卡尔曼哈特Karmanghat
}
    
func (c *拘梨波稽KollipakeCounty) BKollipake拘梨波稽() feud.Barony {
	return c.拘梨波稽Kollipake
}
    
func (c *拘梨波稽KollipakeCounty) BPahla波诃罗() feud.Barony {
	return c.波诃罗Pahla
}
    
func (c *拘梨波稽KollipakeCounty) BRatia罗底耶() feud.Barony {
	return c.罗底耶Ratia
}
    
var CKollipake拘梨波稽 KollipakeCounty = &拘梨波稽KollipakeCounty{}

func init() {
	f := CKollipake拘梨波稽.(*拘梨波稽KollipakeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1144",
		Title:     "kollipake",
		TitleName: "拘梨波稽",
		TitleCode: "c_kollipake",
		Baronies:  map[string]feud.Barony{},
	}

	f.赤库尔巴拉吉Chilkur_balaji = BChilkur_balaji赤库尔巴拉吉
	f.赤库尔巴拉吉Chilkur_balaji.SetParent(f)
	
	f.德瓦尔亚加利Devaryamjal = BDevaryamjal德瓦尔亚加利
	f.德瓦尔亚加利Devaryamjal.SetParent(f)
	
	f.瞿罗军荼Golkonda = BGolkonda瞿罗军荼
	f.瞿罗军荼Golkonda.SetParent(f)
	
	f.卡尔曼哈特Karmanghat = BKarmanghat卡尔曼哈特
	f.卡尔曼哈特Karmanghat.SetParent(f)
	
	f.拘梨波稽Kollipake = BKollipake拘梨波稽
	f.拘梨波稽Kollipake.SetParent(f)
	
	f.波诃罗Pahla = BPahla波诃罗
	f.波诃罗Pahla.SetParent(f)
	
	f.罗底耶Ratia = BRatia罗底耶
	f.罗底耶Ratia.SetParent(f)
	
}
