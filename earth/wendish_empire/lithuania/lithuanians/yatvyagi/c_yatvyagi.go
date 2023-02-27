package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YatvyagiCounty interface {
    feud.County
    BAlytus阿利图斯() 	feud.Barony
    BBirstonas比尔什托纳斯() 	feud.Barony
    BDarsuniskis达尔苏尼斯基斯() 	feud.Barony
    BGailen盖伦() 	feud.Barony
    BKaunas考纳斯() 	feud.Barony
    BRaiziai赖日艾() 	feud.Barony
    BVilkaviskis维尔卡维什基斯() 	feud.Barony
}

type 考纳斯YatvyagiCounty struct {
	feud.BaseCounty
	阿利图斯Alytus 	feud.Barony
	比尔什托纳斯Birstonas 	feud.Barony
	达尔苏尼斯基斯Darsuniskis 	feud.Barony
	盖伦Gailen 	feud.Barony
	考纳斯Kaunas 	feud.Barony
	赖日艾Raiziai 	feud.Barony
	维尔卡维什基斯Vilkaviskis 	feud.Barony
}

func (c *考纳斯YatvyagiCounty) BAlytus阿利图斯() feud.Barony {
	return c.阿利图斯Alytus
}
    
func (c *考纳斯YatvyagiCounty) BBirstonas比尔什托纳斯() feud.Barony {
	return c.比尔什托纳斯Birstonas
}
    
func (c *考纳斯YatvyagiCounty) BDarsuniskis达尔苏尼斯基斯() feud.Barony {
	return c.达尔苏尼斯基斯Darsuniskis
}
    
func (c *考纳斯YatvyagiCounty) BGailen盖伦() feud.Barony {
	return c.盖伦Gailen
}
    
func (c *考纳斯YatvyagiCounty) BKaunas考纳斯() feud.Barony {
	return c.考纳斯Kaunas
}
    
func (c *考纳斯YatvyagiCounty) BRaiziai赖日艾() feud.Barony {
	return c.赖日艾Raiziai
}
    
func (c *考纳斯YatvyagiCounty) BVilkaviskis维尔卡维什基斯() feud.Barony {
	return c.维尔卡维什基斯Vilkaviskis
}
    
var CYatvyagi考纳斯 YatvyagiCounty = &考纳斯YatvyagiCounty{}

func init() {
	f := CYatvyagi考纳斯.(*考纳斯YatvyagiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "426",
		Title:     "yatvyagi",
		TitleName: "考纳斯",
		TitleCode: "c_yatvyagi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿利图斯Alytus = BAlytus阿利图斯
	f.阿利图斯Alytus.SetParent(f)
	
	f.比尔什托纳斯Birstonas = BBirstonas比尔什托纳斯
	f.比尔什托纳斯Birstonas.SetParent(f)
	
	f.达尔苏尼斯基斯Darsuniskis = BDarsuniskis达尔苏尼斯基斯
	f.达尔苏尼斯基斯Darsuniskis.SetParent(f)
	
	f.盖伦Gailen = BGailen盖伦
	f.盖伦Gailen.SetParent(f)
	
	f.考纳斯Kaunas = BKaunas考纳斯
	f.考纳斯Kaunas.SetParent(f)
	
	f.赖日艾Raiziai = BRaiziai赖日艾
	f.赖日艾Raiziai.SetParent(f)
	
	f.维尔卡维什基斯Vilkaviskis = BVilkaviskis维尔卡维什基斯
	f.维尔卡维什基斯Vilkaviskis.SetParent(f)
	
}
