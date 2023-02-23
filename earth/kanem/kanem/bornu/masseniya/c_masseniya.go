package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MasseniyaCounty interface {
	feud.County
	BAtegbe阿特格贝() feud.Barony
	BMalfe马勒非() feud.Barony
	BMasseniya马塞尼亚() feud.Barony
	BOganga奥冈加() feud.Barony
}

type 马塞尼亚MasseniyaCounty struct {
	feud.BaseCounty
	阿特格贝Ategbe    feud.Barony
	马勒非Malfe      feud.Barony
	马塞尼亚Masseniya feud.Barony
	奥冈加Oganga     feud.Barony
}

func (c *马塞尼亚MasseniyaCounty) BAtegbe阿特格贝() feud.Barony {
	return c.阿特格贝Ategbe
}

func (c *马塞尼亚MasseniyaCounty) BMalfe马勒非() feud.Barony {
	return c.马勒非Malfe
}

func (c *马塞尼亚MasseniyaCounty) BMasseniya马塞尼亚() feud.Barony {
	return c.马塞尼亚Masseniya
}

func (c *马塞尼亚MasseniyaCounty) BOganga奥冈加() feud.Barony {
	return c.奥冈加Oganga
}

var CMasseniya马塞尼亚 MasseniyaCounty = &马塞尼亚MasseniyaCounty{}

func init() {
	f := CMasseniya马塞尼亚.(*马塞尼亚MasseniyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1740",
		Title:     "masseniya",
		TitleName: "马塞尼亚",
		TitleCode: "c_masseniya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿特格贝Ategbe = BAtegbe阿特格贝
	f.阿特格贝Ategbe.SetParent(f)

	f.马勒非Malfe = BMalfe马勒非
	f.马勒非Malfe.SetParent(f)

	f.马塞尼亚Masseniya = BMasseniya马塞尼亚
	f.马塞尼亚Masseniya.SetParent(f)

	f.奥冈加Oganga = BOganga奥冈加
	f.奥冈加Oganga.SetParent(f)

}
