package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShorkotCounty interface {
	feud.County
	BBhawana巴哈瓦那() feud.Barony
	BBhirr班尔() feud.Barony
	BChandrod旃陀卢陀() feud.Barony
	BJhang章() feud.Barony
	BKamalia卡马拉() feud.Barony
	BRabwah拉布瓦() feud.Barony
	BShorkot绍科特() feud.Barony
}

type 绍科特ShorkotCounty struct {
	feud.BaseCounty
	巴哈瓦那Bhawana  feud.Barony
	班尔Bhirr      feud.Barony
	旃陀卢陀Chandrod feud.Barony
	章Jhang       feud.Barony
	卡马拉Kamalia   feud.Barony
	拉布瓦Rabwah    feud.Barony
	绍科特Shorkot   feud.Barony
}

func (c *绍科特ShorkotCounty) BBhawana巴哈瓦那() feud.Barony {
	return c.巴哈瓦那Bhawana
}

func (c *绍科特ShorkotCounty) BBhirr班尔() feud.Barony {
	return c.班尔Bhirr
}

func (c *绍科特ShorkotCounty) BChandrod旃陀卢陀() feud.Barony {
	return c.旃陀卢陀Chandrod
}

func (c *绍科特ShorkotCounty) BJhang章() feud.Barony {
	return c.章Jhang
}

func (c *绍科特ShorkotCounty) BKamalia卡马拉() feud.Barony {
	return c.卡马拉Kamalia
}

func (c *绍科特ShorkotCounty) BRabwah拉布瓦() feud.Barony {
	return c.拉布瓦Rabwah
}

func (c *绍科特ShorkotCounty) BShorkot绍科特() feud.Barony {
	return c.绍科特Shorkot
}

var CShorkot绍科特 ShorkotCounty = &绍科特ShorkotCounty{}

func init() {
	f := CShorkot绍科特.(*绍科特ShorkotCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1361",
		Title:     "shorkot",
		TitleName: "绍科特",
		TitleCode: "c_shorkot",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴哈瓦那Bhawana = BBhawana巴哈瓦那
	f.巴哈瓦那Bhawana.SetParent(f)

	f.班尔Bhirr = BBhirr班尔
	f.班尔Bhirr.SetParent(f)

	f.旃陀卢陀Chandrod = BChandrod旃陀卢陀
	f.旃陀卢陀Chandrod.SetParent(f)

	f.章Jhang = BJhang章
	f.章Jhang.SetParent(f)

	f.卡马拉Kamalia = BKamalia卡马拉
	f.卡马拉Kamalia.SetParent(f)

	f.拉布瓦Rabwah = BRabwah拉布瓦
	f.拉布瓦Rabwah.SetParent(f)

	f.绍科特Shorkot = BShorkot绍科特
	f.绍科特Shorkot.SetParent(f)

}
