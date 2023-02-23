package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NarbonneCounty interface {
	feud.County
	BAgde阿格德() feud.Barony
	BAlbi阿尔比() feud.Barony
	BBeziers贝济耶() feud.Barony
	BCastres卡斯特尔() feud.Barony
	BNarbonne纳博讷() feud.Barony
	BPuisserguier皮塞吉耶() feud.Barony
	BQueribus凯里比() feud.Barony
	BStponsdethomieres圣蓬德托米耶尔() feud.Barony
}

type 纳博讷NarbonneCounty struct {
	feud.BaseCounty
	阿格德Agde                  feud.Barony
	阿尔比Albi                  feud.Barony
	贝济耶Beziers               feud.Barony
	卡斯特尔Castres              feud.Barony
	纳博讷Narbonne              feud.Barony
	皮塞吉耶Puisserguier         feud.Barony
	凯里比Queribus              feud.Barony
	圣蓬德托米耶尔Stponsdethomieres feud.Barony
}

func (c *纳博讷NarbonneCounty) BAgde阿格德() feud.Barony {
	return c.阿格德Agde
}

func (c *纳博讷NarbonneCounty) BAlbi阿尔比() feud.Barony {
	return c.阿尔比Albi
}

func (c *纳博讷NarbonneCounty) BBeziers贝济耶() feud.Barony {
	return c.贝济耶Beziers
}

func (c *纳博讷NarbonneCounty) BCastres卡斯特尔() feud.Barony {
	return c.卡斯特尔Castres
}

func (c *纳博讷NarbonneCounty) BNarbonne纳博讷() feud.Barony {
	return c.纳博讷Narbonne
}

func (c *纳博讷NarbonneCounty) BPuisserguier皮塞吉耶() feud.Barony {
	return c.皮塞吉耶Puisserguier
}

func (c *纳博讷NarbonneCounty) BQueribus凯里比() feud.Barony {
	return c.凯里比Queribus
}

func (c *纳博讷NarbonneCounty) BStponsdethomieres圣蓬德托米耶尔() feud.Barony {
	return c.圣蓬德托米耶尔Stponsdethomieres
}

var CNarbonne纳博讷 NarbonneCounty = &纳博讷NarbonneCounty{}

func init() {
	f := CNarbonne纳博讷.(*纳博讷NarbonneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "212",
		Title:     "narbonne",
		TitleName: "纳博讷",
		TitleCode: "c_narbonne",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格德Agde = BAgde阿格德
	f.阿格德Agde.SetParent(f)

	f.阿尔比Albi = BAlbi阿尔比
	f.阿尔比Albi.SetParent(f)

	f.贝济耶Beziers = BBeziers贝济耶
	f.贝济耶Beziers.SetParent(f)

	f.卡斯特尔Castres = BCastres卡斯特尔
	f.卡斯特尔Castres.SetParent(f)

	f.纳博讷Narbonne = BNarbonne纳博讷
	f.纳博讷Narbonne.SetParent(f)

	f.皮塞吉耶Puisserguier = BPuisserguier皮塞吉耶
	f.皮塞吉耶Puisserguier.SetParent(f)

	f.凯里比Queribus = BQueribus凯里比
	f.凯里比Queribus.SetParent(f)

	f.圣蓬德托米耶尔Stponsdethomieres = BStponsdethomieres圣蓬德托米耶尔
	f.圣蓬德托米耶尔Stponsdethomieres.SetParent(f)

}
