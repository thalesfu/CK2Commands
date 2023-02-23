package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LindisfarneCounty interface {
	feud.County
	BCornhill康希尔() feud.Barony
	BEtal伊特尔() feud.Barony
	BLindisfarne林迪斯法恩() feud.Barony
	BLowick洛伊克() feud.Barony
	BNorham诺勒姆() feud.Barony
	BTweedmouth特威德茅斯() feud.Barony
	BYeavering耶弗林() feud.Barony
}

type 林迪斯法恩LindisfarneCounty struct {
	feud.BaseCounty
	康希尔Cornhill      feud.Barony
	伊特尔Etal          feud.Barony
	林迪斯法恩Lindisfarne feud.Barony
	洛伊克Lowick        feud.Barony
	诺勒姆Norham        feud.Barony
	特威德茅斯Tweedmouth  feud.Barony
	耶弗林Yeavering     feud.Barony
}

func (c *林迪斯法恩LindisfarneCounty) BCornhill康希尔() feud.Barony {
	return c.康希尔Cornhill
}

func (c *林迪斯法恩LindisfarneCounty) BEtal伊特尔() feud.Barony {
	return c.伊特尔Etal
}

func (c *林迪斯法恩LindisfarneCounty) BLindisfarne林迪斯法恩() feud.Barony {
	return c.林迪斯法恩Lindisfarne
}

func (c *林迪斯法恩LindisfarneCounty) BLowick洛伊克() feud.Barony {
	return c.洛伊克Lowick
}

func (c *林迪斯法恩LindisfarneCounty) BNorham诺勒姆() feud.Barony {
	return c.诺勒姆Norham
}

func (c *林迪斯法恩LindisfarneCounty) BTweedmouth特威德茅斯() feud.Barony {
	return c.特威德茅斯Tweedmouth
}

func (c *林迪斯法恩LindisfarneCounty) BYeavering耶弗林() feud.Barony {
	return c.耶弗林Yeavering
}

var CLindisfarne林迪斯法恩 LindisfarneCounty = &林迪斯法恩LindisfarneCounty{}

func init() {
	f := CLindisfarne林迪斯法恩.(*林迪斯法恩LindisfarneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1940",
		Title:     "lindisfarne",
		TitleName: "林迪斯法恩",
		TitleCode: "c_lindisfarne",
		Baronies:  map[string]feud.Barony{},
	}

	f.康希尔Cornhill = BCornhill康希尔
	f.康希尔Cornhill.SetParent(f)

	f.伊特尔Etal = BEtal伊特尔
	f.伊特尔Etal.SetParent(f)

	f.林迪斯法恩Lindisfarne = BLindisfarne林迪斯法恩
	f.林迪斯法恩Lindisfarne.SetParent(f)

	f.洛伊克Lowick = BLowick洛伊克
	f.洛伊克Lowick.SetParent(f)

	f.诺勒姆Norham = BNorham诺勒姆
	f.诺勒姆Norham.SetParent(f)

	f.特威德茅斯Tweedmouth = BTweedmouth特威德茅斯
	f.特威德茅斯Tweedmouth.SetParent(f)

	f.耶弗林Yeavering = BYeavering耶弗林
	f.耶弗林Yeavering.SetParent(f)

}
