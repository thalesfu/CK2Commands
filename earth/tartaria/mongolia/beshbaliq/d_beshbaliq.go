package beshbaliq

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/beshbaliq/altay"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/beshbaliq/beshbaliq"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/beshbaliq/dunkheger"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeshbaliqDuke interface {
	feud.Duke
	CAltay兀泷古() altay.AltayCounty
	CBeshbaliq别失八里() beshbaliq.BeshbaliqCounty
	CDunkheger冬和格尔() dunkheger.DunkhegerCounty
}

type 别失八里BeshbaliqDuke struct {
	feud.BaseDuke
	兀泷古Altay      altay.AltayCounty
	别失八里Beshbaliq beshbaliq.BeshbaliqCounty
	冬和格尔Dunkheger dunkheger.DunkhegerCounty
}

func (d *别失八里BeshbaliqDuke) CAltay兀泷古() altay.AltayCounty {
	return d.兀泷古Altay
}

func (d *别失八里BeshbaliqDuke) CBeshbaliq别失八里() beshbaliq.BeshbaliqCounty {
	return d.别失八里Beshbaliq
}

func (d *别失八里BeshbaliqDuke) CDunkheger冬和格尔() dunkheger.DunkhegerCounty {
	return d.冬和格尔Dunkheger
}

var DBeshbaliq别失八里 BeshbaliqDuke = &别失八里BeshbaliqDuke{}

func init() {
	f := DBeshbaliq别失八里.(*别失八里BeshbaliqDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "beshbaliq",
		TitleName: "别失八里",
		TitleCode: "d_beshbaliq",
		Counties:  map[string]feud.County{},
	}

	f.兀泷古Altay = altay.CAltay兀泷古
	f.兀泷古Altay.SetParent(f)

	f.别失八里Beshbaliq = beshbaliq.CBeshbaliq别失八里
	f.别失八里Beshbaliq.SetParent(f)

	f.冬和格尔Dunkheger = dunkheger.CDunkheger冬和格尔
	f.冬和格尔Dunkheger.SetParent(f)

}
