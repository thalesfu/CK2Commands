package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HerefordCounty interface {
	feud.County
	BArchenfield阿肯菲尔德() feud.Barony
	BBrobury布罗伯里() feud.Barony
	BClifford克利福德() feud.Barony
	BHereford赫里福德() feud.Barony
	BLedbury莱德伯里() feud.Barony
	BLeominster莱姆斯特() feud.Barony
}

type 赫里福德HerefordCounty struct {
	feud.BaseCounty
	阿肯菲尔德Archenfield feud.Barony
	布罗伯里Brobury      feud.Barony
	克利福德Clifford     feud.Barony
	赫里福德Hereford     feud.Barony
	莱德伯里Ledbury      feud.Barony
	莱姆斯特Leominster   feud.Barony
}

func (c *赫里福德HerefordCounty) BArchenfield阿肯菲尔德() feud.Barony {
	return c.阿肯菲尔德Archenfield
}

func (c *赫里福德HerefordCounty) BBrobury布罗伯里() feud.Barony {
	return c.布罗伯里Brobury
}

func (c *赫里福德HerefordCounty) BClifford克利福德() feud.Barony {
	return c.克利福德Clifford
}

func (c *赫里福德HerefordCounty) BHereford赫里福德() feud.Barony {
	return c.赫里福德Hereford
}

func (c *赫里福德HerefordCounty) BLedbury莱德伯里() feud.Barony {
	return c.莱德伯里Ledbury
}

func (c *赫里福德HerefordCounty) BLeominster莱姆斯特() feud.Barony {
	return c.莱姆斯特Leominster
}

var CHereford赫里福德 HerefordCounty = &赫里福德HerefordCounty{}

func init() {
	f := CHereford赫里福德.(*赫里福德HerefordCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "17",
		Title:     "hereford",
		TitleName: "赫里福德",
		TitleCode: "c_hereford",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿肯菲尔德Archenfield = BArchenfield阿肯菲尔德
	f.阿肯菲尔德Archenfield.SetParent(f)

	f.布罗伯里Brobury = BBrobury布罗伯里
	f.布罗伯里Brobury.SetParent(f)

	f.克利福德Clifford = BClifford克利福德
	f.克利福德Clifford.SetParent(f)

	f.赫里福德Hereford = BHereford赫里福德
	f.赫里福德Hereford.SetParent(f)

	f.莱德伯里Ledbury = BLedbury莱德伯里
	f.莱德伯里Ledbury.SetParent(f)

	f.莱姆斯特Leominster = BLeominster莱姆斯特
	f.莱姆斯特Leominster.SetParent(f)

}
