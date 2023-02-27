package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AracenaCounty interface {
    feud.County
    BAlajar阿拉哈尔() 	feud.Barony
    BAlmonasterlareal阿尔莫纳斯特拉雷亚尔() 	feud.Barony
    BAracena阿拉塞纳() 	feud.Barony
    BCalanas卡拉尼亚斯() 	feud.Barony
    BCortegana科尔特加纳() 	feud.Barony
    BFacanias法卡尼亚斯() 	feud.Barony
    BGalaroza加拉罗萨() 	feud.Barony
    BItalica伊塔利卡() 	feud.Barony
}

type 阿拉塞纳AracenaCounty struct {
	feud.BaseCounty
	阿拉哈尔Alajar 	feud.Barony
	阿尔莫纳斯特拉雷亚尔Almonasterlareal 	feud.Barony
	阿拉塞纳Aracena 	feud.Barony
	卡拉尼亚斯Calanas 	feud.Barony
	科尔特加纳Cortegana 	feud.Barony
	法卡尼亚斯Facanias 	feud.Barony
	加拉罗萨Galaroza 	feud.Barony
	伊塔利卡Italica 	feud.Barony
}

func (c *阿拉塞纳AracenaCounty) BAlajar阿拉哈尔() feud.Barony {
	return c.阿拉哈尔Alajar
}
    
func (c *阿拉塞纳AracenaCounty) BAlmonasterlareal阿尔莫纳斯特拉雷亚尔() feud.Barony {
	return c.阿尔莫纳斯特拉雷亚尔Almonasterlareal
}
    
func (c *阿拉塞纳AracenaCounty) BAracena阿拉塞纳() feud.Barony {
	return c.阿拉塞纳Aracena
}
    
func (c *阿拉塞纳AracenaCounty) BCalanas卡拉尼亚斯() feud.Barony {
	return c.卡拉尼亚斯Calanas
}
    
func (c *阿拉塞纳AracenaCounty) BCortegana科尔特加纳() feud.Barony {
	return c.科尔特加纳Cortegana
}
    
func (c *阿拉塞纳AracenaCounty) BFacanias法卡尼亚斯() feud.Barony {
	return c.法卡尼亚斯Facanias
}
    
func (c *阿拉塞纳AracenaCounty) BGalaroza加拉罗萨() feud.Barony {
	return c.加拉罗萨Galaroza
}
    
func (c *阿拉塞纳AracenaCounty) BItalica伊塔利卡() feud.Barony {
	return c.伊塔利卡Italica
}
    
var CAracena阿拉塞纳 AracenaCounty = &阿拉塞纳AracenaCounty{}

func init() {
	f := CAracena阿拉塞纳.(*阿拉塞纳AracenaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "183",
		Title:     "aracena",
		TitleName: "阿拉塞纳",
		TitleCode: "c_aracena",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉哈尔Alajar = BAlajar阿拉哈尔
	f.阿拉哈尔Alajar.SetParent(f)
	
	f.阿尔莫纳斯特拉雷亚尔Almonasterlareal = BAlmonasterlareal阿尔莫纳斯特拉雷亚尔
	f.阿尔莫纳斯特拉雷亚尔Almonasterlareal.SetParent(f)
	
	f.阿拉塞纳Aracena = BAracena阿拉塞纳
	f.阿拉塞纳Aracena.SetParent(f)
	
	f.卡拉尼亚斯Calanas = BCalanas卡拉尼亚斯
	f.卡拉尼亚斯Calanas.SetParent(f)
	
	f.科尔特加纳Cortegana = BCortegana科尔特加纳
	f.科尔特加纳Cortegana.SetParent(f)
	
	f.法卡尼亚斯Facanias = BFacanias法卡尼亚斯
	f.法卡尼亚斯Facanias.SetParent(f)
	
	f.加拉罗萨Galaroza = BGalaroza加拉罗萨
	f.加拉罗萨Galaroza.SetParent(f)
	
	f.伊塔利卡Italica = BItalica伊塔利卡
	f.伊塔利卡Italica.SetParent(f)
	
}
