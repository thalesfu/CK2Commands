package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FaramaCounty interface {
	feud.County
	BBirelabd阿卜德井() feud.Barony
	BBirqatia盖提耶井() feud.Barony
	BFarama珀卢西亚() feud.Barony
	BMustabiq穆斯塔比克() feud.Barony
	BRomani罗马尼() feud.Barony
	BSeyan瑟燕() feud.Barony
	BSin训() feud.Barony
	BZaaraniq扎拉尼克() feud.Barony
}

type 珀卢西亚FaramaCounty struct {
	feud.BaseCounty
	阿卜德井Birelabd  feud.Barony
	盖提耶井Birqatia  feud.Barony
	珀卢西亚Farama    feud.Barony
	穆斯塔比克Mustabiq feud.Barony
	罗马尼Romani     feud.Barony
	瑟燕Seyan       feud.Barony
	训Sin          feud.Barony
	扎拉尼克Zaaraniq  feud.Barony
}

func (c *珀卢西亚FaramaCounty) BBirelabd阿卜德井() feud.Barony {
	return c.阿卜德井Birelabd
}

func (c *珀卢西亚FaramaCounty) BBirqatia盖提耶井() feud.Barony {
	return c.盖提耶井Birqatia
}

func (c *珀卢西亚FaramaCounty) BFarama珀卢西亚() feud.Barony {
	return c.珀卢西亚Farama
}

func (c *珀卢西亚FaramaCounty) BMustabiq穆斯塔比克() feud.Barony {
	return c.穆斯塔比克Mustabiq
}

func (c *珀卢西亚FaramaCounty) BRomani罗马尼() feud.Barony {
	return c.罗马尼Romani
}

func (c *珀卢西亚FaramaCounty) BSeyan瑟燕() feud.Barony {
	return c.瑟燕Seyan
}

func (c *珀卢西亚FaramaCounty) BSin训() feud.Barony {
	return c.训Sin
}

func (c *珀卢西亚FaramaCounty) BZaaraniq扎拉尼克() feud.Barony {
	return c.扎拉尼克Zaaraniq
}

var CFarama珀卢西亚 FaramaCounty = &珀卢西亚FaramaCounty{}

func init() {
	f := CFarama珀卢西亚.(*珀卢西亚FaramaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "788",
		Title:     "farama",
		TitleName: "珀卢西亚",
		TitleCode: "c_farama",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卜德井Birelabd = BBirelabd阿卜德井
	f.阿卜德井Birelabd.SetParent(f)

	f.盖提耶井Birqatia = BBirqatia盖提耶井
	f.盖提耶井Birqatia.SetParent(f)

	f.珀卢西亚Farama = BFarama珀卢西亚
	f.珀卢西亚Farama.SetParent(f)

	f.穆斯塔比克Mustabiq = BMustabiq穆斯塔比克
	f.穆斯塔比克Mustabiq.SetParent(f)

	f.罗马尼Romani = BRomani罗马尼
	f.罗马尼Romani.SetParent(f)

	f.瑟燕Seyan = BSeyan瑟燕
	f.瑟燕Seyan.SetParent(f)

	f.训Sin = BSin训
	f.训Sin.SetParent(f)

	f.扎拉尼克Zaaraniq = BZaaraniq扎拉尼克
	f.扎拉尼克Zaaraniq.SetParent(f)

}
