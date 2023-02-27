package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SnassenCounty interface {
    feud.County
    BAhfir艾赫菲尔() 	feud.Barony
    BAinbenimathar艾因贝尼迈特海尔() 	feud.Barony
    BBerkane拜尔坎() 	feud.Barony
    BDebdou德卜杜() 	feud.Barony
    BJerada杰拉代() 	feud.Barony
    BOujda乌季达() 	feud.Barony
    BSaldia萨伊迪耶() 	feud.Barony
    BTendrara坦德拉拉() 	feud.Barony
}

type 瓦伊达SnassenCounty struct {
	feud.BaseCounty
	艾赫菲尔Ahfir 	feud.Barony
	艾因贝尼迈特海尔Ainbenimathar 	feud.Barony
	拜尔坎Berkane 	feud.Barony
	德卜杜Debdou 	feud.Barony
	杰拉代Jerada 	feud.Barony
	乌季达Oujda 	feud.Barony
	萨伊迪耶Saldia 	feud.Barony
	坦德拉拉Tendrara 	feud.Barony
}

func (c *瓦伊达SnassenCounty) BAhfir艾赫菲尔() feud.Barony {
	return c.艾赫菲尔Ahfir
}
    
func (c *瓦伊达SnassenCounty) BAinbenimathar艾因贝尼迈特海尔() feud.Barony {
	return c.艾因贝尼迈特海尔Ainbenimathar
}
    
func (c *瓦伊达SnassenCounty) BBerkane拜尔坎() feud.Barony {
	return c.拜尔坎Berkane
}
    
func (c *瓦伊达SnassenCounty) BDebdou德卜杜() feud.Barony {
	return c.德卜杜Debdou
}
    
func (c *瓦伊达SnassenCounty) BJerada杰拉代() feud.Barony {
	return c.杰拉代Jerada
}
    
func (c *瓦伊达SnassenCounty) BOujda乌季达() feud.Barony {
	return c.乌季达Oujda
}
    
func (c *瓦伊达SnassenCounty) BSaldia萨伊迪耶() feud.Barony {
	return c.萨伊迪耶Saldia
}
    
func (c *瓦伊达SnassenCounty) BTendrara坦德拉拉() feud.Barony {
	return c.坦德拉拉Tendrara
}
    
var CSnassen瓦伊达 SnassenCounty = &瓦伊达SnassenCounty{}

func init() {
	f := CSnassen瓦伊达.(*瓦伊达SnassenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "836",
		Title:     "snassen",
		TitleName: "瓦伊达",
		TitleCode: "c_snassen",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾赫菲尔Ahfir = BAhfir艾赫菲尔
	f.艾赫菲尔Ahfir.SetParent(f)
	
	f.艾因贝尼迈特海尔Ainbenimathar = BAinbenimathar艾因贝尼迈特海尔
	f.艾因贝尼迈特海尔Ainbenimathar.SetParent(f)
	
	f.拜尔坎Berkane = BBerkane拜尔坎
	f.拜尔坎Berkane.SetParent(f)
	
	f.德卜杜Debdou = BDebdou德卜杜
	f.德卜杜Debdou.SetParent(f)
	
	f.杰拉代Jerada = BJerada杰拉代
	f.杰拉代Jerada.SetParent(f)
	
	f.乌季达Oujda = BOujda乌季达
	f.乌季达Oujda.SetParent(f)
	
	f.萨伊迪耶Saldia = BSaldia萨伊迪耶
	f.萨伊迪耶Saldia.SetParent(f)
	
	f.坦德拉拉Tendrara = BTendrara坦德拉拉
	f.坦德拉拉Tendrara.SetParent(f)
	
}
