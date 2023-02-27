package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KeriyaCounty interface {
    feud.County
    BChejiang车江() 	feud.Barony
    BDandan_uilik丹丹乌依里克() 	feud.Barony
    BGaoyuanqiang高院墙() 	feud.Barony
    BKeriya克里雅() 	feud.Barony
    BPimo媲摩() 	feud.Barony
    BYiyi伊夷() 	feud.Barony
    BZhoucun周村() 	feud.Barony
}

type 克里雅KeriyaCounty struct {
	feud.BaseCounty
	车江Chejiang 	feud.Barony
	丹丹乌依里克Dandan_uilik 	feud.Barony
	高院墙Gaoyuanqiang 	feud.Barony
	克里雅Keriya 	feud.Barony
	媲摩Pimo 	feud.Barony
	伊夷Yiyi 	feud.Barony
	周村Zhoucun 	feud.Barony
}

func (c *克里雅KeriyaCounty) BChejiang车江() feud.Barony {
	return c.车江Chejiang
}
    
func (c *克里雅KeriyaCounty) BDandan_uilik丹丹乌依里克() feud.Barony {
	return c.丹丹乌依里克Dandan_uilik
}
    
func (c *克里雅KeriyaCounty) BGaoyuanqiang高院墙() feud.Barony {
	return c.高院墙Gaoyuanqiang
}
    
func (c *克里雅KeriyaCounty) BKeriya克里雅() feud.Barony {
	return c.克里雅Keriya
}
    
func (c *克里雅KeriyaCounty) BPimo媲摩() feud.Barony {
	return c.媲摩Pimo
}
    
func (c *克里雅KeriyaCounty) BYiyi伊夷() feud.Barony {
	return c.伊夷Yiyi
}
    
func (c *克里雅KeriyaCounty) BZhoucun周村() feud.Barony {
	return c.周村Zhoucun
}
    
var CKeriya克里雅 KeriyaCounty = &克里雅KeriyaCounty{}

func init() {
	f := CKeriya克里雅.(*克里雅KeriyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1523",
		Title:     "keriya",
		TitleName: "克里雅",
		TitleCode: "c_keriya",
		Baronies:  map[string]feud.Barony{},
	}

	f.车江Chejiang = BChejiang车江
	f.车江Chejiang.SetParent(f)
	
	f.丹丹乌依里克Dandan_uilik = BDandan_uilik丹丹乌依里克
	f.丹丹乌依里克Dandan_uilik.SetParent(f)
	
	f.高院墙Gaoyuanqiang = BGaoyuanqiang高院墙
	f.高院墙Gaoyuanqiang.SetParent(f)
	
	f.克里雅Keriya = BKeriya克里雅
	f.克里雅Keriya.SetParent(f)
	
	f.媲摩Pimo = BPimo媲摩
	f.媲摩Pimo.SetParent(f)
	
	f.伊夷Yiyi = BYiyi伊夷
	f.伊夷Yiyi.SetParent(f)
	
	f.周村Zhoucun = BZhoucun周村
	f.周村Zhoucun.SetParent(f)
	
}
