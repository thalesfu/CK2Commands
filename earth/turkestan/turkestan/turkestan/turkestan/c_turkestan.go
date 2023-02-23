package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurkestanCounty interface {
	feud.County
	BAkdzulpas阿克_祖尔帕斯() feud.Barony
	BAkespe阿克斯佩() feud.Barony
	BAkshelek阿克舍列克() feud.Barony
	BAralkum阿拉尔库姆() feud.Barony
	BKosskul科斯_库尔() feud.Barony
	BSaksaulskiy萨克绍利斯基() feud.Barony
	BSapak萨帕克() feud.Barony
}

type 阿拉尔TurkestanCounty struct {
	feud.BaseCounty
	阿克_祖尔帕斯Akdzulpas  feud.Barony
	阿克斯佩Akespe        feud.Barony
	阿克舍列克Akshelek     feud.Barony
	阿拉尔库姆Aralkum      feud.Barony
	科斯_库尔Kosskul      feud.Barony
	萨克绍利斯基Saksaulskiy feud.Barony
	萨帕克Sapak          feud.Barony
}

func (c *阿拉尔TurkestanCounty) BAkdzulpas阿克_祖尔帕斯() feud.Barony {
	return c.阿克_祖尔帕斯Akdzulpas
}

func (c *阿拉尔TurkestanCounty) BAkespe阿克斯佩() feud.Barony {
	return c.阿克斯佩Akespe
}

func (c *阿拉尔TurkestanCounty) BAkshelek阿克舍列克() feud.Barony {
	return c.阿克舍列克Akshelek
}

func (c *阿拉尔TurkestanCounty) BAralkum阿拉尔库姆() feud.Barony {
	return c.阿拉尔库姆Aralkum
}

func (c *阿拉尔TurkestanCounty) BKosskul科斯_库尔() feud.Barony {
	return c.科斯_库尔Kosskul
}

func (c *阿拉尔TurkestanCounty) BSaksaulskiy萨克绍利斯基() feud.Barony {
	return c.萨克绍利斯基Saksaulskiy
}

func (c *阿拉尔TurkestanCounty) BSapak萨帕克() feud.Barony {
	return c.萨帕克Sapak
}

var CTurkestan阿拉尔 TurkestanCounty = &阿拉尔TurkestanCounty{}

func init() {
	f := CTurkestan阿拉尔.(*阿拉尔TurkestanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "622",
		Title:     "turkestan",
		TitleName: "阿拉尔",
		TitleCode: "c_turkestan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克_祖尔帕斯Akdzulpas = BAkdzulpas阿克_祖尔帕斯
	f.阿克_祖尔帕斯Akdzulpas.SetParent(f)

	f.阿克斯佩Akespe = BAkespe阿克斯佩
	f.阿克斯佩Akespe.SetParent(f)

	f.阿克舍列克Akshelek = BAkshelek阿克舍列克
	f.阿克舍列克Akshelek.SetParent(f)

	f.阿拉尔库姆Aralkum = BAralkum阿拉尔库姆
	f.阿拉尔库姆Aralkum.SetParent(f)

	f.科斯_库尔Kosskul = BKosskul科斯_库尔
	f.科斯_库尔Kosskul.SetParent(f)

	f.萨克绍利斯基Saksaulskiy = BSaksaulskiy萨克绍利斯基
	f.萨克绍利斯基Saksaulskiy.SetParent(f)

	f.萨帕克Sapak = BSapak萨帕克
	f.萨帕克Sapak.SetParent(f)

}
