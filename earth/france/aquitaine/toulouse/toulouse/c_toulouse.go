package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ToulouseCounty interface {
    feud.County
    BCastelnaudary卡斯泰尔诺达里() 	feud.Barony
    BHautpoul欧特普尔() 	feud.Barony
    BLavaur拉沃尔() 	feud.Barony
    BLombez隆贝() 	feud.Barony
    BMontauban蒙托邦() 	feud.Barony
    BMontgiscard蒙日斯卡尔() 	feud.Barony
    BMuret米雷() 	feud.Barony
    BToulouse图卢兹() 	feud.Barony
}

type 图卢兹ToulouseCounty struct {
	feud.BaseCounty
	卡斯泰尔诺达里Castelnaudary 	feud.Barony
	欧特普尔Hautpoul 	feud.Barony
	拉沃尔Lavaur 	feud.Barony
	隆贝Lombez 	feud.Barony
	蒙托邦Montauban 	feud.Barony
	蒙日斯卡尔Montgiscard 	feud.Barony
	米雷Muret 	feud.Barony
	图卢兹Toulouse 	feud.Barony
}

func (c *图卢兹ToulouseCounty) BCastelnaudary卡斯泰尔诺达里() feud.Barony {
	return c.卡斯泰尔诺达里Castelnaudary
}
    
func (c *图卢兹ToulouseCounty) BHautpoul欧特普尔() feud.Barony {
	return c.欧特普尔Hautpoul
}
    
func (c *图卢兹ToulouseCounty) BLavaur拉沃尔() feud.Barony {
	return c.拉沃尔Lavaur
}
    
func (c *图卢兹ToulouseCounty) BLombez隆贝() feud.Barony {
	return c.隆贝Lombez
}
    
func (c *图卢兹ToulouseCounty) BMontauban蒙托邦() feud.Barony {
	return c.蒙托邦Montauban
}
    
func (c *图卢兹ToulouseCounty) BMontgiscard蒙日斯卡尔() feud.Barony {
	return c.蒙日斯卡尔Montgiscard
}
    
func (c *图卢兹ToulouseCounty) BMuret米雷() feud.Barony {
	return c.米雷Muret
}
    
func (c *图卢兹ToulouseCounty) BToulouse图卢兹() feud.Barony {
	return c.图卢兹Toulouse
}
    
var CToulouse图卢兹 ToulouseCounty = &图卢兹ToulouseCounty{}

func init() {
	f := CToulouse图卢兹.(*图卢兹ToulouseCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "214",
		Title:     "toulouse",
		TitleName: "图卢兹",
		TitleCode: "c_toulouse",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡斯泰尔诺达里Castelnaudary = BCastelnaudary卡斯泰尔诺达里
	f.卡斯泰尔诺达里Castelnaudary.SetParent(f)
	
	f.欧特普尔Hautpoul = BHautpoul欧特普尔
	f.欧特普尔Hautpoul.SetParent(f)
	
	f.拉沃尔Lavaur = BLavaur拉沃尔
	f.拉沃尔Lavaur.SetParent(f)
	
	f.隆贝Lombez = BLombez隆贝
	f.隆贝Lombez.SetParent(f)
	
	f.蒙托邦Montauban = BMontauban蒙托邦
	f.蒙托邦Montauban.SetParent(f)
	
	f.蒙日斯卡尔Montgiscard = BMontgiscard蒙日斯卡尔
	f.蒙日斯卡尔Montgiscard.SetParent(f)
	
	f.米雷Muret = BMuret米雷
	f.米雷Muret.SetParent(f)
	
	f.图卢兹Toulouse = BToulouse图卢兹
	f.图卢兹Toulouse.SetParent(f)
	
}
