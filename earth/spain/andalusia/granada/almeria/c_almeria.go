package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlmeriaCounty interface {
	feud.County
	BAlbox阿尔沃克斯() feud.Barony
	BAlmeria阿尔梅里亚() feud.Barony
	BBaza巴萨() feud.Barony
	BBerja贝尔哈() feud.Barony
	BMotril莫特里尔() feud.Barony
	BPechina佩奇纳() feud.Barony
	BPurchena普尔切纳() feud.Barony
	BVera贝拉() feud.Barony
}

type 阿尔梅里亚AlmeriaCounty struct {
	feud.BaseCounty
	阿尔沃克斯Albox   feud.Barony
	阿尔梅里亚Almeria feud.Barony
	巴萨Baza       feud.Barony
	贝尔哈Berja     feud.Barony
	莫特里尔Motril   feud.Barony
	佩奇纳Pechina   feud.Barony
	普尔切纳Purchena feud.Barony
	贝拉Vera       feud.Barony
}

func (c *阿尔梅里亚AlmeriaCounty) BAlbox阿尔沃克斯() feud.Barony {
	return c.阿尔沃克斯Albox
}

func (c *阿尔梅里亚AlmeriaCounty) BAlmeria阿尔梅里亚() feud.Barony {
	return c.阿尔梅里亚Almeria
}

func (c *阿尔梅里亚AlmeriaCounty) BBaza巴萨() feud.Barony {
	return c.巴萨Baza
}

func (c *阿尔梅里亚AlmeriaCounty) BBerja贝尔哈() feud.Barony {
	return c.贝尔哈Berja
}

func (c *阿尔梅里亚AlmeriaCounty) BMotril莫特里尔() feud.Barony {
	return c.莫特里尔Motril
}

func (c *阿尔梅里亚AlmeriaCounty) BPechina佩奇纳() feud.Barony {
	return c.佩奇纳Pechina
}

func (c *阿尔梅里亚AlmeriaCounty) BPurchena普尔切纳() feud.Barony {
	return c.普尔切纳Purchena
}

func (c *阿尔梅里亚AlmeriaCounty) BVera贝拉() feud.Barony {
	return c.贝拉Vera
}

var CAlmeria阿尔梅里亚 AlmeriaCounty = &阿尔梅里亚AlmeriaCounty{}

func init() {
	f := CAlmeria阿尔梅里亚.(*阿尔梅里亚AlmeriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "168",
		Title:     "almeria",
		TitleName: "阿尔梅里亚",
		TitleCode: "c_almeria",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔沃克斯Albox = BAlbox阿尔沃克斯
	f.阿尔沃克斯Albox.SetParent(f)

	f.阿尔梅里亚Almeria = BAlmeria阿尔梅里亚
	f.阿尔梅里亚Almeria.SetParent(f)

	f.巴萨Baza = BBaza巴萨
	f.巴萨Baza.SetParent(f)

	f.贝尔哈Berja = BBerja贝尔哈
	f.贝尔哈Berja.SetParent(f)

	f.莫特里尔Motril = BMotril莫特里尔
	f.莫特里尔Motril.SetParent(f)

	f.佩奇纳Pechina = BPechina佩奇纳
	f.佩奇纳Pechina.SetParent(f)

	f.普尔切纳Purchena = BPurchena普尔切纳
	f.普尔切纳Purchena.SetParent(f)

	f.贝拉Vera = BVera贝拉
	f.贝拉Vera.SetParent(f)

}
