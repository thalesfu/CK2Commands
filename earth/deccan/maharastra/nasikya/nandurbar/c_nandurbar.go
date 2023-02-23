package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NandurbarCounty interface {
	feud.County
	BBhambhagiri梵婆耆厘() feud.Barony
	BChandor尚多尔() feud.Barony
	BMalegaon摩戾伽罗摩() feud.Barony
	BMayuragiri摩由罗耆厘() feud.Barony
	BNandurbar难豆罗婆罗() feud.Barony
	BSongarh僧姞利呬() feud.Barony
	BSultanpur苏坦普尔() feud.Barony
}

type 难豆罗婆罗NandurbarCounty struct {
	feud.BaseCounty
	梵婆耆厘Bhambhagiri feud.Barony
	尚多尔Chandor      feud.Barony
	摩戾伽罗摩Malegaon   feud.Barony
	摩由罗耆厘Mayuragiri feud.Barony
	难豆罗婆罗Nandurbar  feud.Barony
	僧姞利呬Songarh     feud.Barony
	苏坦普尔Sultanpur   feud.Barony
}

func (c *难豆罗婆罗NandurbarCounty) BBhambhagiri梵婆耆厘() feud.Barony {
	return c.梵婆耆厘Bhambhagiri
}

func (c *难豆罗婆罗NandurbarCounty) BChandor尚多尔() feud.Barony {
	return c.尚多尔Chandor
}

func (c *难豆罗婆罗NandurbarCounty) BMalegaon摩戾伽罗摩() feud.Barony {
	return c.摩戾伽罗摩Malegaon
}

func (c *难豆罗婆罗NandurbarCounty) BMayuragiri摩由罗耆厘() feud.Barony {
	return c.摩由罗耆厘Mayuragiri
}

func (c *难豆罗婆罗NandurbarCounty) BNandurbar难豆罗婆罗() feud.Barony {
	return c.难豆罗婆罗Nandurbar
}

func (c *难豆罗婆罗NandurbarCounty) BSongarh僧姞利呬() feud.Barony {
	return c.僧姞利呬Songarh
}

func (c *难豆罗婆罗NandurbarCounty) BSultanpur苏坦普尔() feud.Barony {
	return c.苏坦普尔Sultanpur
}

var CNandurbar难豆罗婆罗 NandurbarCounty = &难豆罗婆罗NandurbarCounty{}

func init() {
	f := CNandurbar难豆罗婆罗.(*难豆罗婆罗NandurbarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1264",
		Title:     "nandurbar",
		TitleName: "难豆罗婆罗",
		TitleCode: "c_nandurbar",
		Baronies:  map[string]feud.Barony{},
	}

	f.梵婆耆厘Bhambhagiri = BBhambhagiri梵婆耆厘
	f.梵婆耆厘Bhambhagiri.SetParent(f)

	f.尚多尔Chandor = BChandor尚多尔
	f.尚多尔Chandor.SetParent(f)

	f.摩戾伽罗摩Malegaon = BMalegaon摩戾伽罗摩
	f.摩戾伽罗摩Malegaon.SetParent(f)

	f.摩由罗耆厘Mayuragiri = BMayuragiri摩由罗耆厘
	f.摩由罗耆厘Mayuragiri.SetParent(f)

	f.难豆罗婆罗Nandurbar = BNandurbar难豆罗婆罗
	f.难豆罗婆罗Nandurbar.SetParent(f)

	f.僧姞利呬Songarh = BSongarh僧姞利呬
	f.僧姞利呬Songarh.SetParent(f)

	f.苏坦普尔Sultanpur = BSultanpur苏坦普尔
	f.苏坦普尔Sultanpur.SetParent(f)

}
