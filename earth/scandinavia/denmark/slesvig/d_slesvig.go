package slesvig

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/slesvig/abosyssel"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/slesvig/hardsyssel"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/slesvig/himmersyssel"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/slesvig/jylland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SlesvigDuke interface {
    feud.Duke
    CAbosyssel奥胡斯() 	abosyssel.AbosysselCounty
    CHardsyssel瓦德() 	hardsyssel.HardsysselCounty
    CHimmersyssel奥尔堡() 	himmersyssel.HimmersysselCounty
    CJylland斯卡恩() 	jylland.JyllandCounty
}

type 日德兰SlesvigDuke struct {
	feud.BaseDuke
	奥胡斯Abosyssel 	abosyssel.AbosysselCounty
	瓦德Hardsyssel 	hardsyssel.HardsysselCounty
	奥尔堡Himmersyssel 	himmersyssel.HimmersysselCounty
	斯卡恩Jylland 	jylland.JyllandCounty
}

func (d *日德兰SlesvigDuke) CAbosyssel奥胡斯() abosyssel.AbosysselCounty {
	return d.奥胡斯Abosyssel
}
    
func (d *日德兰SlesvigDuke) CHardsyssel瓦德() hardsyssel.HardsysselCounty {
	return d.瓦德Hardsyssel
}
    
func (d *日德兰SlesvigDuke) CHimmersyssel奥尔堡() himmersyssel.HimmersysselCounty {
	return d.奥尔堡Himmersyssel
}
    
func (d *日德兰SlesvigDuke) CJylland斯卡恩() jylland.JyllandCounty {
	return d.斯卡恩Jylland
}
    
var DSlesvig日德兰 SlesvigDuke = &日德兰SlesvigDuke{}

func init() {
	f := DSlesvig日德兰.(*日德兰SlesvigDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "slesvig",
		TitleName: "日德兰",
		TitleCode: "d_slesvig",
		Counties:  map[string]feud.County{},
	}

	f.奥胡斯Abosyssel = abosyssel.CAbosyssel奥胡斯
	f.奥胡斯Abosyssel.SetParent(f)
	
	f.瓦德Hardsyssel = hardsyssel.CHardsyssel瓦德
	f.瓦德Hardsyssel.SetParent(f)
	
	f.奥尔堡Himmersyssel = himmersyssel.CHimmersyssel奥尔堡
	f.奥尔堡Himmersyssel.SetParent(f)
	
	f.斯卡恩Jylland = jylland.CJylland斯卡恩
	f.斯卡恩Jylland.SetParent(f)
	
}
