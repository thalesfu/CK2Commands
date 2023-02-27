package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LapplandCounty interface {
    feud.County
    BArjeplog阿尔耶普卢格() 	feud.Barony
    BArvidsjaur阿尔维斯尧尔() 	feud.Barony
    BAsele奥瑟勒() 	feud.Barony
    BBergvattnet巴依瓦特奈() 	feud.Barony
    BGallivare耶利瓦勒() 	feud.Barony
    BKiruna基律纳() 	feud.Barony
    BLycksele吕克瑟勒() 	feud.Barony
    BSorsele索舍勒() 	feud.Barony
}

type 拉普兰LapplandCounty struct {
	feud.BaseCounty
	阿尔耶普卢格Arjeplog 	feud.Barony
	阿尔维斯尧尔Arvidsjaur 	feud.Barony
	奥瑟勒Asele 	feud.Barony
	巴依瓦特奈Bergvattnet 	feud.Barony
	耶利瓦勒Gallivare 	feud.Barony
	基律纳Kiruna 	feud.Barony
	吕克瑟勒Lycksele 	feud.Barony
	索舍勒Sorsele 	feud.Barony
}

func (c *拉普兰LapplandCounty) BArjeplog阿尔耶普卢格() feud.Barony {
	return c.阿尔耶普卢格Arjeplog
}
    
func (c *拉普兰LapplandCounty) BArvidsjaur阿尔维斯尧尔() feud.Barony {
	return c.阿尔维斯尧尔Arvidsjaur
}
    
func (c *拉普兰LapplandCounty) BAsele奥瑟勒() feud.Barony {
	return c.奥瑟勒Asele
}
    
func (c *拉普兰LapplandCounty) BBergvattnet巴依瓦特奈() feud.Barony {
	return c.巴依瓦特奈Bergvattnet
}
    
func (c *拉普兰LapplandCounty) BGallivare耶利瓦勒() feud.Barony {
	return c.耶利瓦勒Gallivare
}
    
func (c *拉普兰LapplandCounty) BKiruna基律纳() feud.Barony {
	return c.基律纳Kiruna
}
    
func (c *拉普兰LapplandCounty) BLycksele吕克瑟勒() feud.Barony {
	return c.吕克瑟勒Lycksele
}
    
func (c *拉普兰LapplandCounty) BSorsele索舍勒() feud.Barony {
	return c.索舍勒Sorsele
}
    
var CLappland拉普兰 LapplandCounty = &拉普兰LapplandCounty{}

func init() {
	f := CLappland拉普兰.(*拉普兰LapplandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "279",
		Title:     "lappland",
		TitleName: "拉普兰",
		TitleCode: "c_lappland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔耶普卢格Arjeplog = BArjeplog阿尔耶普卢格
	f.阿尔耶普卢格Arjeplog.SetParent(f)
	
	f.阿尔维斯尧尔Arvidsjaur = BArvidsjaur阿尔维斯尧尔
	f.阿尔维斯尧尔Arvidsjaur.SetParent(f)
	
	f.奥瑟勒Asele = BAsele奥瑟勒
	f.奥瑟勒Asele.SetParent(f)
	
	f.巴依瓦特奈Bergvattnet = BBergvattnet巴依瓦特奈
	f.巴依瓦特奈Bergvattnet.SetParent(f)
	
	f.耶利瓦勒Gallivare = BGallivare耶利瓦勒
	f.耶利瓦勒Gallivare.SetParent(f)
	
	f.基律纳Kiruna = BKiruna基律纳
	f.基律纳Kiruna.SetParent(f)
	
	f.吕克瑟勒Lycksele = BLycksele吕克瑟勒
	f.吕克瑟勒Lycksele.SetParent(f)
	
	f.索舍勒Sorsele = BSorsele索舍勒
	f.索舍勒Sorsele.SetParent(f)
	
}
