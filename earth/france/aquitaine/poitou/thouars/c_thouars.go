package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThouarsCounty interface {
	feud.County
	BBressuire布雷叙尔() feud.Barony
	BChatelaillon沙特拉永() feud.Barony
	BFontenay丰特奈() feud.Barony
	BLarochelle拉罗谢尔() feud.Barony
	BLucon吕松() feud.Barony
	BMauleon莫莱翁() feud.Barony
	BOlonne奥洛讷() feud.Barony
	BThouars图阿尔() feud.Barony
}

type 图阿尔ThouarsCounty struct {
	feud.BaseCounty
	布雷叙尔Bressuire    feud.Barony
	沙特拉永Chatelaillon feud.Barony
	丰特奈Fontenay      feud.Barony
	拉罗谢尔Larochelle   feud.Barony
	吕松Lucon          feud.Barony
	莫莱翁Mauleon       feud.Barony
	奥洛讷Olonne        feud.Barony
	图阿尔Thouars       feud.Barony
}

func (c *图阿尔ThouarsCounty) BBressuire布雷叙尔() feud.Barony {
	return c.布雷叙尔Bressuire
}

func (c *图阿尔ThouarsCounty) BChatelaillon沙特拉永() feud.Barony {
	return c.沙特拉永Chatelaillon
}

func (c *图阿尔ThouarsCounty) BFontenay丰特奈() feud.Barony {
	return c.丰特奈Fontenay
}

func (c *图阿尔ThouarsCounty) BLarochelle拉罗谢尔() feud.Barony {
	return c.拉罗谢尔Larochelle
}

func (c *图阿尔ThouarsCounty) BLucon吕松() feud.Barony {
	return c.吕松Lucon
}

func (c *图阿尔ThouarsCounty) BMauleon莫莱翁() feud.Barony {
	return c.莫莱翁Mauleon
}

func (c *图阿尔ThouarsCounty) BOlonne奥洛讷() feud.Barony {
	return c.奥洛讷Olonne
}

func (c *图阿尔ThouarsCounty) BThouars图阿尔() feud.Barony {
	return c.图阿尔Thouars
}

var CThouars图阿尔 ThouarsCounty = &图阿尔ThouarsCounty{}

func init() {
	f := CThouars图阿尔.(*图阿尔ThouarsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "142",
		Title:     "thouars",
		TitleName: "图阿尔",
		TitleCode: "c_thouars",
		Baronies:  map[string]feud.Barony{},
	}

	f.布雷叙尔Bressuire = BBressuire布雷叙尔
	f.布雷叙尔Bressuire.SetParent(f)

	f.沙特拉永Chatelaillon = BChatelaillon沙特拉永
	f.沙特拉永Chatelaillon.SetParent(f)

	f.丰特奈Fontenay = BFontenay丰特奈
	f.丰特奈Fontenay.SetParent(f)

	f.拉罗谢尔Larochelle = BLarochelle拉罗谢尔
	f.拉罗谢尔Larochelle.SetParent(f)

	f.吕松Lucon = BLucon吕松
	f.吕松Lucon.SetParent(f)

	f.莫莱翁Mauleon = BMauleon莫莱翁
	f.莫莱翁Mauleon.SetParent(f)

	f.奥洛讷Olonne = BOlonne奥洛讷
	f.奥洛讷Olonne.SetParent(f)

	f.图阿尔Thouars = BThouars图阿尔
	f.图阿尔Thouars.SetParent(f)

}
