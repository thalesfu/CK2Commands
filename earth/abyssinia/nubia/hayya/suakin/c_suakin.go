package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SuakinCounty interface {
	feud.County
	BDerkin德尔金() feud.Barony
	BDubarua杜巴鲁阿() feud.Barony
	BKomelshokafa科梅尔索卡法() feud.Barony
	BSalum萨鲁姆() feud.Barony
	BSawakin萨瓦金() feud.Barony
	BTaiwi泰威() feud.Barony
	BTisiamti提斯阿姆提() feud.Barony
}

type 萨瓦金SuakinCounty struct {
	feud.BaseCounty
	德尔金Derkin          feud.Barony
	杜巴鲁阿Dubarua        feud.Barony
	科梅尔索卡法Komelshokafa feud.Barony
	萨鲁姆Salum           feud.Barony
	萨瓦金Sawakin         feud.Barony
	泰威Taiwi            feud.Barony
	提斯阿姆提Tisiamti      feud.Barony
}

func (c *萨瓦金SuakinCounty) BDerkin德尔金() feud.Barony {
	return c.德尔金Derkin
}

func (c *萨瓦金SuakinCounty) BDubarua杜巴鲁阿() feud.Barony {
	return c.杜巴鲁阿Dubarua
}

func (c *萨瓦金SuakinCounty) BKomelshokafa科梅尔索卡法() feud.Barony {
	return c.科梅尔索卡法Komelshokafa
}

func (c *萨瓦金SuakinCounty) BSalum萨鲁姆() feud.Barony {
	return c.萨鲁姆Salum
}

func (c *萨瓦金SuakinCounty) BSawakin萨瓦金() feud.Barony {
	return c.萨瓦金Sawakin
}

func (c *萨瓦金SuakinCounty) BTaiwi泰威() feud.Barony {
	return c.泰威Taiwi
}

func (c *萨瓦金SuakinCounty) BTisiamti提斯阿姆提() feud.Barony {
	return c.提斯阿姆提Tisiamti
}

var CSuakin萨瓦金 SuakinCounty = &萨瓦金SuakinCounty{}

func init() {
	f := CSuakin萨瓦金.(*萨瓦金SuakinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1330",
		Title:     "suakin",
		TitleName: "萨瓦金",
		TitleCode: "c_suakin",
		Baronies:  map[string]feud.Barony{},
	}

	f.德尔金Derkin = BDerkin德尔金
	f.德尔金Derkin.SetParent(f)

	f.杜巴鲁阿Dubarua = BDubarua杜巴鲁阿
	f.杜巴鲁阿Dubarua.SetParent(f)

	f.科梅尔索卡法Komelshokafa = BKomelshokafa科梅尔索卡法
	f.科梅尔索卡法Komelshokafa.SetParent(f)

	f.萨鲁姆Salum = BSalum萨鲁姆
	f.萨鲁姆Salum.SetParent(f)

	f.萨瓦金Sawakin = BSawakin萨瓦金
	f.萨瓦金Sawakin.SetParent(f)

	f.泰威Taiwi = BTaiwi泰威
	f.泰威Taiwi.SetParent(f)

	f.提斯阿姆提Tisiamti = BTisiamti提斯阿姆提
	f.提斯阿姆提Tisiamti.SetParent(f)

}
