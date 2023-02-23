package kerman

import (
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kerman/bam"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kerman/hormuz"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kerman/kerman"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kerman/sirjan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kerman/sistan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KermanDuke interface {
	feud.Duke
	CBam巴姆() bam.BamCounty
	CHormuz霍尔木兹() hormuz.HormuzCounty
	CKerman克尔曼() kerman.KermanCounty
	CSirjan锡尔詹() sirjan.SirjanCounty
	CSistan古夫斯山() sistan.SistanCounty
}

type 克尔曼KermanDuke struct {
	feud.BaseDuke
	巴姆Bam      bam.BamCounty
	霍尔木兹Hormuz hormuz.HormuzCounty
	克尔曼Kerman  kerman.KermanCounty
	锡尔詹Sirjan  sirjan.SirjanCounty
	古夫斯山Sistan sistan.SistanCounty
}

func (d *克尔曼KermanDuke) CBam巴姆() bam.BamCounty {
	return d.巴姆Bam
}

func (d *克尔曼KermanDuke) CHormuz霍尔木兹() hormuz.HormuzCounty {
	return d.霍尔木兹Hormuz
}

func (d *克尔曼KermanDuke) CKerman克尔曼() kerman.KermanCounty {
	return d.克尔曼Kerman
}

func (d *克尔曼KermanDuke) CSirjan锡尔詹() sirjan.SirjanCounty {
	return d.锡尔詹Sirjan
}

func (d *克尔曼KermanDuke) CSistan古夫斯山() sistan.SistanCounty {
	return d.古夫斯山Sistan
}

var DKerman克尔曼 KermanDuke = &克尔曼KermanDuke{}

func init() {
	f := DKerman克尔曼.(*克尔曼KermanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kerman",
		TitleName: "克尔曼",
		TitleCode: "d_kerman",
		Counties:  map[string]feud.County{},
	}

	f.巴姆Bam = bam.CBam巴姆
	f.巴姆Bam.SetParent(f)

	f.霍尔木兹Hormuz = hormuz.CHormuz霍尔木兹
	f.霍尔木兹Hormuz.SetParent(f)

	f.克尔曼Kerman = kerman.CKerman克尔曼
	f.克尔曼Kerman.SetParent(f)

	f.锡尔詹Sirjan = sirjan.CSirjan锡尔詹
	f.锡尔詹Sirjan.SetParent(f)

	f.古夫斯山Sistan = sistan.CSistan古夫斯山
	f.古夫斯山Sistan.SetParent(f)

}
