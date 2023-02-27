package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlmaliqCounty interface {
    feud.County
    BAlmaliq阿力麻里() 	feud.Barony
    BKashi喀什() 	feud.Barony
    BNalatizhen那拉提镇() 	feud.Barony
    BNilka尼勒克() 	feud.Barony
    BTaledzhen塔勒德镇() 	feud.Barony
    BXinyuan新源() 	feud.Barony
    BZeketaizhen则克台镇() 	feud.Barony
}

type 阿力麻里AlmaliqCounty struct {
	feud.BaseCounty
	阿力麻里Almaliq 	feud.Barony
	喀什Kashi 	feud.Barony
	那拉提镇Nalatizhen 	feud.Barony
	尼勒克Nilka 	feud.Barony
	塔勒德镇Taledzhen 	feud.Barony
	新源Xinyuan 	feud.Barony
	则克台镇Zeketaizhen 	feud.Barony
}

func (c *阿力麻里AlmaliqCounty) BAlmaliq阿力麻里() feud.Barony {
	return c.阿力麻里Almaliq
}
    
func (c *阿力麻里AlmaliqCounty) BKashi喀什() feud.Barony {
	return c.喀什Kashi
}
    
func (c *阿力麻里AlmaliqCounty) BNalatizhen那拉提镇() feud.Barony {
	return c.那拉提镇Nalatizhen
}
    
func (c *阿力麻里AlmaliqCounty) BNilka尼勒克() feud.Barony {
	return c.尼勒克Nilka
}
    
func (c *阿力麻里AlmaliqCounty) BTaledzhen塔勒德镇() feud.Barony {
	return c.塔勒德镇Taledzhen
}
    
func (c *阿力麻里AlmaliqCounty) BXinyuan新源() feud.Barony {
	return c.新源Xinyuan
}
    
func (c *阿力麻里AlmaliqCounty) BZeketaizhen则克台镇() feud.Barony {
	return c.则克台镇Zeketaizhen
}
    
var CAlmaliq阿力麻里 AlmaliqCounty = &阿力麻里AlmaliqCounty{}

func init() {
	f := CAlmaliq阿力麻里.(*阿力麻里AlmaliqCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1798",
		Title:     "almaliq",
		TitleName: "阿力麻里",
		TitleCode: "c_almaliq",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿力麻里Almaliq = BAlmaliq阿力麻里
	f.阿力麻里Almaliq.SetParent(f)
	
	f.喀什Kashi = BKashi喀什
	f.喀什Kashi.SetParent(f)
	
	f.那拉提镇Nalatizhen = BNalatizhen那拉提镇
	f.那拉提镇Nalatizhen.SetParent(f)
	
	f.尼勒克Nilka = BNilka尼勒克
	f.尼勒克Nilka.SetParent(f)
	
	f.塔勒德镇Taledzhen = BTaledzhen塔勒德镇
	f.塔勒德镇Taledzhen.SetParent(f)
	
	f.新源Xinyuan = BXinyuan新源
	f.新源Xinyuan.SetParent(f)
	
	f.则克台镇Zeketaizhen = BZeketaizhen则克台镇
	f.则克台镇Zeketaizhen.SetParent(f)
	
}
