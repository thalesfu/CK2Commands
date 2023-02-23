package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DamascusCounty interface {
	feud.County
	BAlsanamayn塞奈迈因() feud.Barony
	BDamascus大马士革() feud.Barony
	BDaraa德拉() feud.Barony
	BDuma杜马() feud.Barony
	BIzra伊兹拉() feud.Barony
	BQsarbardawil巴尔达维勒堡() feud.Barony
	BShahba舍赫巴() feud.Barony
	BSuada苏瓦达() feud.Barony
}

type 大马士革DamascusCounty struct {
	feud.BaseCounty
	塞奈迈因Alsanamayn     feud.Barony
	大马士革Damascus       feud.Barony
	德拉Daraa            feud.Barony
	杜马Duma             feud.Barony
	伊兹拉Izra            feud.Barony
	巴尔达维勒堡Qsarbardawil feud.Barony
	舍赫巴Shahba          feud.Barony
	苏瓦达Suada           feud.Barony
}

func (c *大马士革DamascusCounty) BAlsanamayn塞奈迈因() feud.Barony {
	return c.塞奈迈因Alsanamayn
}

func (c *大马士革DamascusCounty) BDamascus大马士革() feud.Barony {
	return c.大马士革Damascus
}

func (c *大马士革DamascusCounty) BDaraa德拉() feud.Barony {
	return c.德拉Daraa
}

func (c *大马士革DamascusCounty) BDuma杜马() feud.Barony {
	return c.杜马Duma
}

func (c *大马士革DamascusCounty) BIzra伊兹拉() feud.Barony {
	return c.伊兹拉Izra
}

func (c *大马士革DamascusCounty) BQsarbardawil巴尔达维勒堡() feud.Barony {
	return c.巴尔达维勒堡Qsarbardawil
}

func (c *大马士革DamascusCounty) BShahba舍赫巴() feud.Barony {
	return c.舍赫巴Shahba
}

func (c *大马士革DamascusCounty) BSuada苏瓦达() feud.Barony {
	return c.苏瓦达Suada
}

var CDamascus大马士革 DamascusCounty = &大马士革DamascusCounty{}

func init() {
	f := CDamascus大马士革.(*大马士革DamascusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "728",
		Title:     "damascus",
		TitleName: "大马士革",
		TitleCode: "c_damascus",
		Baronies:  map[string]feud.Barony{},
	}

	f.塞奈迈因Alsanamayn = BAlsanamayn塞奈迈因
	f.塞奈迈因Alsanamayn.SetParent(f)

	f.大马士革Damascus = BDamascus大马士革
	f.大马士革Damascus.SetParent(f)

	f.德拉Daraa = BDaraa德拉
	f.德拉Daraa.SetParent(f)

	f.杜马Duma = BDuma杜马
	f.杜马Duma.SetParent(f)

	f.伊兹拉Izra = BIzra伊兹拉
	f.伊兹拉Izra.SetParent(f)

	f.巴尔达维勒堡Qsarbardawil = BQsarbardawil巴尔达维勒堡
	f.巴尔达维勒堡Qsarbardawil.SetParent(f)

	f.舍赫巴Shahba = BShahba舍赫巴
	f.舍赫巴Shahba.SetParent(f)

	f.苏瓦达Suada = BSuada苏瓦达
	f.苏瓦达Suada.SetParent(f)

}
