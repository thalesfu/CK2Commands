package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZhetysuCounty interface {
	feud.County
	BKaratal卡拉塔尔() feud.Barony
	BKoksu科克苏() feud.Barony
	BKuraksu库拉克苏() feud.Barony
	BLepsy列普瑟() feud.Barony
	BMatay马泰() feud.Barony
	BSarkand萨尔坎德() feud.Barony
	BTaldykorgan塔尔迪库尔干() feud.Barony
}

type 卡拉塔尔ZhetysuCounty struct {
	feud.BaseCounty
	卡拉塔尔Karatal       feud.Barony
	科克苏Koksu          feud.Barony
	库拉克苏Kuraksu       feud.Barony
	列普瑟Lepsy          feud.Barony
	马泰Matay           feud.Barony
	萨尔坎德Sarkand       feud.Barony
	塔尔迪库尔干Taldykorgan feud.Barony
}

func (c *卡拉塔尔ZhetysuCounty) BKaratal卡拉塔尔() feud.Barony {
	return c.卡拉塔尔Karatal
}

func (c *卡拉塔尔ZhetysuCounty) BKoksu科克苏() feud.Barony {
	return c.科克苏Koksu
}

func (c *卡拉塔尔ZhetysuCounty) BKuraksu库拉克苏() feud.Barony {
	return c.库拉克苏Kuraksu
}

func (c *卡拉塔尔ZhetysuCounty) BLepsy列普瑟() feud.Barony {
	return c.列普瑟Lepsy
}

func (c *卡拉塔尔ZhetysuCounty) BMatay马泰() feud.Barony {
	return c.马泰Matay
}

func (c *卡拉塔尔ZhetysuCounty) BSarkand萨尔坎德() feud.Barony {
	return c.萨尔坎德Sarkand
}

func (c *卡拉塔尔ZhetysuCounty) BTaldykorgan塔尔迪库尔干() feud.Barony {
	return c.塔尔迪库尔干Taldykorgan
}

var CZhetysu卡拉塔尔 ZhetysuCounty = &卡拉塔尔ZhetysuCounty{}

func init() {
	f := CZhetysu卡拉塔尔.(*卡拉塔尔ZhetysuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1426",
		Title:     "zhetysu",
		TitleName: "卡拉塔尔",
		TitleCode: "c_zhetysu",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉塔尔Karatal = BKaratal卡拉塔尔
	f.卡拉塔尔Karatal.SetParent(f)

	f.科克苏Koksu = BKoksu科克苏
	f.科克苏Koksu.SetParent(f)

	f.库拉克苏Kuraksu = BKuraksu库拉克苏
	f.库拉克苏Kuraksu.SetParent(f)

	f.列普瑟Lepsy = BLepsy列普瑟
	f.列普瑟Lepsy.SetParent(f)

	f.马泰Matay = BMatay马泰
	f.马泰Matay.SetParent(f)

	f.萨尔坎德Sarkand = BSarkand萨尔坎德
	f.萨尔坎德Sarkand.SetParent(f)

	f.塔尔迪库尔干Taldykorgan = BTaldykorgan塔尔迪库尔干
	f.塔尔迪库尔干Taldykorgan.SetParent(f)

}
