package zavarot

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/zavarot/kolguyev"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/zavarot/soyma"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/zavarot/zavarot"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZavarotDuke interface {
    feud.Duke
    CKolguyev科尔古耶夫() 	kolguyev.KolguyevCounty
    CSoyma因迪加() 	soyma.SoymaCounty
    CZavarot扎瓦罗特() 	zavarot.ZavarotCounty
}

type 扎瓦罗特ZavarotDuke struct {
	feud.BaseDuke
	科尔古耶夫Kolguyev 	kolguyev.KolguyevCounty
	因迪加Soyma 	soyma.SoymaCounty
	扎瓦罗特Zavarot 	zavarot.ZavarotCounty
}

func (d *扎瓦罗特ZavarotDuke) CKolguyev科尔古耶夫() kolguyev.KolguyevCounty {
	return d.科尔古耶夫Kolguyev
}
    
func (d *扎瓦罗特ZavarotDuke) CSoyma因迪加() soyma.SoymaCounty {
	return d.因迪加Soyma
}
    
func (d *扎瓦罗特ZavarotDuke) CZavarot扎瓦罗特() zavarot.ZavarotCounty {
	return d.扎瓦罗特Zavarot
}
    
var DZavarot扎瓦罗特 ZavarotDuke = &扎瓦罗特ZavarotDuke{}

func init() {
	f := DZavarot扎瓦罗特.(*扎瓦罗特ZavarotDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "zavarot",
		TitleName: "扎瓦罗特",
		TitleCode: "d_zavarot",
		Counties:  map[string]feud.County{},
	}

	f.科尔古耶夫Kolguyev = kolguyev.CKolguyev科尔古耶夫
	f.科尔古耶夫Kolguyev.SetParent(f)
	
	f.因迪加Soyma = soyma.CSoyma因迪加
	f.因迪加Soyma.SetParent(f)
	
	f.扎瓦罗特Zavarot = zavarot.CZavarot扎瓦罗特
	f.扎瓦罗特Zavarot.SetParent(f)
	
}
