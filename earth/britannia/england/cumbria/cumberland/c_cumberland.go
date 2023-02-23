package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CumberlandCounty interface {
	feud.County
	BBurgh巴勒() feud.Barony
	BCarlisle卡莱尔() feud.Barony
	BCockermouth科克茅斯() feud.Barony
	BDacre戴克() feud.Barony
	BEgremont埃格勒蒙特() feud.Barony
	BGilsland吉尔斯兰() feud.Barony
	BPapcastlet帕普卡斯特尔() feud.Barony
	BPenrith彭里斯() feud.Barony
}

type 坎伯兰CumberlandCounty struct {
	feud.BaseCounty
	巴勒Burgh          feud.Barony
	卡莱尔Carlisle      feud.Barony
	科克茅斯Cockermouth  feud.Barony
	戴克Dacre          feud.Barony
	埃格勒蒙特Egremont    feud.Barony
	吉尔斯兰Gilsland     feud.Barony
	帕普卡斯特尔Papcastlet feud.Barony
	彭里斯Penrith       feud.Barony
}

func (c *坎伯兰CumberlandCounty) BBurgh巴勒() feud.Barony {
	return c.巴勒Burgh
}

func (c *坎伯兰CumberlandCounty) BCarlisle卡莱尔() feud.Barony {
	return c.卡莱尔Carlisle
}

func (c *坎伯兰CumberlandCounty) BCockermouth科克茅斯() feud.Barony {
	return c.科克茅斯Cockermouth
}

func (c *坎伯兰CumberlandCounty) BDacre戴克() feud.Barony {
	return c.戴克Dacre
}

func (c *坎伯兰CumberlandCounty) BEgremont埃格勒蒙特() feud.Barony {
	return c.埃格勒蒙特Egremont
}

func (c *坎伯兰CumberlandCounty) BGilsland吉尔斯兰() feud.Barony {
	return c.吉尔斯兰Gilsland
}

func (c *坎伯兰CumberlandCounty) BPapcastlet帕普卡斯特尔() feud.Barony {
	return c.帕普卡斯特尔Papcastlet
}

func (c *坎伯兰CumberlandCounty) BPenrith彭里斯() feud.Barony {
	return c.彭里斯Penrith
}

var CCumberland坎伯兰 CumberlandCounty = &坎伯兰CumberlandCounty{}

func init() {
	f := CCumberland坎伯兰.(*坎伯兰CumberlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "53",
		Title:     "cumberland",
		TitleName: "坎伯兰",
		TitleCode: "c_cumberland",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴勒Burgh = BBurgh巴勒
	f.巴勒Burgh.SetParent(f)

	f.卡莱尔Carlisle = BCarlisle卡莱尔
	f.卡莱尔Carlisle.SetParent(f)

	f.科克茅斯Cockermouth = BCockermouth科克茅斯
	f.科克茅斯Cockermouth.SetParent(f)

	f.戴克Dacre = BDacre戴克
	f.戴克Dacre.SetParent(f)

	f.埃格勒蒙特Egremont = BEgremont埃格勒蒙特
	f.埃格勒蒙特Egremont.SetParent(f)

	f.吉尔斯兰Gilsland = BGilsland吉尔斯兰
	f.吉尔斯兰Gilsland.SetParent(f)

	f.帕普卡斯特尔Papcastlet = BPapcastlet帕普卡斯特尔
	f.帕普卡斯特尔Papcastlet.SetParent(f)

	f.彭里斯Penrith = BPenrith彭里斯
	f.彭里斯Penrith.SetParent(f)

}
