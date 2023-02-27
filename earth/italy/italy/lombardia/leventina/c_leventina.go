package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeventinaCounty interface {
    feud.County
    BBellinzona贝林佐纳() 	feud.Barony
    BBiasca比亚斯卡() 	feud.Barony
    BCastelgrande卡斯特尔格兰德() 	feud.Barony
    BFaedo法埃多() 	feud.Barony
    BLocarno洛迦诺() 	feud.Barony
    BLugano卢加诺() 	feud.Barony
    BMontebello_leventina蒙特贝洛() 	feud.Barony
}

type 贝林佐纳LeventinaCounty struct {
	feud.BaseCounty
	贝林佐纳Bellinzona 	feud.Barony
	比亚斯卡Biasca 	feud.Barony
	卡斯特尔格兰德Castelgrande 	feud.Barony
	法埃多Faedo 	feud.Barony
	洛迦诺Locarno 	feud.Barony
	卢加诺Lugano 	feud.Barony
	蒙特贝洛Montebello_leventina 	feud.Barony
}

func (c *贝林佐纳LeventinaCounty) BBellinzona贝林佐纳() feud.Barony {
	return c.贝林佐纳Bellinzona
}
    
func (c *贝林佐纳LeventinaCounty) BBiasca比亚斯卡() feud.Barony {
	return c.比亚斯卡Biasca
}
    
func (c *贝林佐纳LeventinaCounty) BCastelgrande卡斯特尔格兰德() feud.Barony {
	return c.卡斯特尔格兰德Castelgrande
}
    
func (c *贝林佐纳LeventinaCounty) BFaedo法埃多() feud.Barony {
	return c.法埃多Faedo
}
    
func (c *贝林佐纳LeventinaCounty) BLocarno洛迦诺() feud.Barony {
	return c.洛迦诺Locarno
}
    
func (c *贝林佐纳LeventinaCounty) BLugano卢加诺() feud.Barony {
	return c.卢加诺Lugano
}
    
func (c *贝林佐纳LeventinaCounty) BMontebello_leventina蒙特贝洛() feud.Barony {
	return c.蒙特贝洛Montebello_leventina
}
    
var CLeventina贝林佐纳 LeventinaCounty = &贝林佐纳LeventinaCounty{}

func init() {
	f := CLeventina贝林佐纳.(*贝林佐纳LeventinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1612",
		Title:     "leventina",
		TitleName: "贝林佐纳",
		TitleCode: "c_leventina",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝林佐纳Bellinzona = BBellinzona贝林佐纳
	f.贝林佐纳Bellinzona.SetParent(f)
	
	f.比亚斯卡Biasca = BBiasca比亚斯卡
	f.比亚斯卡Biasca.SetParent(f)
	
	f.卡斯特尔格兰德Castelgrande = BCastelgrande卡斯特尔格兰德
	f.卡斯特尔格兰德Castelgrande.SetParent(f)
	
	f.法埃多Faedo = BFaedo法埃多
	f.法埃多Faedo.SetParent(f)
	
	f.洛迦诺Locarno = BLocarno洛迦诺
	f.洛迦诺Locarno.SetParent(f)
	
	f.卢加诺Lugano = BLugano卢加诺
	f.卢加诺Lugano.SetParent(f)
	
	f.蒙特贝洛Montebello_leventina = BMontebello_leventina蒙特贝洛
	f.蒙特贝洛Montebello_leventina.SetParent(f)
	
}
