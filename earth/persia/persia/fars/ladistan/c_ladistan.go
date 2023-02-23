package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LadistanCounty interface {
	feud.County
	BBandarkhamir班达尔哈米尔() feud.Barony
	BBastak巴斯塔克() feud.Barony
	BEvaz埃瓦兹() feud.Barony
	BForg佛拉格() feud.Barony
	BJask贾斯克() feud.Barony
	BKhonj洪季() feud.Barony
	BLad拉德() feud.Barony
	BMorbagh摩尔巴赫() feud.Barony
}

type 达拉布杰德LadistanCounty struct {
	feud.BaseCounty
	班达尔哈米尔Bandarkhamir feud.Barony
	巴斯塔克Bastak         feud.Barony
	埃瓦兹Evaz            feud.Barony
	佛拉格Forg            feud.Barony
	贾斯克Jask            feud.Barony
	洪季Khonj            feud.Barony
	拉德Lad              feud.Barony
	摩尔巴赫Morbagh        feud.Barony
}

func (c *达拉布杰德LadistanCounty) BBandarkhamir班达尔哈米尔() feud.Barony {
	return c.班达尔哈米尔Bandarkhamir
}

func (c *达拉布杰德LadistanCounty) BBastak巴斯塔克() feud.Barony {
	return c.巴斯塔克Bastak
}

func (c *达拉布杰德LadistanCounty) BEvaz埃瓦兹() feud.Barony {
	return c.埃瓦兹Evaz
}

func (c *达拉布杰德LadistanCounty) BForg佛拉格() feud.Barony {
	return c.佛拉格Forg
}

func (c *达拉布杰德LadistanCounty) BJask贾斯克() feud.Barony {
	return c.贾斯克Jask
}

func (c *达拉布杰德LadistanCounty) BKhonj洪季() feud.Barony {
	return c.洪季Khonj
}

func (c *达拉布杰德LadistanCounty) BLad拉德() feud.Barony {
	return c.拉德Lad
}

func (c *达拉布杰德LadistanCounty) BMorbagh摩尔巴赫() feud.Barony {
	return c.摩尔巴赫Morbagh
}

var CLadistan达拉布杰德 LadistanCounty = &达拉布杰德LadistanCounty{}

func init() {
	f := CLadistan达拉布杰德.(*达拉布杰德LadistanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "642",
		Title:     "ladistan",
		TitleName: "达拉布杰德",
		TitleCode: "c_ladistan",
		Baronies:  map[string]feud.Barony{},
	}

	f.班达尔哈米尔Bandarkhamir = BBandarkhamir班达尔哈米尔
	f.班达尔哈米尔Bandarkhamir.SetParent(f)

	f.巴斯塔克Bastak = BBastak巴斯塔克
	f.巴斯塔克Bastak.SetParent(f)

	f.埃瓦兹Evaz = BEvaz埃瓦兹
	f.埃瓦兹Evaz.SetParent(f)

	f.佛拉格Forg = BForg佛拉格
	f.佛拉格Forg.SetParent(f)

	f.贾斯克Jask = BJask贾斯克
	f.贾斯克Jask.SetParent(f)

	f.洪季Khonj = BKhonj洪季
	f.洪季Khonj.SetParent(f)

	f.拉德Lad = BLad拉德
	f.拉德Lad.SetParent(f)

	f.摩尔巴赫Morbagh = BMorbagh摩尔巴赫
	f.摩尔巴赫Morbagh.SetParent(f)

}
