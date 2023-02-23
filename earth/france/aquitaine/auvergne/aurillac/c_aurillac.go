package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AurillacCounty interface {
	feud.County
	BAurillac欧里亚克() feud.Barony
	BCarlat卡尔拉() feud.Barony
	BIssoire伊苏瓦尔() feud.Barony
	BMauriac莫里亚克() feud.Barony
	BMurat米拉() feud.Barony
}

type 欧里亚克AurillacCounty struct {
	feud.BaseCounty
	欧里亚克Aurillac feud.Barony
	卡尔拉Carlat    feud.Barony
	伊苏瓦尔Issoire  feud.Barony
	莫里亚克Mauriac  feud.Barony
	米拉Murat      feud.Barony
}

func (c *欧里亚克AurillacCounty) BAurillac欧里亚克() feud.Barony {
	return c.欧里亚克Aurillac
}

func (c *欧里亚克AurillacCounty) BCarlat卡尔拉() feud.Barony {
	return c.卡尔拉Carlat
}

func (c *欧里亚克AurillacCounty) BIssoire伊苏瓦尔() feud.Barony {
	return c.伊苏瓦尔Issoire
}

func (c *欧里亚克AurillacCounty) BMauriac莫里亚克() feud.Barony {
	return c.莫里亚克Mauriac
}

func (c *欧里亚克AurillacCounty) BMurat米拉() feud.Barony {
	return c.米拉Murat
}

var CAurillac欧里亚克 AurillacCounty = &欧里亚克AurillacCounty{}

func init() {
	f := CAurillac欧里亚克.(*欧里亚克AurillacCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1966",
		Title:     "aurillac",
		TitleName: "欧里亚克",
		TitleCode: "c_aurillac",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧里亚克Aurillac = BAurillac欧里亚克
	f.欧里亚克Aurillac.SetParent(f)

	f.卡尔拉Carlat = BCarlat卡尔拉
	f.卡尔拉Carlat.SetParent(f)

	f.伊苏瓦尔Issoire = BIssoire伊苏瓦尔
	f.伊苏瓦尔Issoire.SetParent(f)

	f.莫里亚克Mauriac = BMauriac莫里亚克
	f.莫里亚克Mauriac.SetParent(f)

	f.米拉Murat = BMurat米拉
	f.米拉Murat.SetParent(f)

}
