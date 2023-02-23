package somerset

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/somerset/bath"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/somerset/somerset"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SomersetDuke interface {
	feud.Duke
	CBath巴斯() bath.BathCounty
	CSomerset萨默塞特() somerset.SomersetCounty
}

type 萨默塞特SomersetDuke struct {
	feud.BaseDuke
	巴斯Bath       bath.BathCounty
	萨默塞特Somerset somerset.SomersetCounty
}

func (d *萨默塞特SomersetDuke) CBath巴斯() bath.BathCounty {
	return d.巴斯Bath
}

func (d *萨默塞特SomersetDuke) CSomerset萨默塞特() somerset.SomersetCounty {
	return d.萨默塞特Somerset
}

var DSomerset萨默塞特 SomersetDuke = &萨默塞特SomersetDuke{}

func init() {
	f := DSomerset萨默塞特.(*萨默塞特SomersetDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "somerset",
		TitleName: "萨默塞特",
		TitleCode: "d_somerset",
		Counties:  map[string]feud.County{},
	}

	f.巴斯Bath = bath.CBath巴斯
	f.巴斯Bath.SetParent(f)

	f.萨默塞特Somerset = somerset.CSomerset萨默塞特
	f.萨默塞特Somerset.SetParent(f)

}
