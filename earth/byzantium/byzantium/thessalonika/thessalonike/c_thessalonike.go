package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThessalonikeCounty interface {
    feud.County
    BCemren切木兰() 	feud.Barony
    BElasson埃拉索纳() 	feud.Barony
    BHlerin赫雷林() 	feud.Barony
    BServia塞尔维亚() 	feud.Barony
    BThesedessa埃泽萨() 	feud.Barony
    BThessaloniki塞萨洛尼基() 	feud.Barony
    BVeria韦里亚() 	feud.Barony
    BVoden沃泽纳() 	feud.Barony
}

type 塞萨洛尼基ThessalonikeCounty struct {
	feud.BaseCounty
	切木兰Cemren 	feud.Barony
	埃拉索纳Elasson 	feud.Barony
	赫雷林Hlerin 	feud.Barony
	塞尔维亚Servia 	feud.Barony
	埃泽萨Thesedessa 	feud.Barony
	塞萨洛尼基Thessaloniki 	feud.Barony
	韦里亚Veria 	feud.Barony
	沃泽纳Voden 	feud.Barony
}

func (c *塞萨洛尼基ThessalonikeCounty) BCemren切木兰() feud.Barony {
	return c.切木兰Cemren
}
    
func (c *塞萨洛尼基ThessalonikeCounty) BElasson埃拉索纳() feud.Barony {
	return c.埃拉索纳Elasson
}
    
func (c *塞萨洛尼基ThessalonikeCounty) BHlerin赫雷林() feud.Barony {
	return c.赫雷林Hlerin
}
    
func (c *塞萨洛尼基ThessalonikeCounty) BServia塞尔维亚() feud.Barony {
	return c.塞尔维亚Servia
}
    
func (c *塞萨洛尼基ThessalonikeCounty) BThesedessa埃泽萨() feud.Barony {
	return c.埃泽萨Thesedessa
}
    
func (c *塞萨洛尼基ThessalonikeCounty) BThessaloniki塞萨洛尼基() feud.Barony {
	return c.塞萨洛尼基Thessaloniki
}
    
func (c *塞萨洛尼基ThessalonikeCounty) BVeria韦里亚() feud.Barony {
	return c.韦里亚Veria
}
    
func (c *塞萨洛尼基ThessalonikeCounty) BVoden沃泽纳() feud.Barony {
	return c.沃泽纳Voden
}
    
var CThessalonike塞萨洛尼基 ThessalonikeCounty = &塞萨洛尼基ThessalonikeCounty{}

func init() {
	f := CThessalonike塞萨洛尼基.(*塞萨洛尼基ThessalonikeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "490",
		Title:     "thessalonike",
		TitleName: "塞萨洛尼基",
		TitleCode: "c_thessalonike",
		Baronies:  map[string]feud.Barony{},
	}

	f.切木兰Cemren = BCemren切木兰
	f.切木兰Cemren.SetParent(f)
	
	f.埃拉索纳Elasson = BElasson埃拉索纳
	f.埃拉索纳Elasson.SetParent(f)
	
	f.赫雷林Hlerin = BHlerin赫雷林
	f.赫雷林Hlerin.SetParent(f)
	
	f.塞尔维亚Servia = BServia塞尔维亚
	f.塞尔维亚Servia.SetParent(f)
	
	f.埃泽萨Thesedessa = BThesedessa埃泽萨
	f.埃泽萨Thesedessa.SetParent(f)
	
	f.塞萨洛尼基Thessaloniki = BThessaloniki塞萨洛尼基
	f.塞萨洛尼基Thessaloniki.SetParent(f)
	
	f.韦里亚Veria = BVeria韦里亚
	f.韦里亚Veria.SetParent(f)
	
	f.沃泽纳Voden = BVoden沃泽纳
	f.沃泽纳Voden.SetParent(f)
	
}
