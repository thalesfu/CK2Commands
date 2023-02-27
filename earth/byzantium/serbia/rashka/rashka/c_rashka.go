package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RashkaCounty interface {
    feud.County
    BDecani德查尼() 	feud.Barony
    BDioclea迪奥克勒亚() 	feud.Barony
    BDjakovica贾科维察() 	feud.Barony
    BPolog波洛格() 	feud.Barony
    BPrizren普里兹伦() 	feud.Barony
    BSvetispas斯维帖斯帕斯() 	feud.Barony
    BTrepca特雷普查() 	feud.Barony
    BZvecan兹韦钱() 	feud.Barony
}

type 拉什卡RashkaCounty struct {
	feud.BaseCounty
	德查尼Decani 	feud.Barony
	迪奥克勒亚Dioclea 	feud.Barony
	贾科维察Djakovica 	feud.Barony
	波洛格Polog 	feud.Barony
	普里兹伦Prizren 	feud.Barony
	斯维帖斯帕斯Svetispas 	feud.Barony
	特雷普查Trepca 	feud.Barony
	兹韦钱Zvecan 	feud.Barony
}

func (c *拉什卡RashkaCounty) BDecani德查尼() feud.Barony {
	return c.德查尼Decani
}
    
func (c *拉什卡RashkaCounty) BDioclea迪奥克勒亚() feud.Barony {
	return c.迪奥克勒亚Dioclea
}
    
func (c *拉什卡RashkaCounty) BDjakovica贾科维察() feud.Barony {
	return c.贾科维察Djakovica
}
    
func (c *拉什卡RashkaCounty) BPolog波洛格() feud.Barony {
	return c.波洛格Polog
}
    
func (c *拉什卡RashkaCounty) BPrizren普里兹伦() feud.Barony {
	return c.普里兹伦Prizren
}
    
func (c *拉什卡RashkaCounty) BSvetispas斯维帖斯帕斯() feud.Barony {
	return c.斯维帖斯帕斯Svetispas
}
    
func (c *拉什卡RashkaCounty) BTrepca特雷普查() feud.Barony {
	return c.特雷普查Trepca
}
    
func (c *拉什卡RashkaCounty) BZvecan兹韦钱() feud.Barony {
	return c.兹韦钱Zvecan
}
    
var CRashka拉什卡 RashkaCounty = &拉什卡RashkaCounty{}

func init() {
	f := CRashka拉什卡.(*拉什卡RashkaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "502",
		Title:     "rashka",
		TitleName: "拉什卡",
		TitleCode: "c_rashka",
		Baronies:  map[string]feud.Barony{},
	}

	f.德查尼Decani = BDecani德查尼
	f.德查尼Decani.SetParent(f)
	
	f.迪奥克勒亚Dioclea = BDioclea迪奥克勒亚
	f.迪奥克勒亚Dioclea.SetParent(f)
	
	f.贾科维察Djakovica = BDjakovica贾科维察
	f.贾科维察Djakovica.SetParent(f)
	
	f.波洛格Polog = BPolog波洛格
	f.波洛格Polog.SetParent(f)
	
	f.普里兹伦Prizren = BPrizren普里兹伦
	f.普里兹伦Prizren.SetParent(f)
	
	f.斯维帖斯帕斯Svetispas = BSvetispas斯维帖斯帕斯
	f.斯维帖斯帕斯Svetispas.SetParent(f)
	
	f.特雷普查Trepca = BTrepca特雷普查
	f.特雷普查Trepca.SetParent(f)
	
	f.兹韦钱Zvecan = BZvecan兹韦钱
	f.兹韦钱Zvecan.SetParent(f)
	
}
