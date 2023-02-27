package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YegorlykCounty interface {
    feud.County
    BBeshpagir别什帕吉尔() 	feud.Barony
    BErkenshakhar埃尔肯_沙哈尔() 	feud.Barony
    BKiankiz基安基兹() 	feud.Barony
    BKuguty库古特() 	feud.Barony
    BMadjar马佳尔() 	feud.Barony
    BPiatigorie皮亚蒂戈里耶() 	feud.Barony
    BStauropolis斯塔夫罗波利斯() 	feud.Barony
    BYamki亚姆基() 	feud.Barony
}

type 叶戈尔雷克YegorlykCounty struct {
	feud.BaseCounty
	别什帕吉尔Beshpagir 	feud.Barony
	埃尔肯_沙哈尔Erkenshakhar 	feud.Barony
	基安基兹Kiankiz 	feud.Barony
	库古特Kuguty 	feud.Barony
	马佳尔Madjar 	feud.Barony
	皮亚蒂戈里耶Piatigorie 	feud.Barony
	斯塔夫罗波利斯Stauropolis 	feud.Barony
	亚姆基Yamki 	feud.Barony
}

func (c *叶戈尔雷克YegorlykCounty) BBeshpagir别什帕吉尔() feud.Barony {
	return c.别什帕吉尔Beshpagir
}
    
func (c *叶戈尔雷克YegorlykCounty) BErkenshakhar埃尔肯_沙哈尔() feud.Barony {
	return c.埃尔肯_沙哈尔Erkenshakhar
}
    
func (c *叶戈尔雷克YegorlykCounty) BKiankiz基安基兹() feud.Barony {
	return c.基安基兹Kiankiz
}
    
func (c *叶戈尔雷克YegorlykCounty) BKuguty库古特() feud.Barony {
	return c.库古特Kuguty
}
    
func (c *叶戈尔雷克YegorlykCounty) BMadjar马佳尔() feud.Barony {
	return c.马佳尔Madjar
}
    
func (c *叶戈尔雷克YegorlykCounty) BPiatigorie皮亚蒂戈里耶() feud.Barony {
	return c.皮亚蒂戈里耶Piatigorie
}
    
func (c *叶戈尔雷克YegorlykCounty) BStauropolis斯塔夫罗波利斯() feud.Barony {
	return c.斯塔夫罗波利斯Stauropolis
}
    
func (c *叶戈尔雷克YegorlykCounty) BYamki亚姆基() feud.Barony {
	return c.亚姆基Yamki
}
    
var CYegorlyk叶戈尔雷克 YegorlykCounty = &叶戈尔雷克YegorlykCounty{}

func init() {
	f := CYegorlyk叶戈尔雷克.(*叶戈尔雷克YegorlykCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "606",
		Title:     "yegorlyk",
		TitleName: "叶戈尔雷克",
		TitleCode: "c_yegorlyk",
		Baronies:  map[string]feud.Barony{},
	}

	f.别什帕吉尔Beshpagir = BBeshpagir别什帕吉尔
	f.别什帕吉尔Beshpagir.SetParent(f)
	
	f.埃尔肯_沙哈尔Erkenshakhar = BErkenshakhar埃尔肯_沙哈尔
	f.埃尔肯_沙哈尔Erkenshakhar.SetParent(f)
	
	f.基安基兹Kiankiz = BKiankiz基安基兹
	f.基安基兹Kiankiz.SetParent(f)
	
	f.库古特Kuguty = BKuguty库古特
	f.库古特Kuguty.SetParent(f)
	
	f.马佳尔Madjar = BMadjar马佳尔
	f.马佳尔Madjar.SetParent(f)
	
	f.皮亚蒂戈里耶Piatigorie = BPiatigorie皮亚蒂戈里耶
	f.皮亚蒂戈里耶Piatigorie.SetParent(f)
	
	f.斯塔夫罗波利斯Stauropolis = BStauropolis斯塔夫罗波利斯
	f.斯塔夫罗波利斯Stauropolis.SetParent(f)
	
	f.亚姆基Yamki = BYamki亚姆基
	f.亚姆基Yamki.SetParent(f)
	
}
