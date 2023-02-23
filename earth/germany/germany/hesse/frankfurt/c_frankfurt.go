package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FrankfurtCounty interface {
	feud.County
	BBotisphaden博蒂斯法登() feud.Barony
	BDietigheim迪蒂希海姆() feud.Barony
	BFrankfurt法兰克福() feud.Barony
	BKatzenelnbogen卡岑埃尔恩博根() feud.Barony
	BRheinfels莱茵费尔斯() feud.Barony
	BWiesbaden威斯巴登() feud.Barony
}

type 法兰克福FrankfurtCounty struct {
	feud.BaseCounty
	博蒂斯法登Botisphaden      feud.Barony
	迪蒂希海姆Dietigheim       feud.Barony
	法兰克福Frankfurt         feud.Barony
	卡岑埃尔恩博根Katzenelnbogen feud.Barony
	莱茵费尔斯Rheinfels        feud.Barony
	威斯巴登Wiesbaden         feud.Barony
}

func (c *法兰克福FrankfurtCounty) BBotisphaden博蒂斯法登() feud.Barony {
	return c.博蒂斯法登Botisphaden
}

func (c *法兰克福FrankfurtCounty) BDietigheim迪蒂希海姆() feud.Barony {
	return c.迪蒂希海姆Dietigheim
}

func (c *法兰克福FrankfurtCounty) BFrankfurt法兰克福() feud.Barony {
	return c.法兰克福Frankfurt
}

func (c *法兰克福FrankfurtCounty) BKatzenelnbogen卡岑埃尔恩博根() feud.Barony {
	return c.卡岑埃尔恩博根Katzenelnbogen
}

func (c *法兰克福FrankfurtCounty) BRheinfels莱茵费尔斯() feud.Barony {
	return c.莱茵费尔斯Rheinfels
}

func (c *法兰克福FrankfurtCounty) BWiesbaden威斯巴登() feud.Barony {
	return c.威斯巴登Wiesbaden
}

var CFrankfurt法兰克福 FrankfurtCounty = &法兰克福FrankfurtCounty{}

func init() {
	f := CFrankfurt法兰克福.(*法兰克福FrankfurtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1981",
		Title:     "frankfurt",
		TitleName: "法兰克福",
		TitleCode: "c_frankfurt",
		Baronies:  map[string]feud.Barony{},
	}

	f.博蒂斯法登Botisphaden = BBotisphaden博蒂斯法登
	f.博蒂斯法登Botisphaden.SetParent(f)

	f.迪蒂希海姆Dietigheim = BDietigheim迪蒂希海姆
	f.迪蒂希海姆Dietigheim.SetParent(f)

	f.法兰克福Frankfurt = BFrankfurt法兰克福
	f.法兰克福Frankfurt.SetParent(f)

	f.卡岑埃尔恩博根Katzenelnbogen = BKatzenelnbogen卡岑埃尔恩博根
	f.卡岑埃尔恩博根Katzenelnbogen.SetParent(f)

	f.莱茵费尔斯Rheinfels = BRheinfels莱茵费尔斯
	f.莱茵费尔斯Rheinfels.SetParent(f)

	f.威斯巴登Wiesbaden = BWiesbaden威斯巴登
	f.威斯巴登Wiesbaden.SetParent(f)

}
