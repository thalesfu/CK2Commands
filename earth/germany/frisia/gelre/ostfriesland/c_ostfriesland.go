package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OstfrieslandCounty interface {
    feud.County
    BAurich奥里希() 	feud.Barony
    BBorkum博尔库姆() 	feud.Barony
    BEmden埃姆登() 	feud.Barony
    BEsens埃森斯() 	feud.Barony
    BLeer莱尔() 	feud.Barony
    BNorden诺登() 	feud.Barony
    BNorderney诺德奈() 	feud.Barony
    BWittmund维特蒙德() 	feud.Barony
}

type 东弗里斯兰OstfrieslandCounty struct {
	feud.BaseCounty
	奥里希Aurich 	feud.Barony
	博尔库姆Borkum 	feud.Barony
	埃姆登Emden 	feud.Barony
	埃森斯Esens 	feud.Barony
	莱尔Leer 	feud.Barony
	诺登Norden 	feud.Barony
	诺德奈Norderney 	feud.Barony
	维特蒙德Wittmund 	feud.Barony
}

func (c *东弗里斯兰OstfrieslandCounty) BAurich奥里希() feud.Barony {
	return c.奥里希Aurich
}
    
func (c *东弗里斯兰OstfrieslandCounty) BBorkum博尔库姆() feud.Barony {
	return c.博尔库姆Borkum
}
    
func (c *东弗里斯兰OstfrieslandCounty) BEmden埃姆登() feud.Barony {
	return c.埃姆登Emden
}
    
func (c *东弗里斯兰OstfrieslandCounty) BEsens埃森斯() feud.Barony {
	return c.埃森斯Esens
}
    
func (c *东弗里斯兰OstfrieslandCounty) BLeer莱尔() feud.Barony {
	return c.莱尔Leer
}
    
func (c *东弗里斯兰OstfrieslandCounty) BNorden诺登() feud.Barony {
	return c.诺登Norden
}
    
func (c *东弗里斯兰OstfrieslandCounty) BNorderney诺德奈() feud.Barony {
	return c.诺德奈Norderney
}
    
func (c *东弗里斯兰OstfrieslandCounty) BWittmund维特蒙德() feud.Barony {
	return c.维特蒙德Wittmund
}
    
var COstfriesland东弗里斯兰 OstfrieslandCounty = &东弗里斯兰OstfrieslandCounty{}

func init() {
	f := COstfriesland东弗里斯兰.(*东弗里斯兰OstfrieslandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "85",
		Title:     "ostfriesland",
		TitleName: "东弗里斯兰",
		TitleCode: "c_ostfriesland",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥里希Aurich = BAurich奥里希
	f.奥里希Aurich.SetParent(f)
	
	f.博尔库姆Borkum = BBorkum博尔库姆
	f.博尔库姆Borkum.SetParent(f)
	
	f.埃姆登Emden = BEmden埃姆登
	f.埃姆登Emden.SetParent(f)
	
	f.埃森斯Esens = BEsens埃森斯
	f.埃森斯Esens.SetParent(f)
	
	f.莱尔Leer = BLeer莱尔
	f.莱尔Leer.SetParent(f)
	
	f.诺登Norden = BNorden诺登
	f.诺登Norden.SetParent(f)
	
	f.诺德奈Norderney = BNorderney诺德奈
	f.诺德奈Norderney.SetParent(f)
	
	f.维特蒙德Wittmund = BWittmund维特蒙德
	f.维特蒙德Wittmund.SetParent(f)
	
}
