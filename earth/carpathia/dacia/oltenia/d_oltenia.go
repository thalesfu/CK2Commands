package oltenia

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/oltenia/craiova"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/oltenia/severin"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/oltenia/targu_jiu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OlteniaDuke interface {
    feud.Duke
    CCraiova克拉约瓦() 	craiova.CraiovaCounty
    CSeverin塞韦林() 	severin.SeverinCounty
    CTargu_jiu特尔古日乌() 	targu_jiu.Targu_jiuCounty
}

type 奥尔泰尼亚OlteniaDuke struct {
	feud.BaseDuke
	克拉约瓦Craiova 	craiova.CraiovaCounty
	塞韦林Severin 	severin.SeverinCounty
	特尔古日乌Targu_jiu 	targu_jiu.Targu_jiuCounty
}

func (d *奥尔泰尼亚OlteniaDuke) CCraiova克拉约瓦() craiova.CraiovaCounty {
	return d.克拉约瓦Craiova
}
    
func (d *奥尔泰尼亚OlteniaDuke) CSeverin塞韦林() severin.SeverinCounty {
	return d.塞韦林Severin
}
    
func (d *奥尔泰尼亚OlteniaDuke) CTargu_jiu特尔古日乌() targu_jiu.Targu_jiuCounty {
	return d.特尔古日乌Targu_jiu
}
    
var DOltenia奥尔泰尼亚 OlteniaDuke = &奥尔泰尼亚OlteniaDuke{}

func init() {
	f := DOltenia奥尔泰尼亚.(*奥尔泰尼亚OlteniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "oltenia",
		TitleName: "奥尔泰尼亚",
		TitleCode: "d_oltenia",
		Counties:  map[string]feud.County{},
	}

	f.克拉约瓦Craiova = craiova.CCraiova克拉约瓦
	f.克拉约瓦Craiova.SetParent(f)
	
	f.塞韦林Severin = severin.CSeverin塞韦林
	f.塞韦林Severin.SetParent(f)
	
	f.特尔古日乌Targu_jiu = targu_jiu.CTargu_jiu特尔古日乌
	f.特尔古日乌Targu_jiu.SetParent(f)
	
}
