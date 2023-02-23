package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaldurgCounty interface {
	feud.County
	BBaniagaon槃尼耶伽罗摩() feud.Barony
	BDharaseo达拉斯() feud.Barony
	BKarmala卡尔马拉() feud.Barony
	BNaldurg那罗突伽() feud.Barony
	BParanda波烂陀() feud.Barony
	BSonnalagi十六城() feud.Barony
	BTuljapur图尔贾普尔() feud.Barony
}

type 那罗突伽NaldurgCounty struct {
	feud.BaseCounty
	槃尼耶伽罗摩Baniagaon feud.Barony
	达拉斯Dharaseo     feud.Barony
	卡尔马拉Karmala     feud.Barony
	那罗突伽Naldurg     feud.Barony
	波烂陀Paranda      feud.Barony
	十六城Sonnalagi    feud.Barony
	图尔贾普尔Tuljapur   feud.Barony
}

func (c *那罗突伽NaldurgCounty) BBaniagaon槃尼耶伽罗摩() feud.Barony {
	return c.槃尼耶伽罗摩Baniagaon
}

func (c *那罗突伽NaldurgCounty) BDharaseo达拉斯() feud.Barony {
	return c.达拉斯Dharaseo
}

func (c *那罗突伽NaldurgCounty) BKarmala卡尔马拉() feud.Barony {
	return c.卡尔马拉Karmala
}

func (c *那罗突伽NaldurgCounty) BNaldurg那罗突伽() feud.Barony {
	return c.那罗突伽Naldurg
}

func (c *那罗突伽NaldurgCounty) BParanda波烂陀() feud.Barony {
	return c.波烂陀Paranda
}

func (c *那罗突伽NaldurgCounty) BSonnalagi十六城() feud.Barony {
	return c.十六城Sonnalagi
}

func (c *那罗突伽NaldurgCounty) BTuljapur图尔贾普尔() feud.Barony {
	return c.图尔贾普尔Tuljapur
}

var CNaldurg那罗突伽 NaldurgCounty = &那罗突伽NaldurgCounty{}

func init() {
	f := CNaldurg那罗突伽.(*那罗突伽NaldurgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1146",
		Title:     "naldurg",
		TitleName: "那罗突伽",
		TitleCode: "c_naldurg",
		Baronies:  map[string]feud.Barony{},
	}

	f.槃尼耶伽罗摩Baniagaon = BBaniagaon槃尼耶伽罗摩
	f.槃尼耶伽罗摩Baniagaon.SetParent(f)

	f.达拉斯Dharaseo = BDharaseo达拉斯
	f.达拉斯Dharaseo.SetParent(f)

	f.卡尔马拉Karmala = BKarmala卡尔马拉
	f.卡尔马拉Karmala.SetParent(f)

	f.那罗突伽Naldurg = BNaldurg那罗突伽
	f.那罗突伽Naldurg.SetParent(f)

	f.波烂陀Paranda = BParanda波烂陀
	f.波烂陀Paranda.SetParent(f)

	f.十六城Sonnalagi = BSonnalagi十六城
	f.十六城Sonnalagi.SetParent(f)

	f.图尔贾普尔Tuljapur = BTuljapur图尔贾普尔
	f.图尔贾普尔Tuljapur.SetParent(f)

}
