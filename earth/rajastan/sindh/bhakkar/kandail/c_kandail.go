package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KandailCounty interface {
	feud.County
	BDadu达都() feud.Barony
	BJhal切尔() feud.Barony
	BKandail甘代尔() feud.Barony
	BVinjamoor频阇慕罗() feud.Barony
	BWilong威龙() feud.Barony
	BXelvona切符那() feud.Barony
	BYamanpalli耶摩那波梨() feud.Barony
}

type 甘代尔KandailCounty struct {
	feud.BaseCounty
	达都Dadu          feud.Barony
	切尔Jhal          feud.Barony
	甘代尔Kandail      feud.Barony
	频阇慕罗Vinjamoor   feud.Barony
	威龙Wilong        feud.Barony
	切符那Xelvona      feud.Barony
	耶摩那波梨Yamanpalli feud.Barony
}

func (c *甘代尔KandailCounty) BDadu达都() feud.Barony {
	return c.达都Dadu
}

func (c *甘代尔KandailCounty) BJhal切尔() feud.Barony {
	return c.切尔Jhal
}

func (c *甘代尔KandailCounty) BKandail甘代尔() feud.Barony {
	return c.甘代尔Kandail
}

func (c *甘代尔KandailCounty) BVinjamoor频阇慕罗() feud.Barony {
	return c.频阇慕罗Vinjamoor
}

func (c *甘代尔KandailCounty) BWilong威龙() feud.Barony {
	return c.威龙Wilong
}

func (c *甘代尔KandailCounty) BXelvona切符那() feud.Barony {
	return c.切符那Xelvona
}

func (c *甘代尔KandailCounty) BYamanpalli耶摩那波梨() feud.Barony {
	return c.耶摩那波梨Yamanpalli
}

var CKandail甘代尔 KandailCounty = &甘代尔KandailCounty{}

func init() {
	f := CKandail甘代尔.(*甘代尔KandailCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1372",
		Title:     "kandail",
		TitleName: "甘代尔",
		TitleCode: "c_kandail",
		Baronies:  map[string]feud.Barony{},
	}

	f.达都Dadu = BDadu达都
	f.达都Dadu.SetParent(f)

	f.切尔Jhal = BJhal切尔
	f.切尔Jhal.SetParent(f)

	f.甘代尔Kandail = BKandail甘代尔
	f.甘代尔Kandail.SetParent(f)

	f.频阇慕罗Vinjamoor = BVinjamoor频阇慕罗
	f.频阇慕罗Vinjamoor.SetParent(f)

	f.威龙Wilong = BWilong威龙
	f.威龙Wilong.SetParent(f)

	f.切符那Xelvona = BXelvona切符那
	f.切符那Xelvona.SetParent(f)

	f.耶摩那波梨Yamanpalli = BYamanpalli耶摩那波梨
	f.耶摩那波梨Yamanpalli.SetParent(f)

}
