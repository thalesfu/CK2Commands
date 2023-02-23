package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DorsetCounty interface {
	feud.County
	BCorfe科夫() feud.Barony
	BDorchester多切斯特() feud.Barony
	BLyme莱姆() feud.Barony
	BShaftesbury沙夫茨伯里() feud.Barony
	BSherborne舍伯恩() feud.Barony
	BWareham韦勒姆() feud.Barony
	BWeymouth韦茅斯() feud.Barony
	BWimborne温伯恩() feud.Barony
}

type 多塞特DorsetCounty struct {
	feud.BaseCounty
	科夫Corfe          feud.Barony
	多切斯特Dorchester   feud.Barony
	莱姆Lyme           feud.Barony
	沙夫茨伯里Shaftesbury feud.Barony
	舍伯恩Sherborne     feud.Barony
	韦勒姆Wareham       feud.Barony
	韦茅斯Weymouth      feud.Barony
	温伯恩Wimborne      feud.Barony
}

func (c *多塞特DorsetCounty) BCorfe科夫() feud.Barony {
	return c.科夫Corfe
}

func (c *多塞特DorsetCounty) BDorchester多切斯特() feud.Barony {
	return c.多切斯特Dorchester
}

func (c *多塞特DorsetCounty) BLyme莱姆() feud.Barony {
	return c.莱姆Lyme
}

func (c *多塞特DorsetCounty) BShaftesbury沙夫茨伯里() feud.Barony {
	return c.沙夫茨伯里Shaftesbury
}

func (c *多塞特DorsetCounty) BSherborne舍伯恩() feud.Barony {
	return c.舍伯恩Sherborne
}

func (c *多塞特DorsetCounty) BWareham韦勒姆() feud.Barony {
	return c.韦勒姆Wareham
}

func (c *多塞特DorsetCounty) BWeymouth韦茅斯() feud.Barony {
	return c.韦茅斯Weymouth
}

func (c *多塞特DorsetCounty) BWimborne温伯恩() feud.Barony {
	return c.温伯恩Wimborne
}

var CDorset多塞特 DorsetCounty = &多塞特DorsetCounty{}

func init() {
	f := CDorset多塞特.(*多塞特DorsetCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "27",
		Title:     "dorset",
		TitleName: "多塞特",
		TitleCode: "c_dorset",
		Baronies:  map[string]feud.Barony{},
	}

	f.科夫Corfe = BCorfe科夫
	f.科夫Corfe.SetParent(f)

	f.多切斯特Dorchester = BDorchester多切斯特
	f.多切斯特Dorchester.SetParent(f)

	f.莱姆Lyme = BLyme莱姆
	f.莱姆Lyme.SetParent(f)

	f.沙夫茨伯里Shaftesbury = BShaftesbury沙夫茨伯里
	f.沙夫茨伯里Shaftesbury.SetParent(f)

	f.舍伯恩Sherborne = BSherborne舍伯恩
	f.舍伯恩Sherborne.SetParent(f)

	f.韦勒姆Wareham = BWareham韦勒姆
	f.韦勒姆Wareham.SetParent(f)

	f.韦茅斯Weymouth = BWeymouth韦茅斯
	f.韦茅斯Weymouth.SetParent(f)

	f.温伯恩Wimborne = BWimborne温伯恩
	f.温伯恩Wimborne.SetParent(f)

}
