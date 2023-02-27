package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RigaCounty interface {
    feud.County
    BIkskile伊克什基莱() 	feud.Barony
    BLedurga莱杜尔加() 	feud.Barony
    BRiga里加() 	feud.Barony
    BRopazi罗帕日() 	feud.Barony
    BSalaspils萨拉斯皮尔斯() 	feud.Barony
    BSigulda锡古尔达() 	feud.Barony
    BTuraida图赖达() 	feud.Barony
}

type 里加RigaCounty struct {
	feud.BaseCounty
	伊克什基莱Ikskile 	feud.Barony
	莱杜尔加Ledurga 	feud.Barony
	里加Riga 	feud.Barony
	罗帕日Ropazi 	feud.Barony
	萨拉斯皮尔斯Salaspils 	feud.Barony
	锡古尔达Sigulda 	feud.Barony
	图赖达Turaida 	feud.Barony
}

func (c *里加RigaCounty) BIkskile伊克什基莱() feud.Barony {
	return c.伊克什基莱Ikskile
}
    
func (c *里加RigaCounty) BLedurga莱杜尔加() feud.Barony {
	return c.莱杜尔加Ledurga
}
    
func (c *里加RigaCounty) BRiga里加() feud.Barony {
	return c.里加Riga
}
    
func (c *里加RigaCounty) BRopazi罗帕日() feud.Barony {
	return c.罗帕日Ropazi
}
    
func (c *里加RigaCounty) BSalaspils萨拉斯皮尔斯() feud.Barony {
	return c.萨拉斯皮尔斯Salaspils
}
    
func (c *里加RigaCounty) BSigulda锡古尔达() feud.Barony {
	return c.锡古尔达Sigulda
}
    
func (c *里加RigaCounty) BTuraida图赖达() feud.Barony {
	return c.图赖达Turaida
}
    
var CRiga里加 RigaCounty = &里加RigaCounty{}

func init() {
	f := CRiga里加.(*里加RigaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1596",
		Title:     "riga",
		TitleName: "里加",
		TitleCode: "c_riga",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊克什基莱Ikskile = BIkskile伊克什基莱
	f.伊克什基莱Ikskile.SetParent(f)
	
	f.莱杜尔加Ledurga = BLedurga莱杜尔加
	f.莱杜尔加Ledurga.SetParent(f)
	
	f.里加Riga = BRiga里加
	f.里加Riga.SetParent(f)
	
	f.罗帕日Ropazi = BRopazi罗帕日
	f.罗帕日Ropazi.SetParent(f)
	
	f.萨拉斯皮尔斯Salaspils = BSalaspils萨拉斯皮尔斯
	f.萨拉斯皮尔斯Salaspils.SetParent(f)
	
	f.锡古尔达Sigulda = BSigulda锡古尔达
	f.锡古尔达Sigulda.SetParent(f)
	
	f.图赖达Turaida = BTuraida图赖达
	f.图赖达Turaida.SetParent(f)
	
}
