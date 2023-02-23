package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ComoCounty interface {
	feud.County
	BAlbizzate阿尔比扎泰() feud.Barony
	BBesnate贝斯纳泰() feud.Barony
	BComo科莫() feud.Barony
	BLecco莱科() feud.Barony
	BSottoceneri索托切内里() feud.Barony
	BValtellina瓦尔泰利纳() feud.Barony
	BVarese瓦雷泽() feud.Barony
}

type 科莫ComoCounty struct {
	feud.BaseCounty
	阿尔比扎泰Albizzate   feud.Barony
	贝斯纳泰Besnate      feud.Barony
	科莫Como           feud.Barony
	莱科Lecco          feud.Barony
	索托切内里Sottoceneri feud.Barony
	瓦尔泰利纳Valtellina  feud.Barony
	瓦雷泽Varese        feud.Barony
}

func (c *科莫ComoCounty) BAlbizzate阿尔比扎泰() feud.Barony {
	return c.阿尔比扎泰Albizzate
}

func (c *科莫ComoCounty) BBesnate贝斯纳泰() feud.Barony {
	return c.贝斯纳泰Besnate
}

func (c *科莫ComoCounty) BComo科莫() feud.Barony {
	return c.科莫Como
}

func (c *科莫ComoCounty) BLecco莱科() feud.Barony {
	return c.莱科Lecco
}

func (c *科莫ComoCounty) BSottoceneri索托切内里() feud.Barony {
	return c.索托切内里Sottoceneri
}

func (c *科莫ComoCounty) BValtellina瓦尔泰利纳() feud.Barony {
	return c.瓦尔泰利纳Valtellina
}

func (c *科莫ComoCounty) BVarese瓦雷泽() feud.Barony {
	return c.瓦雷泽Varese
}

var CComo科莫 ComoCounty = &科莫ComoCounty{}

func init() {
	f := CComo科莫.(*科莫ComoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1709",
		Title:     "como",
		TitleName: "科莫",
		TitleCode: "c_como",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔比扎泰Albizzate = BAlbizzate阿尔比扎泰
	f.阿尔比扎泰Albizzate.SetParent(f)

	f.贝斯纳泰Besnate = BBesnate贝斯纳泰
	f.贝斯纳泰Besnate.SetParent(f)

	f.科莫Como = BComo科莫
	f.科莫Como.SetParent(f)

	f.莱科Lecco = BLecco莱科
	f.莱科Lecco.SetParent(f)

	f.索托切内里Sottoceneri = BSottoceneri索托切内里
	f.索托切内里Sottoceneri.SetParent(f)

	f.瓦尔泰利纳Valtellina = BValtellina瓦尔泰利纳
	f.瓦尔泰利纳Valtellina.SetParent(f)

	f.瓦雷泽Varese = BVarese瓦雷泽
	f.瓦雷泽Varese.SetParent(f)

}
