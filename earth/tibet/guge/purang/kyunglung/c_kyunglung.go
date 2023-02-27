package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KyunglungCounty interface {
    feud.County
    BDolchu顿久() 	feud.Barony
    BGurugem古入江() 	feud.Barony
    BGyanyima姜叶玛() 	feud.Barony
    BKyunglung琼隆() 	feud.Barony
    BMonicer门士() 	feud.Barony
    BTirthapuri直达布日() 	feud.Barony
    BXalazakung夏拉杂孔() 	feud.Barony
}

type 琼隆KyunglungCounty struct {
	feud.BaseCounty
	顿久Dolchu 	feud.Barony
	古入江Gurugem 	feud.Barony
	姜叶玛Gyanyima 	feud.Barony
	琼隆Kyunglung 	feud.Barony
	门士Monicer 	feud.Barony
	直达布日Tirthapuri 	feud.Barony
	夏拉杂孔Xalazakung 	feud.Barony
}

func (c *琼隆KyunglungCounty) BDolchu顿久() feud.Barony {
	return c.顿久Dolchu
}
    
func (c *琼隆KyunglungCounty) BGurugem古入江() feud.Barony {
	return c.古入江Gurugem
}
    
func (c *琼隆KyunglungCounty) BGyanyima姜叶玛() feud.Barony {
	return c.姜叶玛Gyanyima
}
    
func (c *琼隆KyunglungCounty) BKyunglung琼隆() feud.Barony {
	return c.琼隆Kyunglung
}
    
func (c *琼隆KyunglungCounty) BMonicer门士() feud.Barony {
	return c.门士Monicer
}
    
func (c *琼隆KyunglungCounty) BTirthapuri直达布日() feud.Barony {
	return c.直达布日Tirthapuri
}
    
func (c *琼隆KyunglungCounty) BXalazakung夏拉杂孔() feud.Barony {
	return c.夏拉杂孔Xalazakung
}
    
var CKyunglung琼隆 KyunglungCounty = &琼隆KyunglungCounty{}

func init() {
	f := CKyunglung琼隆.(*琼隆KyunglungCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1490",
		Title:     "kyunglung",
		TitleName: "琼隆",
		TitleCode: "c_kyunglung",
		Baronies:  map[string]feud.Barony{},
	}

	f.顿久Dolchu = BDolchu顿久
	f.顿久Dolchu.SetParent(f)
	
	f.古入江Gurugem = BGurugem古入江
	f.古入江Gurugem.SetParent(f)
	
	f.姜叶玛Gyanyima = BGyanyima姜叶玛
	f.姜叶玛Gyanyima.SetParent(f)
	
	f.琼隆Kyunglung = BKyunglung琼隆
	f.琼隆Kyunglung.SetParent(f)
	
	f.门士Monicer = BMonicer门士
	f.门士Monicer.SetParent(f)
	
	f.直达布日Tirthapuri = BTirthapuri直达布日
	f.直达布日Tirthapuri.SetParent(f)
	
	f.夏拉杂孔Xalazakung = BXalazakung夏拉杂孔
	f.夏拉杂孔Xalazakung.SetParent(f)
	
}
