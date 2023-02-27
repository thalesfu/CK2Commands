package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TanaCounty interface {
    feud.County
    BBataysk巴泰斯克() 	feud.Barony
    BCherkassk切尔卡斯克() 	feud.Barony
    BGundorovka贡多罗夫卡() 	feud.Barony
    BMonastyrsky莫纳斯特尔斯基() 	feud.Barony
    BRostovnadonu顿河畔罗斯托夫() 	feud.Barony
    BSulin苏林() 	feud.Barony
    BTana塔纳() 	feud.Barony
    BUstaksayaskaya乌斯季_阿克赛斯卡亚() 	feud.Barony
}

type 塔纳TanaCounty struct {
	feud.BaseCounty
	巴泰斯克Bataysk 	feud.Barony
	切尔卡斯克Cherkassk 	feud.Barony
	贡多罗夫卡Gundorovka 	feud.Barony
	莫纳斯特尔斯基Monastyrsky 	feud.Barony
	顿河畔罗斯托夫Rostovnadonu 	feud.Barony
	苏林Sulin 	feud.Barony
	塔纳Tana 	feud.Barony
	乌斯季_阿克赛斯卡亚Ustaksayaskaya 	feud.Barony
}

func (c *塔纳TanaCounty) BBataysk巴泰斯克() feud.Barony {
	return c.巴泰斯克Bataysk
}
    
func (c *塔纳TanaCounty) BCherkassk切尔卡斯克() feud.Barony {
	return c.切尔卡斯克Cherkassk
}
    
func (c *塔纳TanaCounty) BGundorovka贡多罗夫卡() feud.Barony {
	return c.贡多罗夫卡Gundorovka
}
    
func (c *塔纳TanaCounty) BMonastyrsky莫纳斯特尔斯基() feud.Barony {
	return c.莫纳斯特尔斯基Monastyrsky
}
    
func (c *塔纳TanaCounty) BRostovnadonu顿河畔罗斯托夫() feud.Barony {
	return c.顿河畔罗斯托夫Rostovnadonu
}
    
func (c *塔纳TanaCounty) BSulin苏林() feud.Barony {
	return c.苏林Sulin
}
    
func (c *塔纳TanaCounty) BTana塔纳() feud.Barony {
	return c.塔纳Tana
}
    
func (c *塔纳TanaCounty) BUstaksayaskaya乌斯季_阿克赛斯卡亚() feud.Barony {
	return c.乌斯季_阿克赛斯卡亚Ustaksayaskaya
}
    
var CTana塔纳 TanaCounty = &塔纳TanaCounty{}

func init() {
	f := CTana塔纳.(*塔纳TanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "596",
		Title:     "tana",
		TitleName: "塔纳",
		TitleCode: "c_tana",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴泰斯克Bataysk = BBataysk巴泰斯克
	f.巴泰斯克Bataysk.SetParent(f)
	
	f.切尔卡斯克Cherkassk = BCherkassk切尔卡斯克
	f.切尔卡斯克Cherkassk.SetParent(f)
	
	f.贡多罗夫卡Gundorovka = BGundorovka贡多罗夫卡
	f.贡多罗夫卡Gundorovka.SetParent(f)
	
	f.莫纳斯特尔斯基Monastyrsky = BMonastyrsky莫纳斯特尔斯基
	f.莫纳斯特尔斯基Monastyrsky.SetParent(f)
	
	f.顿河畔罗斯托夫Rostovnadonu = BRostovnadonu顿河畔罗斯托夫
	f.顿河畔罗斯托夫Rostovnadonu.SetParent(f)
	
	f.苏林Sulin = BSulin苏林
	f.苏林Sulin.SetParent(f)
	
	f.塔纳Tana = BTana塔纳
	f.塔纳Tana.SetParent(f)
	
	f.乌斯季_阿克赛斯卡亚Ustaksayaskaya = BUstaksayaskaya乌斯季_阿克赛斯卡亚
	f.乌斯季_阿克赛斯卡亚Ustaksayaskaya.SetParent(f)
	
}
