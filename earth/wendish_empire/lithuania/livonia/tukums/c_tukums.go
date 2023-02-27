package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TukumsCounty interface {
    feud.County
    BBulla布拉() 	feud.Barony
    BDundaga敦达加() 	feud.Barony
    BKandava坎达瓦() 	feud.Barony
    BKayeva卡耶瓦() 	feud.Barony
    BSabile萨比莱() 	feud.Barony
    BTalsi塔尔西() 	feud.Barony
    BTukums图库姆斯() 	feud.Barony
}

type 图库姆斯TukumsCounty struct {
	feud.BaseCounty
	布拉Bulla 	feud.Barony
	敦达加Dundaga 	feud.Barony
	坎达瓦Kandava 	feud.Barony
	卡耶瓦Kayeva 	feud.Barony
	萨比莱Sabile 	feud.Barony
	塔尔西Talsi 	feud.Barony
	图库姆斯Tukums 	feud.Barony
}

func (c *图库姆斯TukumsCounty) BBulla布拉() feud.Barony {
	return c.布拉Bulla
}
    
func (c *图库姆斯TukumsCounty) BDundaga敦达加() feud.Barony {
	return c.敦达加Dundaga
}
    
func (c *图库姆斯TukumsCounty) BKandava坎达瓦() feud.Barony {
	return c.坎达瓦Kandava
}
    
func (c *图库姆斯TukumsCounty) BKayeva卡耶瓦() feud.Barony {
	return c.卡耶瓦Kayeva
}
    
func (c *图库姆斯TukumsCounty) BSabile萨比莱() feud.Barony {
	return c.萨比莱Sabile
}
    
func (c *图库姆斯TukumsCounty) BTalsi塔尔西() feud.Barony {
	return c.塔尔西Talsi
}
    
func (c *图库姆斯TukumsCounty) BTukums图库姆斯() feud.Barony {
	return c.图库姆斯Tukums
}
    
var CTukums图库姆斯 TukumsCounty = &图库姆斯TukumsCounty{}

func init() {
	f := CTukums图库姆斯.(*图库姆斯TukumsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1592",
		Title:     "tukums",
		TitleName: "图库姆斯",
		TitleCode: "c_tukums",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉Bulla = BBulla布拉
	f.布拉Bulla.SetParent(f)
	
	f.敦达加Dundaga = BDundaga敦达加
	f.敦达加Dundaga.SetParent(f)
	
	f.坎达瓦Kandava = BKandava坎达瓦
	f.坎达瓦Kandava.SetParent(f)
	
	f.卡耶瓦Kayeva = BKayeva卡耶瓦
	f.卡耶瓦Kayeva.SetParent(f)
	
	f.萨比莱Sabile = BSabile萨比莱
	f.萨比莱Sabile.SetParent(f)
	
	f.塔尔西Talsi = BTalsi塔尔西
	f.塔尔西Talsi.SetParent(f)
	
	f.图库姆斯Tukums = BTukums图库姆斯
	f.图库姆斯Tukums.SetParent(f)
	
}
