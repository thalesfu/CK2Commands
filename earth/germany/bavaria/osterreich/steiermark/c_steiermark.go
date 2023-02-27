package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SteiermarkCounty interface {
    feud.County
    BBehamberg贝汉贝格() 	feud.Barony
    BGarstina加斯蒂纳() 	feud.Barony
    BGrobraming大拉明() 	feud.Barony
    BLosenstein洛森施泰因() 	feud.Barony
    BRubinicha鲁宾尼察() 	feud.Barony
    BSteyr施泰尔() 	feud.Barony
    BTernberg滕贝格() 	feud.Barony
}

type 施蒂里亚SteiermarkCounty struct {
	feud.BaseCounty
	贝汉贝格Behamberg 	feud.Barony
	加斯蒂纳Garstina 	feud.Barony
	大拉明Grobraming 	feud.Barony
	洛森施泰因Losenstein 	feud.Barony
	鲁宾尼察Rubinicha 	feud.Barony
	施泰尔Steyr 	feud.Barony
	滕贝格Ternberg 	feud.Barony
}

func (c *施蒂里亚SteiermarkCounty) BBehamberg贝汉贝格() feud.Barony {
	return c.贝汉贝格Behamberg
}
    
func (c *施蒂里亚SteiermarkCounty) BGarstina加斯蒂纳() feud.Barony {
	return c.加斯蒂纳Garstina
}
    
func (c *施蒂里亚SteiermarkCounty) BGrobraming大拉明() feud.Barony {
	return c.大拉明Grobraming
}
    
func (c *施蒂里亚SteiermarkCounty) BLosenstein洛森施泰因() feud.Barony {
	return c.洛森施泰因Losenstein
}
    
func (c *施蒂里亚SteiermarkCounty) BRubinicha鲁宾尼察() feud.Barony {
	return c.鲁宾尼察Rubinicha
}
    
func (c *施蒂里亚SteiermarkCounty) BSteyr施泰尔() feud.Barony {
	return c.施泰尔Steyr
}
    
func (c *施蒂里亚SteiermarkCounty) BTernberg滕贝格() feud.Barony {
	return c.滕贝格Ternberg
}
    
var CSteiermark施蒂里亚 SteiermarkCounty = &施蒂里亚SteiermarkCounty{}

func init() {
	f := CSteiermark施蒂里亚.(*施蒂里亚SteiermarkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "455",
		Title:     "steiermark",
		TitleName: "施蒂里亚",
		TitleCode: "c_steiermark",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝汉贝格Behamberg = BBehamberg贝汉贝格
	f.贝汉贝格Behamberg.SetParent(f)
	
	f.加斯蒂纳Garstina = BGarstina加斯蒂纳
	f.加斯蒂纳Garstina.SetParent(f)
	
	f.大拉明Grobraming = BGrobraming大拉明
	f.大拉明Grobraming.SetParent(f)
	
	f.洛森施泰因Losenstein = BLosenstein洛森施泰因
	f.洛森施泰因Losenstein.SetParent(f)
	
	f.鲁宾尼察Rubinicha = BRubinicha鲁宾尼察
	f.鲁宾尼察Rubinicha.SetParent(f)
	
	f.施泰尔Steyr = BSteyr施泰尔
	f.施泰尔Steyr.SetParent(f)
	
	f.滕贝格Ternberg = BTernberg滕贝格
	f.滕贝格Ternberg.SetParent(f)
	
}
