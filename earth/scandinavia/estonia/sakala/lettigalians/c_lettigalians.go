package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LettigaliansCounty interface {
    feud.County
    BGaujiena高耶纳() 	feud.Barony
    BKokenois库凯诺伊斯() 	feud.Barony
    BLemisele莱米塞莱() 	feud.Barony
    BSeswegen塞斯韦根() 	feud.Barony
    BUxkull伊克什基莱() 	feud.Barony
    BWenden文登() 	feud.Barony
    BWolmer沃尔梅尔() 	feud.Barony
}

type 拉特加利亚LettigaliansCounty struct {
	feud.BaseCounty
	高耶纳Gaujiena 	feud.Barony
	库凯诺伊斯Kokenois 	feud.Barony
	莱米塞莱Lemisele 	feud.Barony
	塞斯韦根Seswegen 	feud.Barony
	伊克什基莱Uxkull 	feud.Barony
	文登Wenden 	feud.Barony
	沃尔梅尔Wolmer 	feud.Barony
}

func (c *拉特加利亚LettigaliansCounty) BGaujiena高耶纳() feud.Barony {
	return c.高耶纳Gaujiena
}
    
func (c *拉特加利亚LettigaliansCounty) BKokenois库凯诺伊斯() feud.Barony {
	return c.库凯诺伊斯Kokenois
}
    
func (c *拉特加利亚LettigaliansCounty) BLemisele莱米塞莱() feud.Barony {
	return c.莱米塞莱Lemisele
}
    
func (c *拉特加利亚LettigaliansCounty) BSeswegen塞斯韦根() feud.Barony {
	return c.塞斯韦根Seswegen
}
    
func (c *拉特加利亚LettigaliansCounty) BUxkull伊克什基莱() feud.Barony {
	return c.伊克什基莱Uxkull
}
    
func (c *拉特加利亚LettigaliansCounty) BWenden文登() feud.Barony {
	return c.文登Wenden
}
    
func (c *拉特加利亚LettigaliansCounty) BWolmer沃尔梅尔() feud.Barony {
	return c.沃尔梅尔Wolmer
}
    
var CLettigalians拉特加利亚 LettigaliansCounty = &拉特加利亚LettigaliansCounty{}

func init() {
	f := CLettigalians拉特加利亚.(*拉特加利亚LettigaliansCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "375",
		Title:     "lettigalians",
		TitleName: "拉特加利亚",
		TitleCode: "c_lettigalians",
		Baronies:  map[string]feud.Barony{},
	}

	f.高耶纳Gaujiena = BGaujiena高耶纳
	f.高耶纳Gaujiena.SetParent(f)
	
	f.库凯诺伊斯Kokenois = BKokenois库凯诺伊斯
	f.库凯诺伊斯Kokenois.SetParent(f)
	
	f.莱米塞莱Lemisele = BLemisele莱米塞莱
	f.莱米塞莱Lemisele.SetParent(f)
	
	f.塞斯韦根Seswegen = BSeswegen塞斯韦根
	f.塞斯韦根Seswegen.SetParent(f)
	
	f.伊克什基莱Uxkull = BUxkull伊克什基莱
	f.伊克什基莱Uxkull.SetParent(f)
	
	f.文登Wenden = BWenden文登
	f.文登Wenden.SetParent(f)
	
	f.沃尔梅尔Wolmer = BWolmer沃尔梅尔
	f.沃尔梅尔Wolmer.SetParent(f)
	
}
