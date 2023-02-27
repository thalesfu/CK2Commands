package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BashkirsCounty interface {
    feud.County
    BBajmaq巴日马克() 	feud.Barony
    BBelebey别列别伊() 	feud.Barony
    BBeloret别洛列特() 	feud.Barony
    BChishmy奇什梅() 	feud.Barony
    BIsembaj伊森巴日() 	feud.Barony
    BMeleus梅列乌兹() 	feud.Barony
    BSterlitamak斯捷尔利塔马克() 	feud.Barony
    BUfa乌法() 	feud.Barony
}

type 巴什基里亚BashkirsCounty struct {
	feud.BaseCounty
	巴日马克Bajmaq 	feud.Barony
	别列别伊Belebey 	feud.Barony
	别洛列特Beloret 	feud.Barony
	奇什梅Chishmy 	feud.Barony
	伊森巴日Isembaj 	feud.Barony
	梅列乌兹Meleus 	feud.Barony
	斯捷尔利塔马克Sterlitamak 	feud.Barony
	乌法Ufa 	feud.Barony
}

func (c *巴什基里亚BashkirsCounty) BBajmaq巴日马克() feud.Barony {
	return c.巴日马克Bajmaq
}
    
func (c *巴什基里亚BashkirsCounty) BBelebey别列别伊() feud.Barony {
	return c.别列别伊Belebey
}
    
func (c *巴什基里亚BashkirsCounty) BBeloret别洛列特() feud.Barony {
	return c.别洛列特Beloret
}
    
func (c *巴什基里亚BashkirsCounty) BChishmy奇什梅() feud.Barony {
	return c.奇什梅Chishmy
}
    
func (c *巴什基里亚BashkirsCounty) BIsembaj伊森巴日() feud.Barony {
	return c.伊森巴日Isembaj
}
    
func (c *巴什基里亚BashkirsCounty) BMeleus梅列乌兹() feud.Barony {
	return c.梅列乌兹Meleus
}
    
func (c *巴什基里亚BashkirsCounty) BSterlitamak斯捷尔利塔马克() feud.Barony {
	return c.斯捷尔利塔马克Sterlitamak
}
    
func (c *巴什基里亚BashkirsCounty) BUfa乌法() feud.Barony {
	return c.乌法Ufa
}
    
var CBashkirs巴什基里亚 BashkirsCounty = &巴什基里亚BashkirsCounty{}

func init() {
	f := CBashkirs巴什基里亚.(*巴什基里亚BashkirsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "615",
		Title:     "bashkirs",
		TitleName: "巴什基里亚",
		TitleCode: "c_bashkirs",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴日马克Bajmaq = BBajmaq巴日马克
	f.巴日马克Bajmaq.SetParent(f)
	
	f.别列别伊Belebey = BBelebey别列别伊
	f.别列别伊Belebey.SetParent(f)
	
	f.别洛列特Beloret = BBeloret别洛列特
	f.别洛列特Beloret.SetParent(f)
	
	f.奇什梅Chishmy = BChishmy奇什梅
	f.奇什梅Chishmy.SetParent(f)
	
	f.伊森巴日Isembaj = BIsembaj伊森巴日
	f.伊森巴日Isembaj.SetParent(f)
	
	f.梅列乌兹Meleus = BMeleus梅列乌兹
	f.梅列乌兹Meleus.SetParent(f)
	
	f.斯捷尔利塔马克Sterlitamak = BSterlitamak斯捷尔利塔马克
	f.斯捷尔利塔马克Sterlitamak.SetParent(f)
	
	f.乌法Ufa = BUfa乌法
	f.乌法Ufa.SetParent(f)
	
}
