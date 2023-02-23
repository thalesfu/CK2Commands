package sauvira

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/sauvira/debul"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/sauvira/mansura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/sauvira/ranikot"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/sauvira/siwistan"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/sauvira/sonda"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SauviraDuke interface {
	feud.Duke
	CDebul提部罗() debul.DebulCounty
	CMansura曼苏拉() mansura.MansuraCounty
	CRanikot兰尼科特() ranikot.RanikotCounty
	CSiwistan室毗湿檀那() siwistan.SiwistanCounty
	CSonda孙陀() sonda.SondaCounty
}

type 粟毗罗SauviraDuke struct {
	feud.BaseDuke
	提部罗Debul      debul.DebulCounty
	曼苏拉Mansura    mansura.MansuraCounty
	兰尼科特Ranikot   ranikot.RanikotCounty
	室毗湿檀那Siwistan siwistan.SiwistanCounty
	孙陀Sonda       sonda.SondaCounty
}

func (d *粟毗罗SauviraDuke) CDebul提部罗() debul.DebulCounty {
	return d.提部罗Debul
}

func (d *粟毗罗SauviraDuke) CMansura曼苏拉() mansura.MansuraCounty {
	return d.曼苏拉Mansura
}

func (d *粟毗罗SauviraDuke) CRanikot兰尼科特() ranikot.RanikotCounty {
	return d.兰尼科特Ranikot
}

func (d *粟毗罗SauviraDuke) CSiwistan室毗湿檀那() siwistan.SiwistanCounty {
	return d.室毗湿檀那Siwistan
}

func (d *粟毗罗SauviraDuke) CSonda孙陀() sonda.SondaCounty {
	return d.孙陀Sonda
}

var DSauvira粟毗罗 SauviraDuke = &粟毗罗SauviraDuke{}

func init() {
	f := DSauvira粟毗罗.(*粟毗罗SauviraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sauvira",
		TitleName: "粟毗罗",
		TitleCode: "d_sauvira",
		Counties:  map[string]feud.County{},
	}

	f.提部罗Debul = debul.CDebul提部罗
	f.提部罗Debul.SetParent(f)

	f.曼苏拉Mansura = mansura.CMansura曼苏拉
	f.曼苏拉Mansura.SetParent(f)

	f.兰尼科特Ranikot = ranikot.CRanikot兰尼科特
	f.兰尼科特Ranikot.SetParent(f)

	f.室毗湿檀那Siwistan = siwistan.CSiwistan室毗湿檀那
	f.室毗湿檀那Siwistan.SetParent(f)

	f.孙陀Sonda = sonda.CSonda孙陀
	f.孙陀Sonda.SetParent(f)

}
