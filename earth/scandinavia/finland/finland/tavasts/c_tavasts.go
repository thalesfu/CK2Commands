package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TavastsCounty interface {
	feud.County
	BHaikonen海科宁() feud.Barony
	BHameenlinna海门林纳() feud.Barony
	BHarviala哈尔维亚拉() feud.Barony
	BHattula哈图拉() feud.Barony
	BLahti拉赫蒂() feud.Barony
	BMattila马蒂拉() feud.Barony
	BVanaja瓦纳亚() feud.Barony
	BVesilahti韦西拉赫蒂() feud.Barony
}

type 塔瓦斯蒂亚TavastsCounty struct {
	feud.BaseCounty
	海科宁Haikonen     feud.Barony
	海门林纳Hameenlinna feud.Barony
	哈尔维亚拉Harviala   feud.Barony
	哈图拉Hattula      feud.Barony
	拉赫蒂Lahti        feud.Barony
	马蒂拉Mattila      feud.Barony
	瓦纳亚Vanaja       feud.Barony
	韦西拉赫蒂Vesilahti  feud.Barony
}

func (c *塔瓦斯蒂亚TavastsCounty) BHaikonen海科宁() feud.Barony {
	return c.海科宁Haikonen
}

func (c *塔瓦斯蒂亚TavastsCounty) BHameenlinna海门林纳() feud.Barony {
	return c.海门林纳Hameenlinna
}

func (c *塔瓦斯蒂亚TavastsCounty) BHarviala哈尔维亚拉() feud.Barony {
	return c.哈尔维亚拉Harviala
}

func (c *塔瓦斯蒂亚TavastsCounty) BHattula哈图拉() feud.Barony {
	return c.哈图拉Hattula
}

func (c *塔瓦斯蒂亚TavastsCounty) BLahti拉赫蒂() feud.Barony {
	return c.拉赫蒂Lahti
}

func (c *塔瓦斯蒂亚TavastsCounty) BMattila马蒂拉() feud.Barony {
	return c.马蒂拉Mattila
}

func (c *塔瓦斯蒂亚TavastsCounty) BVanaja瓦纳亚() feud.Barony {
	return c.瓦纳亚Vanaja
}

func (c *塔瓦斯蒂亚TavastsCounty) BVesilahti韦西拉赫蒂() feud.Barony {
	return c.韦西拉赫蒂Vesilahti
}

var CTavasts塔瓦斯蒂亚 TavastsCounty = &塔瓦斯蒂亚TavastsCounty{}

func init() {
	f := CTavasts塔瓦斯蒂亚.(*塔瓦斯蒂亚TavastsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "383",
		Title:     "tavasts",
		TitleName: "塔瓦斯蒂亚",
		TitleCode: "c_tavasts",
		Baronies:  map[string]feud.Barony{},
	}

	f.海科宁Haikonen = BHaikonen海科宁
	f.海科宁Haikonen.SetParent(f)

	f.海门林纳Hameenlinna = BHameenlinna海门林纳
	f.海门林纳Hameenlinna.SetParent(f)

	f.哈尔维亚拉Harviala = BHarviala哈尔维亚拉
	f.哈尔维亚拉Harviala.SetParent(f)

	f.哈图拉Hattula = BHattula哈图拉
	f.哈图拉Hattula.SetParent(f)

	f.拉赫蒂Lahti = BLahti拉赫蒂
	f.拉赫蒂Lahti.SetParent(f)

	f.马蒂拉Mattila = BMattila马蒂拉
	f.马蒂拉Mattila.SetParent(f)

	f.瓦纳亚Vanaja = BVanaja瓦纳亚
	f.瓦纳亚Vanaja.SetParent(f)

	f.韦西拉赫蒂Vesilahti = BVesilahti韦西拉赫蒂
	f.韦西拉赫蒂Vesilahti.SetParent(f)

}
