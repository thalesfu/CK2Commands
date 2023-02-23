package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AuxerreCounty interface {
	feud.County
	BAuxerre欧塞尔() feud.Barony
	BCravant克拉旺() feud.Barony
	BCrisenon克里瑟农() feud.Barony
	BDruyes德吕耶() feud.Barony
	BMailly马伊() feud.Barony
	BPontigny蓬蒂尼() feud.Barony
	BStsauveurenpuisaye圣索沃尔_昂皮赛() feud.Barony
	BTonnerre托内尔() feud.Barony
}

type 欧塞尔AuxerreCounty struct {
	feud.BaseCounty
	欧塞尔Auxerre                 feud.Barony
	克拉旺Cravant                 feud.Barony
	克里瑟农Crisenon               feud.Barony
	德吕耶Druyes                  feud.Barony
	马伊Mailly                   feud.Barony
	蓬蒂尼Pontigny                feud.Barony
	圣索沃尔_昂皮赛Stsauveurenpuisaye feud.Barony
	托内尔Tonnerre                feud.Barony
}

func (c *欧塞尔AuxerreCounty) BAuxerre欧塞尔() feud.Barony {
	return c.欧塞尔Auxerre
}

func (c *欧塞尔AuxerreCounty) BCravant克拉旺() feud.Barony {
	return c.克拉旺Cravant
}

func (c *欧塞尔AuxerreCounty) BCrisenon克里瑟农() feud.Barony {
	return c.克里瑟农Crisenon
}

func (c *欧塞尔AuxerreCounty) BDruyes德吕耶() feud.Barony {
	return c.德吕耶Druyes
}

func (c *欧塞尔AuxerreCounty) BMailly马伊() feud.Barony {
	return c.马伊Mailly
}

func (c *欧塞尔AuxerreCounty) BPontigny蓬蒂尼() feud.Barony {
	return c.蓬蒂尼Pontigny
}

func (c *欧塞尔AuxerreCounty) BStsauveurenpuisaye圣索沃尔_昂皮赛() feud.Barony {
	return c.圣索沃尔_昂皮赛Stsauveurenpuisaye
}

func (c *欧塞尔AuxerreCounty) BTonnerre托内尔() feud.Barony {
	return c.托内尔Tonnerre
}

var CAuxerre欧塞尔 AuxerreCounty = &欧塞尔AuxerreCounty{}

func init() {
	f := CAuxerre欧塞尔.(*欧塞尔AuxerreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "132",
		Title:     "auxerre",
		TitleName: "欧塞尔",
		TitleCode: "c_auxerre",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧塞尔Auxerre = BAuxerre欧塞尔
	f.欧塞尔Auxerre.SetParent(f)

	f.克拉旺Cravant = BCravant克拉旺
	f.克拉旺Cravant.SetParent(f)

	f.克里瑟农Crisenon = BCrisenon克里瑟农
	f.克里瑟农Crisenon.SetParent(f)

	f.德吕耶Druyes = BDruyes德吕耶
	f.德吕耶Druyes.SetParent(f)

	f.马伊Mailly = BMailly马伊
	f.马伊Mailly.SetParent(f)

	f.蓬蒂尼Pontigny = BPontigny蓬蒂尼
	f.蓬蒂尼Pontigny.SetParent(f)

	f.圣索沃尔_昂皮赛Stsauveurenpuisaye = BStsauveurenpuisaye圣索沃尔_昂皮赛
	f.圣索沃尔_昂皮赛Stsauveurenpuisaye.SetParent(f)

	f.托内尔Tonnerre = BTonnerre托内尔
	f.托内尔Tonnerre.SetParent(f)

}
