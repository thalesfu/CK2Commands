package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BidarCounty interface {
	feud.County
	BAvarwadi阿伐瓦迪() feud.Barony
	BBhalki婆吉() feud.Barony
	BBhatambra婆多摩罗() feud.Barony
	BBidar毗陀罗() feud.Barony
	BJalasangvi阇罗僧毗() feud.Barony
	BKamthana迦摩他那() feud.Barony
	BTekadee陀迦德() feud.Barony
}

type 毗陀罗BidarCounty struct {
	feud.BaseCounty
	阿伐瓦迪Avarwadi   feud.Barony
	婆吉Bhalki       feud.Barony
	婆多摩罗Bhatambra  feud.Barony
	毗陀罗Bidar       feud.Barony
	阇罗僧毗Jalasangvi feud.Barony
	迦摩他那Kamthana   feud.Barony
	陀迦德Tekadee     feud.Barony
}

func (c *毗陀罗BidarCounty) BAvarwadi阿伐瓦迪() feud.Barony {
	return c.阿伐瓦迪Avarwadi
}

func (c *毗陀罗BidarCounty) BBhalki婆吉() feud.Barony {
	return c.婆吉Bhalki
}

func (c *毗陀罗BidarCounty) BBhatambra婆多摩罗() feud.Barony {
	return c.婆多摩罗Bhatambra
}

func (c *毗陀罗BidarCounty) BBidar毗陀罗() feud.Barony {
	return c.毗陀罗Bidar
}

func (c *毗陀罗BidarCounty) BJalasangvi阇罗僧毗() feud.Barony {
	return c.阇罗僧毗Jalasangvi
}

func (c *毗陀罗BidarCounty) BKamthana迦摩他那() feud.Barony {
	return c.迦摩他那Kamthana
}

func (c *毗陀罗BidarCounty) BTekadee陀迦德() feud.Barony {
	return c.陀迦德Tekadee
}

var CBidar毗陀罗 BidarCounty = &毗陀罗BidarCounty{}

func init() {
	f := CBidar毗陀罗.(*毗陀罗BidarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1158",
		Title:     "bidar",
		TitleName: "毗陀罗",
		TitleCode: "c_bidar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伐瓦迪Avarwadi = BAvarwadi阿伐瓦迪
	f.阿伐瓦迪Avarwadi.SetParent(f)

	f.婆吉Bhalki = BBhalki婆吉
	f.婆吉Bhalki.SetParent(f)

	f.婆多摩罗Bhatambra = BBhatambra婆多摩罗
	f.婆多摩罗Bhatambra.SetParent(f)

	f.毗陀罗Bidar = BBidar毗陀罗
	f.毗陀罗Bidar.SetParent(f)

	f.阇罗僧毗Jalasangvi = BJalasangvi阇罗僧毗
	f.阇罗僧毗Jalasangvi.SetParent(f)

	f.迦摩他那Kamthana = BKamthana迦摩他那
	f.迦摩他那Kamthana.SetParent(f)

	f.陀迦德Tekadee = BTekadee陀迦德
	f.陀迦德Tekadee.SetParent(f)

}
