package sakala

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/sakala/dorpat"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/sakala/jarva"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/sakala/lettigalians"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SakalaDuke interface {
    feud.Duke
    CDorpat塔尔图() 	dorpat.DorpatCounty
    CJarva耶尔瓦() 	jarva.JarvaCounty
    CLettigalians拉特加利亚() 	lettigalians.LettigaliansCounty
}

type 萨卡拉SakalaDuke struct {
	feud.BaseDuke
	塔尔图Dorpat 	dorpat.DorpatCounty
	耶尔瓦Jarva 	jarva.JarvaCounty
	拉特加利亚Lettigalians 	lettigalians.LettigaliansCounty
}

func (d *萨卡拉SakalaDuke) CDorpat塔尔图() dorpat.DorpatCounty {
	return d.塔尔图Dorpat
}
    
func (d *萨卡拉SakalaDuke) CJarva耶尔瓦() jarva.JarvaCounty {
	return d.耶尔瓦Jarva
}
    
func (d *萨卡拉SakalaDuke) CLettigalians拉特加利亚() lettigalians.LettigaliansCounty {
	return d.拉特加利亚Lettigalians
}
    
var DSakala萨卡拉 SakalaDuke = &萨卡拉SakalaDuke{}

func init() {
	f := DSakala萨卡拉.(*萨卡拉SakalaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sakala",
		TitleName: "萨卡拉",
		TitleCode: "d_sakala",
		Counties:  map[string]feud.County{},
	}

	f.塔尔图Dorpat = dorpat.CDorpat塔尔图
	f.塔尔图Dorpat.SetParent(f)
	
	f.耶尔瓦Jarva = jarva.CJarva耶尔瓦
	f.耶尔瓦Jarva.SetParent(f)
	
	f.拉特加利亚Lettigalians = lettigalians.CLettigalians拉特加利亚
	f.拉特加利亚Lettigalians.SetParent(f)
	
}
