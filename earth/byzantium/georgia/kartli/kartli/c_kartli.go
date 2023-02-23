package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KartliCounty interface {
	feud.County
	BAkhalkalaki阿哈尔卡拉基() feud.Barony
	BBolnisi博尔尼西() feud.Barony
	BDmanisi德马尼西() feud.Barony
	BNarikala纳里卡拉() feud.Barony
	BSvetitskhoveli斯韦季茨霍韦利() feud.Barony
	BTbilisi第比利斯() feud.Barony
	BTmoghvi特莫格维() feud.Barony
}

type 第比利斯KartliCounty struct {
	feud.BaseCounty
	阿哈尔卡拉基Akhalkalaki     feud.Barony
	博尔尼西Bolnisi           feud.Barony
	德马尼西Dmanisi           feud.Barony
	纳里卡拉Narikala          feud.Barony
	斯韦季茨霍韦利Svetitskhoveli feud.Barony
	第比利斯Tbilisi           feud.Barony
	特莫格维Tmoghvi           feud.Barony
}

func (c *第比利斯KartliCounty) BAkhalkalaki阿哈尔卡拉基() feud.Barony {
	return c.阿哈尔卡拉基Akhalkalaki
}

func (c *第比利斯KartliCounty) BBolnisi博尔尼西() feud.Barony {
	return c.博尔尼西Bolnisi
}

func (c *第比利斯KartliCounty) BDmanisi德马尼西() feud.Barony {
	return c.德马尼西Dmanisi
}

func (c *第比利斯KartliCounty) BNarikala纳里卡拉() feud.Barony {
	return c.纳里卡拉Narikala
}

func (c *第比利斯KartliCounty) BSvetitskhoveli斯韦季茨霍韦利() feud.Barony {
	return c.斯韦季茨霍韦利Svetitskhoveli
}

func (c *第比利斯KartliCounty) BTbilisi第比利斯() feud.Barony {
	return c.第比利斯Tbilisi
}

func (c *第比利斯KartliCounty) BTmoghvi特莫格维() feud.Barony {
	return c.特莫格维Tmoghvi
}

var CKartli第比利斯 KartliCounty = &第比利斯KartliCounty{}

func init() {
	f := CKartli第比利斯.(*第比利斯KartliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "679",
		Title:     "kartli",
		TitleName: "第比利斯",
		TitleCode: "c_kartli",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿哈尔卡拉基Akhalkalaki = BAkhalkalaki阿哈尔卡拉基
	f.阿哈尔卡拉基Akhalkalaki.SetParent(f)

	f.博尔尼西Bolnisi = BBolnisi博尔尼西
	f.博尔尼西Bolnisi.SetParent(f)

	f.德马尼西Dmanisi = BDmanisi德马尼西
	f.德马尼西Dmanisi.SetParent(f)

	f.纳里卡拉Narikala = BNarikala纳里卡拉
	f.纳里卡拉Narikala.SetParent(f)

	f.斯韦季茨霍韦利Svetitskhoveli = BSvetitskhoveli斯韦季茨霍韦利
	f.斯韦季茨霍韦利Svetitskhoveli.SetParent(f)

	f.第比利斯Tbilisi = BTbilisi第比利斯
	f.第比利斯Tbilisi.SetParent(f)

	f.特莫格维Tmoghvi = BTmoghvi特莫格维
	f.特莫格维Tmoghvi.SetParent(f)

}
