package sutiya

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/sutiya/haruppeswara"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/sutiya/kundina"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/sutiya/lhoyu"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/sutiya/monyul"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SutiyaDuke interface {
	feud.Duke
	CHaruppeswara诃卢毗湿伐罗() haruppeswara.HaruppeswaraCounty
	CKundina军稚那() kundina.KundinaCounty
	CLhoyu珞隅() lhoyu.LhoyuCounty
	CMonyul门隅() monyul.MonyulCounty
}

type 苏底耶SutiyaDuke struct {
	feud.BaseDuke
	诃卢毗湿伐罗Haruppeswara haruppeswara.HaruppeswaraCounty
	军稚那Kundina         kundina.KundinaCounty
	珞隅Lhoyu            lhoyu.LhoyuCounty
	门隅Monyul           monyul.MonyulCounty
}

func (d *苏底耶SutiyaDuke) CHaruppeswara诃卢毗湿伐罗() haruppeswara.HaruppeswaraCounty {
	return d.诃卢毗湿伐罗Haruppeswara
}

func (d *苏底耶SutiyaDuke) CKundina军稚那() kundina.KundinaCounty {
	return d.军稚那Kundina
}

func (d *苏底耶SutiyaDuke) CLhoyu珞隅() lhoyu.LhoyuCounty {
	return d.珞隅Lhoyu
}

func (d *苏底耶SutiyaDuke) CMonyul门隅() monyul.MonyulCounty {
	return d.门隅Monyul
}

var DSutiya苏底耶 SutiyaDuke = &苏底耶SutiyaDuke{}

func init() {
	f := DSutiya苏底耶.(*苏底耶SutiyaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sutiya",
		TitleName: "苏底耶",
		TitleCode: "d_sutiya",
		Counties:  map[string]feud.County{},
	}

	f.诃卢毗湿伐罗Haruppeswara = haruppeswara.CHaruppeswara诃卢毗湿伐罗
	f.诃卢毗湿伐罗Haruppeswara.SetParent(f)

	f.军稚那Kundina = kundina.CKundina军稚那
	f.军稚那Kundina.SetParent(f)

	f.珞隅Lhoyu = lhoyu.CLhoyu珞隅
	f.珞隅Lhoyu.SetParent(f)

	f.门隅Monyul = monyul.CMonyul门隅
	f.门隅Monyul.SetParent(f)

}
