package hlynov

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/hlynov/syrj"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/hlynov/vashka"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/hlynov/velsk"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/hlynov/yezhuga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HlynovDuke interface {
    feud.Duke
    CSyrj西里() 	syrj.SyrjCounty
    CVashka瓦什卡() 	vashka.VashkaCounty
    CVelsk韦利斯克() 	velsk.VelskCounty
    CYezhuga皮涅加() 	yezhuga.YezhugaCounty
}

type 韦利斯克HlynovDuke struct {
	feud.BaseDuke
	西里Syrj 	syrj.SyrjCounty
	瓦什卡Vashka 	vashka.VashkaCounty
	韦利斯克Velsk 	velsk.VelskCounty
	皮涅加Yezhuga 	yezhuga.YezhugaCounty
}

func (d *韦利斯克HlynovDuke) CSyrj西里() syrj.SyrjCounty {
	return d.西里Syrj
}
    
func (d *韦利斯克HlynovDuke) CVashka瓦什卡() vashka.VashkaCounty {
	return d.瓦什卡Vashka
}
    
func (d *韦利斯克HlynovDuke) CVelsk韦利斯克() velsk.VelskCounty {
	return d.韦利斯克Velsk
}
    
func (d *韦利斯克HlynovDuke) CYezhuga皮涅加() yezhuga.YezhugaCounty {
	return d.皮涅加Yezhuga
}
    
var DHlynov韦利斯克 HlynovDuke = &韦利斯克HlynovDuke{}

func init() {
	f := DHlynov韦利斯克.(*韦利斯克HlynovDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "hlynov",
		TitleName: "韦利斯克",
		TitleCode: "d_hlynov",
		Counties:  map[string]feud.County{},
	}

	f.西里Syrj = syrj.CSyrj西里
	f.西里Syrj.SetParent(f)
	
	f.瓦什卡Vashka = vashka.CVashka瓦什卡
	f.瓦什卡Vashka.SetParent(f)
	
	f.韦利斯克Velsk = velsk.CVelsk韦利斯克
	f.韦利斯克Velsk.SetParent(f)
	
	f.皮涅加Yezhuga = yezhuga.CYezhuga皮涅加
	f.皮涅加Yezhuga.SetParent(f)
	
}
