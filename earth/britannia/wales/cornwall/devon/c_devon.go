package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DevonCounty interface {
    feud.County
    BAxminster阿克斯明斯特() 	feud.Barony
    BBuckfast巴克法斯特() 	feud.Barony
    BCrediton克雷迪顿() 	feud.Barony
    BDartmouth达特茅斯() 	feud.Barony
    BExeter埃克塞特() 	feud.Barony
    BLydford利德福德() 	feud.Barony
    BTavistock塔维斯托克() 	feud.Barony
    BTotnes托特尼斯() 	feud.Barony
}

type 德文DevonCounty struct {
	feud.BaseCounty
	阿克斯明斯特Axminster 	feud.Barony
	巴克法斯特Buckfast 	feud.Barony
	克雷迪顿Crediton 	feud.Barony
	达特茅斯Dartmouth 	feud.Barony
	埃克塞特Exeter 	feud.Barony
	利德福德Lydford 	feud.Barony
	塔维斯托克Tavistock 	feud.Barony
	托特尼斯Totnes 	feud.Barony
}

func (c *德文DevonCounty) BAxminster阿克斯明斯特() feud.Barony {
	return c.阿克斯明斯特Axminster
}
    
func (c *德文DevonCounty) BBuckfast巴克法斯特() feud.Barony {
	return c.巴克法斯特Buckfast
}
    
func (c *德文DevonCounty) BCrediton克雷迪顿() feud.Barony {
	return c.克雷迪顿Crediton
}
    
func (c *德文DevonCounty) BDartmouth达特茅斯() feud.Barony {
	return c.达特茅斯Dartmouth
}
    
func (c *德文DevonCounty) BExeter埃克塞特() feud.Barony {
	return c.埃克塞特Exeter
}
    
func (c *德文DevonCounty) BLydford利德福德() feud.Barony {
	return c.利德福德Lydford
}
    
func (c *德文DevonCounty) BTavistock塔维斯托克() feud.Barony {
	return c.塔维斯托克Tavistock
}
    
func (c *德文DevonCounty) BTotnes托特尼斯() feud.Barony {
	return c.托特尼斯Totnes
}
    
var CDevon德文 DevonCounty = &德文DevonCounty{}

func init() {
	f := CDevon德文.(*德文DevonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "29",
		Title:     "devon",
		TitleName: "德文",
		TitleCode: "c_devon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克斯明斯特Axminster = BAxminster阿克斯明斯特
	f.阿克斯明斯特Axminster.SetParent(f)
	
	f.巴克法斯特Buckfast = BBuckfast巴克法斯特
	f.巴克法斯特Buckfast.SetParent(f)
	
	f.克雷迪顿Crediton = BCrediton克雷迪顿
	f.克雷迪顿Crediton.SetParent(f)
	
	f.达特茅斯Dartmouth = BDartmouth达特茅斯
	f.达特茅斯Dartmouth.SetParent(f)
	
	f.埃克塞特Exeter = BExeter埃克塞特
	f.埃克塞特Exeter.SetParent(f)
	
	f.利德福德Lydford = BLydford利德福德
	f.利德福德Lydford.SetParent(f)
	
	f.塔维斯托克Tavistock = BTavistock塔维斯托克
	f.塔维斯托克Tavistock.SetParent(f)
	
	f.托特尼斯Totnes = BTotnes托特尼斯
	f.托特尼斯Totnes.SetParent(f)
	
}
