package smolensk

import (
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/smolensk/mstislavl"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/smolensk/roslavl"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/smolensk/smolensk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/smolensk/toropets"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/smolensk/vyazma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SmolenskDuke interface {
	feud.Duke
	CMstislavl姆斯季斯拉夫尔() mstislavl.MstislavlCounty
	CRoslavl罗斯拉夫尔() roslavl.RoslavlCounty
	CSmolensk斯摩棱斯克() smolensk.SmolenskCounty
	CToropets托罗佩茨() toropets.ToropetsCounty
	CVyazma维亚济马() vyazma.VyazmaCounty
}

type 斯摩棱斯克SmolenskDuke struct {
	feud.BaseDuke
	姆斯季斯拉夫尔Mstislavl mstislavl.MstislavlCounty
	罗斯拉夫尔Roslavl     roslavl.RoslavlCounty
	斯摩棱斯克Smolensk    smolensk.SmolenskCounty
	托罗佩茨Toropets     toropets.ToropetsCounty
	维亚济马Vyazma       vyazma.VyazmaCounty
}

func (d *斯摩棱斯克SmolenskDuke) CMstislavl姆斯季斯拉夫尔() mstislavl.MstislavlCounty {
	return d.姆斯季斯拉夫尔Mstislavl
}

func (d *斯摩棱斯克SmolenskDuke) CRoslavl罗斯拉夫尔() roslavl.RoslavlCounty {
	return d.罗斯拉夫尔Roslavl
}

func (d *斯摩棱斯克SmolenskDuke) CSmolensk斯摩棱斯克() smolensk.SmolenskCounty {
	return d.斯摩棱斯克Smolensk
}

func (d *斯摩棱斯克SmolenskDuke) CToropets托罗佩茨() toropets.ToropetsCounty {
	return d.托罗佩茨Toropets
}

func (d *斯摩棱斯克SmolenskDuke) CVyazma维亚济马() vyazma.VyazmaCounty {
	return d.维亚济马Vyazma
}

var DSmolensk斯摩棱斯克 SmolenskDuke = &斯摩棱斯克SmolenskDuke{}

func init() {
	f := DSmolensk斯摩棱斯克.(*斯摩棱斯克SmolenskDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "smolensk",
		TitleName: "斯摩棱斯克",
		TitleCode: "d_smolensk",
		Counties:  map[string]feud.County{},
	}

	f.姆斯季斯拉夫尔Mstislavl = mstislavl.CMstislavl姆斯季斯拉夫尔
	f.姆斯季斯拉夫尔Mstislavl.SetParent(f)

	f.罗斯拉夫尔Roslavl = roslavl.CRoslavl罗斯拉夫尔
	f.罗斯拉夫尔Roslavl.SetParent(f)

	f.斯摩棱斯克Smolensk = smolensk.CSmolensk斯摩棱斯克
	f.斯摩棱斯克Smolensk.SetParent(f)

	f.托罗佩茨Toropets = toropets.CToropets托罗佩茨
	f.托罗佩茨Toropets.SetParent(f)

	f.维亚济马Vyazma = vyazma.CVyazma维亚济马
	f.维亚济马Vyazma.SetParent(f)

}
