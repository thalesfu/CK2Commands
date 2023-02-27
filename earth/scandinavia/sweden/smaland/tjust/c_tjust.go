package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TjustCounty interface {
    feud.County
    BAlmvik阿尔姆维克() 	feud.Barony
    BGunnebo贡讷布() 	feud.Barony
    BJangolsrum杨格尔斯鲁姆() 	feud.Barony
    BStegsholm斯特格斯霍尔姆() 	feud.Barony
    BTornsfall特恩斯法尔() 	feud.Barony
    BVastervik韦斯特维克() 	feud.Barony
    BVinas维奈斯() 	feud.Barony
    BVraka弗罗卡() 	feud.Barony
}

type 修斯特TjustCounty struct {
	feud.BaseCounty
	阿尔姆维克Almvik 	feud.Barony
	贡讷布Gunnebo 	feud.Barony
	杨格尔斯鲁姆Jangolsrum 	feud.Barony
	斯特格斯霍尔姆Stegsholm 	feud.Barony
	特恩斯法尔Tornsfall 	feud.Barony
	韦斯特维克Vastervik 	feud.Barony
	维奈斯Vinas 	feud.Barony
	弗罗卡Vraka 	feud.Barony
}

func (c *修斯特TjustCounty) BAlmvik阿尔姆维克() feud.Barony {
	return c.阿尔姆维克Almvik
}
    
func (c *修斯特TjustCounty) BGunnebo贡讷布() feud.Barony {
	return c.贡讷布Gunnebo
}
    
func (c *修斯特TjustCounty) BJangolsrum杨格尔斯鲁姆() feud.Barony {
	return c.杨格尔斯鲁姆Jangolsrum
}
    
func (c *修斯特TjustCounty) BStegsholm斯特格斯霍尔姆() feud.Barony {
	return c.斯特格斯霍尔姆Stegsholm
}
    
func (c *修斯特TjustCounty) BTornsfall特恩斯法尔() feud.Barony {
	return c.特恩斯法尔Tornsfall
}
    
func (c *修斯特TjustCounty) BVastervik韦斯特维克() feud.Barony {
	return c.韦斯特维克Vastervik
}
    
func (c *修斯特TjustCounty) BVinas维奈斯() feud.Barony {
	return c.维奈斯Vinas
}
    
func (c *修斯特TjustCounty) BVraka弗罗卡() feud.Barony {
	return c.弗罗卡Vraka
}
    
var CTjust修斯特 TjustCounty = &修斯特TjustCounty{}

func init() {
	f := CTjust修斯特.(*修斯特TjustCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "932",
		Title:     "tjust",
		TitleName: "修斯特",
		TitleCode: "c_tjust",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔姆维克Almvik = BAlmvik阿尔姆维克
	f.阿尔姆维克Almvik.SetParent(f)
	
	f.贡讷布Gunnebo = BGunnebo贡讷布
	f.贡讷布Gunnebo.SetParent(f)
	
	f.杨格尔斯鲁姆Jangolsrum = BJangolsrum杨格尔斯鲁姆
	f.杨格尔斯鲁姆Jangolsrum.SetParent(f)
	
	f.斯特格斯霍尔姆Stegsholm = BStegsholm斯特格斯霍尔姆
	f.斯特格斯霍尔姆Stegsholm.SetParent(f)
	
	f.特恩斯法尔Tornsfall = BTornsfall特恩斯法尔
	f.特恩斯法尔Tornsfall.SetParent(f)
	
	f.韦斯特维克Vastervik = BVastervik韦斯特维克
	f.韦斯特维克Vastervik.SetParent(f)
	
	f.维奈斯Vinas = BVinas维奈斯
	f.维奈斯Vinas.SetParent(f)
	
	f.弗罗卡Vraka = BVraka弗罗卡
	f.弗罗卡Vraka.SetParent(f)
	
}
