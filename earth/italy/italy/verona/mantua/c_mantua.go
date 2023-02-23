package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MantuaCounty interface {
	feud.County
	BBagnolosanvito巴尼奥洛圣维托() feud.Barony
	BCastellucchio卡斯泰卢基奥() feud.Barony
	BCurtatone库尔塔托内() feud.Barony
	BGonzaga贡扎加() feud.Barony
	BMantua曼图阿() feud.Barony
	BMarmirolo马尔米罗洛() feud.Barony
	BSerravalle塞拉瓦莱() feud.Barony
	BVirgilio维尔吉利奥() feud.Barony
}

type 曼图阿MantuaCounty struct {
	feud.BaseCounty
	巴尼奥洛圣维托Bagnolosanvito feud.Barony
	卡斯泰卢基奥Castellucchio   feud.Barony
	库尔塔托内Curtatone        feud.Barony
	贡扎加Gonzaga            feud.Barony
	曼图阿Mantua             feud.Barony
	马尔米罗洛Marmirolo        feud.Barony
	塞拉瓦莱Serravalle        feud.Barony
	维尔吉利奥Virgilio         feud.Barony
}

func (c *曼图阿MantuaCounty) BBagnolosanvito巴尼奥洛圣维托() feud.Barony {
	return c.巴尼奥洛圣维托Bagnolosanvito
}

func (c *曼图阿MantuaCounty) BCastellucchio卡斯泰卢基奥() feud.Barony {
	return c.卡斯泰卢基奥Castellucchio
}

func (c *曼图阿MantuaCounty) BCurtatone库尔塔托内() feud.Barony {
	return c.库尔塔托内Curtatone
}

func (c *曼图阿MantuaCounty) BGonzaga贡扎加() feud.Barony {
	return c.贡扎加Gonzaga
}

func (c *曼图阿MantuaCounty) BMantua曼图阿() feud.Barony {
	return c.曼图阿Mantua
}

func (c *曼图阿MantuaCounty) BMarmirolo马尔米罗洛() feud.Barony {
	return c.马尔米罗洛Marmirolo
}

func (c *曼图阿MantuaCounty) BSerravalle塞拉瓦莱() feud.Barony {
	return c.塞拉瓦莱Serravalle
}

func (c *曼图阿MantuaCounty) BVirgilio维尔吉利奥() feud.Barony {
	return c.维尔吉利奥Virgilio
}

var CMantua曼图阿 MantuaCounty = &曼图阿MantuaCounty{}

func init() {
	f := CMantua曼图阿.(*曼图阿MantuaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "354",
		Title:     "mantua",
		TitleName: "曼图阿",
		TitleCode: "c_mantua",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尼奥洛圣维托Bagnolosanvito = BBagnolosanvito巴尼奥洛圣维托
	f.巴尼奥洛圣维托Bagnolosanvito.SetParent(f)

	f.卡斯泰卢基奥Castellucchio = BCastellucchio卡斯泰卢基奥
	f.卡斯泰卢基奥Castellucchio.SetParent(f)

	f.库尔塔托内Curtatone = BCurtatone库尔塔托内
	f.库尔塔托内Curtatone.SetParent(f)

	f.贡扎加Gonzaga = BGonzaga贡扎加
	f.贡扎加Gonzaga.SetParent(f)

	f.曼图阿Mantua = BMantua曼图阿
	f.曼图阿Mantua.SetParent(f)

	f.马尔米罗洛Marmirolo = BMarmirolo马尔米罗洛
	f.马尔米罗洛Marmirolo.SetParent(f)

	f.塞拉瓦莱Serravalle = BSerravalle塞拉瓦莱
	f.塞拉瓦莱Serravalle.SetParent(f)

	f.维尔吉利奥Virgilio = BVirgilio维尔吉利奥
	f.维尔吉利奥Virgilio.SetParent(f)

}
