package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EsnaCounty interface {
    feud.County
    BArmant阿尔曼特() 	feud.Barony
    BEdfu埃德富() 	feud.Barony
    BEgypabydos阿拜多斯() 	feud.Barony
    BEgypthebes底比斯() 	feud.Barony
    BEsna埃斯纳() 	feud.Barony
    BHiw希乌() 	feud.Barony
    BKhnum赫努姆() 	feud.Barony
}

type 埃斯纳EsnaCounty struct {
	feud.BaseCounty
	阿尔曼特Armant 	feud.Barony
	埃德富Edfu 	feud.Barony
	阿拜多斯Egypabydos 	feud.Barony
	底比斯Egypthebes 	feud.Barony
	埃斯纳Esna 	feud.Barony
	希乌Hiw 	feud.Barony
	赫努姆Khnum 	feud.Barony
}

func (c *埃斯纳EsnaCounty) BArmant阿尔曼特() feud.Barony {
	return c.阿尔曼特Armant
}
    
func (c *埃斯纳EsnaCounty) BEdfu埃德富() feud.Barony {
	return c.埃德富Edfu
}
    
func (c *埃斯纳EsnaCounty) BEgypabydos阿拜多斯() feud.Barony {
	return c.阿拜多斯Egypabydos
}
    
func (c *埃斯纳EsnaCounty) BEgypthebes底比斯() feud.Barony {
	return c.底比斯Egypthebes
}
    
func (c *埃斯纳EsnaCounty) BEsna埃斯纳() feud.Barony {
	return c.埃斯纳Esna
}
    
func (c *埃斯纳EsnaCounty) BHiw希乌() feud.Barony {
	return c.希乌Hiw
}
    
func (c *埃斯纳EsnaCounty) BKhnum赫努姆() feud.Barony {
	return c.赫努姆Khnum
}
    
var CEsna埃斯纳 EsnaCounty = &埃斯纳EsnaCounty{}

func init() {
	f := CEsna埃斯纳.(*埃斯纳EsnaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2003",
		Title:     "esna",
		TitleName: "埃斯纳",
		TitleCode: "c_esna",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔曼特Armant = BArmant阿尔曼特
	f.阿尔曼特Armant.SetParent(f)
	
	f.埃德富Edfu = BEdfu埃德富
	f.埃德富Edfu.SetParent(f)
	
	f.阿拜多斯Egypabydos = BEgypabydos阿拜多斯
	f.阿拜多斯Egypabydos.SetParent(f)
	
	f.底比斯Egypthebes = BEgypthebes底比斯
	f.底比斯Egypthebes.SetParent(f)
	
	f.埃斯纳Esna = BEsna埃斯纳
	f.埃斯纳Esna.SetParent(f)
	
	f.希乌Hiw = BHiw希乌
	f.希乌Hiw.SetParent(f)
	
	f.赫努姆Khnum = BKhnum赫努姆
	f.赫努姆Khnum.SetParent(f)
	
}
