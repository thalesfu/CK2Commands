package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AssabCounty interface {
	feud.County
	BAgurto阿古脱() feud.Barony
	BAssab阿萨布() feud.Barony
	BDebaysima德巴伊斯玛() feud.Barony
	BGehare格哈尔() feud.Barony
	BManda曼达() feud.Barony
	BMergebla梅葛布拉() feud.Barony
	BRehayto拉海托() feud.Barony
}

type 阿萨布AssabCounty struct {
	feud.BaseCounty
	阿古脱Agurto      feud.Barony
	阿萨布Assab       feud.Barony
	德巴伊斯玛Debaysima feud.Barony
	格哈尔Gehare      feud.Barony
	曼达Manda        feud.Barony
	梅葛布拉Mergebla   feud.Barony
	拉海托Rehayto     feud.Barony
}

func (c *阿萨布AssabCounty) BAgurto阿古脱() feud.Barony {
	return c.阿古脱Agurto
}

func (c *阿萨布AssabCounty) BAssab阿萨布() feud.Barony {
	return c.阿萨布Assab
}

func (c *阿萨布AssabCounty) BDebaysima德巴伊斯玛() feud.Barony {
	return c.德巴伊斯玛Debaysima
}

func (c *阿萨布AssabCounty) BGehare格哈尔() feud.Barony {
	return c.格哈尔Gehare
}

func (c *阿萨布AssabCounty) BManda曼达() feud.Barony {
	return c.曼达Manda
}

func (c *阿萨布AssabCounty) BMergebla梅葛布拉() feud.Barony {
	return c.梅葛布拉Mergebla
}

func (c *阿萨布AssabCounty) BRehayto拉海托() feud.Barony {
	return c.拉海托Rehayto
}

var CAssab阿萨布 AssabCounty = &阿萨布AssabCounty{}

func init() {
	f := CAssab阿萨布.(*阿萨布AssabCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1373",
		Title:     "assab",
		TitleName: "阿萨布",
		TitleCode: "c_assab",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿古脱Agurto = BAgurto阿古脱
	f.阿古脱Agurto.SetParent(f)

	f.阿萨布Assab = BAssab阿萨布
	f.阿萨布Assab.SetParent(f)

	f.德巴伊斯玛Debaysima = BDebaysima德巴伊斯玛
	f.德巴伊斯玛Debaysima.SetParent(f)

	f.格哈尔Gehare = BGehare格哈尔
	f.格哈尔Gehare.SetParent(f)

	f.曼达Manda = BManda曼达
	f.曼达Manda.SetParent(f)

	f.梅葛布拉Mergebla = BMergebla梅葛布拉
	f.梅葛布拉Mergebla.SetParent(f)

	f.拉海托Rehayto = BRehayto拉海托
	f.拉海托Rehayto.SetParent(f)

}
