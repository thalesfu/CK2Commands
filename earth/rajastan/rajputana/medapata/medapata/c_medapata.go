package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedapataCounty interface {
	feud.County
	BAghata阿伽吒() feud.Barony
	BBagar婆迦尔() feud.Barony
	BEklingji翳迦林伽视() feud.Barony
	BJagadamba社伽淡婆() feud.Barony
	BNagahrada那伽贺逻驮() feud.Barony
	BRishabhdeo勒沙婆提婆() feud.Barony
}

type 迷陀波吒MedapataCounty struct {
	feud.BaseCounty
	阿伽吒Aghata       feud.Barony
	婆迦尔Bagar        feud.Barony
	翳迦林伽视Eklingji   feud.Barony
	社伽淡婆Jagadamba   feud.Barony
	那伽贺逻驮Nagahrada  feud.Barony
	勒沙婆提婆Rishabhdeo feud.Barony
}

func (c *迷陀波吒MedapataCounty) BAghata阿伽吒() feud.Barony {
	return c.阿伽吒Aghata
}

func (c *迷陀波吒MedapataCounty) BBagar婆迦尔() feud.Barony {
	return c.婆迦尔Bagar
}

func (c *迷陀波吒MedapataCounty) BEklingji翳迦林伽视() feud.Barony {
	return c.翳迦林伽视Eklingji
}

func (c *迷陀波吒MedapataCounty) BJagadamba社伽淡婆() feud.Barony {
	return c.社伽淡婆Jagadamba
}

func (c *迷陀波吒MedapataCounty) BNagahrada那伽贺逻驮() feud.Barony {
	return c.那伽贺逻驮Nagahrada
}

func (c *迷陀波吒MedapataCounty) BRishabhdeo勒沙婆提婆() feud.Barony {
	return c.勒沙婆提婆Rishabhdeo
}

var CMedapata迷陀波吒 MedapataCounty = &迷陀波吒MedapataCounty{}

func init() {
	f := CMedapata迷陀波吒.(*迷陀波吒MedapataCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1345",
		Title:     "medapata",
		TitleName: "迷陀波吒",
		TitleCode: "c_medapata",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伽吒Aghata = BAghata阿伽吒
	f.阿伽吒Aghata.SetParent(f)

	f.婆迦尔Bagar = BBagar婆迦尔
	f.婆迦尔Bagar.SetParent(f)

	f.翳迦林伽视Eklingji = BEklingji翳迦林伽视
	f.翳迦林伽视Eklingji.SetParent(f)

	f.社伽淡婆Jagadamba = BJagadamba社伽淡婆
	f.社伽淡婆Jagadamba.SetParent(f)

	f.那伽贺逻驮Nagahrada = BNagahrada那伽贺逻驮
	f.那伽贺逻驮Nagahrada.SetParent(f)

	f.勒沙婆提婆Rishabhdeo = BRishabhdeo勒沙婆提婆
	f.勒沙婆提婆Rishabhdeo.SetParent(f)

}
