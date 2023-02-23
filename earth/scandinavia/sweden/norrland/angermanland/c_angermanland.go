package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AngermanlandCounty interface {
	feud.County
	BBiedsta别德斯塔() feud.Barony
	BBjartra比耶尔特罗() feud.Barony
	BHarnosand海讷桑德() feud.Barony
	BNatra涅特拉() feud.Barony
	BOrnskoldsvik恩舍尔兹维克() feud.Barony
	BSkulnas斯库涅斯() feud.Barony
	BSolleftea索莱夫特奥() feud.Barony
	BUlvo乌尔沃() feud.Barony
}

type 翁厄曼兰AngermanlandCounty struct {
	feud.BaseCounty
	别德斯塔Biedsta        feud.Barony
	比耶尔特罗Bjartra       feud.Barony
	海讷桑德Harnosand      feud.Barony
	涅特拉Natra           feud.Barony
	恩舍尔兹维克Ornskoldsvik feud.Barony
	斯库涅斯Skulnas        feud.Barony
	索莱夫特奥Solleftea     feud.Barony
	乌尔沃Ulvo            feud.Barony
}

func (c *翁厄曼兰AngermanlandCounty) BBiedsta别德斯塔() feud.Barony {
	return c.别德斯塔Biedsta
}

func (c *翁厄曼兰AngermanlandCounty) BBjartra比耶尔特罗() feud.Barony {
	return c.比耶尔特罗Bjartra
}

func (c *翁厄曼兰AngermanlandCounty) BHarnosand海讷桑德() feud.Barony {
	return c.海讷桑德Harnosand
}

func (c *翁厄曼兰AngermanlandCounty) BNatra涅特拉() feud.Barony {
	return c.涅特拉Natra
}

func (c *翁厄曼兰AngermanlandCounty) BOrnskoldsvik恩舍尔兹维克() feud.Barony {
	return c.恩舍尔兹维克Ornskoldsvik
}

func (c *翁厄曼兰AngermanlandCounty) BSkulnas斯库涅斯() feud.Barony {
	return c.斯库涅斯Skulnas
}

func (c *翁厄曼兰AngermanlandCounty) BSolleftea索莱夫特奥() feud.Barony {
	return c.索莱夫特奥Solleftea
}

func (c *翁厄曼兰AngermanlandCounty) BUlvo乌尔沃() feud.Barony {
	return c.乌尔沃Ulvo
}

var CAngermanland翁厄曼兰 AngermanlandCounty = &翁厄曼兰AngermanlandCounty{}

func init() {
	f := CAngermanland翁厄曼兰.(*翁厄曼兰AngermanlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "281",
		Title:     "angermanland",
		TitleName: "翁厄曼兰",
		TitleCode: "c_angermanland",
		Baronies:  map[string]feud.Barony{},
	}

	f.别德斯塔Biedsta = BBiedsta别德斯塔
	f.别德斯塔Biedsta.SetParent(f)

	f.比耶尔特罗Bjartra = BBjartra比耶尔特罗
	f.比耶尔特罗Bjartra.SetParent(f)

	f.海讷桑德Harnosand = BHarnosand海讷桑德
	f.海讷桑德Harnosand.SetParent(f)

	f.涅特拉Natra = BNatra涅特拉
	f.涅特拉Natra.SetParent(f)

	f.恩舍尔兹维克Ornskoldsvik = BOrnskoldsvik恩舍尔兹维克
	f.恩舍尔兹维克Ornskoldsvik.SetParent(f)

	f.斯库涅斯Skulnas = BSkulnas斯库涅斯
	f.斯库涅斯Skulnas.SetParent(f)

	f.索莱夫特奥Solleftea = BSolleftea索莱夫特奥
	f.索莱夫特奥Solleftea.SetParent(f)

	f.乌尔沃Ulvo = BUlvo乌尔沃
	f.乌尔沃Ulvo.SetParent(f)

}
