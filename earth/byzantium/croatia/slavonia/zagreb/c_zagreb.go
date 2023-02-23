package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZagrebCounty interface {
	feud.County
	BCetin采廷() feud.Barony
	BDreznik德雷兹尼克() feud.Barony
	BKarlovac卡尔洛瓦茨() feud.Barony
	BOzalj奥扎利() feud.Barony
	BSisak锡萨克() feud.Barony
	BStenicnjak斯泰尼奇尼亚克() feud.Barony
	BZagreb扎格雷布() feud.Barony
	BZrin兹林() feud.Barony
}

type 扎格雷布ZagrebCounty struct {
	feud.BaseCounty
	采廷Cetin           feud.Barony
	德雷兹尼克Dreznik      feud.Barony
	卡尔洛瓦茨Karlovac     feud.Barony
	奥扎利Ozalj          feud.Barony
	锡萨克Sisak          feud.Barony
	斯泰尼奇尼亚克Stenicnjak feud.Barony
	扎格雷布Zagreb        feud.Barony
	兹林Zrin            feud.Barony
}

func (c *扎格雷布ZagrebCounty) BCetin采廷() feud.Barony {
	return c.采廷Cetin
}

func (c *扎格雷布ZagrebCounty) BDreznik德雷兹尼克() feud.Barony {
	return c.德雷兹尼克Dreznik
}

func (c *扎格雷布ZagrebCounty) BKarlovac卡尔洛瓦茨() feud.Barony {
	return c.卡尔洛瓦茨Karlovac
}

func (c *扎格雷布ZagrebCounty) BOzalj奥扎利() feud.Barony {
	return c.奥扎利Ozalj
}

func (c *扎格雷布ZagrebCounty) BSisak锡萨克() feud.Barony {
	return c.锡萨克Sisak
}

func (c *扎格雷布ZagrebCounty) BStenicnjak斯泰尼奇尼亚克() feud.Barony {
	return c.斯泰尼奇尼亚克Stenicnjak
}

func (c *扎格雷布ZagrebCounty) BZagreb扎格雷布() feud.Barony {
	return c.扎格雷布Zagreb
}

func (c *扎格雷布ZagrebCounty) BZrin兹林() feud.Barony {
	return c.兹林Zrin
}

var CZagreb扎格雷布 ZagrebCounty = &扎格雷布ZagrebCounty{}

func init() {
	f := CZagreb扎格雷布.(*扎格雷布ZagrebCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "461",
		Title:     "zagreb",
		TitleName: "扎格雷布",
		TitleCode: "c_zagreb",
		Baronies:  map[string]feud.Barony{},
	}

	f.采廷Cetin = BCetin采廷
	f.采廷Cetin.SetParent(f)

	f.德雷兹尼克Dreznik = BDreznik德雷兹尼克
	f.德雷兹尼克Dreznik.SetParent(f)

	f.卡尔洛瓦茨Karlovac = BKarlovac卡尔洛瓦茨
	f.卡尔洛瓦茨Karlovac.SetParent(f)

	f.奥扎利Ozalj = BOzalj奥扎利
	f.奥扎利Ozalj.SetParent(f)

	f.锡萨克Sisak = BSisak锡萨克
	f.锡萨克Sisak.SetParent(f)

	f.斯泰尼奇尼亚克Stenicnjak = BStenicnjak斯泰尼奇尼亚克
	f.斯泰尼奇尼亚克Stenicnjak.SetParent(f)

	f.扎格雷布Zagreb = BZagreb扎格雷布
	f.扎格雷布Zagreb.SetParent(f)

	f.兹林Zrin = BZrin兹林
	f.兹林Zrin.SetParent(f)

}
