package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NalsiaCounty interface {
    feud.County
    BAnyksciai阿尼克什奇艾() 	feud.Barony
    BKernave克尔纳韦() 	feud.Barony
    BNemencine内门奇内() 	feud.Barony
    BNerisia内里西亚() 	feud.Barony
    BSvencionys什文乔尼斯() 	feud.Barony
    BUkmerge乌克梅尔盖() 	feud.Barony
    BUtena乌田纳() 	feud.Barony
}

type 纳尔希亚NalsiaCounty struct {
	feud.BaseCounty
	阿尼克什奇艾Anyksciai 	feud.Barony
	克尔纳韦Kernave 	feud.Barony
	内门奇内Nemencine 	feud.Barony
	内里西亚Nerisia 	feud.Barony
	什文乔尼斯Svencionys 	feud.Barony
	乌克梅尔盖Ukmerge 	feud.Barony
	乌田纳Utena 	feud.Barony
}

func (c *纳尔希亚NalsiaCounty) BAnyksciai阿尼克什奇艾() feud.Barony {
	return c.阿尼克什奇艾Anyksciai
}
    
func (c *纳尔希亚NalsiaCounty) BKernave克尔纳韦() feud.Barony {
	return c.克尔纳韦Kernave
}
    
func (c *纳尔希亚NalsiaCounty) BNemencine内门奇内() feud.Barony {
	return c.内门奇内Nemencine
}
    
func (c *纳尔希亚NalsiaCounty) BNerisia内里西亚() feud.Barony {
	return c.内里西亚Nerisia
}
    
func (c *纳尔希亚NalsiaCounty) BSvencionys什文乔尼斯() feud.Barony {
	return c.什文乔尼斯Svencionys
}
    
func (c *纳尔希亚NalsiaCounty) BUkmerge乌克梅尔盖() feud.Barony {
	return c.乌克梅尔盖Ukmerge
}
    
func (c *纳尔希亚NalsiaCounty) BUtena乌田纳() feud.Barony {
	return c.乌田纳Utena
}
    
var CNalsia纳尔希亚 NalsiaCounty = &纳尔希亚NalsiaCounty{}

func init() {
	f := CNalsia纳尔希亚.(*纳尔希亚NalsiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1590",
		Title:     "nalsia",
		TitleName: "纳尔希亚",
		TitleCode: "c_nalsia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尼克什奇艾Anyksciai = BAnyksciai阿尼克什奇艾
	f.阿尼克什奇艾Anyksciai.SetParent(f)
	
	f.克尔纳韦Kernave = BKernave克尔纳韦
	f.克尔纳韦Kernave.SetParent(f)
	
	f.内门奇内Nemencine = BNemencine内门奇内
	f.内门奇内Nemencine.SetParent(f)
	
	f.内里西亚Nerisia = BNerisia内里西亚
	f.内里西亚Nerisia.SetParent(f)
	
	f.什文乔尼斯Svencionys = BSvencionys什文乔尼斯
	f.什文乔尼斯Svencionys.SetParent(f)
	
	f.乌克梅尔盖Ukmerge = BUkmerge乌克梅尔盖
	f.乌克梅尔盖Ukmerge.SetParent(f)
	
	f.乌田纳Utena = BUtena乌田纳
	f.乌田纳Utena.SetParent(f)
	
}
