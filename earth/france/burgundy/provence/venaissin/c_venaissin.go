package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VenaissinCounty interface {
	feud.County
	BAvignon阿维尼翁() feud.Barony
	BCarpentras卡庞特拉() feud.Barony
	BCavaillon卡瓦永() feud.Barony
	BChateauneufdupape教宗新堡() feud.Barony
	BMondragon蒙德拉贡() feud.Barony
	BOrange奥朗热() feud.Barony
	BStpaul圣保罗() feud.Barony
	BVenasque韦纳斯克() feud.Barony
}

type 韦奈桑VenaissinCounty struct {
	feud.BaseCounty
	阿维尼翁Avignon           feud.Barony
	卡庞特拉Carpentras        feud.Barony
	卡瓦永Cavaillon          feud.Barony
	教宗新堡Chateauneufdupape feud.Barony
	蒙德拉贡Mondragon         feud.Barony
	奥朗热Orange             feud.Barony
	圣保罗Stpaul             feud.Barony
	韦纳斯克Venasque          feud.Barony
}

func (c *韦奈桑VenaissinCounty) BAvignon阿维尼翁() feud.Barony {
	return c.阿维尼翁Avignon
}

func (c *韦奈桑VenaissinCounty) BCarpentras卡庞特拉() feud.Barony {
	return c.卡庞特拉Carpentras
}

func (c *韦奈桑VenaissinCounty) BCavaillon卡瓦永() feud.Barony {
	return c.卡瓦永Cavaillon
}

func (c *韦奈桑VenaissinCounty) BChateauneufdupape教宗新堡() feud.Barony {
	return c.教宗新堡Chateauneufdupape
}

func (c *韦奈桑VenaissinCounty) BMondragon蒙德拉贡() feud.Barony {
	return c.蒙德拉贡Mondragon
}

func (c *韦奈桑VenaissinCounty) BOrange奥朗热() feud.Barony {
	return c.奥朗热Orange
}

func (c *韦奈桑VenaissinCounty) BStpaul圣保罗() feud.Barony {
	return c.圣保罗Stpaul
}

func (c *韦奈桑VenaissinCounty) BVenasque韦纳斯克() feud.Barony {
	return c.韦纳斯克Venasque
}

var CVenaissin韦奈桑 VenaissinCounty = &韦奈桑VenaissinCounty{}

func init() {
	f := CVenaissin韦奈桑.(*韦奈桑VenaissinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "222",
		Title:     "venaissin",
		TitleName: "韦奈桑",
		TitleCode: "c_venaissin",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿维尼翁Avignon = BAvignon阿维尼翁
	f.阿维尼翁Avignon.SetParent(f)

	f.卡庞特拉Carpentras = BCarpentras卡庞特拉
	f.卡庞特拉Carpentras.SetParent(f)

	f.卡瓦永Cavaillon = BCavaillon卡瓦永
	f.卡瓦永Cavaillon.SetParent(f)

	f.教宗新堡Chateauneufdupape = BChateauneufdupape教宗新堡
	f.教宗新堡Chateauneufdupape.SetParent(f)

	f.蒙德拉贡Mondragon = BMondragon蒙德拉贡
	f.蒙德拉贡Mondragon.SetParent(f)

	f.奥朗热Orange = BOrange奥朗热
	f.奥朗热Orange.SetParent(f)

	f.圣保罗Stpaul = BStpaul圣保罗
	f.圣保罗Stpaul.SetParent(f)

	f.韦纳斯克Venasque = BVenasque韦纳斯克
	f.韦纳斯克Venasque.SetParent(f)

}
