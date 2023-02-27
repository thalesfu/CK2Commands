package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MarienburgCounty interface {
    feud.County
    BBalga巴尔加() 	feud.Barony
    BBartenstein巴滕施泰因() 	feud.Barony
    BBraunsberg布劳恩斯贝格() 	feud.Barony
    BChristburg基督堡() 	feud.Barony
    BElbing埃尔宾() 	feud.Barony
    BHeilsberg海尔斯堡() 	feud.Barony
    BMarienburg马林堡() 	feud.Barony
    BMarienwerder马林韦尔德() 	feud.Barony
}

type 马林堡MarienburgCounty struct {
	feud.BaseCounty
	巴尔加Balga 	feud.Barony
	巴滕施泰因Bartenstein 	feud.Barony
	布劳恩斯贝格Braunsberg 	feud.Barony
	基督堡Christburg 	feud.Barony
	埃尔宾Elbing 	feud.Barony
	海尔斯堡Heilsberg 	feud.Barony
	马林堡Marienburg 	feud.Barony
	马林韦尔德Marienwerder 	feud.Barony
}

func (c *马林堡MarienburgCounty) BBalga巴尔加() feud.Barony {
	return c.巴尔加Balga
}
    
func (c *马林堡MarienburgCounty) BBartenstein巴滕施泰因() feud.Barony {
	return c.巴滕施泰因Bartenstein
}
    
func (c *马林堡MarienburgCounty) BBraunsberg布劳恩斯贝格() feud.Barony {
	return c.布劳恩斯贝格Braunsberg
}
    
func (c *马林堡MarienburgCounty) BChristburg基督堡() feud.Barony {
	return c.基督堡Christburg
}
    
func (c *马林堡MarienburgCounty) BElbing埃尔宾() feud.Barony {
	return c.埃尔宾Elbing
}
    
func (c *马林堡MarienburgCounty) BHeilsberg海尔斯堡() feud.Barony {
	return c.海尔斯堡Heilsberg
}
    
func (c *马林堡MarienburgCounty) BMarienburg马林堡() feud.Barony {
	return c.马林堡Marienburg
}
    
func (c *马林堡MarienburgCounty) BMarienwerder马林韦尔德() feud.Barony {
	return c.马林韦尔德Marienwerder
}
    
var CMarienburg马林堡 MarienburgCounty = &马林堡MarienburgCounty{}

func init() {
	f := CMarienburg马林堡.(*马林堡MarienburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "370",
		Title:     "marienburg",
		TitleName: "马林堡",
		TitleCode: "c_marienburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔加Balga = BBalga巴尔加
	f.巴尔加Balga.SetParent(f)
	
	f.巴滕施泰因Bartenstein = BBartenstein巴滕施泰因
	f.巴滕施泰因Bartenstein.SetParent(f)
	
	f.布劳恩斯贝格Braunsberg = BBraunsberg布劳恩斯贝格
	f.布劳恩斯贝格Braunsberg.SetParent(f)
	
	f.基督堡Christburg = BChristburg基督堡
	f.基督堡Christburg.SetParent(f)
	
	f.埃尔宾Elbing = BElbing埃尔宾
	f.埃尔宾Elbing.SetParent(f)
	
	f.海尔斯堡Heilsberg = BHeilsberg海尔斯堡
	f.海尔斯堡Heilsberg.SetParent(f)
	
	f.马林堡Marienburg = BMarienburg马林堡
	f.马林堡Marienburg.SetParent(f)
	
	f.马林韦尔德Marienwerder = BMarienwerder马林韦尔德
	f.马林韦尔德Marienwerder.SetParent(f)
	
}
