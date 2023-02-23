package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyrteCounty interface {
	feud.County
	BAbuhadi阿布哈迪() feud.Barony
	BAssidr锡德尔() feud.Barony
	BLanuf拉努夫() feud.Barony
	BLathrun雷斯伦() feud.Barony
	BNawfalliyah瑙费利耶() feud.Barony
	BSinauen西瑙安() feud.Barony
	BSyrte苏尔特() feud.Barony
}

type 苏尔特SyrteCounty struct {
	feud.BaseCounty
	阿布哈迪Abuhadi     feud.Barony
	锡德尔Assidr       feud.Barony
	拉努夫Lanuf        feud.Barony
	雷斯伦Lathrun      feud.Barony
	瑙费利耶Nawfalliyah feud.Barony
	西瑙安Sinauen      feud.Barony
	苏尔特Syrte        feud.Barony
}

func (c *苏尔特SyrteCounty) BAbuhadi阿布哈迪() feud.Barony {
	return c.阿布哈迪Abuhadi
}

func (c *苏尔特SyrteCounty) BAssidr锡德尔() feud.Barony {
	return c.锡德尔Assidr
}

func (c *苏尔特SyrteCounty) BLanuf拉努夫() feud.Barony {
	return c.拉努夫Lanuf
}

func (c *苏尔特SyrteCounty) BLathrun雷斯伦() feud.Barony {
	return c.雷斯伦Lathrun
}

func (c *苏尔特SyrteCounty) BNawfalliyah瑙费利耶() feud.Barony {
	return c.瑙费利耶Nawfalliyah
}

func (c *苏尔特SyrteCounty) BSinauen西瑙安() feud.Barony {
	return c.西瑙安Sinauen
}

func (c *苏尔特SyrteCounty) BSyrte苏尔特() feud.Barony {
	return c.苏尔特Syrte
}

var CSyrte苏尔特 SyrteCounty = &苏尔特SyrteCounty{}

func init() {
	f := CSyrte苏尔特.(*苏尔特SyrteCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "809",
		Title:     "syrte",
		TitleName: "苏尔特",
		TitleCode: "c_syrte",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布哈迪Abuhadi = BAbuhadi阿布哈迪
	f.阿布哈迪Abuhadi.SetParent(f)

	f.锡德尔Assidr = BAssidr锡德尔
	f.锡德尔Assidr.SetParent(f)

	f.拉努夫Lanuf = BLanuf拉努夫
	f.拉努夫Lanuf.SetParent(f)

	f.雷斯伦Lathrun = BLathrun雷斯伦
	f.雷斯伦Lathrun.SetParent(f)

	f.瑙费利耶Nawfalliyah = BNawfalliyah瑙费利耶
	f.瑙费利耶Nawfalliyah.SetParent(f)

	f.西瑙安Sinauen = BSinauen西瑙安
	f.西瑙安Sinauen.SetParent(f)

	f.苏尔特Syrte = BSyrte苏尔特
	f.苏尔特Syrte.SetParent(f)

}
