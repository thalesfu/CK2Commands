package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Castelo_brancoCounty interface {
    feud.County
    BAcores亚速尔() 	feud.Barony
    BAlmeida阿尔梅达() 	feud.Barony
    BGuarda瓜达() 	feud.Barony
    BLamego拉梅戈() 	feud.Barony
    BPinhel皮涅尔() 	feud.Barony
    BSabugal萨布加尔() 	feud.Barony
    BTrancoso特兰科苏() 	feud.Barony
}

type 瓜尔达Castelo_brancoCounty struct {
	feud.BaseCounty
	亚速尔Acores 	feud.Barony
	阿尔梅达Almeida 	feud.Barony
	瓜达Guarda 	feud.Barony
	拉梅戈Lamego 	feud.Barony
	皮涅尔Pinhel 	feud.Barony
	萨布加尔Sabugal 	feud.Barony
	特兰科苏Trancoso 	feud.Barony
}

func (c *瓜尔达Castelo_brancoCounty) BAcores亚速尔() feud.Barony {
	return c.亚速尔Acores
}
    
func (c *瓜尔达Castelo_brancoCounty) BAlmeida阿尔梅达() feud.Barony {
	return c.阿尔梅达Almeida
}
    
func (c *瓜尔达Castelo_brancoCounty) BGuarda瓜达() feud.Barony {
	return c.瓜达Guarda
}
    
func (c *瓜尔达Castelo_brancoCounty) BLamego拉梅戈() feud.Barony {
	return c.拉梅戈Lamego
}
    
func (c *瓜尔达Castelo_brancoCounty) BPinhel皮涅尔() feud.Barony {
	return c.皮涅尔Pinhel
}
    
func (c *瓜尔达Castelo_brancoCounty) BSabugal萨布加尔() feud.Barony {
	return c.萨布加尔Sabugal
}
    
func (c *瓜尔达Castelo_brancoCounty) BTrancoso特兰科苏() feud.Barony {
	return c.特兰科苏Trancoso
}
    
var CCastelo_branco瓜尔达 Castelo_brancoCounty = &瓜尔达Castelo_brancoCounty{}

func init() {
	f := CCastelo_branco瓜尔达.(*瓜尔达Castelo_brancoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "187",
		Title:     "castelo_branco",
		TitleName: "瓜尔达",
		TitleCode: "c_castelo_branco",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚速尔Acores = BAcores亚速尔
	f.亚速尔Acores.SetParent(f)
	
	f.阿尔梅达Almeida = BAlmeida阿尔梅达
	f.阿尔梅达Almeida.SetParent(f)
	
	f.瓜达Guarda = BGuarda瓜达
	f.瓜达Guarda.SetParent(f)
	
	f.拉梅戈Lamego = BLamego拉梅戈
	f.拉梅戈Lamego.SetParent(f)
	
	f.皮涅尔Pinhel = BPinhel皮涅尔
	f.皮涅尔Pinhel.SetParent(f)
	
	f.萨布加尔Sabugal = BSabugal萨布加尔
	f.萨布加尔Sabugal.SetParent(f)
	
	f.特兰科苏Trancoso = BTrancoso特兰科苏
	f.特兰科苏Trancoso.SetParent(f)
	
}
