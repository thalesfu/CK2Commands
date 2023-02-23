package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AzerbaijanCounty interface {
	feud.County
	BBastam巴斯塔姆() feud.Barony
	BChaldoran查尔多兰() feud.Barony
	BDeglane德格拉尼() feud.Barony
	BKalaberd卡拉贝德() feud.Barony
	BKapalak卡拉克卡() feud.Barony
	BMaku马库() feud.Barony
	BTakhtesuleiman塔赫特苏莱曼() feud.Barony
	BUrmiah乌尔米耶() feud.Barony
}

type 阿塞拜疆AzerbaijanCounty struct {
	feud.BaseCounty
	巴斯塔姆Bastam           feud.Barony
	查尔多兰Chaldoran        feud.Barony
	德格拉尼Deglane          feud.Barony
	卡拉贝德Kalaberd         feud.Barony
	卡拉克卡Kapalak          feud.Barony
	马库Maku               feud.Barony
	塔赫特苏莱曼Takhtesuleiman feud.Barony
	乌尔米耶Urmiah           feud.Barony
}

func (c *阿塞拜疆AzerbaijanCounty) BBastam巴斯塔姆() feud.Barony {
	return c.巴斯塔姆Bastam
}

func (c *阿塞拜疆AzerbaijanCounty) BChaldoran查尔多兰() feud.Barony {
	return c.查尔多兰Chaldoran
}

func (c *阿塞拜疆AzerbaijanCounty) BDeglane德格拉尼() feud.Barony {
	return c.德格拉尼Deglane
}

func (c *阿塞拜疆AzerbaijanCounty) BKalaberd卡拉贝德() feud.Barony {
	return c.卡拉贝德Kalaberd
}

func (c *阿塞拜疆AzerbaijanCounty) BKapalak卡拉克卡() feud.Barony {
	return c.卡拉克卡Kapalak
}

func (c *阿塞拜疆AzerbaijanCounty) BMaku马库() feud.Barony {
	return c.马库Maku
}

func (c *阿塞拜疆AzerbaijanCounty) BTakhtesuleiman塔赫特苏莱曼() feud.Barony {
	return c.塔赫特苏莱曼Takhtesuleiman
}

func (c *阿塞拜疆AzerbaijanCounty) BUrmiah乌尔米耶() feud.Barony {
	return c.乌尔米耶Urmiah
}

var CAzerbaijan阿塞拜疆 AzerbaijanCounty = &阿塞拜疆AzerbaijanCounty{}

func init() {
	f := CAzerbaijan阿塞拜疆.(*阿塞拜疆AzerbaijanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "670",
		Title:     "azerbaijan",
		TitleName: "阿塞拜疆",
		TitleCode: "c_azerbaijan",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴斯塔姆Bastam = BBastam巴斯塔姆
	f.巴斯塔姆Bastam.SetParent(f)

	f.查尔多兰Chaldoran = BChaldoran查尔多兰
	f.查尔多兰Chaldoran.SetParent(f)

	f.德格拉尼Deglane = BDeglane德格拉尼
	f.德格拉尼Deglane.SetParent(f)

	f.卡拉贝德Kalaberd = BKalaberd卡拉贝德
	f.卡拉贝德Kalaberd.SetParent(f)

	f.卡拉克卡Kapalak = BKapalak卡拉克卡
	f.卡拉克卡Kapalak.SetParent(f)

	f.马库Maku = BMaku马库
	f.马库Maku.SetParent(f)

	f.塔赫特苏莱曼Takhtesuleiman = BTakhtesuleiman塔赫特苏莱曼
	f.塔赫特苏莱曼Takhtesuleiman.SetParent(f)

	f.乌尔米耶Urmiah = BUrmiah乌尔米耶
	f.乌尔米耶Urmiah.SetParent(f)

}
