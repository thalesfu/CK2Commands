package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type El_rifCounty interface {
    feud.County
    BAgadirene阿加迪朗() 	feud.Barony
    BAlhoceima胡塞马() 	feud.Barony
    BDriouch代尔尤什() 	feud.Barony
    BId_amy伊德阿米() 	feud.Barony
    BKetama克塔玛() 	feud.Barony
    BMelilla梅利利亚() 	feud.Barony
    BMidar米达尔() 	feud.Barony
    BNador纳祖尔() 	feud.Barony
}

type 里夫El_rifCounty struct {
	feud.BaseCounty
	阿加迪朗Agadirene 	feud.Barony
	胡塞马Alhoceima 	feud.Barony
	代尔尤什Driouch 	feud.Barony
	伊德阿米Id_amy 	feud.Barony
	克塔玛Ketama 	feud.Barony
	梅利利亚Melilla 	feud.Barony
	米达尔Midar 	feud.Barony
	纳祖尔Nador 	feud.Barony
}

func (c *里夫El_rifCounty) BAgadirene阿加迪朗() feud.Barony {
	return c.阿加迪朗Agadirene
}
    
func (c *里夫El_rifCounty) BAlhoceima胡塞马() feud.Barony {
	return c.胡塞马Alhoceima
}
    
func (c *里夫El_rifCounty) BDriouch代尔尤什() feud.Barony {
	return c.代尔尤什Driouch
}
    
func (c *里夫El_rifCounty) BId_amy伊德阿米() feud.Barony {
	return c.伊德阿米Id_amy
}
    
func (c *里夫El_rifCounty) BKetama克塔玛() feud.Barony {
	return c.克塔玛Ketama
}
    
func (c *里夫El_rifCounty) BMelilla梅利利亚() feud.Barony {
	return c.梅利利亚Melilla
}
    
func (c *里夫El_rifCounty) BMidar米达尔() feud.Barony {
	return c.米达尔Midar
}
    
func (c *里夫El_rifCounty) BNador纳祖尔() feud.Barony {
	return c.纳祖尔Nador
}
    
var CEl_rif里夫 El_rifCounty = &里夫El_rifCounty{}

func init() {
	f := CEl_rif里夫.(*里夫El_rifCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "838",
		Title:     "el_rif",
		TitleName: "里夫",
		TitleCode: "c_el_rif",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿加迪朗Agadirene = BAgadirene阿加迪朗
	f.阿加迪朗Agadirene.SetParent(f)
	
	f.胡塞马Alhoceima = BAlhoceima胡塞马
	f.胡塞马Alhoceima.SetParent(f)
	
	f.代尔尤什Driouch = BDriouch代尔尤什
	f.代尔尤什Driouch.SetParent(f)
	
	f.伊德阿米Id_amy = BId_amy伊德阿米
	f.伊德阿米Id_amy.SetParent(f)
	
	f.克塔玛Ketama = BKetama克塔玛
	f.克塔玛Ketama.SetParent(f)
	
	f.梅利利亚Melilla = BMelilla梅利利亚
	f.梅利利亚Melilla.SetParent(f)
	
	f.米达尔Midar = BMidar米达尔
	f.米达尔Midar.SetParent(f)
	
	f.纳祖尔Nador = BNador纳祖尔
	f.纳祖尔Nador.SetParent(f)
	
}
