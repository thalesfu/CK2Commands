package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CrimeaCounty interface {
    feud.County
    BAqmescit阿克梅斯吉特() 	feud.Barony
    BBakhchisaray巴赫奇萨赖() 	feud.Barony
    BDzhankoy占科伊() 	feud.Barony
    BKezlev克兹列夫() 	feud.Barony
    BPerekop佩列科普() 	feud.Barony
    BQarasuvbazar卡拉苏夫巴扎尔() 	feud.Barony
    BQurman库尔曼() 	feud.Barony
    BSaq萨克() 	feud.Barony
}

type 克里米亚CrimeaCounty struct {
	feud.BaseCounty
	阿克梅斯吉特Aqmescit 	feud.Barony
	巴赫奇萨赖Bakhchisaray 	feud.Barony
	占科伊Dzhankoy 	feud.Barony
	克兹列夫Kezlev 	feud.Barony
	佩列科普Perekop 	feud.Barony
	卡拉苏夫巴扎尔Qarasuvbazar 	feud.Barony
	库尔曼Qurman 	feud.Barony
	萨克Saq 	feud.Barony
}

func (c *克里米亚CrimeaCounty) BAqmescit阿克梅斯吉特() feud.Barony {
	return c.阿克梅斯吉特Aqmescit
}
    
func (c *克里米亚CrimeaCounty) BBakhchisaray巴赫奇萨赖() feud.Barony {
	return c.巴赫奇萨赖Bakhchisaray
}
    
func (c *克里米亚CrimeaCounty) BDzhankoy占科伊() feud.Barony {
	return c.占科伊Dzhankoy
}
    
func (c *克里米亚CrimeaCounty) BKezlev克兹列夫() feud.Barony {
	return c.克兹列夫Kezlev
}
    
func (c *克里米亚CrimeaCounty) BPerekop佩列科普() feud.Barony {
	return c.佩列科普Perekop
}
    
func (c *克里米亚CrimeaCounty) BQarasuvbazar卡拉苏夫巴扎尔() feud.Barony {
	return c.卡拉苏夫巴扎尔Qarasuvbazar
}
    
func (c *克里米亚CrimeaCounty) BQurman库尔曼() feud.Barony {
	return c.库尔曼Qurman
}
    
func (c *克里米亚CrimeaCounty) BSaq萨克() feud.Barony {
	return c.萨克Saq
}
    
var CCrimea克里米亚 CrimeaCounty = &克里米亚CrimeaCounty{}

func init() {
	f := CCrimea克里米亚.(*克里米亚CrimeaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "559",
		Title:     "crimea",
		TitleName: "克里米亚",
		TitleCode: "c_crimea",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克梅斯吉特Aqmescit = BAqmescit阿克梅斯吉特
	f.阿克梅斯吉特Aqmescit.SetParent(f)
	
	f.巴赫奇萨赖Bakhchisaray = BBakhchisaray巴赫奇萨赖
	f.巴赫奇萨赖Bakhchisaray.SetParent(f)
	
	f.占科伊Dzhankoy = BDzhankoy占科伊
	f.占科伊Dzhankoy.SetParent(f)
	
	f.克兹列夫Kezlev = BKezlev克兹列夫
	f.克兹列夫Kezlev.SetParent(f)
	
	f.佩列科普Perekop = BPerekop佩列科普
	f.佩列科普Perekop.SetParent(f)
	
	f.卡拉苏夫巴扎尔Qarasuvbazar = BQarasuvbazar卡拉苏夫巴扎尔
	f.卡拉苏夫巴扎尔Qarasuvbazar.SetParent(f)
	
	f.库尔曼Qurman = BQurman库尔曼
	f.库尔曼Qurman.SetParent(f)
	
	f.萨克Saq = BSaq萨克
	f.萨克Saq.SetParent(f)
	
}
