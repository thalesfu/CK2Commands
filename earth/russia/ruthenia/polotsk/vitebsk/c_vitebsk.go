package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VitebskCounty interface {
    feud.County
    BBeshankovichy别申科维奇() 	feud.Barony
    BBilieva比列沃() 	feud.Barony
    BBoguchevsk博古舍夫斯克() 	feud.Barony
    BLiozna利奥兹诺() 	feud.Barony
    BRuba鲁巴() 	feud.Barony
    BSianno先诺() 	feud.Barony
    BVitebsk维捷布斯克() 	feud.Barony
}

type 维捷布斯克VitebskCounty struct {
	feud.BaseCounty
	别申科维奇Beshankovichy 	feud.Barony
	比列沃Bilieva 	feud.Barony
	博古舍夫斯克Boguchevsk 	feud.Barony
	利奥兹诺Liozna 	feud.Barony
	鲁巴Ruba 	feud.Barony
	先诺Sianno 	feud.Barony
	维捷布斯克Vitebsk 	feud.Barony
}

func (c *维捷布斯克VitebskCounty) BBeshankovichy别申科维奇() feud.Barony {
	return c.别申科维奇Beshankovichy
}
    
func (c *维捷布斯克VitebskCounty) BBilieva比列沃() feud.Barony {
	return c.比列沃Bilieva
}
    
func (c *维捷布斯克VitebskCounty) BBoguchevsk博古舍夫斯克() feud.Barony {
	return c.博古舍夫斯克Boguchevsk
}
    
func (c *维捷布斯克VitebskCounty) BLiozna利奥兹诺() feud.Barony {
	return c.利奥兹诺Liozna
}
    
func (c *维捷布斯克VitebskCounty) BRuba鲁巴() feud.Barony {
	return c.鲁巴Ruba
}
    
func (c *维捷布斯克VitebskCounty) BSianno先诺() feud.Barony {
	return c.先诺Sianno
}
    
func (c *维捷布斯克VitebskCounty) BVitebsk维捷布斯克() feud.Barony {
	return c.维捷布斯克Vitebsk
}
    
var CVitebsk维捷布斯克 VitebskCounty = &维捷布斯克VitebskCounty{}

func init() {
	f := CVitebsk维捷布斯克.(*维捷布斯克VitebskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "417",
		Title:     "vitebsk",
		TitleName: "维捷布斯克",
		TitleCode: "c_vitebsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.别申科维奇Beshankovichy = BBeshankovichy别申科维奇
	f.别申科维奇Beshankovichy.SetParent(f)
	
	f.比列沃Bilieva = BBilieva比列沃
	f.比列沃Bilieva.SetParent(f)
	
	f.博古舍夫斯克Boguchevsk = BBoguchevsk博古舍夫斯克
	f.博古舍夫斯克Boguchevsk.SetParent(f)
	
	f.利奥兹诺Liozna = BLiozna利奥兹诺
	f.利奥兹诺Liozna.SetParent(f)
	
	f.鲁巴Ruba = BRuba鲁巴
	f.鲁巴Ruba.SetParent(f)
	
	f.先诺Sianno = BSianno先诺
	f.先诺Sianno.SetParent(f)
	
	f.维捷布斯克Vitebsk = BVitebsk维捷布斯克
	f.维捷布斯克Vitebsk.SetParent(f)
	
}
