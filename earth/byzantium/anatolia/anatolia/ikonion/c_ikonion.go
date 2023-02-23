package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IkonionCounty interface {
	feud.County
	BAmblada安姆布拉达() feud.Barony
	BGaspadale加斯帕德里() feud.Barony
	BIkonion以哥念() feud.Barony
	BIsauria伊苏里亚() feud.Barony
	BLaranda拉兰达() feud.Barony
	BLisdra里斯德拉() feud.Barony
	BSauatra扫阿特拉() feud.Barony
	BTerpe特尔普() feud.Barony
}

type 以哥念IkonionCounty struct {
	feud.BaseCounty
	安姆布拉达Amblada   feud.Barony
	加斯帕德里Gaspadale feud.Barony
	以哥念Ikonion     feud.Barony
	伊苏里亚Isauria    feud.Barony
	拉兰达Laranda     feud.Barony
	里斯德拉Lisdra     feud.Barony
	扫阿特拉Sauatra    feud.Barony
	特尔普Terpe       feud.Barony
}

func (c *以哥念IkonionCounty) BAmblada安姆布拉达() feud.Barony {
	return c.安姆布拉达Amblada
}

func (c *以哥念IkonionCounty) BGaspadale加斯帕德里() feud.Barony {
	return c.加斯帕德里Gaspadale
}

func (c *以哥念IkonionCounty) BIkonion以哥念() feud.Barony {
	return c.以哥念Ikonion
}

func (c *以哥念IkonionCounty) BIsauria伊苏里亚() feud.Barony {
	return c.伊苏里亚Isauria
}

func (c *以哥念IkonionCounty) BLaranda拉兰达() feud.Barony {
	return c.拉兰达Laranda
}

func (c *以哥念IkonionCounty) BLisdra里斯德拉() feud.Barony {
	return c.里斯德拉Lisdra
}

func (c *以哥念IkonionCounty) BSauatra扫阿特拉() feud.Barony {
	return c.扫阿特拉Sauatra
}

func (c *以哥念IkonionCounty) BTerpe特尔普() feud.Barony {
	return c.特尔普Terpe
}

var CIkonion以哥念 IkonionCounty = &以哥念IkonionCounty{}

func init() {
	f := CIkonion以哥念.(*以哥念IkonionCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "759",
		Title:     "ikonion",
		TitleName: "以哥念",
		TitleCode: "c_ikonion",
		Baronies:  map[string]feud.Barony{},
	}

	f.安姆布拉达Amblada = BAmblada安姆布拉达
	f.安姆布拉达Amblada.SetParent(f)

	f.加斯帕德里Gaspadale = BGaspadale加斯帕德里
	f.加斯帕德里Gaspadale.SetParent(f)

	f.以哥念Ikonion = BIkonion以哥念
	f.以哥念Ikonion.SetParent(f)

	f.伊苏里亚Isauria = BIsauria伊苏里亚
	f.伊苏里亚Isauria.SetParent(f)

	f.拉兰达Laranda = BLaranda拉兰达
	f.拉兰达Laranda.SetParent(f)

	f.里斯德拉Lisdra = BLisdra里斯德拉
	f.里斯德拉Lisdra.SetParent(f)

	f.扫阿特拉Sauatra = BSauatra扫阿特拉
	f.扫阿特拉Sauatra.SetParent(f)

	f.特尔普Terpe = BTerpe特尔普
	f.特尔普Terpe.SetParent(f)

}
