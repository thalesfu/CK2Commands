package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KermanshahCounty interface {
	feud.County
	BHulwan胡拉万() feud.Barony
	BJavanroud贾万鲁德() feud.Barony
	BKangavar坎加瓦尔() feud.Barony
	BKermanshah克尔曼沙赫() feud.Barony
	BKuzaran库泽兰() feud.Barony
	BMahalkufa马赫库发() feud.Barony
	BPaveh帕韦() feud.Barony
	BRavansar拉万萨尔() feud.Barony
}

type 克尔曼沙赫KermanshahCounty struct {
	feud.BaseCounty
	胡拉万Hulwan       feud.Barony
	贾万鲁德Javanroud   feud.Barony
	坎加瓦尔Kangavar    feud.Barony
	克尔曼沙赫Kermanshah feud.Barony
	库泽兰Kuzaran      feud.Barony
	马赫库发Mahalkufa   feud.Barony
	帕韦Paveh         feud.Barony
	拉万萨尔Ravansar    feud.Barony
}

func (c *克尔曼沙赫KermanshahCounty) BHulwan胡拉万() feud.Barony {
	return c.胡拉万Hulwan
}

func (c *克尔曼沙赫KermanshahCounty) BJavanroud贾万鲁德() feud.Barony {
	return c.贾万鲁德Javanroud
}

func (c *克尔曼沙赫KermanshahCounty) BKangavar坎加瓦尔() feud.Barony {
	return c.坎加瓦尔Kangavar
}

func (c *克尔曼沙赫KermanshahCounty) BKermanshah克尔曼沙赫() feud.Barony {
	return c.克尔曼沙赫Kermanshah
}

func (c *克尔曼沙赫KermanshahCounty) BKuzaran库泽兰() feud.Barony {
	return c.库泽兰Kuzaran
}

func (c *克尔曼沙赫KermanshahCounty) BMahalkufa马赫库发() feud.Barony {
	return c.马赫库发Mahalkufa
}

func (c *克尔曼沙赫KermanshahCounty) BPaveh帕韦() feud.Barony {
	return c.帕韦Paveh
}

func (c *克尔曼沙赫KermanshahCounty) BRavansar拉万萨尔() feud.Barony {
	return c.拉万萨尔Ravansar
}

var CKermanshah克尔曼沙赫 KermanshahCounty = &克尔曼沙赫KermanshahCounty{}

func init() {
	f := CKermanshah克尔曼沙赫.(*克尔曼沙赫KermanshahCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "688",
		Title:     "kermanshah",
		TitleName: "克尔曼沙赫",
		TitleCode: "c_kermanshah",
		Baronies:  map[string]feud.Barony{},
	}

	f.胡拉万Hulwan = BHulwan胡拉万
	f.胡拉万Hulwan.SetParent(f)

	f.贾万鲁德Javanroud = BJavanroud贾万鲁德
	f.贾万鲁德Javanroud.SetParent(f)

	f.坎加瓦尔Kangavar = BKangavar坎加瓦尔
	f.坎加瓦尔Kangavar.SetParent(f)

	f.克尔曼沙赫Kermanshah = BKermanshah克尔曼沙赫
	f.克尔曼沙赫Kermanshah.SetParent(f)

	f.库泽兰Kuzaran = BKuzaran库泽兰
	f.库泽兰Kuzaran.SetParent(f)

	f.马赫库发Mahalkufa = BMahalkufa马赫库发
	f.马赫库发Mahalkufa.SetParent(f)

	f.帕韦Paveh = BPaveh帕韦
	f.帕韦Paveh.SetParent(f)

	f.拉万萨尔Ravansar = BRavansar拉万萨尔
	f.拉万萨尔Ravansar.SetParent(f)

}
