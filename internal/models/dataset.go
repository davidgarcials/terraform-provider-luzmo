package models

type Dataset struct {
	Id                 string
	Name               string
	Description        string
	Subtitle           *string
	Subtype            string
	SourceDataset      string
	SourceSheet        string
	Transformation     *string
	Cache              int64
	UpdateMetadata     bool
	MetaSyncInterval   int32
	MetaSyncInherit    bool
	MetaSyncEnabled    bool
	LastMetadataSyncAt *string
}

type NewDatasetParams struct {
	Id                 string
	Name               string
	Description        string
	SubTitle           *string
	SubType            string
	SourceDataset      string
	SourceSheet        string
	Transformation     *string
	Cache              int64
	UpdateMetadata     bool
	MetaSyncInterval   int32
	MetaSyncInherit    bool
	MetaSyncEnabled    *bool
	LastMetadataSyncAt *string
}

func NewDataset(params NewDatasetParams) *Dataset {
	var metaSyncEnabled *bool = params.MetaSyncEnabled

	dataset := Dataset{
		Id:                 params.Id,
		Name:               params.Name,
		Description:        params.Description,
		Subtitle:           params.SubTitle,
		Subtype:            params.SubType,
		SourceDataset:      params.SourceDataset,
		SourceSheet:        params.SourceSheet,
		Transformation:     params.Transformation,
		Cache:              params.Cache,
		UpdateMetadata:     params.UpdateMetadata,
		MetaSyncInherit:    params.MetaSyncInherit,
		LastMetadataSyncAt: params.LastMetadataSyncAt,
		MetaSyncEnabled:    *metaSyncEnabled,
		MetaSyncInterval:   params.MetaSyncInterval,
	}

	return &dataset
}