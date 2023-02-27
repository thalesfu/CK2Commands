package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VestislandCounty interface {
    feud.County
    BAlftafjordur奥尔塔菲厄泽() 	feud.Barony
    BBorg博尔格() 	feud.Barony
    BHvamm赫瓦姆() 	feud.Barony
    BIsafjordur伊萨菲厄泽() 	feud.Barony
    BReykholar雷克侯拉尔() 	feud.Barony
    BSudureyri叙聚雷里() 	feud.Barony
    BVatnafjordur瓦特纳菲厄泽() 	feud.Barony
}

type 西冰岛VestislandCounty struct {
	feud.BaseCounty
	奥尔塔菲厄泽Alftafjordur 	feud.Barony
	博尔格Borg 	feud.Barony
	赫瓦姆Hvamm 	feud.Barony
	伊萨菲厄泽Isafjordur 	feud.Barony
	雷克侯拉尔Reykholar 	feud.Barony
	叙聚雷里Sudureyri 	feud.Barony
	瓦特纳菲厄泽Vatnafjordur 	feud.Barony
}

func (c *西冰岛VestislandCounty) BAlftafjordur奥尔塔菲厄泽() feud.Barony {
	return c.奥尔塔菲厄泽Alftafjordur
}
    
func (c *西冰岛VestislandCounty) BBorg博尔格() feud.Barony {
	return c.博尔格Borg
}
    
func (c *西冰岛VestislandCounty) BHvamm赫瓦姆() feud.Barony {
	return c.赫瓦姆Hvamm
}
    
func (c *西冰岛VestislandCounty) BIsafjordur伊萨菲厄泽() feud.Barony {
	return c.伊萨菲厄泽Isafjordur
}
    
func (c *西冰岛VestislandCounty) BReykholar雷克侯拉尔() feud.Barony {
	return c.雷克侯拉尔Reykholar
}
    
func (c *西冰岛VestislandCounty) BSudureyri叙聚雷里() feud.Barony {
	return c.叙聚雷里Sudureyri
}
    
func (c *西冰岛VestislandCounty) BVatnafjordur瓦特纳菲厄泽() feud.Barony {
	return c.瓦特纳菲厄泽Vatnafjordur
}
    
var CVestisland西冰岛 VestislandCounty = &西冰岛VestislandCounty{}

func init() {
	f := CVestisland西冰岛.(*西冰岛VestislandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1",
		Title:     "vestisland",
		TitleName: "西冰岛",
		TitleCode: "c_vestisland",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥尔塔菲厄泽Alftafjordur = BAlftafjordur奥尔塔菲厄泽
	f.奥尔塔菲厄泽Alftafjordur.SetParent(f)
	
	f.博尔格Borg = BBorg博尔格
	f.博尔格Borg.SetParent(f)
	
	f.赫瓦姆Hvamm = BHvamm赫瓦姆
	f.赫瓦姆Hvamm.SetParent(f)
	
	f.伊萨菲厄泽Isafjordur = BIsafjordur伊萨菲厄泽
	f.伊萨菲厄泽Isafjordur.SetParent(f)
	
	f.雷克侯拉尔Reykholar = BReykholar雷克侯拉尔
	f.雷克侯拉尔Reykholar.SetParent(f)
	
	f.叙聚雷里Sudureyri = BSudureyri叙聚雷里
	f.叙聚雷里Sudureyri.SetParent(f)
	
	f.瓦特纳菲厄泽Vatnafjordur = BVatnafjordur瓦特纳菲厄泽
	f.瓦特纳菲厄泽Vatnafjordur.SetParent(f)
	
}
