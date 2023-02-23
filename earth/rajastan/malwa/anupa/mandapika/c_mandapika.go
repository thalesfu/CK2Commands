package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MandapikaCounty interface {
	feud.County
	BAvasgarh阿伐迦罗() feud.Barony
	BBawangaja婆伐那迦() feud.Barony
	BBordhoman巴尔达曼() feud.Barony
	BBorgarh浦迦尔() feud.Barony
	BDatrauda陀多罗陀() feud.Barony
	BMaheshwar摩醯湿伐罗() feud.Barony
	BMandapika曼荼卑迦() feud.Barony
}

type 曼荼卑迦MandapikaCounty struct {
	feud.BaseCounty
	阿伐迦罗Avasgarh   feud.Barony
	婆伐那迦Bawangaja  feud.Barony
	巴尔达曼Bordhoman  feud.Barony
	浦迦尔Borgarh     feud.Barony
	陀多罗陀Datrauda   feud.Barony
	摩醯湿伐罗Maheshwar feud.Barony
	曼荼卑迦Mandapika  feud.Barony
}

func (c *曼荼卑迦MandapikaCounty) BAvasgarh阿伐迦罗() feud.Barony {
	return c.阿伐迦罗Avasgarh
}

func (c *曼荼卑迦MandapikaCounty) BBawangaja婆伐那迦() feud.Barony {
	return c.婆伐那迦Bawangaja
}

func (c *曼荼卑迦MandapikaCounty) BBordhoman巴尔达曼() feud.Barony {
	return c.巴尔达曼Bordhoman
}

func (c *曼荼卑迦MandapikaCounty) BBorgarh浦迦尔() feud.Barony {
	return c.浦迦尔Borgarh
}

func (c *曼荼卑迦MandapikaCounty) BDatrauda陀多罗陀() feud.Barony {
	return c.陀多罗陀Datrauda
}

func (c *曼荼卑迦MandapikaCounty) BMaheshwar摩醯湿伐罗() feud.Barony {
	return c.摩醯湿伐罗Maheshwar
}

func (c *曼荼卑迦MandapikaCounty) BMandapika曼荼卑迦() feud.Barony {
	return c.曼荼卑迦Mandapika
}

var CMandapika曼荼卑迦 MandapikaCounty = &曼荼卑迦MandapikaCounty{}

func init() {
	f := CMandapika曼荼卑迦.(*曼荼卑迦MandapikaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1147",
		Title:     "mandapika",
		TitleName: "曼荼卑迦",
		TitleCode: "c_mandapika",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伐迦罗Avasgarh = BAvasgarh阿伐迦罗
	f.阿伐迦罗Avasgarh.SetParent(f)

	f.婆伐那迦Bawangaja = BBawangaja婆伐那迦
	f.婆伐那迦Bawangaja.SetParent(f)

	f.巴尔达曼Bordhoman = BBordhoman巴尔达曼
	f.巴尔达曼Bordhoman.SetParent(f)

	f.浦迦尔Borgarh = BBorgarh浦迦尔
	f.浦迦尔Borgarh.SetParent(f)

	f.陀多罗陀Datrauda = BDatrauda陀多罗陀
	f.陀多罗陀Datrauda.SetParent(f)

	f.摩醯湿伐罗Maheshwar = BMaheshwar摩醯湿伐罗
	f.摩醯湿伐罗Maheshwar.SetParent(f)

	f.曼荼卑迦Mandapika = BMandapika曼荼卑迦
	f.曼荼卑迦Mandapika.SetParent(f)

}
