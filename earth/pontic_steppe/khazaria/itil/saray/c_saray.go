package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarayCounty interface {
    feud.County
    BAkhtuba阿赫图巴() 	feud.Barony
    BBataevka巴塔耶夫卡() 	feud.Barony
    BChernyyar黑峡谷() 	feud.Barony
    BDzhelga杰尔加() 	feud.Barony
    BPokrovka波克罗夫卡() 	feud.Barony
    BSaray萨赖() 	feud.Barony
    BTsaganaman察甘_阿曼() 	feud.Barony
    BUspenka乌斯片卡() 	feud.Barony
}

type 萨赖SarayCounty struct {
	feud.BaseCounty
	阿赫图巴Akhtuba 	feud.Barony
	巴塔耶夫卡Bataevka 	feud.Barony
	黑峡谷Chernyyar 	feud.Barony
	杰尔加Dzhelga 	feud.Barony
	波克罗夫卡Pokrovka 	feud.Barony
	萨赖Saray 	feud.Barony
	察甘_阿曼Tsaganaman 	feud.Barony
	乌斯片卡Uspenka 	feud.Barony
}

func (c *萨赖SarayCounty) BAkhtuba阿赫图巴() feud.Barony {
	return c.阿赫图巴Akhtuba
}
    
func (c *萨赖SarayCounty) BBataevka巴塔耶夫卡() feud.Barony {
	return c.巴塔耶夫卡Bataevka
}
    
func (c *萨赖SarayCounty) BChernyyar黑峡谷() feud.Barony {
	return c.黑峡谷Chernyyar
}
    
func (c *萨赖SarayCounty) BDzhelga杰尔加() feud.Barony {
	return c.杰尔加Dzhelga
}
    
func (c *萨赖SarayCounty) BPokrovka波克罗夫卡() feud.Barony {
	return c.波克罗夫卡Pokrovka
}
    
func (c *萨赖SarayCounty) BSaray萨赖() feud.Barony {
	return c.萨赖Saray
}
    
func (c *萨赖SarayCounty) BTsaganaman察甘_阿曼() feud.Barony {
	return c.察甘_阿曼Tsaganaman
}
    
func (c *萨赖SarayCounty) BUspenka乌斯片卡() feud.Barony {
	return c.乌斯片卡Uspenka
}
    
var CSaray萨赖 SarayCounty = &萨赖SarayCounty{}

func init() {
	f := CSaray萨赖.(*萨赖SarayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "619",
		Title:     "saray",
		TitleName: "萨赖",
		TitleCode: "c_saray",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫图巴Akhtuba = BAkhtuba阿赫图巴
	f.阿赫图巴Akhtuba.SetParent(f)
	
	f.巴塔耶夫卡Bataevka = BBataevka巴塔耶夫卡
	f.巴塔耶夫卡Bataevka.SetParent(f)
	
	f.黑峡谷Chernyyar = BChernyyar黑峡谷
	f.黑峡谷Chernyyar.SetParent(f)
	
	f.杰尔加Dzhelga = BDzhelga杰尔加
	f.杰尔加Dzhelga.SetParent(f)
	
	f.波克罗夫卡Pokrovka = BPokrovka波克罗夫卡
	f.波克罗夫卡Pokrovka.SetParent(f)
	
	f.萨赖Saray = BSaray萨赖
	f.萨赖Saray.SetParent(f)
	
	f.察甘_阿曼Tsaganaman = BTsaganaman察甘_阿曼
	f.察甘_阿曼Tsaganaman.SetParent(f)
	
	f.乌斯片卡Uspenka = BUspenka乌斯片卡
	f.乌斯片卡Uspenka.SetParent(f)
	
}
