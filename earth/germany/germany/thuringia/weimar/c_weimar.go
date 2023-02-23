package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WeimarCounty interface {
	feud.County
	BApoldoa阿波尔达() feud.Barony
	BGera格拉() feud.Barony
	BGotha哥达() feud.Barony
	BJena耶拿() feud.Barony
	BMemelsen梅梅尔森() feud.Barony
	BNordhausen诺德豪森() feud.Barony
	BWalkenried瓦尔肯里德() feud.Barony
	BWeimar魏玛() feud.Barony
}

type 魏玛WeimarCounty struct {
	feud.BaseCounty
	阿波尔达Apoldoa     feud.Barony
	格拉Gera          feud.Barony
	哥达Gotha         feud.Barony
	耶拿Jena          feud.Barony
	梅梅尔森Memelsen    feud.Barony
	诺德豪森Nordhausen  feud.Barony
	瓦尔肯里德Walkenried feud.Barony
	魏玛Weimar        feud.Barony
}

func (c *魏玛WeimarCounty) BApoldoa阿波尔达() feud.Barony {
	return c.阿波尔达Apoldoa
}

func (c *魏玛WeimarCounty) BGera格拉() feud.Barony {
	return c.格拉Gera
}

func (c *魏玛WeimarCounty) BGotha哥达() feud.Barony {
	return c.哥达Gotha
}

func (c *魏玛WeimarCounty) BJena耶拿() feud.Barony {
	return c.耶拿Jena
}

func (c *魏玛WeimarCounty) BMemelsen梅梅尔森() feud.Barony {
	return c.梅梅尔森Memelsen
}

func (c *魏玛WeimarCounty) BNordhausen诺德豪森() feud.Barony {
	return c.诺德豪森Nordhausen
}

func (c *魏玛WeimarCounty) BWalkenried瓦尔肯里德() feud.Barony {
	return c.瓦尔肯里德Walkenried
}

func (c *魏玛WeimarCounty) BWeimar魏玛() feud.Barony {
	return c.魏玛Weimar
}

var CWeimar魏玛 WeimarCounty = &魏玛WeimarCounty{}

func init() {
	f := CWeimar魏玛.(*魏玛WeimarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "256",
		Title:     "weimar",
		TitleName: "魏玛",
		TitleCode: "c_weimar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿波尔达Apoldoa = BApoldoa阿波尔达
	f.阿波尔达Apoldoa.SetParent(f)

	f.格拉Gera = BGera格拉
	f.格拉Gera.SetParent(f)

	f.哥达Gotha = BGotha哥达
	f.哥达Gotha.SetParent(f)

	f.耶拿Jena = BJena耶拿
	f.耶拿Jena.SetParent(f)

	f.梅梅尔森Memelsen = BMemelsen梅梅尔森
	f.梅梅尔森Memelsen.SetParent(f)

	f.诺德豪森Nordhausen = BNordhausen诺德豪森
	f.诺德豪森Nordhausen.SetParent(f)

	f.瓦尔肯里德Walkenried = BWalkenried瓦尔肯里德
	f.瓦尔肯里德Walkenried.SetParent(f)

	f.魏玛Weimar = BWeimar魏玛
	f.魏玛Weimar.SetParent(f)

}
