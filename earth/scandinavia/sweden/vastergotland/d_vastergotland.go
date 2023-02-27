package vastergotland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/vastergotland/dal"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/vastergotland/narke"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/vastergotland/skara"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/vastergotland/vastergotland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VastergotlandDuke interface {
    feud.Duke
    CDal达尔() 	dal.DalCounty
    CNarke内尔克() 	narke.NarkeCounty
    CSkara斯卡拉() 	skara.SkaraCounty
    CVastergotland西约特兰() 	vastergotland.VastergotlandCounty
}

type 西约特兰VastergotlandDuke struct {
	feud.BaseDuke
	达尔Dal 	dal.DalCounty
	内尔克Narke 	narke.NarkeCounty
	斯卡拉Skara 	skara.SkaraCounty
	西约特兰Vastergotland 	vastergotland.VastergotlandCounty
}

func (d *西约特兰VastergotlandDuke) CDal达尔() dal.DalCounty {
	return d.达尔Dal
}
    
func (d *西约特兰VastergotlandDuke) CNarke内尔克() narke.NarkeCounty {
	return d.内尔克Narke
}
    
func (d *西约特兰VastergotlandDuke) CSkara斯卡拉() skara.SkaraCounty {
	return d.斯卡拉Skara
}
    
func (d *西约特兰VastergotlandDuke) CVastergotland西约特兰() vastergotland.VastergotlandCounty {
	return d.西约特兰Vastergotland
}
    
var DVastergotland西约特兰 VastergotlandDuke = &西约特兰VastergotlandDuke{}

func init() {
	f := DVastergotland西约特兰.(*西约特兰VastergotlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vastergotland",
		TitleName: "西约特兰",
		TitleCode: "d_vastergotland",
		Counties:  map[string]feud.County{},
	}

	f.达尔Dal = dal.CDal达尔
	f.达尔Dal.SetParent(f)
	
	f.内尔克Narke = narke.CNarke内尔克
	f.内尔克Narke.SetParent(f)
	
	f.斯卡拉Skara = skara.CSkara斯卡拉
	f.斯卡拉Skara.SetParent(f)
	
	f.西约特兰Vastergotland = vastergotland.CVastergotland西约特兰
	f.西约特兰Vastergotland.SetParent(f)
	
}
