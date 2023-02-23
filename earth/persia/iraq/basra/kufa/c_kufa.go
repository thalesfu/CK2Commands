package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KufaCounty interface {
	feud.County
	BAlqurnah古尔奈() feud.Barony
	BBussayyah布赛亚() feud.Barony
	BChibayish吉拜什() feud.Barony
	BHammar哈马尔() feud.Barony
	BKufa库法() feud.Barony
	BRagai拉戈依() feud.Barony
	BShuyukh舒尤赫() feud.Barony
	BSuqash苏格阿赫() feud.Barony
}

type 库法KufaCounty struct {
	feud.BaseCounty
	古尔奈Alqurnah  feud.Barony
	布赛亚Bussayyah feud.Barony
	吉拜什Chibayish feud.Barony
	哈马尔Hammar    feud.Barony
	库法Kufa       feud.Barony
	拉戈依Ragai     feud.Barony
	舒尤赫Shuyukh   feud.Barony
	苏格阿赫Suqash   feud.Barony
}

func (c *库法KufaCounty) BAlqurnah古尔奈() feud.Barony {
	return c.古尔奈Alqurnah
}

func (c *库法KufaCounty) BBussayyah布赛亚() feud.Barony {
	return c.布赛亚Bussayyah
}

func (c *库法KufaCounty) BChibayish吉拜什() feud.Barony {
	return c.吉拜什Chibayish
}

func (c *库法KufaCounty) BHammar哈马尔() feud.Barony {
	return c.哈马尔Hammar
}

func (c *库法KufaCounty) BKufa库法() feud.Barony {
	return c.库法Kufa
}

func (c *库法KufaCounty) BRagai拉戈依() feud.Barony {
	return c.拉戈依Ragai
}

func (c *库法KufaCounty) BShuyukh舒尤赫() feud.Barony {
	return c.舒尤赫Shuyukh
}

func (c *库法KufaCounty) BSuqash苏格阿赫() feud.Barony {
	return c.苏格阿赫Suqash
}

var CKufa库法 KufaCounty = &库法KufaCounty{}

func init() {
	f := CKufa库法.(*库法KufaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "655",
		Title:     "kufa",
		TitleName: "库法",
		TitleCode: "c_kufa",
		Baronies:  map[string]feud.Barony{},
	}

	f.古尔奈Alqurnah = BAlqurnah古尔奈
	f.古尔奈Alqurnah.SetParent(f)

	f.布赛亚Bussayyah = BBussayyah布赛亚
	f.布赛亚Bussayyah.SetParent(f)

	f.吉拜什Chibayish = BChibayish吉拜什
	f.吉拜什Chibayish.SetParent(f)

	f.哈马尔Hammar = BHammar哈马尔
	f.哈马尔Hammar.SetParent(f)

	f.库法Kufa = BKufa库法
	f.库法Kufa.SetParent(f)

	f.拉戈依Ragai = BRagai拉戈依
	f.拉戈依Ragai.SetParent(f)

	f.舒尤赫Shuyukh = BShuyukh舒尤赫
	f.舒尤赫Shuyukh.SetParent(f)

	f.苏格阿赫Suqash = BSuqash苏格阿赫
	f.苏格阿赫Suqash.SetParent(f)

}
