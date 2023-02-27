package ohrid

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/ohrid/kastoria"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/ohrid/lyncestis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/ohrid/ochrid"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OhridDuke interface {
    feud.Duke
    CKastoria卡斯托里亚() 	kastoria.KastoriaCounty
    CLyncestis比托拉() 	lyncestis.LyncestisCounty
    COchrid奥赫里德() 	ochrid.OchridCounty
}

type 奥赫里德OhridDuke struct {
	feud.BaseDuke
	卡斯托里亚Kastoria 	kastoria.KastoriaCounty
	比托拉Lyncestis 	lyncestis.LyncestisCounty
	奥赫里德Ochrid 	ochrid.OchridCounty
}

func (d *奥赫里德OhridDuke) CKastoria卡斯托里亚() kastoria.KastoriaCounty {
	return d.卡斯托里亚Kastoria
}
    
func (d *奥赫里德OhridDuke) CLyncestis比托拉() lyncestis.LyncestisCounty {
	return d.比托拉Lyncestis
}
    
func (d *奥赫里德OhridDuke) COchrid奥赫里德() ochrid.OchridCounty {
	return d.奥赫里德Ochrid
}
    
var DOhrid奥赫里德 OhridDuke = &奥赫里德OhridDuke{}

func init() {
	f := DOhrid奥赫里德.(*奥赫里德OhridDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ohrid",
		TitleName: "奥赫里德",
		TitleCode: "d_ohrid",
		Counties:  map[string]feud.County{},
	}

	f.卡斯托里亚Kastoria = kastoria.CKastoria卡斯托里亚
	f.卡斯托里亚Kastoria.SetParent(f)
	
	f.比托拉Lyncestis = lyncestis.CLyncestis比托拉
	f.比托拉Lyncestis.SetParent(f)
	
	f.奥赫里德Ochrid = ochrid.COchrid奥赫里德
	f.奥赫里德Ochrid.SetParent(f)
	
}
