package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LivsCounty interface {
	feud.County
	BHaapsalu哈普萨卢() feud.Barony
	BKaruse卡鲁瑟() feud.Barony
	BKullamaa库尔拉玛() feud.Barony
	BLihula利胡拉() feud.Barony
	BLivs利夫斯() feud.Barony
	BMartna马尔特纳() feud.Barony
	BParnu派尔努() feud.Barony
	BPiirsalu皮尔萨鲁() feud.Barony
	BPooravere珀拉瓦尔() feud.Barony
}

type 利沃尼亚LivsCounty struct {
	feud.BaseCounty
	哈普萨卢Haapsalu  feud.Barony
	卡鲁瑟Karuse     feud.Barony
	库尔拉玛Kullamaa  feud.Barony
	利胡拉Lihula     feud.Barony
	利夫斯Livs       feud.Barony
	马尔特纳Martna    feud.Barony
	派尔努Parnu      feud.Barony
	皮尔萨鲁Piirsalu  feud.Barony
	珀拉瓦尔Pooravere feud.Barony
}

func (c *利沃尼亚LivsCounty) BHaapsalu哈普萨卢() feud.Barony {
	return c.哈普萨卢Haapsalu
}

func (c *利沃尼亚LivsCounty) BKaruse卡鲁瑟() feud.Barony {
	return c.卡鲁瑟Karuse
}

func (c *利沃尼亚LivsCounty) BKullamaa库尔拉玛() feud.Barony {
	return c.库尔拉玛Kullamaa
}

func (c *利沃尼亚LivsCounty) BLihula利胡拉() feud.Barony {
	return c.利胡拉Lihula
}

func (c *利沃尼亚LivsCounty) BLivs利夫斯() feud.Barony {
	return c.利夫斯Livs
}

func (c *利沃尼亚LivsCounty) BMartna马尔特纳() feud.Barony {
	return c.马尔特纳Martna
}

func (c *利沃尼亚LivsCounty) BParnu派尔努() feud.Barony {
	return c.派尔努Parnu
}

func (c *利沃尼亚LivsCounty) BPiirsalu皮尔萨鲁() feud.Barony {
	return c.皮尔萨鲁Piirsalu
}

func (c *利沃尼亚LivsCounty) BPooravere珀拉瓦尔() feud.Barony {
	return c.珀拉瓦尔Pooravere
}

var CLivs利沃尼亚 LivsCounty = &利沃尼亚LivsCounty{}

func init() {
	f := CLivs利沃尼亚.(*利沃尼亚LivsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "377",
		Title:     "livs",
		TitleName: "利沃尼亚",
		TitleCode: "c_livs",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈普萨卢Haapsalu = BHaapsalu哈普萨卢
	f.哈普萨卢Haapsalu.SetParent(f)

	f.卡鲁瑟Karuse = BKaruse卡鲁瑟
	f.卡鲁瑟Karuse.SetParent(f)

	f.库尔拉玛Kullamaa = BKullamaa库尔拉玛
	f.库尔拉玛Kullamaa.SetParent(f)

	f.利胡拉Lihula = BLihula利胡拉
	f.利胡拉Lihula.SetParent(f)

	f.利夫斯Livs = BLivs利夫斯
	f.利夫斯Livs.SetParent(f)

	f.马尔特纳Martna = BMartna马尔特纳
	f.马尔特纳Martna.SetParent(f)

	f.派尔努Parnu = BParnu派尔努
	f.派尔努Parnu.SetParent(f)

	f.皮尔萨鲁Piirsalu = BPiirsalu皮尔萨鲁
	f.皮尔萨鲁Piirsalu.SetParent(f)

	f.珀拉瓦尔Pooravere = BPooravere珀拉瓦尔
	f.珀拉瓦尔Pooravere.SetParent(f)

}
