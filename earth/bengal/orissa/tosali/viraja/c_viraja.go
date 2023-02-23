package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VirajaCounty interface {
	feud.County
	BAnandpur阿难陀补罗() feud.Barony
	BBaleshvara婆丽湿伐罗() feud.Barony
	BBhadrak跋提城() feud.Barony
	BMahavinayak摩诃毗那夜迦() feud.Barony
	BRaibania拉巴尼亚() feud.Barony
	BSoro索罗() feud.Barony
	BViraja毗罗阇() feud.Barony
}

type 毗罗阇VirajaCounty struct {
	feud.BaseCounty
	阿难陀补罗Anandpur     feud.Barony
	婆丽湿伐罗Baleshvara   feud.Barony
	跋提城Bhadrak        feud.Barony
	摩诃毗那夜迦Mahavinayak feud.Barony
	拉巴尼亚Raibania      feud.Barony
	索罗Soro            feud.Barony
	毗罗阇Viraja         feud.Barony
}

func (c *毗罗阇VirajaCounty) BAnandpur阿难陀补罗() feud.Barony {
	return c.阿难陀补罗Anandpur
}

func (c *毗罗阇VirajaCounty) BBaleshvara婆丽湿伐罗() feud.Barony {
	return c.婆丽湿伐罗Baleshvara
}

func (c *毗罗阇VirajaCounty) BBhadrak跋提城() feud.Barony {
	return c.跋提城Bhadrak
}

func (c *毗罗阇VirajaCounty) BMahavinayak摩诃毗那夜迦() feud.Barony {
	return c.摩诃毗那夜迦Mahavinayak
}

func (c *毗罗阇VirajaCounty) BRaibania拉巴尼亚() feud.Barony {
	return c.拉巴尼亚Raibania
}

func (c *毗罗阇VirajaCounty) BSoro索罗() feud.Barony {
	return c.索罗Soro
}

func (c *毗罗阇VirajaCounty) BViraja毗罗阇() feud.Barony {
	return c.毗罗阇Viraja
}

var CViraja毗罗阇 VirajaCounty = &毗罗阇VirajaCounty{}

func init() {
	f := CViraja毗罗阇.(*毗罗阇VirajaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1231",
		Title:     "viraja",
		TitleName: "毗罗阇",
		TitleCode: "c_viraja",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿难陀补罗Anandpur = BAnandpur阿难陀补罗
	f.阿难陀补罗Anandpur.SetParent(f)

	f.婆丽湿伐罗Baleshvara = BBaleshvara婆丽湿伐罗
	f.婆丽湿伐罗Baleshvara.SetParent(f)

	f.跋提城Bhadrak = BBhadrak跋提城
	f.跋提城Bhadrak.SetParent(f)

	f.摩诃毗那夜迦Mahavinayak = BMahavinayak摩诃毗那夜迦
	f.摩诃毗那夜迦Mahavinayak.SetParent(f)

	f.拉巴尼亚Raibania = BRaibania拉巴尼亚
	f.拉巴尼亚Raibania.SetParent(f)

	f.索罗Soro = BSoro索罗
	f.索罗Soro.SetParent(f)

	f.毗罗阇Viraja = BViraja毗罗阇
	f.毗罗阇Viraja.SetParent(f)

}
