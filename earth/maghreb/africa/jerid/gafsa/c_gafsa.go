package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GafsaCounty interface {
	feud.County
	BGafsa加夫萨() feud.Barony
	BGourine古林() feud.Barony
	BHamma哈马() feud.Barony
	BMadas玛达斯() feud.Barony
	BNafta纳夫塔() feud.Barony
	BTaqyus泰格尤斯() feud.Barony
	BTozeur托泽尔() feud.Barony
}

type 加夫萨GafsaCounty struct {
	feud.BaseCounty
	加夫萨Gafsa   feud.Barony
	古林Gourine  feud.Barony
	哈马Hamma    feud.Barony
	玛达斯Madas   feud.Barony
	纳夫塔Nafta   feud.Barony
	泰格尤斯Taqyus feud.Barony
	托泽尔Tozeur  feud.Barony
}

func (c *加夫萨GafsaCounty) BGafsa加夫萨() feud.Barony {
	return c.加夫萨Gafsa
}

func (c *加夫萨GafsaCounty) BGourine古林() feud.Barony {
	return c.古林Gourine
}

func (c *加夫萨GafsaCounty) BHamma哈马() feud.Barony {
	return c.哈马Hamma
}

func (c *加夫萨GafsaCounty) BMadas玛达斯() feud.Barony {
	return c.玛达斯Madas
}

func (c *加夫萨GafsaCounty) BNafta纳夫塔() feud.Barony {
	return c.纳夫塔Nafta
}

func (c *加夫萨GafsaCounty) BTaqyus泰格尤斯() feud.Barony {
	return c.泰格尤斯Taqyus
}

func (c *加夫萨GafsaCounty) BTozeur托泽尔() feud.Barony {
	return c.托泽尔Tozeur
}

var CGafsa加夫萨 GafsaCounty = &加夫萨GafsaCounty{}

func init() {
	f := CGafsa加夫萨.(*加夫萨GafsaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1726",
		Title:     "gafsa",
		TitleName: "加夫萨",
		TitleCode: "c_gafsa",
		Baronies:  map[string]feud.Barony{},
	}

	f.加夫萨Gafsa = BGafsa加夫萨
	f.加夫萨Gafsa.SetParent(f)

	f.古林Gourine = BGourine古林
	f.古林Gourine.SetParent(f)

	f.哈马Hamma = BHamma哈马
	f.哈马Hamma.SetParent(f)

	f.玛达斯Madas = BMadas玛达斯
	f.玛达斯Madas.SetParent(f)

	f.纳夫塔Nafta = BNafta纳夫塔
	f.纳夫塔Nafta.SetParent(f)

	f.泰格尤斯Taqyus = BTaqyus泰格尤斯
	f.泰格尤斯Taqyus.SetParent(f)

	f.托泽尔Tozeur = BTozeur托泽尔
	f.托泽尔Tozeur.SetParent(f)

}
