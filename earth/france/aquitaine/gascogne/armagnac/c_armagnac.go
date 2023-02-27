package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArmagnacCounty interface {
    feud.County
    BAuch欧什() 	feud.Barony
    BCastelnau卡斯泰尔诺() 	feud.Barony
    BCondom孔东() 	feud.Barony
    BEauze欧兹() 	feud.Barony
    BLaplume拉普吕姆() 	feud.Barony
    BLectoure莱克图尔() 	feud.Barony
    BLisle利勒() 	feud.Barony
    BMirande米朗德() 	feud.Barony
}

type 阿马尼亚克ArmagnacCounty struct {
	feud.BaseCounty
	欧什Auch 	feud.Barony
	卡斯泰尔诺Castelnau 	feud.Barony
	孔东Condom 	feud.Barony
	欧兹Eauze 	feud.Barony
	拉普吕姆Laplume 	feud.Barony
	莱克图尔Lectoure 	feud.Barony
	利勒Lisle 	feud.Barony
	米朗德Mirande 	feud.Barony
}

func (c *阿马尼亚克ArmagnacCounty) BAuch欧什() feud.Barony {
	return c.欧什Auch
}
    
func (c *阿马尼亚克ArmagnacCounty) BCastelnau卡斯泰尔诺() feud.Barony {
	return c.卡斯泰尔诺Castelnau
}
    
func (c *阿马尼亚克ArmagnacCounty) BCondom孔东() feud.Barony {
	return c.孔东Condom
}
    
func (c *阿马尼亚克ArmagnacCounty) BEauze欧兹() feud.Barony {
	return c.欧兹Eauze
}
    
func (c *阿马尼亚克ArmagnacCounty) BLaplume拉普吕姆() feud.Barony {
	return c.拉普吕姆Laplume
}
    
func (c *阿马尼亚克ArmagnacCounty) BLectoure莱克图尔() feud.Barony {
	return c.莱克图尔Lectoure
}
    
func (c *阿马尼亚克ArmagnacCounty) BLisle利勒() feud.Barony {
	return c.利勒Lisle
}
    
func (c *阿马尼亚克ArmagnacCounty) BMirande米朗德() feud.Barony {
	return c.米朗德Mirande
}
    
var CArmagnac阿马尼亚克 ArmagnacCounty = &阿马尼亚克ArmagnacCounty{}

func init() {
	f := CArmagnac阿马尼亚克.(*阿马尼亚克ArmagnacCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "209",
		Title:     "armagnac",
		TitleName: "阿马尼亚克",
		TitleCode: "c_armagnac",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧什Auch = BAuch欧什
	f.欧什Auch.SetParent(f)
	
	f.卡斯泰尔诺Castelnau = BCastelnau卡斯泰尔诺
	f.卡斯泰尔诺Castelnau.SetParent(f)
	
	f.孔东Condom = BCondom孔东
	f.孔东Condom.SetParent(f)
	
	f.欧兹Eauze = BEauze欧兹
	f.欧兹Eauze.SetParent(f)
	
	f.拉普吕姆Laplume = BLaplume拉普吕姆
	f.拉普吕姆Laplume.SetParent(f)
	
	f.莱克图尔Lectoure = BLectoure莱克图尔
	f.莱克图尔Lectoure.SetParent(f)
	
	f.利勒Lisle = BLisle利勒
	f.利勒Lisle.SetParent(f)
	
	f.米朗德Mirande = BMirande米朗德
	f.米朗德Mirande.SetParent(f)
	
}
