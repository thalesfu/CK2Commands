package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OberbayernCounty interface {
    feud.County
    BAndechs安德希斯() 	feud.Barony
    BDachau达豪() 	feud.Barony
    BDiessen迪森() 	feud.Barony
    BDingolfing丁戈尔芬() 	feud.Barony
    BEbersberg埃伯斯贝格() 	feud.Barony
    BFreising弗赖辛() 	feud.Barony
    BMunchen慕尼黑() 	feud.Barony
    BPfaffenberg普法芬贝格() 	feud.Barony
}

type 弗赖辛OberbayernCounty struct {
	feud.BaseCounty
	安德希斯Andechs 	feud.Barony
	达豪Dachau 	feud.Barony
	迪森Diessen 	feud.Barony
	丁戈尔芬Dingolfing 	feud.Barony
	埃伯斯贝格Ebersberg 	feud.Barony
	弗赖辛Freising 	feud.Barony
	慕尼黑Munchen 	feud.Barony
	普法芬贝格Pfaffenberg 	feud.Barony
}

func (c *弗赖辛OberbayernCounty) BAndechs安德希斯() feud.Barony {
	return c.安德希斯Andechs
}
    
func (c *弗赖辛OberbayernCounty) BDachau达豪() feud.Barony {
	return c.达豪Dachau
}
    
func (c *弗赖辛OberbayernCounty) BDiessen迪森() feud.Barony {
	return c.迪森Diessen
}
    
func (c *弗赖辛OberbayernCounty) BDingolfing丁戈尔芬() feud.Barony {
	return c.丁戈尔芬Dingolfing
}
    
func (c *弗赖辛OberbayernCounty) BEbersberg埃伯斯贝格() feud.Barony {
	return c.埃伯斯贝格Ebersberg
}
    
func (c *弗赖辛OberbayernCounty) BFreising弗赖辛() feud.Barony {
	return c.弗赖辛Freising
}
    
func (c *弗赖辛OberbayernCounty) BMunchen慕尼黑() feud.Barony {
	return c.慕尼黑Munchen
}
    
func (c *弗赖辛OberbayernCounty) BPfaffenberg普法芬贝格() feud.Barony {
	return c.普法芬贝格Pfaffenberg
}
    
var COberbayern弗赖辛 OberbayernCounty = &弗赖辛OberbayernCounty{}

func init() {
	f := COberbayern弗赖辛.(*弗赖辛OberbayernCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "360",
		Title:     "oberbayern",
		TitleName: "弗赖辛",
		TitleCode: "c_oberbayern",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德希斯Andechs = BAndechs安德希斯
	f.安德希斯Andechs.SetParent(f)
	
	f.达豪Dachau = BDachau达豪
	f.达豪Dachau.SetParent(f)
	
	f.迪森Diessen = BDiessen迪森
	f.迪森Diessen.SetParent(f)
	
	f.丁戈尔芬Dingolfing = BDingolfing丁戈尔芬
	f.丁戈尔芬Dingolfing.SetParent(f)
	
	f.埃伯斯贝格Ebersberg = BEbersberg埃伯斯贝格
	f.埃伯斯贝格Ebersberg.SetParent(f)
	
	f.弗赖辛Freising = BFreising弗赖辛
	f.弗赖辛Freising.SetParent(f)
	
	f.慕尼黑Munchen = BMunchen慕尼黑
	f.慕尼黑Munchen.SetParent(f)
	
	f.普法芬贝格Pfaffenberg = BPfaffenberg普法芬贝格
	f.普法芬贝格Pfaffenberg.SetParent(f)
	
}
