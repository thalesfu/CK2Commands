package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AostaCounty interface {
	feud.County
	BAosta奥斯塔() feud.Barony
	BAostachatillon沙蒂永() feud.Barony
	BArnad阿尔纳德() feud.Barony
	BFenis费尼斯() feud.Barony
}

type 奥斯塔AostaCounty struct {
	feud.BaseCounty
	奥斯塔Aosta          feud.Barony
	沙蒂永Aostachatillon feud.Barony
	阿尔纳德Arnad         feud.Barony
	费尼斯Fenis          feud.Barony
}

func (c *奥斯塔AostaCounty) BAosta奥斯塔() feud.Barony {
	return c.奥斯塔Aosta
}

func (c *奥斯塔AostaCounty) BAostachatillon沙蒂永() feud.Barony {
	return c.沙蒂永Aostachatillon
}

func (c *奥斯塔AostaCounty) BArnad阿尔纳德() feud.Barony {
	return c.阿尔纳德Arnad
}

func (c *奥斯塔AostaCounty) BFenis费尼斯() feud.Barony {
	return c.费尼斯Fenis
}

var CAosta奥斯塔 AostaCounty = &奥斯塔AostaCounty{}

func init() {
	f := CAosta奥斯塔.(*奥斯塔AostaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1610",
		Title:     "aosta",
		TitleName: "奥斯塔",
		TitleCode: "c_aosta",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥斯塔Aosta = BAosta奥斯塔
	f.奥斯塔Aosta.SetParent(f)

	f.沙蒂永Aostachatillon = BAostachatillon沙蒂永
	f.沙蒂永Aostachatillon.SetParent(f)

	f.阿尔纳德Arnad = BArnad阿尔纳德
	f.阿尔纳德Arnad.SetParent(f)

	f.费尼斯Fenis = BFenis费尼斯
	f.费尼斯Fenis.SetParent(f)

}
