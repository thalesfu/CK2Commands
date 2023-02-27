package qinghai

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/qinghai/delingha"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/qinghai/dulan"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/qinghai/fuqi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QinghaiDuke interface {
    feud.Duke
    CDelingha德令哈() 	delingha.DelinghaCounty
    CDulan都兰() 	dulan.DulanCounty
    CFuqi伏俟() 	fuqi.FuqiCounty
}

type 青海QinghaiDuke struct {
	feud.BaseDuke
	德令哈Delingha 	delingha.DelinghaCounty
	都兰Dulan 	dulan.DulanCounty
	伏俟Fuqi 	fuqi.FuqiCounty
}

func (d *青海QinghaiDuke) CDelingha德令哈() delingha.DelinghaCounty {
	return d.德令哈Delingha
}
    
func (d *青海QinghaiDuke) CDulan都兰() dulan.DulanCounty {
	return d.都兰Dulan
}
    
func (d *青海QinghaiDuke) CFuqi伏俟() fuqi.FuqiCounty {
	return d.伏俟Fuqi
}
    
var DQinghai青海 QinghaiDuke = &青海QinghaiDuke{}

func init() {
	f := DQinghai青海.(*青海QinghaiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "qinghai",
		TitleName: "青海",
		TitleCode: "d_qinghai",
		Counties:  map[string]feud.County{},
	}

	f.德令哈Delingha = delingha.CDelingha德令哈
	f.德令哈Delingha.SetParent(f)
	
	f.都兰Dulan = dulan.CDulan都兰
	f.都兰Dulan.SetParent(f)
	
	f.伏俟Fuqi = fuqi.CFuqi伏俟
	f.伏俟Fuqi.SetParent(f)
	
}
