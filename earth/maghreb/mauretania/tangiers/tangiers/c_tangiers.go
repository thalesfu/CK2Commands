package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TangiersCounty interface {
	feud.County
	BAinbenamar艾因本阿马尔() feud.Barony
	BAlcazarquivir阿尔卡萨基维尔() feud.Barony
	BAsilah艾西拉() feud.Barony
	BBarrha巴拉哈() feud.Barony
	BCharkia哈拉克亚() feud.Barony
	BLaraiche拉腊什() feud.Barony
	BMulaybuselham穆莱布瑟罕() feud.Barony
	BTangiers丹吉尔() feud.Barony
}

type 丹吉尔TangiersCounty struct {
	feud.BaseCounty
	艾因本阿马尔Ainbenamar     feud.Barony
	阿尔卡萨基维尔Alcazarquivir feud.Barony
	艾西拉Asilah            feud.Barony
	巴拉哈Barrha            feud.Barony
	哈拉克亚Charkia          feud.Barony
	拉腊什Laraiche          feud.Barony
	穆莱布瑟罕Mulaybuselham   feud.Barony
	丹吉尔Tangiers          feud.Barony
}

func (c *丹吉尔TangiersCounty) BAinbenamar艾因本阿马尔() feud.Barony {
	return c.艾因本阿马尔Ainbenamar
}

func (c *丹吉尔TangiersCounty) BAlcazarquivir阿尔卡萨基维尔() feud.Barony {
	return c.阿尔卡萨基维尔Alcazarquivir
}

func (c *丹吉尔TangiersCounty) BAsilah艾西拉() feud.Barony {
	return c.艾西拉Asilah
}

func (c *丹吉尔TangiersCounty) BBarrha巴拉哈() feud.Barony {
	return c.巴拉哈Barrha
}

func (c *丹吉尔TangiersCounty) BCharkia哈拉克亚() feud.Barony {
	return c.哈拉克亚Charkia
}

func (c *丹吉尔TangiersCounty) BLaraiche拉腊什() feud.Barony {
	return c.拉腊什Laraiche
}

func (c *丹吉尔TangiersCounty) BMulaybuselham穆莱布瑟罕() feud.Barony {
	return c.穆莱布瑟罕Mulaybuselham
}

func (c *丹吉尔TangiersCounty) BTangiers丹吉尔() feud.Barony {
	return c.丹吉尔Tangiers
}

var CTangiers丹吉尔 TangiersCounty = &丹吉尔TangiersCounty{}

func init() {
	f := CTangiers丹吉尔.(*丹吉尔TangiersCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "841",
		Title:     "tangiers",
		TitleName: "丹吉尔",
		TitleCode: "c_tangiers",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因本阿马尔Ainbenamar = BAinbenamar艾因本阿马尔
	f.艾因本阿马尔Ainbenamar.SetParent(f)

	f.阿尔卡萨基维尔Alcazarquivir = BAlcazarquivir阿尔卡萨基维尔
	f.阿尔卡萨基维尔Alcazarquivir.SetParent(f)

	f.艾西拉Asilah = BAsilah艾西拉
	f.艾西拉Asilah.SetParent(f)

	f.巴拉哈Barrha = BBarrha巴拉哈
	f.巴拉哈Barrha.SetParent(f)

	f.哈拉克亚Charkia = BCharkia哈拉克亚
	f.哈拉克亚Charkia.SetParent(f)

	f.拉腊什Laraiche = BLaraiche拉腊什
	f.拉腊什Laraiche.SetParent(f)

	f.穆莱布瑟罕Mulaybuselham = BMulaybuselham穆莱布瑟罕
	f.穆莱布瑟罕Mulaybuselham.SetParent(f)

	f.丹吉尔Tangiers = BTangiers丹吉尔
	f.丹吉尔Tangiers.SetParent(f)

}
