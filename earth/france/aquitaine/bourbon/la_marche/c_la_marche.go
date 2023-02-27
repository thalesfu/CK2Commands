package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type La_marcheCounty interface {
    feud.County
    BAubusson欧比松() 	feud.Barony
    BBellac贝拉克() 	feud.Barony
    BBourganeuf布尔加讷夫() 	feud.Barony
    BBoussac布萨克() 	feud.Barony
    BCrozant克罗藏() 	feud.Barony
    BGueret盖雷() 	feud.Barony
    BJouillat茹雅() 	feud.Barony
    BLasouterraine拉苏泰赖讷() 	feud.Barony
}

type 拉马什La_marcheCounty struct {
	feud.BaseCounty
	欧比松Aubusson 	feud.Barony
	贝拉克Bellac 	feud.Barony
	布尔加讷夫Bourganeuf 	feud.Barony
	布萨克Boussac 	feud.Barony
	克罗藏Crozant 	feud.Barony
	盖雷Gueret 	feud.Barony
	茹雅Jouillat 	feud.Barony
	拉苏泰赖讷Lasouterraine 	feud.Barony
}

func (c *拉马什La_marcheCounty) BAubusson欧比松() feud.Barony {
	return c.欧比松Aubusson
}
    
func (c *拉马什La_marcheCounty) BBellac贝拉克() feud.Barony {
	return c.贝拉克Bellac
}
    
func (c *拉马什La_marcheCounty) BBourganeuf布尔加讷夫() feud.Barony {
	return c.布尔加讷夫Bourganeuf
}
    
func (c *拉马什La_marcheCounty) BBoussac布萨克() feud.Barony {
	return c.布萨克Boussac
}
    
func (c *拉马什La_marcheCounty) BCrozant克罗藏() feud.Barony {
	return c.克罗藏Crozant
}
    
func (c *拉马什La_marcheCounty) BGueret盖雷() feud.Barony {
	return c.盖雷Gueret
}
    
func (c *拉马什La_marcheCounty) BJouillat茹雅() feud.Barony {
	return c.茹雅Jouillat
}
    
func (c *拉马什La_marcheCounty) BLasouterraine拉苏泰赖讷() feud.Barony {
	return c.拉苏泰赖讷Lasouterraine
}
    
var CLa_marche拉马什 La_marcheCounty = &拉马什La_marcheCounty{}

func init() {
	f := CLa_marche拉马什.(*拉马什La_marcheCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "145",
		Title:     "la_marche",
		TitleName: "拉马什",
		TitleCode: "c_la_marche",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧比松Aubusson = BAubusson欧比松
	f.欧比松Aubusson.SetParent(f)
	
	f.贝拉克Bellac = BBellac贝拉克
	f.贝拉克Bellac.SetParent(f)
	
	f.布尔加讷夫Bourganeuf = BBourganeuf布尔加讷夫
	f.布尔加讷夫Bourganeuf.SetParent(f)
	
	f.布萨克Boussac = BBoussac布萨克
	f.布萨克Boussac.SetParent(f)
	
	f.克罗藏Crozant = BCrozant克罗藏
	f.克罗藏Crozant.SetParent(f)
	
	f.盖雷Gueret = BGueret盖雷
	f.盖雷Gueret.SetParent(f)
	
	f.茹雅Jouillat = BJouillat茹雅
	f.茹雅Jouillat.SetParent(f)
	
	f.拉苏泰赖讷Lasouterraine = BLasouterraine拉苏泰赖讷
	f.拉苏泰赖讷Lasouterraine.SetParent(f)
	
}
