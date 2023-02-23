package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShetlandCounty interface {
	feud.County
	BCunningsburgh坎宁斯堡() feud.Barony
	BMuness穆尼斯() feud.Barony
	BNorthmavine诺赫马因() feud.Barony
	BScalloway斯卡洛韦() feud.Barony
	BSound桑德() feud.Barony
	BSumburgh萨姆堡() feud.Barony
	BTingwall廷沃尔() feud.Barony
	BYell耶尔() feud.Barony
}

type 设得兰ShetlandCounty struct {
	feud.BaseCounty
	坎宁斯堡Cunningsburgh feud.Barony
	穆尼斯Muness         feud.Barony
	诺赫马因Northmavine   feud.Barony
	斯卡洛韦Scalloway     feud.Barony
	桑德Sound           feud.Barony
	萨姆堡Sumburgh       feud.Barony
	廷沃尔Tingwall       feud.Barony
	耶尔Yell            feud.Barony
}

func (c *设得兰ShetlandCounty) BCunningsburgh坎宁斯堡() feud.Barony {
	return c.坎宁斯堡Cunningsburgh
}

func (c *设得兰ShetlandCounty) BMuness穆尼斯() feud.Barony {
	return c.穆尼斯Muness
}

func (c *设得兰ShetlandCounty) BNorthmavine诺赫马因() feud.Barony {
	return c.诺赫马因Northmavine
}

func (c *设得兰ShetlandCounty) BScalloway斯卡洛韦() feud.Barony {
	return c.斯卡洛韦Scalloway
}

func (c *设得兰ShetlandCounty) BSound桑德() feud.Barony {
	return c.桑德Sound
}

func (c *设得兰ShetlandCounty) BSumburgh萨姆堡() feud.Barony {
	return c.萨姆堡Sumburgh
}

func (c *设得兰ShetlandCounty) BTingwall廷沃尔() feud.Barony {
	return c.廷沃尔Tingwall
}

func (c *设得兰ShetlandCounty) BYell耶尔() feud.Barony {
	return c.耶尔Yell
}

var CShetland设得兰 ShetlandCounty = &设得兰ShetlandCounty{}

func init() {
	f := CShetland设得兰.(*设得兰ShetlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "34",
		Title:     "shetland",
		TitleName: "设得兰",
		TitleCode: "c_shetland",
		Baronies:  map[string]feud.Barony{},
	}

	f.坎宁斯堡Cunningsburgh = BCunningsburgh坎宁斯堡
	f.坎宁斯堡Cunningsburgh.SetParent(f)

	f.穆尼斯Muness = BMuness穆尼斯
	f.穆尼斯Muness.SetParent(f)

	f.诺赫马因Northmavine = BNorthmavine诺赫马因
	f.诺赫马因Northmavine.SetParent(f)

	f.斯卡洛韦Scalloway = BScalloway斯卡洛韦
	f.斯卡洛韦Scalloway.SetParent(f)

	f.桑德Sound = BSound桑德
	f.桑德Sound.SetParent(f)

	f.萨姆堡Sumburgh = BSumburgh萨姆堡
	f.萨姆堡Sumburgh.SetParent(f)

	f.廷沃尔Tingwall = BTingwall廷沃尔
	f.廷沃尔Tingwall.SetParent(f)

	f.耶尔Yell = BYell耶尔
	f.耶尔Yell.SetParent(f)

}
