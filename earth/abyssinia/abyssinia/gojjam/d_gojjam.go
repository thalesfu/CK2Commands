package gojjam

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/gojjam/gojjam"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/gojjam/matamma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GojjamDuke interface {
	feud.Duke
	CGojjam戈贾姆() gojjam.GojjamCounty
	CMatamma米提玛() matamma.MatammaCounty
}

type 戈贾姆GojjamDuke struct {
	feud.BaseDuke
	戈贾姆Gojjam  gojjam.GojjamCounty
	米提玛Matamma matamma.MatammaCounty
}

func (d *戈贾姆GojjamDuke) CGojjam戈贾姆() gojjam.GojjamCounty {
	return d.戈贾姆Gojjam
}

func (d *戈贾姆GojjamDuke) CMatamma米提玛() matamma.MatammaCounty {
	return d.米提玛Matamma
}

var DGojjam戈贾姆 GojjamDuke = &戈贾姆GojjamDuke{}

func init() {
	f := DGojjam戈贾姆.(*戈贾姆GojjamDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gojjam",
		TitleName: "戈贾姆",
		TitleCode: "d_gojjam",
		Counties:  map[string]feud.County{},
	}

	f.戈贾姆Gojjam = gojjam.CGojjam戈贾姆
	f.戈贾姆Gojjam.SetParent(f)

	f.米提玛Matamma = matamma.CMatamma米提玛
	f.米提玛Matamma.SetParent(f)

}
