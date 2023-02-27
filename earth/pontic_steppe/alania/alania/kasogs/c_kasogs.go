package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KasogsCounty interface {
    feud.County
    BBakhtarakhtar阿赫塔尔巴赫塔() 	feud.Barony
    BEltarkan厄尔塔拉克() 	feud.Barony
    BKoshkhab科舍尔哈布利() 	feud.Barony
    BMyequape迈科普() 	feud.Barony
    BSamiran萨米兰() 	feud.Barony
    BSarytyuz萨雷_秋兹() 	feud.Barony
    BUstdzeghuta乌斯季_杰古塔() 	feud.Barony
    BZikh济赫() 	feud.Barony
}

type 卡索吉亚KasogsCounty struct {
	feud.BaseCounty
	阿赫塔尔巴赫塔Bakhtarakhtar 	feud.Barony
	厄尔塔拉克Eltarkan 	feud.Barony
	科舍尔哈布利Koshkhab 	feud.Barony
	迈科普Myequape 	feud.Barony
	萨米兰Samiran 	feud.Barony
	萨雷_秋兹Sarytyuz 	feud.Barony
	乌斯季_杰古塔Ustdzeghuta 	feud.Barony
	济赫Zikh 	feud.Barony
}

func (c *卡索吉亚KasogsCounty) BBakhtarakhtar阿赫塔尔巴赫塔() feud.Barony {
	return c.阿赫塔尔巴赫塔Bakhtarakhtar
}
    
func (c *卡索吉亚KasogsCounty) BEltarkan厄尔塔拉克() feud.Barony {
	return c.厄尔塔拉克Eltarkan
}
    
func (c *卡索吉亚KasogsCounty) BKoshkhab科舍尔哈布利() feud.Barony {
	return c.科舍尔哈布利Koshkhab
}
    
func (c *卡索吉亚KasogsCounty) BMyequape迈科普() feud.Barony {
	return c.迈科普Myequape
}
    
func (c *卡索吉亚KasogsCounty) BSamiran萨米兰() feud.Barony {
	return c.萨米兰Samiran
}
    
func (c *卡索吉亚KasogsCounty) BSarytyuz萨雷_秋兹() feud.Barony {
	return c.萨雷_秋兹Sarytyuz
}
    
func (c *卡索吉亚KasogsCounty) BUstdzeghuta乌斯季_杰古塔() feud.Barony {
	return c.乌斯季_杰古塔Ustdzeghuta
}
    
func (c *卡索吉亚KasogsCounty) BZikh济赫() feud.Barony {
	return c.济赫Zikh
}
    
var CKasogs卡索吉亚 KasogsCounty = &卡索吉亚KasogsCounty{}

func init() {
	f := CKasogs卡索吉亚.(*卡索吉亚KasogsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "602",
		Title:     "kasogs",
		TitleName: "卡索吉亚",
		TitleCode: "c_kasogs",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫塔尔巴赫塔Bakhtarakhtar = BBakhtarakhtar阿赫塔尔巴赫塔
	f.阿赫塔尔巴赫塔Bakhtarakhtar.SetParent(f)
	
	f.厄尔塔拉克Eltarkan = BEltarkan厄尔塔拉克
	f.厄尔塔拉克Eltarkan.SetParent(f)
	
	f.科舍尔哈布利Koshkhab = BKoshkhab科舍尔哈布利
	f.科舍尔哈布利Koshkhab.SetParent(f)
	
	f.迈科普Myequape = BMyequape迈科普
	f.迈科普Myequape.SetParent(f)
	
	f.萨米兰Samiran = BSamiran萨米兰
	f.萨米兰Samiran.SetParent(f)
	
	f.萨雷_秋兹Sarytyuz = BSarytyuz萨雷_秋兹
	f.萨雷_秋兹Sarytyuz.SetParent(f)
	
	f.乌斯季_杰古塔Ustdzeghuta = BUstdzeghuta乌斯季_杰古塔
	f.乌斯季_杰古塔Ustdzeghuta.SetParent(f)
	
	f.济赫Zikh = BZikh济赫
	f.济赫Zikh.SetParent(f)
	
}
