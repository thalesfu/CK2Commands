package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhangaiCounty interface {
	feud.County
	BErdenet额尔登特() feud.Barony
	BKhangai杭海() feud.Barony
	BOgii额吉() feud.Barony
	BOrgil敖尔吉勒() feud.Barony
	BSangiin桑根() feud.Barony
	BTsetserleg车车尔勒格() feud.Barony
	BUliastai乌里雅苏台() feud.Barony
}

type 杭海KhangaiCounty struct {
	feud.BaseCounty
	额尔登特Erdenet     feud.Barony
	杭海Khangai       feud.Barony
	额吉Ogii          feud.Barony
	敖尔吉勒Orgil       feud.Barony
	桑根Sangiin       feud.Barony
	车车尔勒格Tsetserleg feud.Barony
	乌里雅苏台Uliastai   feud.Barony
}

func (c *杭海KhangaiCounty) BErdenet额尔登特() feud.Barony {
	return c.额尔登特Erdenet
}

func (c *杭海KhangaiCounty) BKhangai杭海() feud.Barony {
	return c.杭海Khangai
}

func (c *杭海KhangaiCounty) BOgii额吉() feud.Barony {
	return c.额吉Ogii
}

func (c *杭海KhangaiCounty) BOrgil敖尔吉勒() feud.Barony {
	return c.敖尔吉勒Orgil
}

func (c *杭海KhangaiCounty) BSangiin桑根() feud.Barony {
	return c.桑根Sangiin
}

func (c *杭海KhangaiCounty) BTsetserleg车车尔勒格() feud.Barony {
	return c.车车尔勒格Tsetserleg
}

func (c *杭海KhangaiCounty) BUliastai乌里雅苏台() feud.Barony {
	return c.乌里雅苏台Uliastai
}

var CKhangai杭海 KhangaiCounty = &杭海KhangaiCounty{}

func init() {
	f := CKhangai杭海.(*杭海KhangaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1458",
		Title:     "khangai",
		TitleName: "杭海",
		TitleCode: "c_khangai",
		Baronies:  map[string]feud.Barony{},
	}

	f.额尔登特Erdenet = BErdenet额尔登特
	f.额尔登特Erdenet.SetParent(f)

	f.杭海Khangai = BKhangai杭海
	f.杭海Khangai.SetParent(f)

	f.额吉Ogii = BOgii额吉
	f.额吉Ogii.SetParent(f)

	f.敖尔吉勒Orgil = BOrgil敖尔吉勒
	f.敖尔吉勒Orgil.SetParent(f)

	f.桑根Sangiin = BSangiin桑根
	f.桑根Sangiin.SetParent(f)

	f.车车尔勒格Tsetserleg = BTsetserleg车车尔勒格
	f.车车尔勒格Tsetserleg.SetParent(f)

	f.乌里雅苏台Uliastai = BUliastai乌里雅苏台
	f.乌里雅苏台Uliastai.SetParent(f)

}
