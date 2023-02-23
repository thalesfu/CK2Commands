package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YamaliaCounty interface {
	feud.County
	BBaygul拜古尔() feud.Barony
	BItuyakha伊图_亚哈() feud.Barony
	BKaek卡耶克() feud.Barony
	BLapytnangk拉贝特南吉() feud.Barony
	BNazym纳济姆() feud.Barony
	BObdorsk鄂毕多尔斯克() feud.Barony
	BPolnovatvozh波尔诺瓦特_沃日() feud.Barony
	BUrengoi乌连戈伊() feud.Barony
}

type 亚马利亚YamaliaCounty struct {
	feud.BaseCounty
	拜古尔Baygul            feud.Barony
	伊图_亚哈Ituyakha        feud.Barony
	卡耶克Kaek              feud.Barony
	拉贝特南吉Lapytnangk      feud.Barony
	纳济姆Nazym             feud.Barony
	鄂毕多尔斯克Obdorsk        feud.Barony
	波尔诺瓦特_沃日Polnovatvozh feud.Barony
	乌连戈伊Urengoi          feud.Barony
}

func (c *亚马利亚YamaliaCounty) BBaygul拜古尔() feud.Barony {
	return c.拜古尔Baygul
}

func (c *亚马利亚YamaliaCounty) BItuyakha伊图_亚哈() feud.Barony {
	return c.伊图_亚哈Ituyakha
}

func (c *亚马利亚YamaliaCounty) BKaek卡耶克() feud.Barony {
	return c.卡耶克Kaek
}

func (c *亚马利亚YamaliaCounty) BLapytnangk拉贝特南吉() feud.Barony {
	return c.拉贝特南吉Lapytnangk
}

func (c *亚马利亚YamaliaCounty) BNazym纳济姆() feud.Barony {
	return c.纳济姆Nazym
}

func (c *亚马利亚YamaliaCounty) BObdorsk鄂毕多尔斯克() feud.Barony {
	return c.鄂毕多尔斯克Obdorsk
}

func (c *亚马利亚YamaliaCounty) BPolnovatvozh波尔诺瓦特_沃日() feud.Barony {
	return c.波尔诺瓦特_沃日Polnovatvozh
}

func (c *亚马利亚YamaliaCounty) BUrengoi乌连戈伊() feud.Barony {
	return c.乌连戈伊Urengoi
}

var CYamalia亚马利亚 YamaliaCounty = &亚马利亚YamaliaCounty{}

func init() {
	f := CYamalia亚马利亚.(*亚马利亚YamaliaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "887",
		Title:     "yamalia",
		TitleName: "亚马利亚",
		TitleCode: "c_yamalia",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜古尔Baygul = BBaygul拜古尔
	f.拜古尔Baygul.SetParent(f)

	f.伊图_亚哈Ituyakha = BItuyakha伊图_亚哈
	f.伊图_亚哈Ituyakha.SetParent(f)

	f.卡耶克Kaek = BKaek卡耶克
	f.卡耶克Kaek.SetParent(f)

	f.拉贝特南吉Lapytnangk = BLapytnangk拉贝特南吉
	f.拉贝特南吉Lapytnangk.SetParent(f)

	f.纳济姆Nazym = BNazym纳济姆
	f.纳济姆Nazym.SetParent(f)

	f.鄂毕多尔斯克Obdorsk = BObdorsk鄂毕多尔斯克
	f.鄂毕多尔斯克Obdorsk.SetParent(f)

	f.波尔诺瓦特_沃日Polnovatvozh = BPolnovatvozh波尔诺瓦特_沃日
	f.波尔诺瓦特_沃日Polnovatvozh.SetParent(f)

	f.乌连戈伊Urengoi = BUrengoi乌连戈伊
	f.乌连戈伊Urengoi.SetParent(f)

}
