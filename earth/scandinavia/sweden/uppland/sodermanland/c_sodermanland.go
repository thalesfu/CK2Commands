package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SodermanlandCounty interface {
    feud.County
    BEskilstuna埃斯基尔斯图纳() 	feud.Barony
    BGripsholm格利普霍姆() 	feud.Barony
    BHundhamra洪德哈姆拉() 	feud.Barony
    BNykoping尼雪平() 	feud.Barony
    BStrangnas斯特兰奈斯() 	feud.Barony
    BTelge特尔赫() 	feud.Barony
    BTorshalla托什海拉() 	feud.Barony
    BVaderbrunn韦德布鲁恩() 	feud.Barony
}

type 南曼兰SodermanlandCounty struct {
	feud.BaseCounty
	埃斯基尔斯图纳Eskilstuna 	feud.Barony
	格利普霍姆Gripsholm 	feud.Barony
	洪德哈姆拉Hundhamra 	feud.Barony
	尼雪平Nykoping 	feud.Barony
	斯特兰奈斯Strangnas 	feud.Barony
	特尔赫Telge 	feud.Barony
	托什海拉Torshalla 	feud.Barony
	韦德布鲁恩Vaderbrunn 	feud.Barony
}

func (c *南曼兰SodermanlandCounty) BEskilstuna埃斯基尔斯图纳() feud.Barony {
	return c.埃斯基尔斯图纳Eskilstuna
}
    
func (c *南曼兰SodermanlandCounty) BGripsholm格利普霍姆() feud.Barony {
	return c.格利普霍姆Gripsholm
}
    
func (c *南曼兰SodermanlandCounty) BHundhamra洪德哈姆拉() feud.Barony {
	return c.洪德哈姆拉Hundhamra
}
    
func (c *南曼兰SodermanlandCounty) BNykoping尼雪平() feud.Barony {
	return c.尼雪平Nykoping
}
    
func (c *南曼兰SodermanlandCounty) BStrangnas斯特兰奈斯() feud.Barony {
	return c.斯特兰奈斯Strangnas
}
    
func (c *南曼兰SodermanlandCounty) BTelge特尔赫() feud.Barony {
	return c.特尔赫Telge
}
    
func (c *南曼兰SodermanlandCounty) BTorshalla托什海拉() feud.Barony {
	return c.托什海拉Torshalla
}
    
func (c *南曼兰SodermanlandCounty) BVaderbrunn韦德布鲁恩() feud.Barony {
	return c.韦德布鲁恩Vaderbrunn
}
    
var CSodermanland南曼兰 SodermanlandCounty = &南曼兰SodermanlandCounty{}

func init() {
	f := CSodermanland南曼兰.(*南曼兰SodermanlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "292",
		Title:     "sodermanland",
		TitleName: "南曼兰",
		TitleCode: "c_sodermanland",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃斯基尔斯图纳Eskilstuna = BEskilstuna埃斯基尔斯图纳
	f.埃斯基尔斯图纳Eskilstuna.SetParent(f)
	
	f.格利普霍姆Gripsholm = BGripsholm格利普霍姆
	f.格利普霍姆Gripsholm.SetParent(f)
	
	f.洪德哈姆拉Hundhamra = BHundhamra洪德哈姆拉
	f.洪德哈姆拉Hundhamra.SetParent(f)
	
	f.尼雪平Nykoping = BNykoping尼雪平
	f.尼雪平Nykoping.SetParent(f)
	
	f.斯特兰奈斯Strangnas = BStrangnas斯特兰奈斯
	f.斯特兰奈斯Strangnas.SetParent(f)
	
	f.特尔赫Telge = BTelge特尔赫
	f.特尔赫Telge.SetParent(f)
	
	f.托什海拉Torshalla = BTorshalla托什海拉
	f.托什海拉Torshalla.SetParent(f)
	
	f.韦德布鲁恩Vaderbrunn = BVaderbrunn韦德布鲁恩
	f.韦德布鲁恩Vaderbrunn.SetParent(f)
	
}
