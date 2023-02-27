package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbydosCounty interface {
    feud.County
    BAbydos阿卑多斯() 	feud.Barony
    BAegae埃盖() 	feud.Barony
    BAlexandriatroas亚历山大特罗亚() 	feud.Barony
    BAllianoi阿里亚诺伊() 	feud.Barony
    BCebrene色布里尼() 	feud.Barony
    BElaia埃莱亚() 	feud.Barony
    BLampsakos兰姆萨科洛斯() 	feud.Barony
    BPigai皮盖伊() 	feud.Barony
}

type 阿卑多斯AbydosCounty struct {
	feud.BaseCounty
	阿卑多斯Abydos 	feud.Barony
	埃盖Aegae 	feud.Barony
	亚历山大特罗亚Alexandriatroas 	feud.Barony
	阿里亚诺伊Allianoi 	feud.Barony
	色布里尼Cebrene 	feud.Barony
	埃莱亚Elaia 	feud.Barony
	兰姆萨科洛斯Lampsakos 	feud.Barony
	皮盖伊Pigai 	feud.Barony
}

func (c *阿卑多斯AbydosCounty) BAbydos阿卑多斯() feud.Barony {
	return c.阿卑多斯Abydos
}
    
func (c *阿卑多斯AbydosCounty) BAegae埃盖() feud.Barony {
	return c.埃盖Aegae
}
    
func (c *阿卑多斯AbydosCounty) BAlexandriatroas亚历山大特罗亚() feud.Barony {
	return c.亚历山大特罗亚Alexandriatroas
}
    
func (c *阿卑多斯AbydosCounty) BAllianoi阿里亚诺伊() feud.Barony {
	return c.阿里亚诺伊Allianoi
}
    
func (c *阿卑多斯AbydosCounty) BCebrene色布里尼() feud.Barony {
	return c.色布里尼Cebrene
}
    
func (c *阿卑多斯AbydosCounty) BElaia埃莱亚() feud.Barony {
	return c.埃莱亚Elaia
}
    
func (c *阿卑多斯AbydosCounty) BLampsakos兰姆萨科洛斯() feud.Barony {
	return c.兰姆萨科洛斯Lampsakos
}
    
func (c *阿卑多斯AbydosCounty) BPigai皮盖伊() feud.Barony {
	return c.皮盖伊Pigai
}
    
var CAbydos阿卑多斯 AbydosCounty = &阿卑多斯AbydosCounty{}

func init() {
	f := CAbydos阿卑多斯.(*阿卑多斯AbydosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "744",
		Title:     "abydos",
		TitleName: "阿卑多斯",
		TitleCode: "c_abydos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卑多斯Abydos = BAbydos阿卑多斯
	f.阿卑多斯Abydos.SetParent(f)
	
	f.埃盖Aegae = BAegae埃盖
	f.埃盖Aegae.SetParent(f)
	
	f.亚历山大特罗亚Alexandriatroas = BAlexandriatroas亚历山大特罗亚
	f.亚历山大特罗亚Alexandriatroas.SetParent(f)
	
	f.阿里亚诺伊Allianoi = BAllianoi阿里亚诺伊
	f.阿里亚诺伊Allianoi.SetParent(f)
	
	f.色布里尼Cebrene = BCebrene色布里尼
	f.色布里尼Cebrene.SetParent(f)
	
	f.埃莱亚Elaia = BElaia埃莱亚
	f.埃莱亚Elaia.SetParent(f)
	
	f.兰姆萨科洛斯Lampsakos = BLampsakos兰姆萨科洛斯
	f.兰姆萨科洛斯Lampsakos.SetParent(f)
	
	f.皮盖伊Pigai = BPigai皮盖伊
	f.皮盖伊Pigai.SetParent(f)
	
}
