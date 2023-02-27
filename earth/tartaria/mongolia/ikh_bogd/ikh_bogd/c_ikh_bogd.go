package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Ikh_bogdCounty interface {
    feud.County
    BArguut阿尔古特() 	feud.Barony
    BGobi戈壁() 	feud.Barony
    BIkh_bogd也客孛黑答() 	feud.Barony
    BKhereid克烈() 	feud.Barony
    BKhuld呼勒德() 	feud.Barony
    BNuuruudyn_hondij湖谷() 	feud.Barony
    BZubu阻卜() 	feud.Barony
}

type 也客孛黑答Ikh_bogdCounty struct {
	feud.BaseCounty
	阿尔古特Arguut 	feud.Barony
	戈壁Gobi 	feud.Barony
	也客孛黑答Ikh_bogd 	feud.Barony
	克烈Khereid 	feud.Barony
	呼勒德Khuld 	feud.Barony
	湖谷Nuuruudyn_hondij 	feud.Barony
	阻卜Zubu 	feud.Barony
}

func (c *也客孛黑答Ikh_bogdCounty) BArguut阿尔古特() feud.Barony {
	return c.阿尔古特Arguut
}
    
func (c *也客孛黑答Ikh_bogdCounty) BGobi戈壁() feud.Barony {
	return c.戈壁Gobi
}
    
func (c *也客孛黑答Ikh_bogdCounty) BIkh_bogd也客孛黑答() feud.Barony {
	return c.也客孛黑答Ikh_bogd
}
    
func (c *也客孛黑答Ikh_bogdCounty) BKhereid克烈() feud.Barony {
	return c.克烈Khereid
}
    
func (c *也客孛黑答Ikh_bogdCounty) BKhuld呼勒德() feud.Barony {
	return c.呼勒德Khuld
}
    
func (c *也客孛黑答Ikh_bogdCounty) BNuuruudyn_hondij湖谷() feud.Barony {
	return c.湖谷Nuuruudyn_hondij
}
    
func (c *也客孛黑答Ikh_bogdCounty) BZubu阻卜() feud.Barony {
	return c.阻卜Zubu
}
    
var CIkh_bogd也客孛黑答 Ikh_bogdCounty = &也客孛黑答Ikh_bogdCounty{}

func init() {
	f := CIkh_bogd也客孛黑答.(*也客孛黑答Ikh_bogdCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1456",
		Title:     "ikh_bogd",
		TitleName: "也客孛黑答",
		TitleCode: "c_ikh_bogd",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔古特Arguut = BArguut阿尔古特
	f.阿尔古特Arguut.SetParent(f)
	
	f.戈壁Gobi = BGobi戈壁
	f.戈壁Gobi.SetParent(f)
	
	f.也客孛黑答Ikh_bogd = BIkh_bogd也客孛黑答
	f.也客孛黑答Ikh_bogd.SetParent(f)
	
	f.克烈Khereid = BKhereid克烈
	f.克烈Khereid.SetParent(f)
	
	f.呼勒德Khuld = BKhuld呼勒德
	f.呼勒德Khuld.SetParent(f)
	
	f.湖谷Nuuruudyn_hondij = BNuuruudyn_hondij湖谷
	f.湖谷Nuuruudyn_hondij.SetParent(f)
	
	f.阻卜Zubu = BZubu阻卜
	f.阻卜Zubu.SetParent(f)
	
}
