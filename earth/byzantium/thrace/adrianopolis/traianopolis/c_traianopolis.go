package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TraianopolisCounty interface {
	feud.County
	BAinos爱诺斯() feud.Barony
	BCypsela库普塞拉() feud.Barony
	BEvros埃夫罗斯() feud.Barony
	BFeres费雷斯() feud.Barony
	BSamothrace萨莫色雷斯() feud.Barony
	BTraianopolis图拉真波利斯() feud.Barony
	BTychero蒂海罗() feud.Barony
}

type 图拉真波利斯TraianopolisCounty struct {
	feud.BaseCounty
	爱诺斯Ainos           feud.Barony
	库普塞拉Cypsela        feud.Barony
	埃夫罗斯Evros          feud.Barony
	费雷斯Feres           feud.Barony
	萨莫色雷斯Samothrace    feud.Barony
	图拉真波利斯Traianopolis feud.Barony
	蒂海罗Tychero         feud.Barony
}

func (c *图拉真波利斯TraianopolisCounty) BAinos爱诺斯() feud.Barony {
	return c.爱诺斯Ainos
}

func (c *图拉真波利斯TraianopolisCounty) BCypsela库普塞拉() feud.Barony {
	return c.库普塞拉Cypsela
}

func (c *图拉真波利斯TraianopolisCounty) BEvros埃夫罗斯() feud.Barony {
	return c.埃夫罗斯Evros
}

func (c *图拉真波利斯TraianopolisCounty) BFeres费雷斯() feud.Barony {
	return c.费雷斯Feres
}

func (c *图拉真波利斯TraianopolisCounty) BSamothrace萨莫色雷斯() feud.Barony {
	return c.萨莫色雷斯Samothrace
}

func (c *图拉真波利斯TraianopolisCounty) BTraianopolis图拉真波利斯() feud.Barony {
	return c.图拉真波利斯Traianopolis
}

func (c *图拉真波利斯TraianopolisCounty) BTychero蒂海罗() feud.Barony {
	return c.蒂海罗Tychero
}

var CTraianopolis图拉真波利斯 TraianopolisCounty = &图拉真波利斯TraianopolisCounty{}

func init() {
	f := CTraianopolis图拉真波利斯.(*图拉真波利斯TraianopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1884",
		Title:     "traianopolis",
		TitleName: "图拉真波利斯",
		TitleCode: "c_traianopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.爱诺斯Ainos = BAinos爱诺斯
	f.爱诺斯Ainos.SetParent(f)

	f.库普塞拉Cypsela = BCypsela库普塞拉
	f.库普塞拉Cypsela.SetParent(f)

	f.埃夫罗斯Evros = BEvros埃夫罗斯
	f.埃夫罗斯Evros.SetParent(f)

	f.费雷斯Feres = BFeres费雷斯
	f.费雷斯Feres.SetParent(f)

	f.萨莫色雷斯Samothrace = BSamothrace萨莫色雷斯
	f.萨莫色雷斯Samothrace.SetParent(f)

	f.图拉真波利斯Traianopolis = BTraianopolis图拉真波利斯
	f.图拉真波利斯Traianopolis.SetParent(f)

	f.蒂海罗Tychero = BTychero蒂海罗
	f.蒂海罗Tychero.SetParent(f)

}
