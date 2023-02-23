package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MakranCounty interface {
	feud.County
	BChhajian查绛() feud.Barony
	BKannazbun坎纳兹本() feud.Barony
	BKiz基兹() feud.Barony
	BMan曼恩() feud.Barony
	BMughal莫卧儿() feud.Barony
	BOrmara奥尔马拉() feud.Barony
}

type 弥兰MakranCounty struct {
	feud.BaseCounty
	查绛Chhajian    feud.Barony
	坎纳兹本Kannazbun feud.Barony
	基兹Kiz         feud.Barony
	曼恩Man         feud.Barony
	莫卧儿Mughal     feud.Barony
	奥尔马拉Ormara    feud.Barony
}

func (c *弥兰MakranCounty) BChhajian查绛() feud.Barony {
	return c.查绛Chhajian
}

func (c *弥兰MakranCounty) BKannazbun坎纳兹本() feud.Barony {
	return c.坎纳兹本Kannazbun
}

func (c *弥兰MakranCounty) BKiz基兹() feud.Barony {
	return c.基兹Kiz
}

func (c *弥兰MakranCounty) BMan曼恩() feud.Barony {
	return c.曼恩Man
}

func (c *弥兰MakranCounty) BMughal莫卧儿() feud.Barony {
	return c.莫卧儿Mughal
}

func (c *弥兰MakranCounty) BOrmara奥尔马拉() feud.Barony {
	return c.奥尔马拉Ormara
}

var CMakran弥兰 MakranCounty = &弥兰MakranCounty{}

func init() {
	f := CMakran弥兰.(*弥兰MakranCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1139",
		Title:     "makran",
		TitleName: "弥兰",
		TitleCode: "c_makran",
		Baronies:  map[string]feud.Barony{},
	}

	f.查绛Chhajian = BChhajian查绛
	f.查绛Chhajian.SetParent(f)

	f.坎纳兹本Kannazbun = BKannazbun坎纳兹本
	f.坎纳兹本Kannazbun.SetParent(f)

	f.基兹Kiz = BKiz基兹
	f.基兹Kiz.SetParent(f)

	f.曼恩Man = BMan曼恩
	f.曼恩Man.SetParent(f)

	f.莫卧儿Mughal = BMughal莫卧儿
	f.莫卧儿Mughal.SetParent(f)

	f.奥尔马拉Ormara = BOrmara奥尔马拉
	f.奥尔马拉Ormara.SetParent(f)

}
