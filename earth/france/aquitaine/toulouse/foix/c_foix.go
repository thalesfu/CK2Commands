package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FoixCounty interface {
	feud.County
	BFoix富瓦() feud.Barony
	BMirepoix米尔普瓦() feud.Barony
	BMontsegur蒙塞居尔() feud.Barony
	BPamiers帕米耶() feud.Barony
	BRoquefeuil罗克弗伊() feud.Barony
	BStbertrand圣贝特朗() feud.Barony
	BStgaudens圣戈当() feud.Barony
	BUsson于松() feud.Barony
}

type 富瓦FoixCounty struct {
	feud.BaseCounty
	富瓦Foix         feud.Barony
	米尔普瓦Mirepoix   feud.Barony
	蒙塞居尔Montsegur  feud.Barony
	帕米耶Pamiers     feud.Barony
	罗克弗伊Roquefeuil feud.Barony
	圣贝特朗Stbertrand feud.Barony
	圣戈当Stgaudens   feud.Barony
	于松Usson        feud.Barony
}

func (c *富瓦FoixCounty) BFoix富瓦() feud.Barony {
	return c.富瓦Foix
}

func (c *富瓦FoixCounty) BMirepoix米尔普瓦() feud.Barony {
	return c.米尔普瓦Mirepoix
}

func (c *富瓦FoixCounty) BMontsegur蒙塞居尔() feud.Barony {
	return c.蒙塞居尔Montsegur
}

func (c *富瓦FoixCounty) BPamiers帕米耶() feud.Barony {
	return c.帕米耶Pamiers
}

func (c *富瓦FoixCounty) BRoquefeuil罗克弗伊() feud.Barony {
	return c.罗克弗伊Roquefeuil
}

func (c *富瓦FoixCounty) BStbertrand圣贝特朗() feud.Barony {
	return c.圣贝特朗Stbertrand
}

func (c *富瓦FoixCounty) BStgaudens圣戈当() feud.Barony {
	return c.圣戈当Stgaudens
}

func (c *富瓦FoixCounty) BUsson于松() feud.Barony {
	return c.于松Usson
}

var CFoix富瓦 FoixCounty = &富瓦FoixCounty{}

func init() {
	f := CFoix富瓦.(*富瓦FoixCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "210",
		Title:     "foix",
		TitleName: "富瓦",
		TitleCode: "c_foix",
		Baronies:  map[string]feud.Barony{},
	}

	f.富瓦Foix = BFoix富瓦
	f.富瓦Foix.SetParent(f)

	f.米尔普瓦Mirepoix = BMirepoix米尔普瓦
	f.米尔普瓦Mirepoix.SetParent(f)

	f.蒙塞居尔Montsegur = BMontsegur蒙塞居尔
	f.蒙塞居尔Montsegur.SetParent(f)

	f.帕米耶Pamiers = BPamiers帕米耶
	f.帕米耶Pamiers.SetParent(f)

	f.罗克弗伊Roquefeuil = BRoquefeuil罗克弗伊
	f.罗克弗伊Roquefeuil.SetParent(f)

	f.圣贝特朗Stbertrand = BStbertrand圣贝特朗
	f.圣贝特朗Stbertrand.SetParent(f)

	f.圣戈当Stgaudens = BStgaudens圣戈当
	f.圣戈当Stgaudens.SetParent(f)

	f.于松Usson = BUsson于松
	f.于松Usson.SetParent(f)

}
