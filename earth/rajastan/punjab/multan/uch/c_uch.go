package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UchCounty interface {
	feud.County
	BChachran卡克兰() feud.Barony
	BChaudari招陀利() feud.Barony
	BDerawar德拉瓦() feud.Barony
	BMithankot米坦科特() feud.Barony
	BRardhu罗菟() feud.Barony
	BSitpur悉多补罗() feud.Barony
	BUch邬脂() feud.Barony
}

type 邬脂UchCounty struct {
	feud.BaseCounty
	卡克兰Chachran   feud.Barony
	招陀利Chaudari   feud.Barony
	德拉瓦Derawar    feud.Barony
	米坦科特Mithankot feud.Barony
	罗菟Rardhu      feud.Barony
	悉多补罗Sitpur    feud.Barony
	邬脂Uch         feud.Barony
}

func (c *邬脂UchCounty) BChachran卡克兰() feud.Barony {
	return c.卡克兰Chachran
}

func (c *邬脂UchCounty) BChaudari招陀利() feud.Barony {
	return c.招陀利Chaudari
}

func (c *邬脂UchCounty) BDerawar德拉瓦() feud.Barony {
	return c.德拉瓦Derawar
}

func (c *邬脂UchCounty) BMithankot米坦科特() feud.Barony {
	return c.米坦科特Mithankot
}

func (c *邬脂UchCounty) BRardhu罗菟() feud.Barony {
	return c.罗菟Rardhu
}

func (c *邬脂UchCounty) BSitpur悉多补罗() feud.Barony {
	return c.悉多补罗Sitpur
}

func (c *邬脂UchCounty) BUch邬脂() feud.Barony {
	return c.邬脂Uch
}

var CUch邬脂 UchCounty = &邬脂UchCounty{}

func init() {
	f := CUch邬脂.(*邬脂UchCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1337",
		Title:     "uch",
		TitleName: "邬脂",
		TitleCode: "c_uch",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡克兰Chachran = BChachran卡克兰
	f.卡克兰Chachran.SetParent(f)

	f.招陀利Chaudari = BChaudari招陀利
	f.招陀利Chaudari.SetParent(f)

	f.德拉瓦Derawar = BDerawar德拉瓦
	f.德拉瓦Derawar.SetParent(f)

	f.米坦科特Mithankot = BMithankot米坦科特
	f.米坦科特Mithankot.SetParent(f)

	f.罗菟Rardhu = BRardhu罗菟
	f.罗菟Rardhu.SetParent(f)

	f.悉多补罗Sitpur = BSitpur悉多补罗
	f.悉多补罗Sitpur.SetParent(f)

	f.邬脂Uch = BUch邬脂
	f.邬脂Uch.SetParent(f)

}
