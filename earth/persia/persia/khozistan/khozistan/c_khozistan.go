package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhozistanCounty interface {
	feud.County
	BAbadan阿巴丹() feud.Barony
	BAhvaz阿瓦士() feud.Barony
	BDora达尔拉() feud.Barony
	BFao法奥() feud.Barony
	BHoveizeh霍韦伊泽() feud.Barony
	BIzaj伊泽() feud.Barony
	BKhorramshahr霍拉姆沙赫尔() feud.Barony
	BShadegan沙代甘() feud.Barony
}

type 胡齐斯坦KhozistanCounty struct {
	feud.BaseCounty
	阿巴丹Abadan          feud.Barony
	阿瓦士Ahvaz           feud.Barony
	达尔拉Dora            feud.Barony
	法奥Fao              feud.Barony
	霍韦伊泽Hoveizeh       feud.Barony
	伊泽Izaj             feud.Barony
	霍拉姆沙赫尔Khorramshahr feud.Barony
	沙代甘Shadegan        feud.Barony
}

func (c *胡齐斯坦KhozistanCounty) BAbadan阿巴丹() feud.Barony {
	return c.阿巴丹Abadan
}

func (c *胡齐斯坦KhozistanCounty) BAhvaz阿瓦士() feud.Barony {
	return c.阿瓦士Ahvaz
}

func (c *胡齐斯坦KhozistanCounty) BDora达尔拉() feud.Barony {
	return c.达尔拉Dora
}

func (c *胡齐斯坦KhozistanCounty) BFao法奥() feud.Barony {
	return c.法奥Fao
}

func (c *胡齐斯坦KhozistanCounty) BHoveizeh霍韦伊泽() feud.Barony {
	return c.霍韦伊泽Hoveizeh
}

func (c *胡齐斯坦KhozistanCounty) BIzaj伊泽() feud.Barony {
	return c.伊泽Izaj
}

func (c *胡齐斯坦KhozistanCounty) BKhorramshahr霍拉姆沙赫尔() feud.Barony {
	return c.霍拉姆沙赫尔Khorramshahr
}

func (c *胡齐斯坦KhozistanCounty) BShadegan沙代甘() feud.Barony {
	return c.沙代甘Shadegan
}

var CKhozistan胡齐斯坦 KhozistanCounty = &胡齐斯坦KhozistanCounty{}

func init() {
	f := CKhozistan胡齐斯坦.(*胡齐斯坦KhozistanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "648",
		Title:     "khozistan",
		TitleName: "胡齐斯坦",
		TitleCode: "c_khozistan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴丹Abadan = BAbadan阿巴丹
	f.阿巴丹Abadan.SetParent(f)

	f.阿瓦士Ahvaz = BAhvaz阿瓦士
	f.阿瓦士Ahvaz.SetParent(f)

	f.达尔拉Dora = BDora达尔拉
	f.达尔拉Dora.SetParent(f)

	f.法奥Fao = BFao法奥
	f.法奥Fao.SetParent(f)

	f.霍韦伊泽Hoveizeh = BHoveizeh霍韦伊泽
	f.霍韦伊泽Hoveizeh.SetParent(f)

	f.伊泽Izaj = BIzaj伊泽
	f.伊泽Izaj.SetParent(f)

	f.霍拉姆沙赫尔Khorramshahr = BKhorramshahr霍拉姆沙赫尔
	f.霍拉姆沙赫尔Khorramshahr.SetParent(f)

	f.沙代甘Shadegan = BShadegan沙代甘
	f.沙代甘Shadegan.SetParent(f)

}
