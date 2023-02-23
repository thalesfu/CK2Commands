package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BredaCounty interface {
	feud.County
	BBergenopzoom贝亨奥普佐姆() feud.Barony
	BBreda布雷达() feud.Barony
	BCuijck克伊克() feud.Barony
	BHorn霍恩() feud.Barony
	BRavenstein拉芬斯泰因() feud.Barony
	BShertogenbosch斯海尔托亨博思() feud.Barony
	BTilburg蒂尔堡() feud.Barony
	BWillemstad威廉斯塔德() feud.Barony
}

type 布雷达BredaCounty struct {
	feud.BaseCounty
	贝亨奥普佐姆Bergenopzoom    feud.Barony
	布雷达Breda              feud.Barony
	克伊克Cuijck             feud.Barony
	霍恩Horn                feud.Barony
	拉芬斯泰因Ravenstein       feud.Barony
	斯海尔托亨博思Shertogenbosch feud.Barony
	蒂尔堡Tilburg            feud.Barony
	威廉斯塔德Willemstad       feud.Barony
}

func (c *布雷达BredaCounty) BBergenopzoom贝亨奥普佐姆() feud.Barony {
	return c.贝亨奥普佐姆Bergenopzoom
}

func (c *布雷达BredaCounty) BBreda布雷达() feud.Barony {
	return c.布雷达Breda
}

func (c *布雷达BredaCounty) BCuijck克伊克() feud.Barony {
	return c.克伊克Cuijck
}

func (c *布雷达BredaCounty) BHorn霍恩() feud.Barony {
	return c.霍恩Horn
}

func (c *布雷达BredaCounty) BRavenstein拉芬斯泰因() feud.Barony {
	return c.拉芬斯泰因Ravenstein
}

func (c *布雷达BredaCounty) BShertogenbosch斯海尔托亨博思() feud.Barony {
	return c.斯海尔托亨博思Shertogenbosch
}

func (c *布雷达BredaCounty) BTilburg蒂尔堡() feud.Barony {
	return c.蒂尔堡Tilburg
}

func (c *布雷达BredaCounty) BWillemstad威廉斯塔德() feud.Barony {
	return c.威廉斯塔德Willemstad
}

var CBreda布雷达 BredaCounty = &布雷达BredaCounty{}

func init() {
	f := CBreda布雷达.(*布雷达BredaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "92",
		Title:     "breda",
		TitleName: "布雷达",
		TitleCode: "c_breda",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝亨奥普佐姆Bergenopzoom = BBergenopzoom贝亨奥普佐姆
	f.贝亨奥普佐姆Bergenopzoom.SetParent(f)

	f.布雷达Breda = BBreda布雷达
	f.布雷达Breda.SetParent(f)

	f.克伊克Cuijck = BCuijck克伊克
	f.克伊克Cuijck.SetParent(f)

	f.霍恩Horn = BHorn霍恩
	f.霍恩Horn.SetParent(f)

	f.拉芬斯泰因Ravenstein = BRavenstein拉芬斯泰因
	f.拉芬斯泰因Ravenstein.SetParent(f)

	f.斯海尔托亨博思Shertogenbosch = BShertogenbosch斯海尔托亨博思
	f.斯海尔托亨博思Shertogenbosch.SetParent(f)

	f.蒂尔堡Tilburg = BTilburg蒂尔堡
	f.蒂尔堡Tilburg.SetParent(f)

	f.威廉斯塔德Willemstad = BWillemstad威廉斯塔德
	f.威廉斯塔德Willemstad.SetParent(f)

}
