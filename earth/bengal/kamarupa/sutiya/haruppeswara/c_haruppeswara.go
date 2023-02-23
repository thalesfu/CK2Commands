package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HaruppeswaraCounty interface {
	feud.County
	BCharaideo查莱碉() feud.Barony
	BDibarumukh提婆卢目佉() feud.Barony
	BHaruppeswara诃卢毗湿伐罗() feud.Barony
	BKunnatur贡那兜() feud.Barony
	BKuppar鸠波() feud.Barony
	BMalabar麻啰拔() feud.Barony
	BNumaligarh奴摩梨堡() feud.Barony
}

type 诃卢毗湿伐罗HaruppeswaraCounty struct {
	feud.BaseCounty
	查莱碉Charaideo       feud.Barony
	提婆卢目佉Dibarumukh    feud.Barony
	诃卢毗湿伐罗Haruppeswara feud.Barony
	贡那兜Kunnatur        feud.Barony
	鸠波Kuppar           feud.Barony
	麻啰拔Malabar         feud.Barony
	奴摩梨堡Numaligarh     feud.Barony
}

func (c *诃卢毗湿伐罗HaruppeswaraCounty) BCharaideo查莱碉() feud.Barony {
	return c.查莱碉Charaideo
}

func (c *诃卢毗湿伐罗HaruppeswaraCounty) BDibarumukh提婆卢目佉() feud.Barony {
	return c.提婆卢目佉Dibarumukh
}

func (c *诃卢毗湿伐罗HaruppeswaraCounty) BHaruppeswara诃卢毗湿伐罗() feud.Barony {
	return c.诃卢毗湿伐罗Haruppeswara
}

func (c *诃卢毗湿伐罗HaruppeswaraCounty) BKunnatur贡那兜() feud.Barony {
	return c.贡那兜Kunnatur
}

func (c *诃卢毗湿伐罗HaruppeswaraCounty) BKuppar鸠波() feud.Barony {
	return c.鸠波Kuppar
}

func (c *诃卢毗湿伐罗HaruppeswaraCounty) BMalabar麻啰拔() feud.Barony {
	return c.麻啰拔Malabar
}

func (c *诃卢毗湿伐罗HaruppeswaraCounty) BNumaligarh奴摩梨堡() feud.Barony {
	return c.奴摩梨堡Numaligarh
}

var CHaruppeswara诃卢毗湿伐罗 HaruppeswaraCounty = &诃卢毗湿伐罗HaruppeswaraCounty{}

func init() {
	f := CHaruppeswara诃卢毗湿伐罗.(*诃卢毗湿伐罗HaruppeswaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1418",
		Title:     "haruppeswara",
		TitleName: "诃卢毗湿伐罗",
		TitleCode: "c_haruppeswara",
		Baronies:  map[string]feud.Barony{},
	}

	f.查莱碉Charaideo = BCharaideo查莱碉
	f.查莱碉Charaideo.SetParent(f)

	f.提婆卢目佉Dibarumukh = BDibarumukh提婆卢目佉
	f.提婆卢目佉Dibarumukh.SetParent(f)

	f.诃卢毗湿伐罗Haruppeswara = BHaruppeswara诃卢毗湿伐罗
	f.诃卢毗湿伐罗Haruppeswara.SetParent(f)

	f.贡那兜Kunnatur = BKunnatur贡那兜
	f.贡那兜Kunnatur.SetParent(f)

	f.鸠波Kuppar = BKuppar鸠波
	f.鸠波Kuppar.SetParent(f)

	f.麻啰拔Malabar = BMalabar麻啰拔
	f.麻啰拔Malabar.SetParent(f)

	f.奴摩梨堡Numaligarh = BNumaligarh奴摩梨堡
	f.奴摩梨堡Numaligarh.SetParent(f)

}
