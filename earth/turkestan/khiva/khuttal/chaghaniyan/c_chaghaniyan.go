package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChaghaniyanCounty interface {
	feud.County
	BChaghaniyan斫汗那() feud.Barony
	BKalif卡利甫() feud.Barony
	BSapoltepa萨波特帕() feud.Barony
	BShuman数瞒() feud.Barony
	BTermez呾蜜() feud.Barony
}

type 斫汗那ChaghaniyanCounty struct {
	feud.BaseCounty
	斫汗那Chaghaniyan feud.Barony
	卡利甫Kalif       feud.Barony
	萨波特帕Sapoltepa  feud.Barony
	数瞒Shuman       feud.Barony
	呾蜜Termez       feud.Barony
}

func (c *斫汗那ChaghaniyanCounty) BChaghaniyan斫汗那() feud.Barony {
	return c.斫汗那Chaghaniyan
}

func (c *斫汗那ChaghaniyanCounty) BKalif卡利甫() feud.Barony {
	return c.卡利甫Kalif
}

func (c *斫汗那ChaghaniyanCounty) BSapoltepa萨波特帕() feud.Barony {
	return c.萨波特帕Sapoltepa
}

func (c *斫汗那ChaghaniyanCounty) BShuman数瞒() feud.Barony {
	return c.数瞒Shuman
}

func (c *斫汗那ChaghaniyanCounty) BTermez呾蜜() feud.Barony {
	return c.呾蜜Termez
}

var CChaghaniyan斫汗那 ChaghaniyanCounty = &斫汗那ChaghaniyanCounty{}

func init() {
	f := CChaghaniyan斫汗那.(*斫汗那ChaghaniyanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1544",
		Title:     "chaghaniyan",
		TitleName: "斫汗那",
		TitleCode: "c_chaghaniyan",
		Baronies:  map[string]feud.Barony{},
	}

	f.斫汗那Chaghaniyan = BChaghaniyan斫汗那
	f.斫汗那Chaghaniyan.SetParent(f)

	f.卡利甫Kalif = BKalif卡利甫
	f.卡利甫Kalif.SetParent(f)

	f.萨波特帕Sapoltepa = BSapoltepa萨波特帕
	f.萨波特帕Sapoltepa.SetParent(f)

	f.数瞒Shuman = BShuman数瞒
	f.数瞒Shuman.SetParent(f)

	f.呾蜜Termez = BTermez呾蜜
	f.呾蜜Termez.SetParent(f)

}
