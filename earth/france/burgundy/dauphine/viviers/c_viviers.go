package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ViviersCounty interface {
	feud.County
	BAlbalaromaine阿尔巴拉罗曼() feud.Barony
	BAubenas欧布纳() feud.Barony
	BJoyeuse茹瓦约斯() feud.Barony
	BLargentiere拉让蒂耶尔() feud.Barony
	BLecheylard勒谢拉尔() feud.Barony
	BPrivas普里瓦() feud.Barony
	BTournon图尔农() feud.Barony
	BViviers维维耶() feud.Barony
}

type 维瓦赖ViviersCounty struct {
	feud.BaseCounty
	阿尔巴拉罗曼Albalaromaine feud.Barony
	欧布纳Aubenas          feud.Barony
	茹瓦约斯Joyeuse         feud.Barony
	拉让蒂耶尔Largentiere    feud.Barony
	勒谢拉尔Lecheylard      feud.Barony
	普里瓦Privas           feud.Barony
	图尔农Tournon          feud.Barony
	维维耶Viviers          feud.Barony
}

func (c *维瓦赖ViviersCounty) BAlbalaromaine阿尔巴拉罗曼() feud.Barony {
	return c.阿尔巴拉罗曼Albalaromaine
}

func (c *维瓦赖ViviersCounty) BAubenas欧布纳() feud.Barony {
	return c.欧布纳Aubenas
}

func (c *维瓦赖ViviersCounty) BJoyeuse茹瓦约斯() feud.Barony {
	return c.茹瓦约斯Joyeuse
}

func (c *维瓦赖ViviersCounty) BLargentiere拉让蒂耶尔() feud.Barony {
	return c.拉让蒂耶尔Largentiere
}

func (c *维瓦赖ViviersCounty) BLecheylard勒谢拉尔() feud.Barony {
	return c.勒谢拉尔Lecheylard
}

func (c *维瓦赖ViviersCounty) BPrivas普里瓦() feud.Barony {
	return c.普里瓦Privas
}

func (c *维瓦赖ViviersCounty) BTournon图尔农() feud.Barony {
	return c.图尔农Tournon
}

func (c *维瓦赖ViviersCounty) BViviers维维耶() feud.Barony {
	return c.维维耶Viviers
}

var CViviers维瓦赖 ViviersCounty = &维瓦赖ViviersCounty{}

func init() {
	f := CViviers维瓦赖.(*维瓦赖ViviersCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "223",
		Title:     "viviers",
		TitleName: "维瓦赖",
		TitleCode: "c_viviers",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔巴拉罗曼Albalaromaine = BAlbalaromaine阿尔巴拉罗曼
	f.阿尔巴拉罗曼Albalaromaine.SetParent(f)

	f.欧布纳Aubenas = BAubenas欧布纳
	f.欧布纳Aubenas.SetParent(f)

	f.茹瓦约斯Joyeuse = BJoyeuse茹瓦约斯
	f.茹瓦约斯Joyeuse.SetParent(f)

	f.拉让蒂耶尔Largentiere = BLargentiere拉让蒂耶尔
	f.拉让蒂耶尔Largentiere.SetParent(f)

	f.勒谢拉尔Lecheylard = BLecheylard勒谢拉尔
	f.勒谢拉尔Lecheylard.SetParent(f)

	f.普里瓦Privas = BPrivas普里瓦
	f.普里瓦Privas.SetParent(f)

	f.图尔农Tournon = BTournon图尔农
	f.图尔农Tournon.SetParent(f)

	f.维维耶Viviers = BViviers维维耶
	f.维维耶Viviers.SetParent(f)

}
