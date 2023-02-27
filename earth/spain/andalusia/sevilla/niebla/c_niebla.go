package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NieblaCounty interface {
    feud.County
    BAlmonte阿尔蒙特() 	feud.Barony
    BAroche阿罗谢() 	feud.Barony
    BLascrocinas克罗西纳斯() 	feud.Barony
    BNerva内尔瓦() 	feud.Barony
    BNiebla涅夫拉() 	feud.Barony
    BOstilia奥斯蒂利亚() 	feud.Barony
    BPilas皮拉斯() 	feud.Barony
}

type 涅夫拉NieblaCounty struct {
	feud.BaseCounty
	阿尔蒙特Almonte 	feud.Barony
	阿罗谢Aroche 	feud.Barony
	克罗西纳斯Lascrocinas 	feud.Barony
	内尔瓦Nerva 	feud.Barony
	涅夫拉Niebla 	feud.Barony
	奥斯蒂利亚Ostilia 	feud.Barony
	皮拉斯Pilas 	feud.Barony
}

func (c *涅夫拉NieblaCounty) BAlmonte阿尔蒙特() feud.Barony {
	return c.阿尔蒙特Almonte
}
    
func (c *涅夫拉NieblaCounty) BAroche阿罗谢() feud.Barony {
	return c.阿罗谢Aroche
}
    
func (c *涅夫拉NieblaCounty) BLascrocinas克罗西纳斯() feud.Barony {
	return c.克罗西纳斯Lascrocinas
}
    
func (c *涅夫拉NieblaCounty) BNerva内尔瓦() feud.Barony {
	return c.内尔瓦Nerva
}
    
func (c *涅夫拉NieblaCounty) BNiebla涅夫拉() feud.Barony {
	return c.涅夫拉Niebla
}
    
func (c *涅夫拉NieblaCounty) BOstilia奥斯蒂利亚() feud.Barony {
	return c.奥斯蒂利亚Ostilia
}
    
func (c *涅夫拉NieblaCounty) BPilas皮拉斯() feud.Barony {
	return c.皮拉斯Pilas
}
    
var CNiebla涅夫拉 NieblaCounty = &涅夫拉NieblaCounty{}

func init() {
	f := CNiebla涅夫拉.(*涅夫拉NieblaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "164",
		Title:     "niebla",
		TitleName: "涅夫拉",
		TitleCode: "c_niebla",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔蒙特Almonte = BAlmonte阿尔蒙特
	f.阿尔蒙特Almonte.SetParent(f)
	
	f.阿罗谢Aroche = BAroche阿罗谢
	f.阿罗谢Aroche.SetParent(f)
	
	f.克罗西纳斯Lascrocinas = BLascrocinas克罗西纳斯
	f.克罗西纳斯Lascrocinas.SetParent(f)
	
	f.内尔瓦Nerva = BNerva内尔瓦
	f.内尔瓦Nerva.SetParent(f)
	
	f.涅夫拉Niebla = BNiebla涅夫拉
	f.涅夫拉Niebla.SetParent(f)
	
	f.奥斯蒂利亚Ostilia = BOstilia奥斯蒂利亚
	f.奥斯蒂利亚Ostilia.SetParent(f)
	
	f.皮拉斯Pilas = BPilas皮拉斯
	f.皮拉斯Pilas.SetParent(f)
	
}
