package models

import (
	"github.com/layer5io/meshkit/models/meshmodel/core/v1beta1"
	"github.com/layer5io/meshkit/models/meshmodel/entity"
)

// API response model for meshmodel models API
type MeshmodelsAPIResponse struct {
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Count    int64           `json:"total_count"`
	Models   []v1beta1.Model `json:"models"`
}

// API response model for meshmodel models API that contains the number of duplicates for each model
type MeshmodelsDuplicateAPIResponse struct {
	Page     int                       `json:"page"`
	PageSize int                       `json:"page_size"`
	Count    int64                     `json:"total_count"`
	Models   []DuplicateResponseModels `json:"models"`
}

// API response model for meshmodel components API
type MeshmodelComponentsAPIResponse struct {
	Page       int                           `json:"page"`
	PageSize   int                           `json:"page_size"`
	Count      int64                         `json:"total_count"`
	Components []v1beta1.ComponentDefinition `json:"components"`
}

// API response model for meshmodel components API that contains the number of duplicates for each component
type MeshmodelComponentsDuplicateAPIResponse struct {
	Page       int                          `json:"page"`
	PageSize   int                          `json:"page_size"`
	Count      int64                        `json:"total_count"`
	Components []DuplicateResponseComponent `json:"components"`
}

// API response model for meshmodel relationships API
type MeshmodelRelationshipsAPIResponse struct {
	Page          int             `json:"page"`
	PageSize      int             `json:"page_size"`
	Count         int64           `json:"total_count"`
	Relationships []entity.Entity `json:"relationships"`
}
type EntityCount struct {
	CompCount     int `json:"comp_count"`
	RelCount      int `json:"rel_count"`
	ModelCount    int `json:"model_count"`
	ErrCompCount  int `json:"err_comp_count"`
	ErrRelCount   int `json:"err_rel_count"`
	ErrModelCount int `json:"err_model_count"`
	TotalErrCount int `json:"total_err_count"`
}
type EntityTypeSummary struct {
	SuccessfulComponents            []map[string]interface{} `json:"successful_components"`
	SuccessfulRelationships         []map[string]interface{} `json:"successful_relationships"`
	SuccessfulModels                []map[string]interface{} `json:"successful_models"`
	UnsuccessfulEntityNameWithError []interface{}            `json:"unsuccessful_component_names"`
}
type RegisterMeshmodelAPIResponse struct {
	EntityCount       EntityCount       `json:"entity_count"`
	ErrMsg            string            `json:"err_msg"`
	EntityTypeSummary EntityTypeSummary `json:"entity_type_summary"`
	ModelName         []string          `json:"model_name"`
}

// API response model for meshmodel categories API
type MeshmodelCategoriesAPIResponse struct {
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	Count      int64           `json:"total_count"`
	Categories []entity.Entity `json:"categories"`
}

// API response model for meshmodel relationships API
type MeshmodelPoliciesAPIResponse struct {
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Count    int64           `json:"total_count"`
	Policies []entity.Entity `json:"policies"`
}

type DuplicateResponseComponent struct {
	v1beta1.ComponentDefinition
	Duplicates int `json:"duplicates"`
}

type DuplicateResponseModels struct {
	v1beta1.Model
	Duplicates int `json:"duplicates"`
}

type MeshmodelRegistrantsAPIResponse struct {
	Page        int                                       `json:"page"`
	PageSize    int                                       `json:"page_size"`
	Count       int64                                     `json:"total_count"`
	Registrants []v1beta1.MeshModelHostsWithEntitySummary `json:"registrants"`
}

func FindDuplicateComponents(components []v1beta1.ComponentDefinition) []DuplicateResponseComponent {
	set := make(map[string]int)

	for _, comp := range components {
		key := comp.Component.Kind + "@" + comp.Component.Version + "@" + comp.Model.Name
		set[key]++
	}

	var comps []DuplicateResponseComponent

	for _, comp := range components {
		key := comp.Component.Kind + "@" + comp.Component.Version + "@" + comp.Model.Name

		comps = append(comps, DuplicateResponseComponent{
			ComponentDefinition: comp,
			Duplicates:          set[key] - 1,
		})
	}

	return comps
}

func FindDuplicateModels(models []v1beta1.Model) []DuplicateResponseModels {
	set := make(map[string]int)

	for _, model := range models {
		key := model.Name + "@" + model.Version
		set[key]++
	}

	var mods []DuplicateResponseModels

	for _, model := range models {
		key := model.Name + "@" + model.Version

		mods = append(mods, DuplicateResponseModels{
			Model:      model,
			Duplicates: set[key] - 1,
		})
	}

	return mods
}
