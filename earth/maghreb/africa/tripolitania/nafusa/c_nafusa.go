package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NafusaCounty interface {
    feud.County
    BDar_dred达尔德雷德() 	feud.Barony
    BDar_jdida达尔吉达() 	feud.Barony
    BDechret_bahirine代什雷巴希林() 	feud.Barony
    BErruda埃鲁达() 	feud.Barony
    BKroussia克鲁西亚() 	feud.Barony
    BNafusa奈富塞() 	feud.Barony
    BSejenane塞杰南() 	feud.Barony
}

type 奈富塞NafusaCounty struct {
	feud.BaseCounty
	达尔德雷德Dar_dred 	feud.Barony
	达尔吉达Dar_jdida 	feud.Barony
	代什雷巴希林Dechret_bahirine 	feud.Barony
	埃鲁达Erruda 	feud.Barony
	克鲁西亚Kroussia 	feud.Barony
	奈富塞Nafusa 	feud.Barony
	塞杰南Sejenane 	feud.Barony
}

func (c *奈富塞NafusaCounty) BDar_dred达尔德雷德() feud.Barony {
	return c.达尔德雷德Dar_dred
}
    
func (c *奈富塞NafusaCounty) BDar_jdida达尔吉达() feud.Barony {
	return c.达尔吉达Dar_jdida
}
    
func (c *奈富塞NafusaCounty) BDechret_bahirine代什雷巴希林() feud.Barony {
	return c.代什雷巴希林Dechret_bahirine
}
    
func (c *奈富塞NafusaCounty) BErruda埃鲁达() feud.Barony {
	return c.埃鲁达Erruda
}
    
func (c *奈富塞NafusaCounty) BKroussia克鲁西亚() feud.Barony {
	return c.克鲁西亚Kroussia
}
    
func (c *奈富塞NafusaCounty) BNafusa奈富塞() feud.Barony {
	return c.奈富塞Nafusa
}
    
func (c *奈富塞NafusaCounty) BSejenane塞杰南() feud.Barony {
	return c.塞杰南Sejenane
}
    
var CNafusa奈富塞 NafusaCounty = &奈富塞NafusaCounty{}

func init() {
	f := CNafusa奈富塞.(*奈富塞NafusaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1723",
		Title:     "nafusa",
		TitleName: "奈富塞",
		TitleCode: "c_nafusa",
		Baronies:  map[string]feud.Barony{},
	}

	f.达尔德雷德Dar_dred = BDar_dred达尔德雷德
	f.达尔德雷德Dar_dred.SetParent(f)
	
	f.达尔吉达Dar_jdida = BDar_jdida达尔吉达
	f.达尔吉达Dar_jdida.SetParent(f)
	
	f.代什雷巴希林Dechret_bahirine = BDechret_bahirine代什雷巴希林
	f.代什雷巴希林Dechret_bahirine.SetParent(f)
	
	f.埃鲁达Erruda = BErruda埃鲁达
	f.埃鲁达Erruda.SetParent(f)
	
	f.克鲁西亚Kroussia = BKroussia克鲁西亚
	f.克鲁西亚Kroussia.SetParent(f)
	
	f.奈富塞Nafusa = BNafusa奈富塞
	f.奈富塞Nafusa.SetParent(f)
	
	f.塞杰南Sejenane = BSejenane塞杰南
	f.塞杰南Sejenane.SetParent(f)
	
}
