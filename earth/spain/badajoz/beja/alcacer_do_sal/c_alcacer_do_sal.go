package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Alcacer_do_salCounty interface {
    feud.County
    BAlcacerdosal萨尔堡() 	feud.Barony
    BAljustrel阿尔茹什特雷尔() 	feud.Barony
    BAlvalade阿尔瓦拉德() 	feud.Barony
    BCastroverde韦迪堡() 	feud.Barony
    BOdemira奥德米拉() 	feud.Barony
    BOurique欧里基() 	feud.Barony
    BSantiagodocacem圣地亚哥杜卡森() 	feud.Barony
}

type 萨尔堡Alcacer_do_salCounty struct {
	feud.BaseCounty
	萨尔堡Alcacerdosal 	feud.Barony
	阿尔茹什特雷尔Aljustrel 	feud.Barony
	阿尔瓦拉德Alvalade 	feud.Barony
	韦迪堡Castroverde 	feud.Barony
	奥德米拉Odemira 	feud.Barony
	欧里基Ourique 	feud.Barony
	圣地亚哥杜卡森Santiagodocacem 	feud.Barony
}

func (c *萨尔堡Alcacer_do_salCounty) BAlcacerdosal萨尔堡() feud.Barony {
	return c.萨尔堡Alcacerdosal
}
    
func (c *萨尔堡Alcacer_do_salCounty) BAljustrel阿尔茹什特雷尔() feud.Barony {
	return c.阿尔茹什特雷尔Aljustrel
}
    
func (c *萨尔堡Alcacer_do_salCounty) BAlvalade阿尔瓦拉德() feud.Barony {
	return c.阿尔瓦拉德Alvalade
}
    
func (c *萨尔堡Alcacer_do_salCounty) BCastroverde韦迪堡() feud.Barony {
	return c.韦迪堡Castroverde
}
    
func (c *萨尔堡Alcacer_do_salCounty) BOdemira奥德米拉() feud.Barony {
	return c.奥德米拉Odemira
}
    
func (c *萨尔堡Alcacer_do_salCounty) BOurique欧里基() feud.Barony {
	return c.欧里基Ourique
}
    
func (c *萨尔堡Alcacer_do_salCounty) BSantiagodocacem圣地亚哥杜卡森() feud.Barony {
	return c.圣地亚哥杜卡森Santiagodocacem
}
    
var CAlcacer_do_sal萨尔堡 Alcacer_do_salCounty = &萨尔堡Alcacer_do_salCounty{}

func init() {
	f := CAlcacer_do_sal萨尔堡.(*萨尔堡Alcacer_do_salCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "161",
		Title:     "alcacer_do_sal",
		TitleName: "萨尔堡",
		TitleCode: "c_alcacer_do_sal",
		Baronies:  map[string]feud.Barony{},
	}

	f.萨尔堡Alcacerdosal = BAlcacerdosal萨尔堡
	f.萨尔堡Alcacerdosal.SetParent(f)
	
	f.阿尔茹什特雷尔Aljustrel = BAljustrel阿尔茹什特雷尔
	f.阿尔茹什特雷尔Aljustrel.SetParent(f)
	
	f.阿尔瓦拉德Alvalade = BAlvalade阿尔瓦拉德
	f.阿尔瓦拉德Alvalade.SetParent(f)
	
	f.韦迪堡Castroverde = BCastroverde韦迪堡
	f.韦迪堡Castroverde.SetParent(f)
	
	f.奥德米拉Odemira = BOdemira奥德米拉
	f.奥德米拉Odemira.SetParent(f)
	
	f.欧里基Ourique = BOurique欧里基
	f.欧里基Ourique.SetParent(f)
	
	f.圣地亚哥杜卡森Santiagodocacem = BSantiagodocacem圣地亚哥杜卡森
	f.圣地亚哥杜卡森Santiagodocacem.SetParent(f)
	
}
