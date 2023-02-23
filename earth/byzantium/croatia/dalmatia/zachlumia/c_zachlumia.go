package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZachlumiaCounty interface {
	feud.County
	BBrac布拉奇() feud.Barony
	BKorcula科尔丘拉() feud.Barony
	BLapcan拉普钱() feud.Barony
	BOmis奥米什() feud.Barony
	BSirokibrijeg希罗基布里耶格() feud.Barony
	BStolac斯托拉茨() feud.Barony
	BSton斯通() feud.Barony
	BZavala扎瓦拉() feud.Barony
}

type 扎库卢米亚ZachlumiaCounty struct {
	feud.BaseCounty
	布拉奇Brac             feud.Barony
	科尔丘拉Korcula         feud.Barony
	拉普钱Lapcan           feud.Barony
	奥米什Omis             feud.Barony
	希罗基布里耶格Sirokibrijeg feud.Barony
	斯托拉茨Stolac          feud.Barony
	斯通Ston              feud.Barony
	扎瓦拉Zavala           feud.Barony
}

func (c *扎库卢米亚ZachlumiaCounty) BBrac布拉奇() feud.Barony {
	return c.布拉奇Brac
}

func (c *扎库卢米亚ZachlumiaCounty) BKorcula科尔丘拉() feud.Barony {
	return c.科尔丘拉Korcula
}

func (c *扎库卢米亚ZachlumiaCounty) BLapcan拉普钱() feud.Barony {
	return c.拉普钱Lapcan
}

func (c *扎库卢米亚ZachlumiaCounty) BOmis奥米什() feud.Barony {
	return c.奥米什Omis
}

func (c *扎库卢米亚ZachlumiaCounty) BSirokibrijeg希罗基布里耶格() feud.Barony {
	return c.希罗基布里耶格Sirokibrijeg
}

func (c *扎库卢米亚ZachlumiaCounty) BStolac斯托拉茨() feud.Barony {
	return c.斯托拉茨Stolac
}

func (c *扎库卢米亚ZachlumiaCounty) BSton斯通() feud.Barony {
	return c.斯通Ston
}

func (c *扎库卢米亚ZachlumiaCounty) BZavala扎瓦拉() feud.Barony {
	return c.扎瓦拉Zavala
}

var CZachlumia扎库卢米亚 ZachlumiaCounty = &扎库卢米亚ZachlumiaCounty{}

func init() {
	f := CZachlumia扎库卢米亚.(*扎库卢米亚ZachlumiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "466",
		Title:     "zachlumia",
		TitleName: "扎库卢米亚",
		TitleCode: "c_zachlumia",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉奇Brac = BBrac布拉奇
	f.布拉奇Brac.SetParent(f)

	f.科尔丘拉Korcula = BKorcula科尔丘拉
	f.科尔丘拉Korcula.SetParent(f)

	f.拉普钱Lapcan = BLapcan拉普钱
	f.拉普钱Lapcan.SetParent(f)

	f.奥米什Omis = BOmis奥米什
	f.奥米什Omis.SetParent(f)

	f.希罗基布里耶格Sirokibrijeg = BSirokibrijeg希罗基布里耶格
	f.希罗基布里耶格Sirokibrijeg.SetParent(f)

	f.斯托拉茨Stolac = BStolac斯托拉茨
	f.斯托拉茨Stolac.SetParent(f)

	f.斯通Ston = BSton斯通
	f.斯通Ston.SetParent(f)

	f.扎瓦拉Zavala = BZavala扎瓦拉
	f.扎瓦拉Zavala.SetParent(f)

}
