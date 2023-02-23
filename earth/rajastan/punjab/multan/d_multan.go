package multan

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/multan/kafirkot"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/multan/karor"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/multan/karur"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/multan/multan"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/multan/uch"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/multan/wana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MultanDuke interface {
	feud.Duke
	CKafirkot卡菲尔科特() kafirkot.KafirkotCounty
	CKaror迦卢罗() karor.KarorCounty
	CKarur伽卢罗() karur.KarurCounty
	CMultan木尔坦() multan.MultanCounty
	CUch邬脂() uch.UchCounty
	CWana弯那() wana.WanaCounty
}

type 木尔坦MultanDuke struct {
	feud.BaseDuke
	卡菲尔科特Kafirkot kafirkot.KafirkotCounty
	迦卢罗Karor      karor.KarorCounty
	伽卢罗Karur      karur.KarurCounty
	木尔坦Multan     multan.MultanCounty
	邬脂Uch         uch.UchCounty
	弯那Wana        wana.WanaCounty
}

func (d *木尔坦MultanDuke) CKafirkot卡菲尔科特() kafirkot.KafirkotCounty {
	return d.卡菲尔科特Kafirkot
}

func (d *木尔坦MultanDuke) CKaror迦卢罗() karor.KarorCounty {
	return d.迦卢罗Karor
}

func (d *木尔坦MultanDuke) CKarur伽卢罗() karur.KarurCounty {
	return d.伽卢罗Karur
}

func (d *木尔坦MultanDuke) CMultan木尔坦() multan.MultanCounty {
	return d.木尔坦Multan
}

func (d *木尔坦MultanDuke) CUch邬脂() uch.UchCounty {
	return d.邬脂Uch
}

func (d *木尔坦MultanDuke) CWana弯那() wana.WanaCounty {
	return d.弯那Wana
}

var DMultan木尔坦 MultanDuke = &木尔坦MultanDuke{}

func init() {
	f := DMultan木尔坦.(*木尔坦MultanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "multan",
		TitleName: "木尔坦",
		TitleCode: "d_multan",
		Counties:  map[string]feud.County{},
	}

	f.卡菲尔科特Kafirkot = kafirkot.CKafirkot卡菲尔科特
	f.卡菲尔科特Kafirkot.SetParent(f)

	f.迦卢罗Karor = karor.CKaror迦卢罗
	f.迦卢罗Karor.SetParent(f)

	f.伽卢罗Karur = karur.CKarur伽卢罗
	f.伽卢罗Karur.SetParent(f)

	f.木尔坦Multan = multan.CMultan木尔坦
	f.木尔坦Multan.SetParent(f)

	f.邬脂Uch = uch.CUch邬脂
	f.邬脂Uch.SetParent(f)

	f.弯那Wana = wana.CWana弯那
	f.弯那Wana.SetParent(f)

}
