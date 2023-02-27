package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LiegeCounty interface {
    feud.County
    BBastogne巴斯托涅() 	feud.Barony
    BCine西内() 	feud.Barony
    BHuy于伊() 	feud.Barony
    BLaroche拉罗什() 	feud.Barony
    BLiege列日() 	feud.Barony
    BNamur那慕尔() 	feud.Barony
    BSalm萨尔姆() 	feud.Barony
    BStavelot斯塔沃洛() 	feud.Barony
}

type 列日LiegeCounty struct {
	feud.BaseCounty
	巴斯托涅Bastogne 	feud.Barony
	西内Cine 	feud.Barony
	于伊Huy 	feud.Barony
	拉罗什Laroche 	feud.Barony
	列日Liege 	feud.Barony
	那慕尔Namur 	feud.Barony
	萨尔姆Salm 	feud.Barony
	斯塔沃洛Stavelot 	feud.Barony
}

func (c *列日LiegeCounty) BBastogne巴斯托涅() feud.Barony {
	return c.巴斯托涅Bastogne
}
    
func (c *列日LiegeCounty) BCine西内() feud.Barony {
	return c.西内Cine
}
    
func (c *列日LiegeCounty) BHuy于伊() feud.Barony {
	return c.于伊Huy
}
    
func (c *列日LiegeCounty) BLaroche拉罗什() feud.Barony {
	return c.拉罗什Laroche
}
    
func (c *列日LiegeCounty) BLiege列日() feud.Barony {
	return c.列日Liege
}
    
func (c *列日LiegeCounty) BNamur那慕尔() feud.Barony {
	return c.那慕尔Namur
}
    
func (c *列日LiegeCounty) BSalm萨尔姆() feud.Barony {
	return c.萨尔姆Salm
}
    
func (c *列日LiegeCounty) BStavelot斯塔沃洛() feud.Barony {
	return c.斯塔沃洛Stavelot
}
    
var CLiege列日 LiegeCounty = &列日LiegeCounty{}

func init() {
	f := CLiege列日.(*列日LiegeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "116",
		Title:     "liege",
		TitleName: "列日",
		TitleCode: "c_liege",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴斯托涅Bastogne = BBastogne巴斯托涅
	f.巴斯托涅Bastogne.SetParent(f)
	
	f.西内Cine = BCine西内
	f.西内Cine.SetParent(f)
	
	f.于伊Huy = BHuy于伊
	f.于伊Huy.SetParent(f)
	
	f.拉罗什Laroche = BLaroche拉罗什
	f.拉罗什Laroche.SetParent(f)
	
	f.列日Liege = BLiege列日
	f.列日Liege.SetParent(f)
	
	f.那慕尔Namur = BNamur那慕尔
	f.那慕尔Namur.SetParent(f)
	
	f.萨尔姆Salm = BSalm萨尔姆
	f.萨尔姆Salm.SetParent(f)
	
	f.斯塔沃洛Stavelot = BStavelot斯塔沃洛
	f.斯塔沃洛Stavelot.SetParent(f)
	
}
