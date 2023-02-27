package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LaukaaCounty interface {
    feud.County
    BHankasalmi汉卡萨尔米() 	feud.Barony
    BJyvaskyla于韦斯屈莱() 	feud.Barony
    BKeuruu凯乌鲁() 	feud.Barony
    BKuusvesi库斯韦西() 	feud.Barony
    BPeurunka佩乌伦卡() 	feud.Barony
    BSaraavesi萨拉韦西() 	feud.Barony
    BSuolahti苏奥拉赫蒂() 	feud.Barony
}

type 凯泰莱LaukaaCounty struct {
	feud.BaseCounty
	汉卡萨尔米Hankasalmi 	feud.Barony
	于韦斯屈莱Jyvaskyla 	feud.Barony
	凯乌鲁Keuruu 	feud.Barony
	库斯韦西Kuusvesi 	feud.Barony
	佩乌伦卡Peurunka 	feud.Barony
	萨拉韦西Saraavesi 	feud.Barony
	苏奥拉赫蒂Suolahti 	feud.Barony
}

func (c *凯泰莱LaukaaCounty) BHankasalmi汉卡萨尔米() feud.Barony {
	return c.汉卡萨尔米Hankasalmi
}
    
func (c *凯泰莱LaukaaCounty) BJyvaskyla于韦斯屈莱() feud.Barony {
	return c.于韦斯屈莱Jyvaskyla
}
    
func (c *凯泰莱LaukaaCounty) BKeuruu凯乌鲁() feud.Barony {
	return c.凯乌鲁Keuruu
}
    
func (c *凯泰莱LaukaaCounty) BKuusvesi库斯韦西() feud.Barony {
	return c.库斯韦西Kuusvesi
}
    
func (c *凯泰莱LaukaaCounty) BPeurunka佩乌伦卡() feud.Barony {
	return c.佩乌伦卡Peurunka
}
    
func (c *凯泰莱LaukaaCounty) BSaraavesi萨拉韦西() feud.Barony {
	return c.萨拉韦西Saraavesi
}
    
func (c *凯泰莱LaukaaCounty) BSuolahti苏奥拉赫蒂() feud.Barony {
	return c.苏奥拉赫蒂Suolahti
}
    
var CLaukaa凯泰莱 LaukaaCounty = &凯泰莱LaukaaCounty{}

func init() {
	f := CLaukaa凯泰莱.(*凯泰莱LaukaaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1607",
		Title:     "laukaa",
		TitleName: "凯泰莱",
		TitleCode: "c_laukaa",
		Baronies:  map[string]feud.Barony{},
	}

	f.汉卡萨尔米Hankasalmi = BHankasalmi汉卡萨尔米
	f.汉卡萨尔米Hankasalmi.SetParent(f)
	
	f.于韦斯屈莱Jyvaskyla = BJyvaskyla于韦斯屈莱
	f.于韦斯屈莱Jyvaskyla.SetParent(f)
	
	f.凯乌鲁Keuruu = BKeuruu凯乌鲁
	f.凯乌鲁Keuruu.SetParent(f)
	
	f.库斯韦西Kuusvesi = BKuusvesi库斯韦西
	f.库斯韦西Kuusvesi.SetParent(f)
	
	f.佩乌伦卡Peurunka = BPeurunka佩乌伦卡
	f.佩乌伦卡Peurunka.SetParent(f)
	
	f.萨拉韦西Saraavesi = BSaraavesi萨拉韦西
	f.萨拉韦西Saraavesi.SetParent(f)
	
	f.苏奥拉赫蒂Suolahti = BSuolahti苏奥拉赫蒂
	f.苏奥拉赫蒂Suolahti.SetParent(f)
	
}
