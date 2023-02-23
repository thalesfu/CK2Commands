package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NelloreCounty interface {
	feud.County
	BAddanki阿登吉() feud.Barony
	BGundigapuri郡稚伽补梨() feud.Barony
	BKalahasti迦罗诃悉帝() feud.Barony
	BKandukura坎杜古拉() feud.Barony
	BMutapali穆塔帕里() feud.Barony
	BNellore内卢楼() feud.Barony
	BTripurantakam帝利补兰多迦摩() feud.Barony
}

type 内卢楼NelloreCounty struct {
	feud.BaseCounty
	阿登吉Addanki           feud.Barony
	郡稚伽补梨Gundigapuri     feud.Barony
	迦罗诃悉帝Kalahasti       feud.Barony
	坎杜古拉Kandukura        feud.Barony
	穆塔帕里Mutapali         feud.Barony
	内卢楼Nellore           feud.Barony
	帝利补兰多迦摩Tripurantakam feud.Barony
}

func (c *内卢楼NelloreCounty) BAddanki阿登吉() feud.Barony {
	return c.阿登吉Addanki
}

func (c *内卢楼NelloreCounty) BGundigapuri郡稚伽补梨() feud.Barony {
	return c.郡稚伽补梨Gundigapuri
}

func (c *内卢楼NelloreCounty) BKalahasti迦罗诃悉帝() feud.Barony {
	return c.迦罗诃悉帝Kalahasti
}

func (c *内卢楼NelloreCounty) BKandukura坎杜古拉() feud.Barony {
	return c.坎杜古拉Kandukura
}

func (c *内卢楼NelloreCounty) BMutapali穆塔帕里() feud.Barony {
	return c.穆塔帕里Mutapali
}

func (c *内卢楼NelloreCounty) BNellore内卢楼() feud.Barony {
	return c.内卢楼Nellore
}

func (c *内卢楼NelloreCounty) BTripurantakam帝利补兰多迦摩() feud.Barony {
	return c.帝利补兰多迦摩Tripurantakam
}

var CNellore内卢楼 NelloreCounty = &内卢楼NelloreCounty{}

func init() {
	f := CNellore内卢楼.(*内卢楼NelloreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1204",
		Title:     "nellore",
		TitleName: "内卢楼",
		TitleCode: "c_nellore",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿登吉Addanki = BAddanki阿登吉
	f.阿登吉Addanki.SetParent(f)

	f.郡稚伽补梨Gundigapuri = BGundigapuri郡稚伽补梨
	f.郡稚伽补梨Gundigapuri.SetParent(f)

	f.迦罗诃悉帝Kalahasti = BKalahasti迦罗诃悉帝
	f.迦罗诃悉帝Kalahasti.SetParent(f)

	f.坎杜古拉Kandukura = BKandukura坎杜古拉
	f.坎杜古拉Kandukura.SetParent(f)

	f.穆塔帕里Mutapali = BMutapali穆塔帕里
	f.穆塔帕里Mutapali.SetParent(f)

	f.内卢楼Nellore = BNellore内卢楼
	f.内卢楼Nellore.SetParent(f)

	f.帝利补兰多迦摩Tripurantakam = BTripurantakam帝利补兰多迦摩
	f.帝利补兰多迦摩Tripurantakam.SetParent(f)

}
