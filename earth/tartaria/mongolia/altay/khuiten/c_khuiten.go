package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhuitenCounty interface {
	feud.County
	BArshaty阿尔沙特() feud.Barony
	BBuyant布彦特() feud.Barony
	BChikhertei奇赫尔特() feud.Barony
	BKhoton和屯() feud.Barony
	BKhuiten辉腾() feud.Barony
	BKhushuut赫舍特() feud.Barony
	BSorvenok索尔韦诺克() feud.Barony
}

type 辉腾KhuitenCounty struct {
	feud.BaseCounty
	阿尔沙特Arshaty    feud.Barony
	布彦特Buyant      feud.Barony
	奇赫尔特Chikhertei feud.Barony
	和屯Khoton       feud.Barony
	辉腾Khuiten      feud.Barony
	赫舍特Khushuut    feud.Barony
	索尔韦诺克Sorvenok  feud.Barony
}

func (c *辉腾KhuitenCounty) BArshaty阿尔沙特() feud.Barony {
	return c.阿尔沙特Arshaty
}

func (c *辉腾KhuitenCounty) BBuyant布彦特() feud.Barony {
	return c.布彦特Buyant
}

func (c *辉腾KhuitenCounty) BChikhertei奇赫尔特() feud.Barony {
	return c.奇赫尔特Chikhertei
}

func (c *辉腾KhuitenCounty) BKhoton和屯() feud.Barony {
	return c.和屯Khoton
}

func (c *辉腾KhuitenCounty) BKhuiten辉腾() feud.Barony {
	return c.辉腾Khuiten
}

func (c *辉腾KhuitenCounty) BKhushuut赫舍特() feud.Barony {
	return c.赫舍特Khushuut
}

func (c *辉腾KhuitenCounty) BSorvenok索尔韦诺克() feud.Barony {
	return c.索尔韦诺克Sorvenok
}

var CKhuiten辉腾 KhuitenCounty = &辉腾KhuitenCounty{}

func init() {
	f := CKhuiten辉腾.(*辉腾KhuitenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1896",
		Title:     "khuiten",
		TitleName: "辉腾",
		TitleCode: "c_khuiten",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔沙特Arshaty = BArshaty阿尔沙特
	f.阿尔沙特Arshaty.SetParent(f)

	f.布彦特Buyant = BBuyant布彦特
	f.布彦特Buyant.SetParent(f)

	f.奇赫尔特Chikhertei = BChikhertei奇赫尔特
	f.奇赫尔特Chikhertei.SetParent(f)

	f.和屯Khoton = BKhoton和屯
	f.和屯Khoton.SetParent(f)

	f.辉腾Khuiten = BKhuiten辉腾
	f.辉腾Khuiten.SetParent(f)

	f.赫舍特Khushuut = BKhushuut赫舍特
	f.赫舍特Khushuut.SetParent(f)

	f.索尔韦诺克Sorvenok = BSorvenok索尔韦诺克
	f.索尔韦诺克Sorvenok.SetParent(f)

}
