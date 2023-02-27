package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlmatyCounty interface {
    feud.County
    BAkzhar_almaty阿克札尔() 	feud.Barony
    BAlmaty_saryesik阿拉木图() 	feud.Barony
    BKapchagai卡普恰盖() 	feud.Barony
    BKaraoy_almaty卡拉奥伊() 	feud.Barony
    BKargal卡尔加尔() 	feud.Barony
    BKoktal_almaty科克塔尔() 	feud.Barony
    BSaryesik萨雷耶西克() 	feud.Barony
}

type 阿拉木图AlmatyCounty struct {
	feud.BaseCounty
	阿克札尔Akzhar_almaty 	feud.Barony
	阿拉木图Almaty_saryesik 	feud.Barony
	卡普恰盖Kapchagai 	feud.Barony
	卡拉奥伊Karaoy_almaty 	feud.Barony
	卡尔加尔Kargal 	feud.Barony
	科克塔尔Koktal_almaty 	feud.Barony
	萨雷耶西克Saryesik 	feud.Barony
}

func (c *阿拉木图AlmatyCounty) BAkzhar_almaty阿克札尔() feud.Barony {
	return c.阿克札尔Akzhar_almaty
}
    
func (c *阿拉木图AlmatyCounty) BAlmaty_saryesik阿拉木图() feud.Barony {
	return c.阿拉木图Almaty_saryesik
}
    
func (c *阿拉木图AlmatyCounty) BKapchagai卡普恰盖() feud.Barony {
	return c.卡普恰盖Kapchagai
}
    
func (c *阿拉木图AlmatyCounty) BKaraoy_almaty卡拉奥伊() feud.Barony {
	return c.卡拉奥伊Karaoy_almaty
}
    
func (c *阿拉木图AlmatyCounty) BKargal卡尔加尔() feud.Barony {
	return c.卡尔加尔Kargal
}
    
func (c *阿拉木图AlmatyCounty) BKoktal_almaty科克塔尔() feud.Barony {
	return c.科克塔尔Koktal_almaty
}
    
func (c *阿拉木图AlmatyCounty) BSaryesik萨雷耶西克() feud.Barony {
	return c.萨雷耶西克Saryesik
}
    
var CAlmaty阿拉木图 AlmatyCounty = &阿拉木图AlmatyCounty{}

func init() {
	f := CAlmaty阿拉木图.(*阿拉木图AlmatyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1797",
		Title:     "almaty",
		TitleName: "阿拉木图",
		TitleCode: "c_almaty",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克札尔Akzhar_almaty = BAkzhar_almaty阿克札尔
	f.阿克札尔Akzhar_almaty.SetParent(f)
	
	f.阿拉木图Almaty_saryesik = BAlmaty_saryesik阿拉木图
	f.阿拉木图Almaty_saryesik.SetParent(f)
	
	f.卡普恰盖Kapchagai = BKapchagai卡普恰盖
	f.卡普恰盖Kapchagai.SetParent(f)
	
	f.卡拉奥伊Karaoy_almaty = BKaraoy_almaty卡拉奥伊
	f.卡拉奥伊Karaoy_almaty.SetParent(f)
	
	f.卡尔加尔Kargal = BKargal卡尔加尔
	f.卡尔加尔Kargal.SetParent(f)
	
	f.科克塔尔Koktal_almaty = BKoktal_almaty科克塔尔
	f.科克塔尔Koktal_almaty.SetParent(f)
	
	f.萨雷耶西克Saryesik = BSaryesik萨雷耶西克
	f.萨雷耶西克Saryesik.SetParent(f)
	
}
