package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KusinagaraCounty interface {
	feud.County
	BAlampura阿蓝补罗() feud.Barony
	BChapra遮波罗() feud.Barony
	BKusinagara拘尸那揭罗() feud.Barony
	BPadrauna钵陀劳那() feud.Barony
	BPava波伐() feud.Barony
	BRamagrama罗摩伽() feud.Barony
	BTrinabad帝利那跋陀() feud.Barony
}

type 拘尸那揭罗KusinagaraCounty struct {
	feud.BaseCounty
	阿蓝补罗Alampura    feud.Barony
	遮波罗Chapra       feud.Barony
	拘尸那揭罗Kusinagara feud.Barony
	钵陀劳那Padrauna    feud.Barony
	波伐Pava          feud.Barony
	罗摩伽Ramagrama    feud.Barony
	帝利那跋陀Trinabad   feud.Barony
}

func (c *拘尸那揭罗KusinagaraCounty) BAlampura阿蓝补罗() feud.Barony {
	return c.阿蓝补罗Alampura
}

func (c *拘尸那揭罗KusinagaraCounty) BChapra遮波罗() feud.Barony {
	return c.遮波罗Chapra
}

func (c *拘尸那揭罗KusinagaraCounty) BKusinagara拘尸那揭罗() feud.Barony {
	return c.拘尸那揭罗Kusinagara
}

func (c *拘尸那揭罗KusinagaraCounty) BPadrauna钵陀劳那() feud.Barony {
	return c.钵陀劳那Padrauna
}

func (c *拘尸那揭罗KusinagaraCounty) BPava波伐() feud.Barony {
	return c.波伐Pava
}

func (c *拘尸那揭罗KusinagaraCounty) BRamagrama罗摩伽() feud.Barony {
	return c.罗摩伽Ramagrama
}

func (c *拘尸那揭罗KusinagaraCounty) BTrinabad帝利那跋陀() feud.Barony {
	return c.帝利那跋陀Trinabad
}

var CKusinagara拘尸那揭罗 KusinagaraCounty = &拘尸那揭罗KusinagaraCounty{}

func init() {
	f := CKusinagara拘尸那揭罗.(*拘尸那揭罗KusinagaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1162",
		Title:     "kusinagara",
		TitleName: "拘尸那揭罗",
		TitleCode: "c_kusinagara",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿蓝补罗Alampura = BAlampura阿蓝补罗
	f.阿蓝补罗Alampura.SetParent(f)

	f.遮波罗Chapra = BChapra遮波罗
	f.遮波罗Chapra.SetParent(f)

	f.拘尸那揭罗Kusinagara = BKusinagara拘尸那揭罗
	f.拘尸那揭罗Kusinagara.SetParent(f)

	f.钵陀劳那Padrauna = BPadrauna钵陀劳那
	f.钵陀劳那Padrauna.SetParent(f)

	f.波伐Pava = BPava波伐
	f.波伐Pava.SetParent(f)

	f.罗摩伽Ramagrama = BRamagrama罗摩伽
	f.罗摩伽Ramagrama.SetParent(f)

	f.帝利那跋陀Trinabad = BTrinabad帝利那跋陀
	f.帝利那跋陀Trinabad.SetParent(f)

}
