package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CornouailleCounty interface {
    feud.County
    BChateaulin沙托兰() 	feud.Barony
    BConcarneau孔卡诺() 	feud.Barony
    BDouarnenez杜瓦讷内() 	feud.Barony
    BLandevennec朗代韦内克() 	feud.Barony
    BLezergue莱泽尔盖() 	feud.Barony
    BQuimper坎佩尔() 	feud.Barony
    BQuimperle坎佩莱() 	feud.Barony
}

type 科努瓦耶CornouailleCounty struct {
	feud.BaseCounty
	沙托兰Chateaulin 	feud.Barony
	孔卡诺Concarneau 	feud.Barony
	杜瓦讷内Douarnenez 	feud.Barony
	朗代韦内克Landevennec 	feud.Barony
	莱泽尔盖Lezergue 	feud.Barony
	坎佩尔Quimper 	feud.Barony
	坎佩莱Quimperle 	feud.Barony
}

func (c *科努瓦耶CornouailleCounty) BChateaulin沙托兰() feud.Barony {
	return c.沙托兰Chateaulin
}
    
func (c *科努瓦耶CornouailleCounty) BConcarneau孔卡诺() feud.Barony {
	return c.孔卡诺Concarneau
}
    
func (c *科努瓦耶CornouailleCounty) BDouarnenez杜瓦讷内() feud.Barony {
	return c.杜瓦讷内Douarnenez
}
    
func (c *科努瓦耶CornouailleCounty) BLandevennec朗代韦内克() feud.Barony {
	return c.朗代韦内克Landevennec
}
    
func (c *科努瓦耶CornouailleCounty) BLezergue莱泽尔盖() feud.Barony {
	return c.莱泽尔盖Lezergue
}
    
func (c *科努瓦耶CornouailleCounty) BQuimper坎佩尔() feud.Barony {
	return c.坎佩尔Quimper
}
    
func (c *科努瓦耶CornouailleCounty) BQuimperle坎佩莱() feud.Barony {
	return c.坎佩莱Quimperle
}
    
var CCornouaille科努瓦耶 CornouailleCounty = &科努瓦耶CornouailleCounty{}

func init() {
	f := CCornouaille科努瓦耶.(*科努瓦耶CornouailleCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "104",
		Title:     "cornouaille",
		TitleName: "科努瓦耶",
		TitleCode: "c_cornouaille",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙托兰Chateaulin = BChateaulin沙托兰
	f.沙托兰Chateaulin.SetParent(f)
	
	f.孔卡诺Concarneau = BConcarneau孔卡诺
	f.孔卡诺Concarneau.SetParent(f)
	
	f.杜瓦讷内Douarnenez = BDouarnenez杜瓦讷内
	f.杜瓦讷内Douarnenez.SetParent(f)
	
	f.朗代韦内克Landevennec = BLandevennec朗代韦内克
	f.朗代韦内克Landevennec.SetParent(f)
	
	f.莱泽尔盖Lezergue = BLezergue莱泽尔盖
	f.莱泽尔盖Lezergue.SetParent(f)
	
	f.坎佩尔Quimper = BQuimper坎佩尔
	f.坎佩尔Quimper.SetParent(f)
	
	f.坎佩莱Quimperle = BQuimperle坎佩莱
	f.坎佩莱Quimperle.SetParent(f)
	
}
