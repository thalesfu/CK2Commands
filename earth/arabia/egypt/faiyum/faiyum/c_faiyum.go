package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FaiyumCounty interface {
    feud.County
    BAhnas阿赫纳斯() 	feud.Barony
    BArsinoe_faiyum阿耳西倪() 	feud.Barony
    BBahnasa巴赫纳萨() 	feud.Barony
    BBeni_suef贝尼苏埃夫() 	feud.Barony
    BFaiyum法尤姆() 	feud.Barony
    BHeracleus赫拉克琉斯() 	feud.Barony
    BWadi_el_rayan雷延谷() 	feud.Barony
}

type 法尤姆FaiyumCounty struct {
	feud.BaseCounty
	阿赫纳斯Ahnas 	feud.Barony
	阿耳西倪Arsinoe_faiyum 	feud.Barony
	巴赫纳萨Bahnasa 	feud.Barony
	贝尼苏埃夫Beni_suef 	feud.Barony
	法尤姆Faiyum 	feud.Barony
	赫拉克琉斯Heracleus 	feud.Barony
	雷延谷Wadi_el_rayan 	feud.Barony
}

func (c *法尤姆FaiyumCounty) BAhnas阿赫纳斯() feud.Barony {
	return c.阿赫纳斯Ahnas
}
    
func (c *法尤姆FaiyumCounty) BArsinoe_faiyum阿耳西倪() feud.Barony {
	return c.阿耳西倪Arsinoe_faiyum
}
    
func (c *法尤姆FaiyumCounty) BBahnasa巴赫纳萨() feud.Barony {
	return c.巴赫纳萨Bahnasa
}
    
func (c *法尤姆FaiyumCounty) BBeni_suef贝尼苏埃夫() feud.Barony {
	return c.贝尼苏埃夫Beni_suef
}
    
func (c *法尤姆FaiyumCounty) BFaiyum法尤姆() feud.Barony {
	return c.法尤姆Faiyum
}
    
func (c *法尤姆FaiyumCounty) BHeracleus赫拉克琉斯() feud.Barony {
	return c.赫拉克琉斯Heracleus
}
    
func (c *法尤姆FaiyumCounty) BWadi_el_rayan雷延谷() feud.Barony {
	return c.雷延谷Wadi_el_rayan
}
    
var CFaiyum法尤姆 FaiyumCounty = &法尤姆FaiyumCounty{}

func init() {
	f := CFaiyum法尤姆.(*法尤姆FaiyumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2001",
		Title:     "faiyum",
		TitleName: "法尤姆",
		TitleCode: "c_faiyum",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫纳斯Ahnas = BAhnas阿赫纳斯
	f.阿赫纳斯Ahnas.SetParent(f)
	
	f.阿耳西倪Arsinoe_faiyum = BArsinoe_faiyum阿耳西倪
	f.阿耳西倪Arsinoe_faiyum.SetParent(f)
	
	f.巴赫纳萨Bahnasa = BBahnasa巴赫纳萨
	f.巴赫纳萨Bahnasa.SetParent(f)
	
	f.贝尼苏埃夫Beni_suef = BBeni_suef贝尼苏埃夫
	f.贝尼苏埃夫Beni_suef.SetParent(f)
	
	f.法尤姆Faiyum = BFaiyum法尤姆
	f.法尤姆Faiyum.SetParent(f)
	
	f.赫拉克琉斯Heracleus = BHeracleus赫拉克琉斯
	f.赫拉克琉斯Heracleus.SetParent(f)
	
	f.雷延谷Wadi_el_rayan = BWadi_el_rayan雷延谷
	f.雷延谷Wadi_el_rayan.SetParent(f)
	
}
