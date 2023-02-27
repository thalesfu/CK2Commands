package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TunisCounty interface {
    feud.County
    BAriana艾尔亚奈() 	feud.Barony
    BBenarous本阿鲁斯() 	feud.Barony
    BCartaghe迦太基() 	feud.Barony
    BHammamet哈马马特() 	feud.Barony
    BKelibia古莱比耶() 	feud.Barony
    BSousse苏塞() 	feud.Barony
    BTunis突尼斯() 	feud.Barony
    BZaghouan宰格万() 	feud.Barony
}

type 突尼斯TunisCounty struct {
	feud.BaseCounty
	艾尔亚奈Ariana 	feud.Barony
	本阿鲁斯Benarous 	feud.Barony
	迦太基Cartaghe 	feud.Barony
	哈马马特Hammamet 	feud.Barony
	古莱比耶Kelibia 	feud.Barony
	苏塞Sousse 	feud.Barony
	突尼斯Tunis 	feud.Barony
	宰格万Zaghouan 	feud.Barony
}

func (c *突尼斯TunisCounty) BAriana艾尔亚奈() feud.Barony {
	return c.艾尔亚奈Ariana
}
    
func (c *突尼斯TunisCounty) BBenarous本阿鲁斯() feud.Barony {
	return c.本阿鲁斯Benarous
}
    
func (c *突尼斯TunisCounty) BCartaghe迦太基() feud.Barony {
	return c.迦太基Cartaghe
}
    
func (c *突尼斯TunisCounty) BHammamet哈马马特() feud.Barony {
	return c.哈马马特Hammamet
}
    
func (c *突尼斯TunisCounty) BKelibia古莱比耶() feud.Barony {
	return c.古莱比耶Kelibia
}
    
func (c *突尼斯TunisCounty) BSousse苏塞() feud.Barony {
	return c.苏塞Sousse
}
    
func (c *突尼斯TunisCounty) BTunis突尼斯() feud.Barony {
	return c.突尼斯Tunis
}
    
func (c *突尼斯TunisCounty) BZaghouan宰格万() feud.Barony {
	return c.宰格万Zaghouan
}
    
var CTunis突尼斯 TunisCounty = &突尼斯TunisCounty{}

func init() {
	f := CTunis突尼斯.(*突尼斯TunisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "817",
		Title:     "tunis",
		TitleName: "突尼斯",
		TitleCode: "c_tunis",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾尔亚奈Ariana = BAriana艾尔亚奈
	f.艾尔亚奈Ariana.SetParent(f)
	
	f.本阿鲁斯Benarous = BBenarous本阿鲁斯
	f.本阿鲁斯Benarous.SetParent(f)
	
	f.迦太基Cartaghe = BCartaghe迦太基
	f.迦太基Cartaghe.SetParent(f)
	
	f.哈马马特Hammamet = BHammamet哈马马特
	f.哈马马特Hammamet.SetParent(f)
	
	f.古莱比耶Kelibia = BKelibia古莱比耶
	f.古莱比耶Kelibia.SetParent(f)
	
	f.苏塞Sousse = BSousse苏塞
	f.苏塞Sousse.SetParent(f)
	
	f.突尼斯Tunis = BTunis突尼斯
	f.突尼斯Tunis.SetParent(f)
	
	f.宰格万Zaghouan = BZaghouan宰格万
	f.宰格万Zaghouan.SetParent(f)
	
}
