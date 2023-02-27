package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NikomedeiaCounty interface {
    feud.County
    BAdapazari阿达帕扎里() 	feud.Barony
    BChalkedon迦克墩() 	feud.Barony
    BChrysopolis克里索波利斯() 	feud.Barony
    BMalagina马拉吉纳() 	feud.Barony
    BNikomedeia尼科米底亚() 	feud.Barony
    BPalodes帕洛德斯() 	feud.Barony
    BPraenetos普莱() 	feud.Barony
}

type 尼科米底亚NikomedeiaCounty struct {
	feud.BaseCounty
	阿达帕扎里Adapazari 	feud.Barony
	迦克墩Chalkedon 	feud.Barony
	克里索波利斯Chrysopolis 	feud.Barony
	马拉吉纳Malagina 	feud.Barony
	尼科米底亚Nikomedeia 	feud.Barony
	帕洛德斯Palodes 	feud.Barony
	普莱Praenetos 	feud.Barony
}

func (c *尼科米底亚NikomedeiaCounty) BAdapazari阿达帕扎里() feud.Barony {
	return c.阿达帕扎里Adapazari
}
    
func (c *尼科米底亚NikomedeiaCounty) BChalkedon迦克墩() feud.Barony {
	return c.迦克墩Chalkedon
}
    
func (c *尼科米底亚NikomedeiaCounty) BChrysopolis克里索波利斯() feud.Barony {
	return c.克里索波利斯Chrysopolis
}
    
func (c *尼科米底亚NikomedeiaCounty) BMalagina马拉吉纳() feud.Barony {
	return c.马拉吉纳Malagina
}
    
func (c *尼科米底亚NikomedeiaCounty) BNikomedeia尼科米底亚() feud.Barony {
	return c.尼科米底亚Nikomedeia
}
    
func (c *尼科米底亚NikomedeiaCounty) BPalodes帕洛德斯() feud.Barony {
	return c.帕洛德斯Palodes
}
    
func (c *尼科米底亚NikomedeiaCounty) BPraenetos普莱() feud.Barony {
	return c.普莱Praenetos
}
    
var CNikomedeia尼科米底亚 NikomedeiaCounty = &尼科米底亚NikomedeiaCounty{}

func init() {
	f := CNikomedeia尼科米底亚.(*尼科米底亚NikomedeiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "741",
		Title:     "nikomedeia",
		TitleName: "尼科米底亚",
		TitleCode: "c_nikomedeia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿达帕扎里Adapazari = BAdapazari阿达帕扎里
	f.阿达帕扎里Adapazari.SetParent(f)
	
	f.迦克墩Chalkedon = BChalkedon迦克墩
	f.迦克墩Chalkedon.SetParent(f)
	
	f.克里索波利斯Chrysopolis = BChrysopolis克里索波利斯
	f.克里索波利斯Chrysopolis.SetParent(f)
	
	f.马拉吉纳Malagina = BMalagina马拉吉纳
	f.马拉吉纳Malagina.SetParent(f)
	
	f.尼科米底亚Nikomedeia = BNikomedeia尼科米底亚
	f.尼科米底亚Nikomedeia.SetParent(f)
	
	f.帕洛德斯Palodes = BPalodes帕洛德斯
	f.帕洛德斯Palodes.SetParent(f)
	
	f.普莱Praenetos = BPraenetos普莱
	f.普莱Praenetos.SetParent(f)
	
}
