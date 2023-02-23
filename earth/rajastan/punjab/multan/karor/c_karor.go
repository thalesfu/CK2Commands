package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarorCounty interface {
	feud.County
	BKaror迦卢罗() feud.Barony
	BMankera门盖拉() feud.Barony
	BMurwa牟瓦() feud.Barony
	BNaithra奈多罗() feud.Barony
	BNamchang南旃俱() feud.Barony
	BShadia沙迪亚() feud.Barony
}

type 迦卢罗KarorCounty struct {
	feud.BaseCounty
	迦卢罗Karor    feud.Barony
	门盖拉Mankera  feud.Barony
	牟瓦Murwa     feud.Barony
	奈多罗Naithra  feud.Barony
	南旃俱Namchang feud.Barony
	沙迪亚Shadia   feud.Barony
}

func (c *迦卢罗KarorCounty) BKaror迦卢罗() feud.Barony {
	return c.迦卢罗Karor
}

func (c *迦卢罗KarorCounty) BMankera门盖拉() feud.Barony {
	return c.门盖拉Mankera
}

func (c *迦卢罗KarorCounty) BMurwa牟瓦() feud.Barony {
	return c.牟瓦Murwa
}

func (c *迦卢罗KarorCounty) BNaithra奈多罗() feud.Barony {
	return c.奈多罗Naithra
}

func (c *迦卢罗KarorCounty) BNamchang南旃俱() feud.Barony {
	return c.南旃俱Namchang
}

func (c *迦卢罗KarorCounty) BShadia沙迪亚() feud.Barony {
	return c.沙迪亚Shadia
}

var CKaror迦卢罗 KarorCounty = &迦卢罗KarorCounty{}

func init() {
	f := CKaror迦卢罗.(*迦卢罗KarorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1340",
		Title:     "karor",
		TitleName: "迦卢罗",
		TitleCode: "c_karor",
		Baronies:  map[string]feud.Barony{},
	}

	f.迦卢罗Karor = BKaror迦卢罗
	f.迦卢罗Karor.SetParent(f)

	f.门盖拉Mankera = BMankera门盖拉
	f.门盖拉Mankera.SetParent(f)

	f.牟瓦Murwa = BMurwa牟瓦
	f.牟瓦Murwa.SetParent(f)

	f.奈多罗Naithra = BNaithra奈多罗
	f.奈多罗Naithra.SetParent(f)

	f.南旃俱Namchang = BNamchang南旃俱
	f.南旃俱Namchang.SetParent(f)

	f.沙迪亚Shadia = BShadia沙迪亚
	f.沙迪亚Shadia.SetParent(f)

}
