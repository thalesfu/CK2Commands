package tunis

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tunis/bizerte"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tunis/kairwan"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tunis/mahdia"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tunis/medjerda"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tunis/tunis"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TunisDuke interface {
	feud.Duke
	CBizerte比塞大() bizerte.BizerteCounty
	CKairwan凯鲁万() kairwan.KairwanCounty
	CMahdia马赫迪耶() mahdia.MahdiaCounty
	CMedjerda迈杰尔达() medjerda.MedjerdaCounty
	CTunis突尼斯() tunis.TunisCounty
}

type 突尼斯TunisDuke struct {
	feud.BaseDuke
	比塞大Bizerte   bizerte.BizerteCounty
	凯鲁万Kairwan   kairwan.KairwanCounty
	马赫迪耶Mahdia   mahdia.MahdiaCounty
	迈杰尔达Medjerda medjerda.MedjerdaCounty
	突尼斯Tunis     tunis.TunisCounty
}

func (d *突尼斯TunisDuke) CBizerte比塞大() bizerte.BizerteCounty {
	return d.比塞大Bizerte
}

func (d *突尼斯TunisDuke) CKairwan凯鲁万() kairwan.KairwanCounty {
	return d.凯鲁万Kairwan
}

func (d *突尼斯TunisDuke) CMahdia马赫迪耶() mahdia.MahdiaCounty {
	return d.马赫迪耶Mahdia
}

func (d *突尼斯TunisDuke) CMedjerda迈杰尔达() medjerda.MedjerdaCounty {
	return d.迈杰尔达Medjerda
}

func (d *突尼斯TunisDuke) CTunis突尼斯() tunis.TunisCounty {
	return d.突尼斯Tunis
}

var DTunis突尼斯 TunisDuke = &突尼斯TunisDuke{}

func init() {
	f := DTunis突尼斯.(*突尼斯TunisDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tunis",
		TitleName: "突尼斯",
		TitleCode: "d_tunis",
		Counties:  map[string]feud.County{},
	}

	f.比塞大Bizerte = bizerte.CBizerte比塞大
	f.比塞大Bizerte.SetParent(f)

	f.凯鲁万Kairwan = kairwan.CKairwan凯鲁万
	f.凯鲁万Kairwan.SetParent(f)

	f.马赫迪耶Mahdia = mahdia.CMahdia马赫迪耶
	f.马赫迪耶Mahdia.SetParent(f)

	f.迈杰尔达Medjerda = medjerda.CMedjerda迈杰尔达
	f.迈杰尔达Medjerda.SetParent(f)

	f.突尼斯Tunis = tunis.CTunis突尼斯
	f.突尼斯Tunis.SetParent(f)

}
