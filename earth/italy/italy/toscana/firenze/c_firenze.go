package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FirenzeCounty interface {
	feud.County
	BCastello卡斯泰洛() feud.Barony
	BCertaldo切塔尔多() feud.Barony
	BEmpoli恩波利() feud.Barony
	BFiesole菲耶索莱() feud.Barony
	BFirenze佛罗伦萨() feud.Barony
	BImpruneta因普鲁内塔() feud.Barony
	BPrato普拉托() feud.Barony
}

type 佛罗伦萨FirenzeCounty struct {
	feud.BaseCounty
	卡斯泰洛Castello   feud.Barony
	切塔尔多Certaldo   feud.Barony
	恩波利Empoli      feud.Barony
	菲耶索莱Fiesole    feud.Barony
	佛罗伦萨Firenze    feud.Barony
	因普鲁内塔Impruneta feud.Barony
	普拉托Prato       feud.Barony
}

func (c *佛罗伦萨FirenzeCounty) BCastello卡斯泰洛() feud.Barony {
	return c.卡斯泰洛Castello
}

func (c *佛罗伦萨FirenzeCounty) BCertaldo切塔尔多() feud.Barony {
	return c.切塔尔多Certaldo
}

func (c *佛罗伦萨FirenzeCounty) BEmpoli恩波利() feud.Barony {
	return c.恩波利Empoli
}

func (c *佛罗伦萨FirenzeCounty) BFiesole菲耶索莱() feud.Barony {
	return c.菲耶索莱Fiesole
}

func (c *佛罗伦萨FirenzeCounty) BFirenze佛罗伦萨() feud.Barony {
	return c.佛罗伦萨Firenze
}

func (c *佛罗伦萨FirenzeCounty) BImpruneta因普鲁内塔() feud.Barony {
	return c.因普鲁内塔Impruneta
}

func (c *佛罗伦萨FirenzeCounty) BPrato普拉托() feud.Barony {
	return c.普拉托Prato
}

var CFirenze佛罗伦萨 FirenzeCounty = &佛罗伦萨FirenzeCounty{}

func init() {
	f := CFirenze佛罗伦萨.(*佛罗伦萨FirenzeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "328",
		Title:     "firenze",
		TitleName: "佛罗伦萨",
		TitleCode: "c_firenze",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡斯泰洛Castello = BCastello卡斯泰洛
	f.卡斯泰洛Castello.SetParent(f)

	f.切塔尔多Certaldo = BCertaldo切塔尔多
	f.切塔尔多Certaldo.SetParent(f)

	f.恩波利Empoli = BEmpoli恩波利
	f.恩波利Empoli.SetParent(f)

	f.菲耶索莱Fiesole = BFiesole菲耶索莱
	f.菲耶索莱Fiesole.SetParent(f)

	f.佛罗伦萨Firenze = BFirenze佛罗伦萨
	f.佛罗伦萨Firenze.SetParent(f)

	f.因普鲁内塔Impruneta = BImpruneta因普鲁内塔
	f.因普鲁内塔Impruneta.SetParent(f)

	f.普拉托Prato = BPrato普拉托
	f.普拉托Prato.SetParent(f)

}
