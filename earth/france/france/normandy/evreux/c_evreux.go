package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EvreuxCounty interface {
	feud.County
	BAlencon阿朗松() feud.Barony
	BArgentan阿让唐() feud.Barony
	BEvreux埃夫勒() feud.Barony
	BFalaise法莱斯() feud.Barony
	BHonfleur翁弗勒尔() feud.Barony
	BLisieux利雪() feud.Barony
	BSees塞() feud.Barony
}

type 埃夫勒EvreuxCounty struct {
	feud.BaseCounty
	阿朗松Alencon   feud.Barony
	阿让唐Argentan  feud.Barony
	埃夫勒Evreux    feud.Barony
	法莱斯Falaise   feud.Barony
	翁弗勒尔Honfleur feud.Barony
	利雪Lisieux    feud.Barony
	塞Sees        feud.Barony
}

func (c *埃夫勒EvreuxCounty) BAlencon阿朗松() feud.Barony {
	return c.阿朗松Alencon
}

func (c *埃夫勒EvreuxCounty) BArgentan阿让唐() feud.Barony {
	return c.阿让唐Argentan
}

func (c *埃夫勒EvreuxCounty) BEvreux埃夫勒() feud.Barony {
	return c.埃夫勒Evreux
}

func (c *埃夫勒EvreuxCounty) BFalaise法莱斯() feud.Barony {
	return c.法莱斯Falaise
}

func (c *埃夫勒EvreuxCounty) BHonfleur翁弗勒尔() feud.Barony {
	return c.翁弗勒尔Honfleur
}

func (c *埃夫勒EvreuxCounty) BLisieux利雪() feud.Barony {
	return c.利雪Lisieux
}

func (c *埃夫勒EvreuxCounty) BSees塞() feud.Barony {
	return c.塞Sees
}

var CEvreux埃夫勒 EvreuxCounty = &埃夫勒EvreuxCounty{}

func init() {
	f := CEvreux埃夫勒.(*埃夫勒EvreuxCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "99",
		Title:     "evreux",
		TitleName: "埃夫勒",
		TitleCode: "c_evreux",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿朗松Alencon = BAlencon阿朗松
	f.阿朗松Alencon.SetParent(f)

	f.阿让唐Argentan = BArgentan阿让唐
	f.阿让唐Argentan.SetParent(f)

	f.埃夫勒Evreux = BEvreux埃夫勒
	f.埃夫勒Evreux.SetParent(f)

	f.法莱斯Falaise = BFalaise法莱斯
	f.法莱斯Falaise.SetParent(f)

	f.翁弗勒尔Honfleur = BHonfleur翁弗勒尔
	f.翁弗勒尔Honfleur.SetParent(f)

	f.利雪Lisieux = BLisieux利雪
	f.利雪Lisieux.SetParent(f)

	f.塞Sees = BSees塞
	f.塞Sees.SetParent(f)

}
