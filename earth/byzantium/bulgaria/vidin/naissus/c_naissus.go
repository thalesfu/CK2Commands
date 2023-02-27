package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaissusCounty interface {
    feud.County
    BBrdo布尔多() 	feud.Barony
    BKambelevac坎贝莱瓦茨() 	feud.Barony
    BKnjazevac克尼亚热瓦茨() 	feud.Barony
    BKoprijan科普里扬() 	feud.Barony
    BKumanovo库马诺沃() 	feud.Barony
    BLesnovo莱斯诺沃() 	feud.Barony
    BNish尼什() 	feud.Barony
    BVranje弗拉涅() 	feud.Barony
}

type 奈索斯NaissusCounty struct {
	feud.BaseCounty
	布尔多Brdo 	feud.Barony
	坎贝莱瓦茨Kambelevac 	feud.Barony
	克尼亚热瓦茨Knjazevac 	feud.Barony
	科普里扬Koprijan 	feud.Barony
	库马诺沃Kumanovo 	feud.Barony
	莱斯诺沃Lesnovo 	feud.Barony
	尼什Nish 	feud.Barony
	弗拉涅Vranje 	feud.Barony
}

func (c *奈索斯NaissusCounty) BBrdo布尔多() feud.Barony {
	return c.布尔多Brdo
}
    
func (c *奈索斯NaissusCounty) BKambelevac坎贝莱瓦茨() feud.Barony {
	return c.坎贝莱瓦茨Kambelevac
}
    
func (c *奈索斯NaissusCounty) BKnjazevac克尼亚热瓦茨() feud.Barony {
	return c.克尼亚热瓦茨Knjazevac
}
    
func (c *奈索斯NaissusCounty) BKoprijan科普里扬() feud.Barony {
	return c.科普里扬Koprijan
}
    
func (c *奈索斯NaissusCounty) BKumanovo库马诺沃() feud.Barony {
	return c.库马诺沃Kumanovo
}
    
func (c *奈索斯NaissusCounty) BLesnovo莱斯诺沃() feud.Barony {
	return c.莱斯诺沃Lesnovo
}
    
func (c *奈索斯NaissusCounty) BNish尼什() feud.Barony {
	return c.尼什Nish
}
    
func (c *奈索斯NaissusCounty) BVranje弗拉涅() feud.Barony {
	return c.弗拉涅Vranje
}
    
var CNaissus奈索斯 NaissusCounty = &奈索斯NaissusCounty{}

func init() {
	f := CNaissus奈索斯.(*奈索斯NaissusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "501",
		Title:     "naissus",
		TitleName: "奈索斯",
		TitleCode: "c_naissus",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔多Brdo = BBrdo布尔多
	f.布尔多Brdo.SetParent(f)
	
	f.坎贝莱瓦茨Kambelevac = BKambelevac坎贝莱瓦茨
	f.坎贝莱瓦茨Kambelevac.SetParent(f)
	
	f.克尼亚热瓦茨Knjazevac = BKnjazevac克尼亚热瓦茨
	f.克尼亚热瓦茨Knjazevac.SetParent(f)
	
	f.科普里扬Koprijan = BKoprijan科普里扬
	f.科普里扬Koprijan.SetParent(f)
	
	f.库马诺沃Kumanovo = BKumanovo库马诺沃
	f.库马诺沃Kumanovo.SetParent(f)
	
	f.莱斯诺沃Lesnovo = BLesnovo莱斯诺沃
	f.莱斯诺沃Lesnovo.SetParent(f)
	
	f.尼什Nish = BNish尼什
	f.尼什Nish.SetParent(f)
	
	f.弗拉涅Vranje = BVranje弗拉涅
	f.弗拉涅Vranje.SetParent(f)
	
}
