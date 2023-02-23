package haritanaka

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/haritanaka/hisar"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/haritanaka/sarasvati"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/haritanaka/tribandapura"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HaritanakaDuke interface {
	feud.Duke
	CHisar醯娑罗() hisar.HisarCounty
	CSarasvati萨罗萨伐底() sarasvati.SarasvatiCounty
	CTribandapura帝利畔荼城() tribandapura.TribandapuraCounty
}

type 诃利多那迦HaritanakaDuke struct {
	feud.BaseDuke
	醯娑罗Hisar          hisar.HisarCounty
	萨罗萨伐底Sarasvati    sarasvati.SarasvatiCounty
	帝利畔荼城Tribandapura tribandapura.TribandapuraCounty
}

func (d *诃利多那迦HaritanakaDuke) CHisar醯娑罗() hisar.HisarCounty {
	return d.醯娑罗Hisar
}

func (d *诃利多那迦HaritanakaDuke) CSarasvati萨罗萨伐底() sarasvati.SarasvatiCounty {
	return d.萨罗萨伐底Sarasvati
}

func (d *诃利多那迦HaritanakaDuke) CTribandapura帝利畔荼城() tribandapura.TribandapuraCounty {
	return d.帝利畔荼城Tribandapura
}

var DHaritanaka诃利多那迦 HaritanakaDuke = &诃利多那迦HaritanakaDuke{}

func init() {
	f := DHaritanaka诃利多那迦.(*诃利多那迦HaritanakaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "haritanaka",
		TitleName: "诃利多那迦",
		TitleCode: "d_haritanaka",
		Counties:  map[string]feud.County{},
	}

	f.醯娑罗Hisar = hisar.CHisar醯娑罗
	f.醯娑罗Hisar.SetParent(f)

	f.萨罗萨伐底Sarasvati = sarasvati.CSarasvati萨罗萨伐底
	f.萨罗萨伐底Sarasvati.SetParent(f)

	f.帝利畔荼城Tribandapura = tribandapura.CTribandapura帝利畔荼城
	f.帝利畔荼城Tribandapura.SetParent(f)

}
