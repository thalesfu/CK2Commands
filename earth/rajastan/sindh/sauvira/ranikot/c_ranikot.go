package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RanikotCounty interface {
	feud.County
	BAbu阿布() feud.Barony
	BJandar旬陀() feud.Barony
	BKhudabad丘陀跋() feud.Barony
	BManhabari曼诃婆尼() feud.Barony
	BPaat波罗() feud.Barony
	BRanikot兰尼科特() feud.Barony
	BSharusan舍楼桑() feud.Barony
}

type 兰尼科特RanikotCounty struct {
	feud.BaseCounty
	阿布Abu         feud.Barony
	旬陀Jandar      feud.Barony
	丘陀跋Khudabad   feud.Barony
	曼诃婆尼Manhabari feud.Barony
	波罗Paat        feud.Barony
	兰尼科特Ranikot   feud.Barony
	舍楼桑Sharusan   feud.Barony
}

func (c *兰尼科特RanikotCounty) BAbu阿布() feud.Barony {
	return c.阿布Abu
}

func (c *兰尼科特RanikotCounty) BJandar旬陀() feud.Barony {
	return c.旬陀Jandar
}

func (c *兰尼科特RanikotCounty) BKhudabad丘陀跋() feud.Barony {
	return c.丘陀跋Khudabad
}

func (c *兰尼科特RanikotCounty) BManhabari曼诃婆尼() feud.Barony {
	return c.曼诃婆尼Manhabari
}

func (c *兰尼科特RanikotCounty) BPaat波罗() feud.Barony {
	return c.波罗Paat
}

func (c *兰尼科特RanikotCounty) BRanikot兰尼科特() feud.Barony {
	return c.兰尼科特Ranikot
}

func (c *兰尼科特RanikotCounty) BSharusan舍楼桑() feud.Barony {
	return c.舍楼桑Sharusan
}

var CRanikot兰尼科特 RanikotCounty = &兰尼科特RanikotCounty{}

func init() {
	f := CRanikot兰尼科特.(*兰尼科特RanikotCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1303",
		Title:     "ranikot",
		TitleName: "兰尼科特",
		TitleCode: "c_ranikot",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布Abu = BAbu阿布
	f.阿布Abu.SetParent(f)

	f.旬陀Jandar = BJandar旬陀
	f.旬陀Jandar.SetParent(f)

	f.丘陀跋Khudabad = BKhudabad丘陀跋
	f.丘陀跋Khudabad.SetParent(f)

	f.曼诃婆尼Manhabari = BManhabari曼诃婆尼
	f.曼诃婆尼Manhabari.SetParent(f)

	f.波罗Paat = BPaat波罗
	f.波罗Paat.SetParent(f)

	f.兰尼科特Ranikot = BRanikot兰尼科特
	f.兰尼科特Ranikot.SetParent(f)

	f.舍楼桑Sharusan = BSharusan舍楼桑
	f.舍楼桑Sharusan.SetParent(f)

}
