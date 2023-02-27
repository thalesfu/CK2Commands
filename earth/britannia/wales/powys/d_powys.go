package powys

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/powys/builth"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/powys/hereford"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/powys/powys"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/powys/shrewsbury"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PowysDuke interface {
    feud.Duke
    CBuilth比尔斯() 	builth.BuilthCounty
    CHereford赫里福德() 	hereford.HerefordCounty
    CPowys波伊斯() 	powys.PowysCounty
    CShrewsbury什鲁斯伯里() 	shrewsbury.ShrewsburyCounty
}

type 波伊斯PowysDuke struct {
	feud.BaseDuke
	比尔斯Builth 	builth.BuilthCounty
	赫里福德Hereford 	hereford.HerefordCounty
	波伊斯Powys 	powys.PowysCounty
	什鲁斯伯里Shrewsbury 	shrewsbury.ShrewsburyCounty
}

func (d *波伊斯PowysDuke) CBuilth比尔斯() builth.BuilthCounty {
	return d.比尔斯Builth
}
    
func (d *波伊斯PowysDuke) CHereford赫里福德() hereford.HerefordCounty {
	return d.赫里福德Hereford
}
    
func (d *波伊斯PowysDuke) CPowys波伊斯() powys.PowysCounty {
	return d.波伊斯Powys
}
    
func (d *波伊斯PowysDuke) CShrewsbury什鲁斯伯里() shrewsbury.ShrewsburyCounty {
	return d.什鲁斯伯里Shrewsbury
}
    
var DPowys波伊斯 PowysDuke = &波伊斯PowysDuke{}

func init() {
	f := DPowys波伊斯.(*波伊斯PowysDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "powys",
		TitleName: "波伊斯",
		TitleCode: "d_powys",
		Counties:  map[string]feud.County{},
	}

	f.比尔斯Builth = builth.CBuilth比尔斯
	f.比尔斯Builth.SetParent(f)
	
	f.赫里福德Hereford = hereford.CHereford赫里福德
	f.赫里福德Hereford.SetParent(f)
	
	f.波伊斯Powys = powys.CPowys波伊斯
	f.波伊斯Powys.SetParent(f)
	
	f.什鲁斯伯里Shrewsbury = shrewsbury.CShrewsbury什鲁斯伯里
	f.什鲁斯伯里Shrewsbury.SetParent(f)
	
}
