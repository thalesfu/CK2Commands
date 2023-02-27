package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PoherCounty interface {
    feud.County
    BCarhaix卡赖() 	feud.Barony
    BChateauneuf_du_faou法乌新堡() 	feud.Barony
    BCorlay科尔莱() 	feud.Barony
    BHuelgoat于埃尔戈阿() 	feud.Barony
    BLangonnet朗戈内() 	feud.Barony
    BPlouguer普卢盖尔() 	feud.Barony
    BRostrenen罗斯特雷南() 	feud.Barony
}

type 波埃PoherCounty struct {
	feud.BaseCounty
	卡赖Carhaix 	feud.Barony
	法乌新堡Chateauneuf_du_faou 	feud.Barony
	科尔莱Corlay 	feud.Barony
	于埃尔戈阿Huelgoat 	feud.Barony
	朗戈内Langonnet 	feud.Barony
	普卢盖尔Plouguer 	feud.Barony
	罗斯特雷南Rostrenen 	feud.Barony
}

func (c *波埃PoherCounty) BCarhaix卡赖() feud.Barony {
	return c.卡赖Carhaix
}
    
func (c *波埃PoherCounty) BChateauneuf_du_faou法乌新堡() feud.Barony {
	return c.法乌新堡Chateauneuf_du_faou
}
    
func (c *波埃PoherCounty) BCorlay科尔莱() feud.Barony {
	return c.科尔莱Corlay
}
    
func (c *波埃PoherCounty) BHuelgoat于埃尔戈阿() feud.Barony {
	return c.于埃尔戈阿Huelgoat
}
    
func (c *波埃PoherCounty) BLangonnet朗戈内() feud.Barony {
	return c.朗戈内Langonnet
}
    
func (c *波埃PoherCounty) BPlouguer普卢盖尔() feud.Barony {
	return c.普卢盖尔Plouguer
}
    
func (c *波埃PoherCounty) BRostrenen罗斯特雷南() feud.Barony {
	return c.罗斯特雷南Rostrenen
}
    
var CPoher波埃 PoherCounty = &波埃PoherCounty{}

func init() {
	f := CPoher波埃.(*波埃PoherCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1937",
		Title:     "poher",
		TitleName: "波埃",
		TitleCode: "c_poher",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡赖Carhaix = BCarhaix卡赖
	f.卡赖Carhaix.SetParent(f)
	
	f.法乌新堡Chateauneuf_du_faou = BChateauneuf_du_faou法乌新堡
	f.法乌新堡Chateauneuf_du_faou.SetParent(f)
	
	f.科尔莱Corlay = BCorlay科尔莱
	f.科尔莱Corlay.SetParent(f)
	
	f.于埃尔戈阿Huelgoat = BHuelgoat于埃尔戈阿
	f.于埃尔戈阿Huelgoat.SetParent(f)
	
	f.朗戈内Langonnet = BLangonnet朗戈内
	f.朗戈内Langonnet.SetParent(f)
	
	f.普卢盖尔Plouguer = BPlouguer普卢盖尔
	f.普卢盖尔Plouguer.SetParent(f)
	
	f.罗斯特雷南Rostrenen = BRostrenen罗斯特雷南
	f.罗斯特雷南Rostrenen.SetParent(f)
	
}
