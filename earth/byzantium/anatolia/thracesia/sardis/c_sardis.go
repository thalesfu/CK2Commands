package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SardisCounty interface {
    feud.County
    BAezani艾赞尼() 	feud.Barony
    BHermus赫耳墨斯() 	feud.Barony
    BMacentus马肯托斯() 	feud.Barony
    BMaeander迈安德() 	feud.Barony
    BMagnesia马格涅西亚() 	feud.Barony
    BPhiladelphia菲拉德尔菲亚() 	feud.Barony
    BSardis萨第斯() 	feud.Barony
}

type 萨第斯SardisCounty struct {
	feud.BaseCounty
	艾赞尼Aezani 	feud.Barony
	赫耳墨斯Hermus 	feud.Barony
	马肯托斯Macentus 	feud.Barony
	迈安德Maeander 	feud.Barony
	马格涅西亚Magnesia 	feud.Barony
	菲拉德尔菲亚Philadelphia 	feud.Barony
	萨第斯Sardis 	feud.Barony
}

func (c *萨第斯SardisCounty) BAezani艾赞尼() feud.Barony {
	return c.艾赞尼Aezani
}
    
func (c *萨第斯SardisCounty) BHermus赫耳墨斯() feud.Barony {
	return c.赫耳墨斯Hermus
}
    
func (c *萨第斯SardisCounty) BMacentus马肯托斯() feud.Barony {
	return c.马肯托斯Macentus
}
    
func (c *萨第斯SardisCounty) BMaeander迈安德() feud.Barony {
	return c.迈安德Maeander
}
    
func (c *萨第斯SardisCounty) BMagnesia马格涅西亚() feud.Barony {
	return c.马格涅西亚Magnesia
}
    
func (c *萨第斯SardisCounty) BPhiladelphia菲拉德尔菲亚() feud.Barony {
	return c.菲拉德尔菲亚Philadelphia
}
    
func (c *萨第斯SardisCounty) BSardis萨第斯() feud.Barony {
	return c.萨第斯Sardis
}
    
var CSardis萨第斯 SardisCounty = &萨第斯SardisCounty{}

func init() {
	f := CSardis萨第斯.(*萨第斯SardisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1935",
		Title:     "sardis",
		TitleName: "萨第斯",
		TitleCode: "c_sardis",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾赞尼Aezani = BAezani艾赞尼
	f.艾赞尼Aezani.SetParent(f)
	
	f.赫耳墨斯Hermus = BHermus赫耳墨斯
	f.赫耳墨斯Hermus.SetParent(f)
	
	f.马肯托斯Macentus = BMacentus马肯托斯
	f.马肯托斯Macentus.SetParent(f)
	
	f.迈安德Maeander = BMaeander迈安德
	f.迈安德Maeander.SetParent(f)
	
	f.马格涅西亚Magnesia = BMagnesia马格涅西亚
	f.马格涅西亚Magnesia.SetParent(f)
	
	f.菲拉德尔菲亚Philadelphia = BPhiladelphia菲拉德尔菲亚
	f.菲拉德尔菲亚Philadelphia.SetParent(f)
	
	f.萨第斯Sardis = BSardis萨第斯
	f.萨第斯Sardis.SetParent(f)
	
}
