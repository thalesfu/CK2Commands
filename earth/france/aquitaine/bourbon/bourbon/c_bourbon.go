package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BourbonCounty interface {
	feud.County
	BBourbon波旁() feud.Barony
	BLancy朗西() feud.Barony
	BMontlucon蒙吕松() feud.Barony
	BMontpensier蒙庞谢() feud.Barony
	BMoulins穆兰() feud.Barony
	BSouvigny苏维尼() feud.Barony
	BVichy维希() feud.Barony
}

type 波旁BourbonCounty struct {
	feud.BaseCounty
	波旁Bourbon      feud.Barony
	朗西Lancy        feud.Barony
	蒙吕松Montlucon   feud.Barony
	蒙庞谢Montpensier feud.Barony
	穆兰Moulins      feud.Barony
	苏维尼Souvigny    feud.Barony
	维希Vichy        feud.Barony
}

func (c *波旁BourbonCounty) BBourbon波旁() feud.Barony {
	return c.波旁Bourbon
}

func (c *波旁BourbonCounty) BLancy朗西() feud.Barony {
	return c.朗西Lancy
}

func (c *波旁BourbonCounty) BMontlucon蒙吕松() feud.Barony {
	return c.蒙吕松Montlucon
}

func (c *波旁BourbonCounty) BMontpensier蒙庞谢() feud.Barony {
	return c.蒙庞谢Montpensier
}

func (c *波旁BourbonCounty) BMoulins穆兰() feud.Barony {
	return c.穆兰Moulins
}

func (c *波旁BourbonCounty) BSouvigny苏维尼() feud.Barony {
	return c.苏维尼Souvigny
}

func (c *波旁BourbonCounty) BVichy维希() feud.Barony {
	return c.维希Vichy
}

var CBourbon波旁 BourbonCounty = &波旁BourbonCounty{}

func init() {
	f := CBourbon波旁.(*波旁BourbonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "146",
		Title:     "bourbon",
		TitleName: "波旁",
		TitleCode: "c_bourbon",
		Baronies:  map[string]feud.Barony{},
	}

	f.波旁Bourbon = BBourbon波旁
	f.波旁Bourbon.SetParent(f)

	f.朗西Lancy = BLancy朗西
	f.朗西Lancy.SetParent(f)

	f.蒙吕松Montlucon = BMontlucon蒙吕松
	f.蒙吕松Montlucon.SetParent(f)

	f.蒙庞谢Montpensier = BMontpensier蒙庞谢
	f.蒙庞谢Montpensier.SetParent(f)

	f.穆兰Moulins = BMoulins穆兰
	f.穆兰Moulins.SetParent(f)

	f.苏维尼Souvigny = BSouvigny苏维尼
	f.苏维尼Souvigny.SetParent(f)

	f.维希Vichy = BVichy维希
	f.维希Vichy.SetParent(f)

}
