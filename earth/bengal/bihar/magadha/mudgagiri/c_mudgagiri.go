package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MudgagiriCounty interface {
	feud.County
	BCampa瞻波() feud.Barony
	BJahnugiri阇喜奴山() feud.Barony
	BKahalgaon迦诃罗伽罗摩() feud.Barony
	BKarnagarh羯罗拏姞利呬() feud.Barony
	BLakhisarai罗吉娑罗伊() feud.Barony
	BMudgagiri勿伽耆厘() feud.Barony
	BVikramasila超戒寺() feud.Barony
}

type 勿伽耆厘MudgagiriCounty struct {
	feud.BaseCounty
	瞻波Campa         feud.Barony
	阇喜奴山Jahnugiri   feud.Barony
	迦诃罗伽罗摩Kahalgaon feud.Barony
	羯罗拏姞利呬Karnagarh feud.Barony
	罗吉娑罗伊Lakhisarai feud.Barony
	勿伽耆厘Mudgagiri   feud.Barony
	超戒寺Vikramasila  feud.Barony
}

func (c *勿伽耆厘MudgagiriCounty) BCampa瞻波() feud.Barony {
	return c.瞻波Campa
}

func (c *勿伽耆厘MudgagiriCounty) BJahnugiri阇喜奴山() feud.Barony {
	return c.阇喜奴山Jahnugiri
}

func (c *勿伽耆厘MudgagiriCounty) BKahalgaon迦诃罗伽罗摩() feud.Barony {
	return c.迦诃罗伽罗摩Kahalgaon
}

func (c *勿伽耆厘MudgagiriCounty) BKarnagarh羯罗拏姞利呬() feud.Barony {
	return c.羯罗拏姞利呬Karnagarh
}

func (c *勿伽耆厘MudgagiriCounty) BLakhisarai罗吉娑罗伊() feud.Barony {
	return c.罗吉娑罗伊Lakhisarai
}

func (c *勿伽耆厘MudgagiriCounty) BMudgagiri勿伽耆厘() feud.Barony {
	return c.勿伽耆厘Mudgagiri
}

func (c *勿伽耆厘MudgagiriCounty) BVikramasila超戒寺() feud.Barony {
	return c.超戒寺Vikramasila
}

var CMudgagiri勿伽耆厘 MudgagiriCounty = &勿伽耆厘MudgagiriCounty{}

func init() {
	f := CMudgagiri勿伽耆厘.(*勿伽耆厘MudgagiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1152",
		Title:     "mudgagiri",
		TitleName: "勿伽耆厘",
		TitleCode: "c_mudgagiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.瞻波Campa = BCampa瞻波
	f.瞻波Campa.SetParent(f)

	f.阇喜奴山Jahnugiri = BJahnugiri阇喜奴山
	f.阇喜奴山Jahnugiri.SetParent(f)

	f.迦诃罗伽罗摩Kahalgaon = BKahalgaon迦诃罗伽罗摩
	f.迦诃罗伽罗摩Kahalgaon.SetParent(f)

	f.羯罗拏姞利呬Karnagarh = BKarnagarh羯罗拏姞利呬
	f.羯罗拏姞利呬Karnagarh.SetParent(f)

	f.罗吉娑罗伊Lakhisarai = BLakhisarai罗吉娑罗伊
	f.罗吉娑罗伊Lakhisarai.SetParent(f)

	f.勿伽耆厘Mudgagiri = BMudgagiri勿伽耆厘
	f.勿伽耆厘Mudgagiri.SetParent(f)

	f.超戒寺Vikramasila = BVikramasila超戒寺
	f.超戒寺Vikramasila.SetParent(f)

}
