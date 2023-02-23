package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaliCounty interface {
	feud.County
	BBaguineda巴吉奈达() feud.Barony
	BBougouni布古尼() feud.Barony
	BDialakoro迪亚拉科罗() feud.Barony
	BKamaro卡马罗() feud.Barony
	BNiani尼亚尼() feud.Barony
	BSafo萨佛() feud.Barony
	BSelegon塞莱贡() feud.Barony
	BSoussa苏萨() feud.Barony
	BTabou塔布() feud.Barony
}

type 尼亚尼MaliCounty struct {
	feud.BaseCounty
	巴吉奈达Baguineda  feud.Barony
	布古尼Bougouni    feud.Barony
	迪亚拉科罗Dialakoro feud.Barony
	卡马罗Kamaro      feud.Barony
	尼亚尼Niani       feud.Barony
	萨佛Safo         feud.Barony
	塞莱贡Selegon     feud.Barony
	苏萨Soussa       feud.Barony
	塔布Tabou        feud.Barony
}

func (c *尼亚尼MaliCounty) BBaguineda巴吉奈达() feud.Barony {
	return c.巴吉奈达Baguineda
}

func (c *尼亚尼MaliCounty) BBougouni布古尼() feud.Barony {
	return c.布古尼Bougouni
}

func (c *尼亚尼MaliCounty) BDialakoro迪亚拉科罗() feud.Barony {
	return c.迪亚拉科罗Dialakoro
}

func (c *尼亚尼MaliCounty) BKamaro卡马罗() feud.Barony {
	return c.卡马罗Kamaro
}

func (c *尼亚尼MaliCounty) BNiani尼亚尼() feud.Barony {
	return c.尼亚尼Niani
}

func (c *尼亚尼MaliCounty) BSafo萨佛() feud.Barony {
	return c.萨佛Safo
}

func (c *尼亚尼MaliCounty) BSelegon塞莱贡() feud.Barony {
	return c.塞莱贡Selegon
}

func (c *尼亚尼MaliCounty) BSoussa苏萨() feud.Barony {
	return c.苏萨Soussa
}

func (c *尼亚尼MaliCounty) BTabou塔布() feud.Barony {
	return c.塔布Tabou
}

var CMali尼亚尼 MaliCounty = &尼亚尼MaliCounty{}

func init() {
	f := CMali尼亚尼.(*尼亚尼MaliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "925",
		Title:     "mali",
		TitleName: "尼亚尼",
		TitleCode: "c_mali",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴吉奈达Baguineda = BBaguineda巴吉奈达
	f.巴吉奈达Baguineda.SetParent(f)

	f.布古尼Bougouni = BBougouni布古尼
	f.布古尼Bougouni.SetParent(f)

	f.迪亚拉科罗Dialakoro = BDialakoro迪亚拉科罗
	f.迪亚拉科罗Dialakoro.SetParent(f)

	f.卡马罗Kamaro = BKamaro卡马罗
	f.卡马罗Kamaro.SetParent(f)

	f.尼亚尼Niani = BNiani尼亚尼
	f.尼亚尼Niani.SetParent(f)

	f.萨佛Safo = BSafo萨佛
	f.萨佛Safo.SetParent(f)

	f.塞莱贡Selegon = BSelegon塞莱贡
	f.塞莱贡Selegon.SetParent(f)

	f.苏萨Soussa = BSoussa苏萨
	f.苏萨Soussa.SetParent(f)

	f.塔布Tabou = BTabou塔布
	f.塔布Tabou.SetParent(f)

}
