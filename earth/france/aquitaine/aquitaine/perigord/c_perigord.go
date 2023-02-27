package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PerigordCounty interface {
    feud.County
    BAuberoche欧布罗什() 	feud.Barony
    BBaneuil巴纳伊() 	feud.Barony
    BBergerac贝尔热拉克() 	feud.Barony
    BBiron比龙() 	feud.Barony
    BBonaguil博纳吉耶() 	feud.Barony
    BChancelade尚瑟拉德() 	feud.Barony
    BPerigueux佩里格() 	feud.Barony
    BSarlat萨尔拉() 	feud.Barony
}

type 佩里戈尔PerigordCounty struct {
	feud.BaseCounty
	欧布罗什Auberoche 	feud.Barony
	巴纳伊Baneuil 	feud.Barony
	贝尔热拉克Bergerac 	feud.Barony
	比龙Biron 	feud.Barony
	博纳吉耶Bonaguil 	feud.Barony
	尚瑟拉德Chancelade 	feud.Barony
	佩里格Perigueux 	feud.Barony
	萨尔拉Sarlat 	feud.Barony
}

func (c *佩里戈尔PerigordCounty) BAuberoche欧布罗什() feud.Barony {
	return c.欧布罗什Auberoche
}
    
func (c *佩里戈尔PerigordCounty) BBaneuil巴纳伊() feud.Barony {
	return c.巴纳伊Baneuil
}
    
func (c *佩里戈尔PerigordCounty) BBergerac贝尔热拉克() feud.Barony {
	return c.贝尔热拉克Bergerac
}
    
func (c *佩里戈尔PerigordCounty) BBiron比龙() feud.Barony {
	return c.比龙Biron
}
    
func (c *佩里戈尔PerigordCounty) BBonaguil博纳吉耶() feud.Barony {
	return c.博纳吉耶Bonaguil
}
    
func (c *佩里戈尔PerigordCounty) BChancelade尚瑟拉德() feud.Barony {
	return c.尚瑟拉德Chancelade
}
    
func (c *佩里戈尔PerigordCounty) BPerigueux佩里格() feud.Barony {
	return c.佩里格Perigueux
}
    
func (c *佩里戈尔PerigordCounty) BSarlat萨尔拉() feud.Barony {
	return c.萨尔拉Sarlat
}
    
var CPerigord佩里戈尔 PerigordCounty = &佩里戈尔PerigordCounty{}

func init() {
	f := CPerigord佩里戈尔.(*佩里戈尔PerigordCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "216",
		Title:     "perigord",
		TitleName: "佩里戈尔",
		TitleCode: "c_perigord",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧布罗什Auberoche = BAuberoche欧布罗什
	f.欧布罗什Auberoche.SetParent(f)
	
	f.巴纳伊Baneuil = BBaneuil巴纳伊
	f.巴纳伊Baneuil.SetParent(f)
	
	f.贝尔热拉克Bergerac = BBergerac贝尔热拉克
	f.贝尔热拉克Bergerac.SetParent(f)
	
	f.比龙Biron = BBiron比龙
	f.比龙Biron.SetParent(f)
	
	f.博纳吉耶Bonaguil = BBonaguil博纳吉耶
	f.博纳吉耶Bonaguil.SetParent(f)
	
	f.尚瑟拉德Chancelade = BChancelade尚瑟拉德
	f.尚瑟拉德Chancelade.SetParent(f)
	
	f.佩里格Perigueux = BPerigueux佩里格
	f.佩里格Perigueux.SetParent(f)
	
	f.萨尔拉Sarlat = BSarlat萨尔拉
	f.萨尔拉Sarlat.SetParent(f)
	
}
