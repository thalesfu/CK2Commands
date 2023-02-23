package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BamCounty interface {
	feud.County
	BAbeshkan阿比沙坎() feud.Barony
	BBampur班布尔() feud.Barony
	BBaravat巴拉瓦特() feud.Barony
	BFahraj法赫拉季() feud.Barony
	BKaj卡吉() feud.Barony
	BNukjub努克祖() feud.Barony
}

type 巴姆BamCounty struct {
	feud.BaseCounty
	阿比沙坎Abeshkan feud.Barony
	班布尔Bampur    feud.Barony
	巴拉瓦特Baravat  feud.Barony
	法赫拉季Fahraj   feud.Barony
	卡吉Kaj        feud.Barony
	努克祖Nukjub    feud.Barony
}

func (c *巴姆BamCounty) BAbeshkan阿比沙坎() feud.Barony {
	return c.阿比沙坎Abeshkan
}

func (c *巴姆BamCounty) BBampur班布尔() feud.Barony {
	return c.班布尔Bampur
}

func (c *巴姆BamCounty) BBaravat巴拉瓦特() feud.Barony {
	return c.巴拉瓦特Baravat
}

func (c *巴姆BamCounty) BFahraj法赫拉季() feud.Barony {
	return c.法赫拉季Fahraj
}

func (c *巴姆BamCounty) BKaj卡吉() feud.Barony {
	return c.卡吉Kaj
}

func (c *巴姆BamCounty) BNukjub努克祖() feud.Barony {
	return c.努克祖Nukjub
}

var CBam巴姆 BamCounty = &巴姆BamCounty{}

func init() {
	f := CBam巴姆.(*巴姆BamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "853",
		Title:     "bam",
		TitleName: "巴姆",
		TitleCode: "c_bam",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿比沙坎Abeshkan = BAbeshkan阿比沙坎
	f.阿比沙坎Abeshkan.SetParent(f)

	f.班布尔Bampur = BBampur班布尔
	f.班布尔Bampur.SetParent(f)

	f.巴拉瓦特Baravat = BBaravat巴拉瓦特
	f.巴拉瓦特Baravat.SetParent(f)

	f.法赫拉季Fahraj = BFahraj法赫拉季
	f.法赫拉季Fahraj.SetParent(f)

	f.卡吉Kaj = BKaj卡吉
	f.卡吉Kaj.SetParent(f)

	f.努克祖Nukjub = BNukjub努克祖
	f.努克祖Nukjub.SetParent(f)

}
