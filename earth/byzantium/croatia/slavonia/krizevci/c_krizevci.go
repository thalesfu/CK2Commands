package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KrizevciCounty interface {
	feud.County
	BDurdevac久尔杰瓦茨() feud.Barony
	BKoprivnica科普里夫尼察() feud.Barony
	BKrizevci克里热夫齐() feud.Barony
	BOsijek奥西耶克() feud.Barony
	BPozega波热加() feud.Barony
	BVinkovci温科夫齐() feud.Barony
	BVirovitica维特罗维蒂察() feud.Barony
}

type 克里热夫齐KrizevciCounty struct {
	feud.BaseCounty
	久尔杰瓦茨Durdevac    feud.Barony
	科普里夫尼察Koprivnica feud.Barony
	克里热夫齐Krizevci    feud.Barony
	奥西耶克Osijek       feud.Barony
	波热加Pozega        feud.Barony
	温科夫齐Vinkovci     feud.Barony
	维特罗维蒂察Virovitica feud.Barony
}

func (c *克里热夫齐KrizevciCounty) BDurdevac久尔杰瓦茨() feud.Barony {
	return c.久尔杰瓦茨Durdevac
}

func (c *克里热夫齐KrizevciCounty) BKoprivnica科普里夫尼察() feud.Barony {
	return c.科普里夫尼察Koprivnica
}

func (c *克里热夫齐KrizevciCounty) BKrizevci克里热夫齐() feud.Barony {
	return c.克里热夫齐Krizevci
}

func (c *克里热夫齐KrizevciCounty) BOsijek奥西耶克() feud.Barony {
	return c.奥西耶克Osijek
}

func (c *克里热夫齐KrizevciCounty) BPozega波热加() feud.Barony {
	return c.波热加Pozega
}

func (c *克里热夫齐KrizevciCounty) BVinkovci温科夫齐() feud.Barony {
	return c.温科夫齐Vinkovci
}

func (c *克里热夫齐KrizevciCounty) BVirovitica维特罗维蒂察() feud.Barony {
	return c.维特罗维蒂察Virovitica
}

var CKrizevci克里热夫齐 KrizevciCounty = &克里热夫齐KrizevciCounty{}

func init() {
	f := CKrizevci克里热夫齐.(*克里热夫齐KrizevciCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "462",
		Title:     "krizevci",
		TitleName: "克里热夫齐",
		TitleCode: "c_krizevci",
		Baronies:  map[string]feud.Barony{},
	}

	f.久尔杰瓦茨Durdevac = BDurdevac久尔杰瓦茨
	f.久尔杰瓦茨Durdevac.SetParent(f)

	f.科普里夫尼察Koprivnica = BKoprivnica科普里夫尼察
	f.科普里夫尼察Koprivnica.SetParent(f)

	f.克里热夫齐Krizevci = BKrizevci克里热夫齐
	f.克里热夫齐Krizevci.SetParent(f)

	f.奥西耶克Osijek = BOsijek奥西耶克
	f.奥西耶克Osijek.SetParent(f)

	f.波热加Pozega = BPozega波热加
	f.波热加Pozega.SetParent(f)

	f.温科夫齐Vinkovci = BVinkovci温科夫齐
	f.温科夫齐Vinkovci.SetParent(f)

	f.维特罗维蒂察Virovitica = BVirovitica维特罗维蒂察
	f.维特罗维蒂察Virovitica.SetParent(f)

}
