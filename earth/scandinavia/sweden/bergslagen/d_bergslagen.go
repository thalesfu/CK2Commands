package bergslagen

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/bergslagen/dalarna"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/bergslagen/varmland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/bergslagen/vastmanland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BergslagenDuke interface {
    feud.Duke
    CDalarna耶恩贝拉兰() 	dalarna.DalarnaCounty
    CVarmland韦姆兰() 	varmland.VarmlandCounty
    CVastmanland西曼兰() 	vastmanland.VastmanlandCounty
}

type 贝里斯拉根BergslagenDuke struct {
	feud.BaseDuke
	耶恩贝拉兰Dalarna 	dalarna.DalarnaCounty
	韦姆兰Varmland 	varmland.VarmlandCounty
	西曼兰Vastmanland 	vastmanland.VastmanlandCounty
}

func (d *贝里斯拉根BergslagenDuke) CDalarna耶恩贝拉兰() dalarna.DalarnaCounty {
	return d.耶恩贝拉兰Dalarna
}
    
func (d *贝里斯拉根BergslagenDuke) CVarmland韦姆兰() varmland.VarmlandCounty {
	return d.韦姆兰Varmland
}
    
func (d *贝里斯拉根BergslagenDuke) CVastmanland西曼兰() vastmanland.VastmanlandCounty {
	return d.西曼兰Vastmanland
}
    
var DBergslagen贝里斯拉根 BergslagenDuke = &贝里斯拉根BergslagenDuke{}

func init() {
	f := DBergslagen贝里斯拉根.(*贝里斯拉根BergslagenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bergslagen",
		TitleName: "贝里斯拉根",
		TitleCode: "d_bergslagen",
		Counties:  map[string]feud.County{},
	}

	f.耶恩贝拉兰Dalarna = dalarna.CDalarna耶恩贝拉兰
	f.耶恩贝拉兰Dalarna.SetParent(f)
	
	f.韦姆兰Varmland = varmland.CVarmland韦姆兰
	f.韦姆兰Varmland.SetParent(f)
	
	f.西曼兰Vastmanland = vastmanland.CVastmanland西曼兰
	f.西曼兰Vastmanland.SetParent(f)
	
}
