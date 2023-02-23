package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OldenburgCounty interface {
	feud.County
	BCloppenburg克洛彭堡() feud.Barony
	BDelmenhorst代尔门霍斯特() feud.Barony
	BJever耶弗尔() feud.Barony
	BKniphausen克尼普豪森() feud.Barony
	BLoningen勒宁根() feud.Barony
	BNordenham诺登哈姆() feud.Barony
	BOldenburg奥尔登堡() feud.Barony
	BVarel法勒尔() feud.Barony
}

type 奥尔登堡OldenburgCounty struct {
	feud.BaseCounty
	克洛彭堡Cloppenburg   feud.Barony
	代尔门霍斯特Delmenhorst feud.Barony
	耶弗尔Jever          feud.Barony
	克尼普豪森Kniphausen   feud.Barony
	勒宁根Loningen       feud.Barony
	诺登哈姆Nordenham     feud.Barony
	奥尔登堡Oldenburg     feud.Barony
	法勒尔Varel          feud.Barony
}

func (c *奥尔登堡OldenburgCounty) BCloppenburg克洛彭堡() feud.Barony {
	return c.克洛彭堡Cloppenburg
}

func (c *奥尔登堡OldenburgCounty) BDelmenhorst代尔门霍斯特() feud.Barony {
	return c.代尔门霍斯特Delmenhorst
}

func (c *奥尔登堡OldenburgCounty) BJever耶弗尔() feud.Barony {
	return c.耶弗尔Jever
}

func (c *奥尔登堡OldenburgCounty) BKniphausen克尼普豪森() feud.Barony {
	return c.克尼普豪森Kniphausen
}

func (c *奥尔登堡OldenburgCounty) BLoningen勒宁根() feud.Barony {
	return c.勒宁根Loningen
}

func (c *奥尔登堡OldenburgCounty) BNordenham诺登哈姆() feud.Barony {
	return c.诺登哈姆Nordenham
}

func (c *奥尔登堡OldenburgCounty) BOldenburg奥尔登堡() feud.Barony {
	return c.奥尔登堡Oldenburg
}

func (c *奥尔登堡OldenburgCounty) BVarel法勒尔() feud.Barony {
	return c.法勒尔Varel
}

var COldenburg奥尔登堡 OldenburgCounty = &奥尔登堡OldenburgCounty{}

func init() {
	f := COldenburg奥尔登堡.(*奥尔登堡OldenburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "86",
		Title:     "oldenburg",
		TitleName: "奥尔登堡",
		TitleCode: "c_oldenburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.克洛彭堡Cloppenburg = BCloppenburg克洛彭堡
	f.克洛彭堡Cloppenburg.SetParent(f)

	f.代尔门霍斯特Delmenhorst = BDelmenhorst代尔门霍斯特
	f.代尔门霍斯特Delmenhorst.SetParent(f)

	f.耶弗尔Jever = BJever耶弗尔
	f.耶弗尔Jever.SetParent(f)

	f.克尼普豪森Kniphausen = BKniphausen克尼普豪森
	f.克尼普豪森Kniphausen.SetParent(f)

	f.勒宁根Loningen = BLoningen勒宁根
	f.勒宁根Loningen.SetParent(f)

	f.诺登哈姆Nordenham = BNordenham诺登哈姆
	f.诺登哈姆Nordenham.SetParent(f)

	f.奥尔登堡Oldenburg = BOldenburg奥尔登堡
	f.奥尔登堡Oldenburg.SetParent(f)

	f.法勒尔Varel = BVarel法勒尔
	f.法勒尔Varel.SetParent(f)

}
