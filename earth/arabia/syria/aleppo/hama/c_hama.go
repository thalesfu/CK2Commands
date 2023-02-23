package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HamaCounty interface {
	feud.County
	BBara巴拉() feud.Barony
	BHama哈马() feud.Barony
	BHamath哈马斯() feud.Barony
	BKharsan哈拉桑() feud.Barony
	BMhardeh玛哈德() feud.Barony
	BQarqar甲加() feud.Barony
	BSalamiyah塞莱米耶() feud.Barony
	BSerjilla塞吉拉() feud.Barony
}

type 哈马HamaCounty struct {
	feud.BaseCounty
	巴拉Bara        feud.Barony
	哈马Hama        feud.Barony
	哈马斯Hamath     feud.Barony
	哈拉桑Kharsan    feud.Barony
	玛哈德Mhardeh    feud.Barony
	甲加Qarqar      feud.Barony
	塞莱米耶Salamiyah feud.Barony
	塞吉拉Serjilla   feud.Barony
}

func (c *哈马HamaCounty) BBara巴拉() feud.Barony {
	return c.巴拉Bara
}

func (c *哈马HamaCounty) BHama哈马() feud.Barony {
	return c.哈马Hama
}

func (c *哈马HamaCounty) BHamath哈马斯() feud.Barony {
	return c.哈马斯Hamath
}

func (c *哈马HamaCounty) BKharsan哈拉桑() feud.Barony {
	return c.哈拉桑Kharsan
}

func (c *哈马HamaCounty) BMhardeh玛哈德() feud.Barony {
	return c.玛哈德Mhardeh
}

func (c *哈马HamaCounty) BQarqar甲加() feud.Barony {
	return c.甲加Qarqar
}

func (c *哈马HamaCounty) BSalamiyah塞莱米耶() feud.Barony {
	return c.塞莱米耶Salamiyah
}

func (c *哈马HamaCounty) BSerjilla塞吉拉() feud.Barony {
	return c.塞吉拉Serjilla
}

var CHama哈马 HamaCounty = &哈马HamaCounty{}

func init() {
	f := CHama哈马.(*哈马HamaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "732",
		Title:     "hama",
		TitleName: "哈马",
		TitleCode: "c_hama",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉Bara = BBara巴拉
	f.巴拉Bara.SetParent(f)

	f.哈马Hama = BHama哈马
	f.哈马Hama.SetParent(f)

	f.哈马斯Hamath = BHamath哈马斯
	f.哈马斯Hamath.SetParent(f)

	f.哈拉桑Kharsan = BKharsan哈拉桑
	f.哈拉桑Kharsan.SetParent(f)

	f.玛哈德Mhardeh = BMhardeh玛哈德
	f.玛哈德Mhardeh.SetParent(f)

	f.甲加Qarqar = BQarqar甲加
	f.甲加Qarqar.SetParent(f)

	f.塞莱米耶Salamiyah = BSalamiyah塞莱米耶
	f.塞莱米耶Salamiyah.SetParent(f)

	f.塞吉拉Serjilla = BSerjilla塞吉拉
	f.塞吉拉Serjilla.SetParent(f)

}
