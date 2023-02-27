package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PortoCounty interface {
    feud.County
    BArcosdevaldevez阿尔库什迪瓦尔德维什() 	feud.Barony
    BBarcelos巴塞卢什() 	feud.Barony
    BBraga布拉加() 	feud.Barony
    BGuimaraes吉马朗伊斯() 	feud.Barony
    BMoncao蒙桑() 	feud.Barony
    BPontedelima蓬蒂迪利马() 	feud.Barony
    BPorto波尔图() 	feud.Barony
    BVianadocastelo维亚纳堡() 	feud.Barony
}

type 波尔图PortoCounty struct {
	feud.BaseCounty
	阿尔库什迪瓦尔德维什Arcosdevaldevez 	feud.Barony
	巴塞卢什Barcelos 	feud.Barony
	布拉加Braga 	feud.Barony
	吉马朗伊斯Guimaraes 	feud.Barony
	蒙桑Moncao 	feud.Barony
	蓬蒂迪利马Pontedelima 	feud.Barony
	波尔图Porto 	feud.Barony
	维亚纳堡Vianadocastelo 	feud.Barony
}

func (c *波尔图PortoCounty) BArcosdevaldevez阿尔库什迪瓦尔德维什() feud.Barony {
	return c.阿尔库什迪瓦尔德维什Arcosdevaldevez
}
    
func (c *波尔图PortoCounty) BBarcelos巴塞卢什() feud.Barony {
	return c.巴塞卢什Barcelos
}
    
func (c *波尔图PortoCounty) BBraga布拉加() feud.Barony {
	return c.布拉加Braga
}
    
func (c *波尔图PortoCounty) BGuimaraes吉马朗伊斯() feud.Barony {
	return c.吉马朗伊斯Guimaraes
}
    
func (c *波尔图PortoCounty) BMoncao蒙桑() feud.Barony {
	return c.蒙桑Moncao
}
    
func (c *波尔图PortoCounty) BPontedelima蓬蒂迪利马() feud.Barony {
	return c.蓬蒂迪利马Pontedelima
}
    
func (c *波尔图PortoCounty) BPorto波尔图() feud.Barony {
	return c.波尔图Porto
}
    
func (c *波尔图PortoCounty) BVianadocastelo维亚纳堡() feud.Barony {
	return c.维亚纳堡Vianadocastelo
}
    
var CPorto波尔图 PortoCounty = &波尔图PortoCounty{}

func init() {
	f := CPorto波尔图.(*波尔图PortoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "158",
		Title:     "porto",
		TitleName: "波尔图",
		TitleCode: "c_porto",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔库什迪瓦尔德维什Arcosdevaldevez = BArcosdevaldevez阿尔库什迪瓦尔德维什
	f.阿尔库什迪瓦尔德维什Arcosdevaldevez.SetParent(f)
	
	f.巴塞卢什Barcelos = BBarcelos巴塞卢什
	f.巴塞卢什Barcelos.SetParent(f)
	
	f.布拉加Braga = BBraga布拉加
	f.布拉加Braga.SetParent(f)
	
	f.吉马朗伊斯Guimaraes = BGuimaraes吉马朗伊斯
	f.吉马朗伊斯Guimaraes.SetParent(f)
	
	f.蒙桑Moncao = BMoncao蒙桑
	f.蒙桑Moncao.SetParent(f)
	
	f.蓬蒂迪利马Pontedelima = BPontedelima蓬蒂迪利马
	f.蓬蒂迪利马Pontedelima.SetParent(f)
	
	f.波尔图Porto = BPorto波尔图
	f.波尔图Porto.SetParent(f)
	
	f.维亚纳堡Vianadocastelo = BVianadocastelo维亚纳堡
	f.维亚纳堡Vianadocastelo.SetParent(f)
	
}
