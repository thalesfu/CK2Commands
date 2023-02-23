package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThessaliaCounty interface {
	feud.County
	BDamasis达玛斯() feud.Barony
	BKastri卡斯特里() feud.Barony
	BLarissa拉里萨() feud.Barony
	BNeopetra纽佩特拉() feud.Barony
	BPharsalos法尔萨洛斯() feud.Barony
	BStagi斯塔基() feud.Barony
	BTrikkala特里卡拉() feud.Barony
	BVolos沃洛斯() feud.Barony
}

type 色萨利ThessaliaCounty struct {
	feud.BaseCounty
	达玛斯Damasis     feud.Barony
	卡斯特里Kastri     feud.Barony
	拉里萨Larissa     feud.Barony
	纽佩特拉Neopetra   feud.Barony
	法尔萨洛斯Pharsalos feud.Barony
	斯塔基Stagi       feud.Barony
	特里卡拉Trikkala   feud.Barony
	沃洛斯Volos       feud.Barony
}

func (c *色萨利ThessaliaCounty) BDamasis达玛斯() feud.Barony {
	return c.达玛斯Damasis
}

func (c *色萨利ThessaliaCounty) BKastri卡斯特里() feud.Barony {
	return c.卡斯特里Kastri
}

func (c *色萨利ThessaliaCounty) BLarissa拉里萨() feud.Barony {
	return c.拉里萨Larissa
}

func (c *色萨利ThessaliaCounty) BNeopetra纽佩特拉() feud.Barony {
	return c.纽佩特拉Neopetra
}

func (c *色萨利ThessaliaCounty) BPharsalos法尔萨洛斯() feud.Barony {
	return c.法尔萨洛斯Pharsalos
}

func (c *色萨利ThessaliaCounty) BStagi斯塔基() feud.Barony {
	return c.斯塔基Stagi
}

func (c *色萨利ThessaliaCounty) BTrikkala特里卡拉() feud.Barony {
	return c.特里卡拉Trikkala
}

func (c *色萨利ThessaliaCounty) BVolos沃洛斯() feud.Barony {
	return c.沃洛斯Volos
}

var CThessalia色萨利 ThessaliaCounty = &色萨利ThessaliaCounty{}

func init() {
	f := CThessalia色萨利.(*色萨利ThessaliaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "489",
		Title:     "thessalia",
		TitleName: "色萨利",
		TitleCode: "c_thessalia",
		Baronies:  map[string]feud.Barony{},
	}

	f.达玛斯Damasis = BDamasis达玛斯
	f.达玛斯Damasis.SetParent(f)

	f.卡斯特里Kastri = BKastri卡斯特里
	f.卡斯特里Kastri.SetParent(f)

	f.拉里萨Larissa = BLarissa拉里萨
	f.拉里萨Larissa.SetParent(f)

	f.纽佩特拉Neopetra = BNeopetra纽佩特拉
	f.纽佩特拉Neopetra.SetParent(f)

	f.法尔萨洛斯Pharsalos = BPharsalos法尔萨洛斯
	f.法尔萨洛斯Pharsalos.SetParent(f)

	f.斯塔基Stagi = BStagi斯塔基
	f.斯塔基Stagi.SetParent(f)

	f.特里卡拉Trikkala = BTrikkala特里卡拉
	f.特里卡拉Trikkala.SetParent(f)

	f.沃洛斯Volos = BVolos沃洛斯
	f.沃洛斯Volos.SetParent(f)

}
