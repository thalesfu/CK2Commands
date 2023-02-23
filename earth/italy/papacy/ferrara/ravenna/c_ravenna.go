package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RavennaCounty interface {
	feud.County
	BAlfonsine阿方西内() feud.Barony
	BCervia切尔维亚() feud.Barony
	BCesena切塞纳() feud.Barony
	BCesenatico切塞纳蒂科() feud.Barony
	BGambettola甘贝托拉() feud.Barony
	BRavenna拉文纳() feud.Barony
	BRussi鲁西() feud.Barony
}

type 拉文纳RavennaCounty struct {
	feud.BaseCounty
	阿方西内Alfonsine   feud.Barony
	切尔维亚Cervia      feud.Barony
	切塞纳Cesena       feud.Barony
	切塞纳蒂科Cesenatico feud.Barony
	甘贝托拉Gambettola  feud.Barony
	拉文纳Ravenna      feud.Barony
	鲁西Russi         feud.Barony
}

func (c *拉文纳RavennaCounty) BAlfonsine阿方西内() feud.Barony {
	return c.阿方西内Alfonsine
}

func (c *拉文纳RavennaCounty) BCervia切尔维亚() feud.Barony {
	return c.切尔维亚Cervia
}

func (c *拉文纳RavennaCounty) BCesena切塞纳() feud.Barony {
	return c.切塞纳Cesena
}

func (c *拉文纳RavennaCounty) BCesenatico切塞纳蒂科() feud.Barony {
	return c.切塞纳蒂科Cesenatico
}

func (c *拉文纳RavennaCounty) BGambettola甘贝托拉() feud.Barony {
	return c.甘贝托拉Gambettola
}

func (c *拉文纳RavennaCounty) BRavenna拉文纳() feud.Barony {
	return c.拉文纳Ravenna
}

func (c *拉文纳RavennaCounty) BRussi鲁西() feud.Barony {
	return c.鲁西Russi
}

var CRavenna拉文纳 RavennaCounty = &拉文纳RavennaCounty{}

func init() {
	f := CRavenna拉文纳.(*拉文纳RavennaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "351",
		Title:     "ravenna",
		TitleName: "拉文纳",
		TitleCode: "c_ravenna",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿方西内Alfonsine = BAlfonsine阿方西内
	f.阿方西内Alfonsine.SetParent(f)

	f.切尔维亚Cervia = BCervia切尔维亚
	f.切尔维亚Cervia.SetParent(f)

	f.切塞纳Cesena = BCesena切塞纳
	f.切塞纳Cesena.SetParent(f)

	f.切塞纳蒂科Cesenatico = BCesenatico切塞纳蒂科
	f.切塞纳蒂科Cesenatico.SetParent(f)

	f.甘贝托拉Gambettola = BGambettola甘贝托拉
	f.甘贝托拉Gambettola.SetParent(f)

	f.拉文纳Ravenna = BRavenna拉文纳
	f.拉文纳Ravenna.SetParent(f)

	f.鲁西Russi = BRussi鲁西
	f.鲁西Russi.SetParent(f)

}
