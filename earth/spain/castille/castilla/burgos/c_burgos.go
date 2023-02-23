package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BurgosCounty interface {
	feud.County
	BAguilardecampo阿吉拉尔德坎波奥() feud.Barony
	BArandadeduero杜罗河畔阿兰达() feud.Barony
	BBurgos布尔戈斯() feud.Barony
	BCarrion卡里翁() feud.Barony
	BCastrobarte卡斯特罗巴尔特() feud.Barony
	BMirandadeebro埃布罗河畔米兰达() feud.Barony
	BSilos锡洛斯() feud.Barony
}

type 布尔戈斯BurgosCounty struct {
	feud.BaseCounty
	阿吉拉尔德坎波奥Aguilardecampo feud.Barony
	杜罗河畔阿兰达Arandadeduero   feud.Barony
	布尔戈斯Burgos             feud.Barony
	卡里翁Carrion             feud.Barony
	卡斯特罗巴尔特Castrobarte     feud.Barony
	埃布罗河畔米兰达Mirandadeebro  feud.Barony
	锡洛斯Silos               feud.Barony
}

func (c *布尔戈斯BurgosCounty) BAguilardecampo阿吉拉尔德坎波奥() feud.Barony {
	return c.阿吉拉尔德坎波奥Aguilardecampo
}

func (c *布尔戈斯BurgosCounty) BArandadeduero杜罗河畔阿兰达() feud.Barony {
	return c.杜罗河畔阿兰达Arandadeduero
}

func (c *布尔戈斯BurgosCounty) BBurgos布尔戈斯() feud.Barony {
	return c.布尔戈斯Burgos
}

func (c *布尔戈斯BurgosCounty) BCarrion卡里翁() feud.Barony {
	return c.卡里翁Carrion
}

func (c *布尔戈斯BurgosCounty) BCastrobarte卡斯特罗巴尔特() feud.Barony {
	return c.卡斯特罗巴尔特Castrobarte
}

func (c *布尔戈斯BurgosCounty) BMirandadeebro埃布罗河畔米兰达() feud.Barony {
	return c.埃布罗河畔米兰达Mirandadeebro
}

func (c *布尔戈斯BurgosCounty) BSilos锡洛斯() feud.Barony {
	return c.锡洛斯Silos
}

var CBurgos布尔戈斯 BurgosCounty = &布尔戈斯BurgosCounty{}

func init() {
	f := CBurgos布尔戈斯.(*布尔戈斯BurgosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "199",
		Title:     "burgos",
		TitleName: "布尔戈斯",
		TitleCode: "c_burgos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿吉拉尔德坎波奥Aguilardecampo = BAguilardecampo阿吉拉尔德坎波奥
	f.阿吉拉尔德坎波奥Aguilardecampo.SetParent(f)

	f.杜罗河畔阿兰达Arandadeduero = BArandadeduero杜罗河畔阿兰达
	f.杜罗河畔阿兰达Arandadeduero.SetParent(f)

	f.布尔戈斯Burgos = BBurgos布尔戈斯
	f.布尔戈斯Burgos.SetParent(f)

	f.卡里翁Carrion = BCarrion卡里翁
	f.卡里翁Carrion.SetParent(f)

	f.卡斯特罗巴尔特Castrobarte = BCastrobarte卡斯特罗巴尔特
	f.卡斯特罗巴尔特Castrobarte.SetParent(f)

	f.埃布罗河畔米兰达Mirandadeebro = BMirandadeebro埃布罗河畔米兰达
	f.埃布罗河畔米兰达Mirandadeebro.SetParent(f)

	f.锡洛斯Silos = BSilos锡洛斯
	f.锡洛斯Silos.SetParent(f)

}
