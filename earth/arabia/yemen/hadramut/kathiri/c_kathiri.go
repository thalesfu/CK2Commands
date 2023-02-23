package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KathiriCounty interface {
	feud.County
	BAzzan阿赞() feud.Barony
	BLodar劳代尔() feud.Barony
	BMukalla穆卡拉() feud.Barony
	BNisab尼萨卜() feud.Barony
	BQana喀纳() feud.Barony
	BSeiyun赛永() feud.Barony
	BShabwa夏卜瓦() feud.Barony
	BShihar希赫尔() feud.Barony
}

type 卡迪里KathiriCounty struct {
	feud.BaseCounty
	阿赞Azzan    feud.Barony
	劳代尔Lodar   feud.Barony
	穆卡拉Mukalla feud.Barony
	尼萨卜Nisab   feud.Barony
	喀纳Qana     feud.Barony
	赛永Seiyun   feud.Barony
	夏卜瓦Shabwa  feud.Barony
	希赫尔Shihar  feud.Barony
}

func (c *卡迪里KathiriCounty) BAzzan阿赞() feud.Barony {
	return c.阿赞Azzan
}

func (c *卡迪里KathiriCounty) BLodar劳代尔() feud.Barony {
	return c.劳代尔Lodar
}

func (c *卡迪里KathiriCounty) BMukalla穆卡拉() feud.Barony {
	return c.穆卡拉Mukalla
}

func (c *卡迪里KathiriCounty) BNisab尼萨卜() feud.Barony {
	return c.尼萨卜Nisab
}

func (c *卡迪里KathiriCounty) BQana喀纳() feud.Barony {
	return c.喀纳Qana
}

func (c *卡迪里KathiriCounty) BSeiyun赛永() feud.Barony {
	return c.赛永Seiyun
}

func (c *卡迪里KathiriCounty) BShabwa夏卜瓦() feud.Barony {
	return c.夏卜瓦Shabwa
}

func (c *卡迪里KathiriCounty) BShihar希赫尔() feud.Barony {
	return c.希赫尔Shihar
}

var CKathiri卡迪里 KathiriCounty = &卡迪里KathiriCounty{}

func init() {
	f := CKathiri卡迪里.(*卡迪里KathiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "856",
		Title:     "kathiri",
		TitleName: "卡迪里",
		TitleCode: "c_kathiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赞Azzan = BAzzan阿赞
	f.阿赞Azzan.SetParent(f)

	f.劳代尔Lodar = BLodar劳代尔
	f.劳代尔Lodar.SetParent(f)

	f.穆卡拉Mukalla = BMukalla穆卡拉
	f.穆卡拉Mukalla.SetParent(f)

	f.尼萨卜Nisab = BNisab尼萨卜
	f.尼萨卜Nisab.SetParent(f)

	f.喀纳Qana = BQana喀纳
	f.喀纳Qana.SetParent(f)

	f.赛永Seiyun = BSeiyun赛永
	f.赛永Seiyun.SetParent(f)

	f.夏卜瓦Shabwa = BShabwa夏卜瓦
	f.夏卜瓦Shabwa.SetParent(f)

	f.希赫尔Shihar = BShihar希赫尔
	f.希赫尔Shihar.SetParent(f)

}
