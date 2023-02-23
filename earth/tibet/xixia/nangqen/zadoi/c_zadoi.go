package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZadoiCounty interface {
	feud.County
	BAduo阿多() feud.Barony
	BAngsai昂赛() feud.Barony
	BChadan查旦() feud.Barony
	BMoyun莫云() feud.Barony
	BQapugtang萨呼腾() feud.Barony
	BZadoi杂多() feud.Barony
	BZhaqing扎青() feud.Barony
}

type 杂多ZadoiCounty struct {
	feud.BaseCounty
	阿多Aduo       feud.Barony
	昂赛Angsai     feud.Barony
	查旦Chadan     feud.Barony
	莫云Moyun      feud.Barony
	萨呼腾Qapugtang feud.Barony
	杂多Zadoi      feud.Barony
	扎青Zhaqing    feud.Barony
}

func (c *杂多ZadoiCounty) BAduo阿多() feud.Barony {
	return c.阿多Aduo
}

func (c *杂多ZadoiCounty) BAngsai昂赛() feud.Barony {
	return c.昂赛Angsai
}

func (c *杂多ZadoiCounty) BChadan查旦() feud.Barony {
	return c.查旦Chadan
}

func (c *杂多ZadoiCounty) BMoyun莫云() feud.Barony {
	return c.莫云Moyun
}

func (c *杂多ZadoiCounty) BQapugtang萨呼腾() feud.Barony {
	return c.萨呼腾Qapugtang
}

func (c *杂多ZadoiCounty) BZadoi杂多() feud.Barony {
	return c.杂多Zadoi
}

func (c *杂多ZadoiCounty) BZhaqing扎青() feud.Barony {
	return c.扎青Zhaqing
}

var CZadoi杂多 ZadoiCounty = &杂多ZadoiCounty{}

func init() {
	f := CZadoi杂多.(*杂多ZadoiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1575",
		Title:     "zadoi",
		TitleName: "杂多",
		TitleCode: "c_zadoi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿多Aduo = BAduo阿多
	f.阿多Aduo.SetParent(f)

	f.昂赛Angsai = BAngsai昂赛
	f.昂赛Angsai.SetParent(f)

	f.查旦Chadan = BChadan查旦
	f.查旦Chadan.SetParent(f)

	f.莫云Moyun = BMoyun莫云
	f.莫云Moyun.SetParent(f)

	f.萨呼腾Qapugtang = BQapugtang萨呼腾
	f.萨呼腾Qapugtang.SetParent(f)

	f.杂多Zadoi = BZadoi杂多
	f.杂多Zadoi.SetParent(f)

	f.扎青Zhaqing = BZhaqing扎青
	f.扎青Zhaqing.SetParent(f)

}
