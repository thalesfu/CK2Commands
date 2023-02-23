package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DeltaCounty interface {
	feud.County
	BBaramun巴拉穆恩() feud.Barony
	BBurah巴甫洛() feud.Barony
	BBurlus布鲁卢斯() feud.Barony
	BDamietta达米埃塔() feud.Barony
	BFareskur法里斯库尔() feud.Barony
	BSaramsah塞慕斯() feud.Barony
	BShirbin希尔宾() feud.Barony
	BTanis塔尼斯() feud.Barony
	BXois索伊斯() feud.Barony
}

type 三角洲DeltaCounty struct {
	feud.BaseCounty
	巴拉穆恩Baramun   feud.Barony
	巴甫洛Burah      feud.Barony
	布鲁卢斯Burlus    feud.Barony
	达米埃塔Damietta  feud.Barony
	法里斯库尔Fareskur feud.Barony
	塞慕斯Saramsah   feud.Barony
	希尔宾Shirbin    feud.Barony
	塔尼斯Tanis      feud.Barony
	索伊斯Xois       feud.Barony
}

func (c *三角洲DeltaCounty) BBaramun巴拉穆恩() feud.Barony {
	return c.巴拉穆恩Baramun
}

func (c *三角洲DeltaCounty) BBurah巴甫洛() feud.Barony {
	return c.巴甫洛Burah
}

func (c *三角洲DeltaCounty) BBurlus布鲁卢斯() feud.Barony {
	return c.布鲁卢斯Burlus
}

func (c *三角洲DeltaCounty) BDamietta达米埃塔() feud.Barony {
	return c.达米埃塔Damietta
}

func (c *三角洲DeltaCounty) BFareskur法里斯库尔() feud.Barony {
	return c.法里斯库尔Fareskur
}

func (c *三角洲DeltaCounty) BSaramsah塞慕斯() feud.Barony {
	return c.塞慕斯Saramsah
}

func (c *三角洲DeltaCounty) BShirbin希尔宾() feud.Barony {
	return c.希尔宾Shirbin
}

func (c *三角洲DeltaCounty) BTanis塔尼斯() feud.Barony {
	return c.塔尼斯Tanis
}

func (c *三角洲DeltaCounty) BXois索伊斯() feud.Barony {
	return c.索伊斯Xois
}

var CDelta三角洲 DeltaCounty = &三角洲DeltaCounty{}

func init() {
	f := CDelta三角洲.(*三角洲DeltaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "798",
		Title:     "delta",
		TitleName: "三角洲",
		TitleCode: "c_delta",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉穆恩Baramun = BBaramun巴拉穆恩
	f.巴拉穆恩Baramun.SetParent(f)

	f.巴甫洛Burah = BBurah巴甫洛
	f.巴甫洛Burah.SetParent(f)

	f.布鲁卢斯Burlus = BBurlus布鲁卢斯
	f.布鲁卢斯Burlus.SetParent(f)

	f.达米埃塔Damietta = BDamietta达米埃塔
	f.达米埃塔Damietta.SetParent(f)

	f.法里斯库尔Fareskur = BFareskur法里斯库尔
	f.法里斯库尔Fareskur.SetParent(f)

	f.塞慕斯Saramsah = BSaramsah塞慕斯
	f.塞慕斯Saramsah.SetParent(f)

	f.希尔宾Shirbin = BShirbin希尔宾
	f.希尔宾Shirbin.SetParent(f)

	f.塔尼斯Tanis = BTanis塔尼斯
	f.塔尼斯Tanis.SetParent(f)

	f.索伊斯Xois = BXois索伊斯
	f.索伊斯Xois.SetParent(f)

}
