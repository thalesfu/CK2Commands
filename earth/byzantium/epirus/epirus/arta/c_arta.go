package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArtaCounty interface {
	feud.County
	BAgnanta阿格南塔() feud.Barony
	BAngelokastron安格洛卡斯特罗() feud.Barony
	BArta阿尔塔() feud.Barony
	BPreveza普雷韦扎() feud.Barony
	BRogoi罗戈伊() feud.Barony
	BThomokastron多莫卡斯特朗() feud.Barony
	BVlacherna弗拉海尔纳() feud.Barony
	BVonitsza沃尼察() feud.Barony
}

type 阿尔塔ArtaCounty struct {
	feud.BaseCounty
	阿格南塔Agnanta          feud.Barony
	安格洛卡斯特罗Angelokastron feud.Barony
	阿尔塔Arta              feud.Barony
	普雷韦扎Preveza          feud.Barony
	罗戈伊Rogoi             feud.Barony
	多莫卡斯特朗Thomokastron   feud.Barony
	弗拉海尔纳Vlacherna       feud.Barony
	沃尼察Vonitsza          feud.Barony
}

func (c *阿尔塔ArtaCounty) BAgnanta阿格南塔() feud.Barony {
	return c.阿格南塔Agnanta
}

func (c *阿尔塔ArtaCounty) BAngelokastron安格洛卡斯特罗() feud.Barony {
	return c.安格洛卡斯特罗Angelokastron
}

func (c *阿尔塔ArtaCounty) BArta阿尔塔() feud.Barony {
	return c.阿尔塔Arta
}

func (c *阿尔塔ArtaCounty) BPreveza普雷韦扎() feud.Barony {
	return c.普雷韦扎Preveza
}

func (c *阿尔塔ArtaCounty) BRogoi罗戈伊() feud.Barony {
	return c.罗戈伊Rogoi
}

func (c *阿尔塔ArtaCounty) BThomokastron多莫卡斯特朗() feud.Barony {
	return c.多莫卡斯特朗Thomokastron
}

func (c *阿尔塔ArtaCounty) BVlacherna弗拉海尔纳() feud.Barony {
	return c.弗拉海尔纳Vlacherna
}

func (c *阿尔塔ArtaCounty) BVonitsza沃尼察() feud.Barony {
	return c.沃尼察Vonitsza
}

var CArta阿尔塔 ArtaCounty = &阿尔塔ArtaCounty{}

func init() {
	f := CArta阿尔塔.(*阿尔塔ArtaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "473",
		Title:     "arta",
		TitleName: "阿尔塔",
		TitleCode: "c_arta",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格南塔Agnanta = BAgnanta阿格南塔
	f.阿格南塔Agnanta.SetParent(f)

	f.安格洛卡斯特罗Angelokastron = BAngelokastron安格洛卡斯特罗
	f.安格洛卡斯特罗Angelokastron.SetParent(f)

	f.阿尔塔Arta = BArta阿尔塔
	f.阿尔塔Arta.SetParent(f)

	f.普雷韦扎Preveza = BPreveza普雷韦扎
	f.普雷韦扎Preveza.SetParent(f)

	f.罗戈伊Rogoi = BRogoi罗戈伊
	f.罗戈伊Rogoi.SetParent(f)

	f.多莫卡斯特朗Thomokastron = BThomokastron多莫卡斯特朗
	f.多莫卡斯特朗Thomokastron.SetParent(f)

	f.弗拉海尔纳Vlacherna = BVlacherna弗拉海尔纳
	f.弗拉海尔纳Vlacherna.SetParent(f)

	f.沃尼察Vonitsza = BVonitsza沃尼察
	f.沃尼察Vonitsza.SetParent(f)

}
