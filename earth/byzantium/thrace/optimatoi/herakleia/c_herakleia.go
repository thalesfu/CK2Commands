package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HerakleiaCounty interface {
    feud.County
    BAmastrine阿马斯特里尼() 	feud.Barony
    BCalpe卡尔佩() 	feud.Barony
    BFlaviopolis弗拉维奥波利斯() 	feud.Barony
    BHerakleia赫拉克利亚() 	feud.Barony
    BMariandyni马里安底尼() 	feud.Barony
    BPolis波利斯() 	feud.Barony
    BZephyropoli泽菲洛波利() 	feud.Barony
}

type 赫拉克利亚HerakleiaCounty struct {
	feud.BaseCounty
	阿马斯特里尼Amastrine 	feud.Barony
	卡尔佩Calpe 	feud.Barony
	弗拉维奥波利斯Flaviopolis 	feud.Barony
	赫拉克利亚Herakleia 	feud.Barony
	马里安底尼Mariandyni 	feud.Barony
	波利斯Polis 	feud.Barony
	泽菲洛波利Zephyropoli 	feud.Barony
}

func (c *赫拉克利亚HerakleiaCounty) BAmastrine阿马斯特里尼() feud.Barony {
	return c.阿马斯特里尼Amastrine
}
    
func (c *赫拉克利亚HerakleiaCounty) BCalpe卡尔佩() feud.Barony {
	return c.卡尔佩Calpe
}
    
func (c *赫拉克利亚HerakleiaCounty) BFlaviopolis弗拉维奥波利斯() feud.Barony {
	return c.弗拉维奥波利斯Flaviopolis
}
    
func (c *赫拉克利亚HerakleiaCounty) BHerakleia赫拉克利亚() feud.Barony {
	return c.赫拉克利亚Herakleia
}
    
func (c *赫拉克利亚HerakleiaCounty) BMariandyni马里安底尼() feud.Barony {
	return c.马里安底尼Mariandyni
}
    
func (c *赫拉克利亚HerakleiaCounty) BPolis波利斯() feud.Barony {
	return c.波利斯Polis
}
    
func (c *赫拉克利亚HerakleiaCounty) BZephyropoli泽菲洛波利() feud.Barony {
	return c.泽菲洛波利Zephyropoli
}
    
var CHerakleia赫拉克利亚 HerakleiaCounty = &赫拉克利亚HerakleiaCounty{}

func init() {
	f := CHerakleia赫拉克利亚.(*赫拉克利亚HerakleiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "740",
		Title:     "herakleia",
		TitleName: "赫拉克利亚",
		TitleCode: "c_herakleia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿马斯特里尼Amastrine = BAmastrine阿马斯特里尼
	f.阿马斯特里尼Amastrine.SetParent(f)
	
	f.卡尔佩Calpe = BCalpe卡尔佩
	f.卡尔佩Calpe.SetParent(f)
	
	f.弗拉维奥波利斯Flaviopolis = BFlaviopolis弗拉维奥波利斯
	f.弗拉维奥波利斯Flaviopolis.SetParent(f)
	
	f.赫拉克利亚Herakleia = BHerakleia赫拉克利亚
	f.赫拉克利亚Herakleia.SetParent(f)
	
	f.马里安底尼Mariandyni = BMariandyni马里安底尼
	f.马里安底尼Mariandyni.SetParent(f)
	
	f.波利斯Polis = BPolis波利斯
	f.波利斯Polis.SetParent(f)
	
	f.泽菲洛波利Zephyropoli = BZephyropoli泽菲洛波利
	f.泽菲洛波利Zephyropoli.SetParent(f)
	
}
