package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AylikCounty interface {
    feud.County
    BAylik艾里克() 	feud.Barony
    BBaijiantan白碱滩() 	feud.Barony
    BKaramay克拉玛依() 	feud.Barony
    BKuytun奎屯() 	feud.Barony
    BShawan沙湾() 	feud.Barony
    BUrho乌尔禾() 	feud.Barony
    BUsu乌苏() 	feud.Barony
}

type 艾里克AylikCounty struct {
	feud.BaseCounty
	艾里克Aylik 	feud.Barony
	白碱滩Baijiantan 	feud.Barony
	克拉玛依Karamay 	feud.Barony
	奎屯Kuytun 	feud.Barony
	沙湾Shawan 	feud.Barony
	乌尔禾Urho 	feud.Barony
	乌苏Usu 	feud.Barony
}

func (c *艾里克AylikCounty) BAylik艾里克() feud.Barony {
	return c.艾里克Aylik
}
    
func (c *艾里克AylikCounty) BBaijiantan白碱滩() feud.Barony {
	return c.白碱滩Baijiantan
}
    
func (c *艾里克AylikCounty) BKaramay克拉玛依() feud.Barony {
	return c.克拉玛依Karamay
}
    
func (c *艾里克AylikCounty) BKuytun奎屯() feud.Barony {
	return c.奎屯Kuytun
}
    
func (c *艾里克AylikCounty) BShawan沙湾() feud.Barony {
	return c.沙湾Shawan
}
    
func (c *艾里克AylikCounty) BUrho乌尔禾() feud.Barony {
	return c.乌尔禾Urho
}
    
func (c *艾里克AylikCounty) BUsu乌苏() feud.Barony {
	return c.乌苏Usu
}
    
var CAylik艾里克 AylikCounty = &艾里克AylikCounty{}

func init() {
	f := CAylik艾里克.(*艾里克AylikCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1875",
		Title:     "aylik",
		TitleName: "艾里克",
		TitleCode: "c_aylik",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾里克Aylik = BAylik艾里克
	f.艾里克Aylik.SetParent(f)
	
	f.白碱滩Baijiantan = BBaijiantan白碱滩
	f.白碱滩Baijiantan.SetParent(f)
	
	f.克拉玛依Karamay = BKaramay克拉玛依
	f.克拉玛依Karamay.SetParent(f)
	
	f.奎屯Kuytun = BKuytun奎屯
	f.奎屯Kuytun.SetParent(f)
	
	f.沙湾Shawan = BShawan沙湾
	f.沙湾Shawan.SetParent(f)
	
	f.乌尔禾Urho = BUrho乌尔禾
	f.乌尔禾Urho.SetParent(f)
	
	f.乌苏Usu = BUsu乌苏
	f.乌苏Usu.SetParent(f)
	
}
