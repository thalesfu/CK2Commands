package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BarskhanCounty interface {
	feud.County
	BBarskhan拔塞干() feud.Barony
	BLipenka利片卡() feud.Barony
	BTosor托索尔() feud.Barony
	BTyup蒂普() feud.Barony
}

type 拔塞干BarskhanCounty struct {
	feud.BaseCounty
	拔塞干Barskhan feud.Barony
	利片卡Lipenka  feud.Barony
	托索尔Tosor    feud.Barony
	蒂普Tyup      feud.Barony
}

func (c *拔塞干BarskhanCounty) BBarskhan拔塞干() feud.Barony {
	return c.拔塞干Barskhan
}

func (c *拔塞干BarskhanCounty) BLipenka利片卡() feud.Barony {
	return c.利片卡Lipenka
}

func (c *拔塞干BarskhanCounty) BTosor托索尔() feud.Barony {
	return c.托索尔Tosor
}

func (c *拔塞干BarskhanCounty) BTyup蒂普() feud.Barony {
	return c.蒂普Tyup
}

var CBarskhan拔塞干 BarskhanCounty = &拔塞干BarskhanCounty{}

func init() {
	f := CBarskhan拔塞干.(*拔塞干BarskhanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1796",
		Title:     "barskhan",
		TitleName: "拔塞干",
		TitleCode: "c_barskhan",
		Baronies:  map[string]feud.Barony{},
	}

	f.拔塞干Barskhan = BBarskhan拔塞干
	f.拔塞干Barskhan.SetParent(f)

	f.利片卡Lipenka = BLipenka利片卡
	f.利片卡Lipenka.SetParent(f)

	f.托索尔Tosor = BTosor托索尔
	f.托索尔Tosor.SetParent(f)

	f.蒂普Tyup = BTyup蒂普
	f.蒂普Tyup.SetParent(f)

}
