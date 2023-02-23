package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SomnathCounty interface {
	feud.County
	BDelvada提伐陀() feud.Barony
	BDiu刁元() feud.Barony
	BKanjetar甘祇陀() feud.Barony
	BMoherak摩诃罗卡() feud.Barony
	BSomnath苏摩那他() feud.Barony
	BVejalkot吠阇罗拘吒() feud.Barony
	BVeraval韦拉沃尔() feud.Barony
}

type 须门那SomnathCounty struct {
	feud.BaseCounty
	提伐陀Delvada    feud.Barony
	刁元Diu         feud.Barony
	甘祇陀Kanjetar   feud.Barony
	摩诃罗卡Moherak   feud.Barony
	苏摩那他Somnath   feud.Barony
	吠阇罗拘吒Vejalkot feud.Barony
	韦拉沃尔Veraval   feud.Barony
}

func (c *须门那SomnathCounty) BDelvada提伐陀() feud.Barony {
	return c.提伐陀Delvada
}

func (c *须门那SomnathCounty) BDiu刁元() feud.Barony {
	return c.刁元Diu
}

func (c *须门那SomnathCounty) BKanjetar甘祇陀() feud.Barony {
	return c.甘祇陀Kanjetar
}

func (c *须门那SomnathCounty) BMoherak摩诃罗卡() feud.Barony {
	return c.摩诃罗卡Moherak
}

func (c *须门那SomnathCounty) BSomnath苏摩那他() feud.Barony {
	return c.苏摩那他Somnath
}

func (c *须门那SomnathCounty) BVejalkot吠阇罗拘吒() feud.Barony {
	return c.吠阇罗拘吒Vejalkot
}

func (c *须门那SomnathCounty) BVeraval韦拉沃尔() feud.Barony {
	return c.韦拉沃尔Veraval
}

var CSomnath须门那 SomnathCounty = &须门那SomnathCounty{}

func init() {
	f := CSomnath须门那.(*须门那SomnathCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1266",
		Title:     "somnath",
		TitleName: "须门那",
		TitleCode: "c_somnath",
		Baronies:  map[string]feud.Barony{},
	}

	f.提伐陀Delvada = BDelvada提伐陀
	f.提伐陀Delvada.SetParent(f)

	f.刁元Diu = BDiu刁元
	f.刁元Diu.SetParent(f)

	f.甘祇陀Kanjetar = BKanjetar甘祇陀
	f.甘祇陀Kanjetar.SetParent(f)

	f.摩诃罗卡Moherak = BMoherak摩诃罗卡
	f.摩诃罗卡Moherak.SetParent(f)

	f.苏摩那他Somnath = BSomnath苏摩那他
	f.苏摩那他Somnath.SetParent(f)

	f.吠阇罗拘吒Vejalkot = BVejalkot吠阇罗拘吒
	f.吠阇罗拘吒Vejalkot.SetParent(f)

	f.韦拉沃尔Veraval = BVeraval韦拉沃尔
	f.韦拉沃尔Veraval.SetParent(f)

}
