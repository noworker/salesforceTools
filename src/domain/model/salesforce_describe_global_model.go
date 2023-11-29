package model

type SalesforceDescribeGlobalModel struct {
	Encoding     string `json:"encoding"`
	MaxBatchSize int    `json:"maxBatchSize"`
	Sobjects     []struct {
		Activateable          string `json:"activateable"`
		AssociateEntityType   string `json:"associateEntityType"`
		AssociateParentEntity string `json:"associateParentEntity"`
		Createable            string `json:"createable"`
		Custom                string `json:"custom"`
		CustomSetting         string `json:"customSetting"`
		DeepCloneable         string `json:"deepCloneable"`
		Deletable             string `json:"deletable"`
		DeprecatedAndHidden   string `json:"deprecatedAndHidden"`
		FeedEnabled           string `json:"feedEnabled"`
		HasSubtypes           string `json:"hasSubtypes"`
		IsInterface           string `json:"isInterface"`
		IsSubtype             string `json:"isSubtype"`
		KeyPrefix             string `json:"keyPrefix"`
		Label                 string `json:"label"`
		LabelPlural           string `json:"labelPlural"`
		Layoutable            string `json:"layoutable"`
		Mergeable             string `json:"mergeable"`
		MruEnabled            string `json:"mruEnabled"`
		Name                  string `json:"name"`
		Queryable             string `json:"queryable"`
		Replicateable         string `json:"replicateable"`
		Retrieveable          string `json:"retrieveable"`
		Searchable            string `json:"searchable"`
		Triggerable           string `json:"triggerable"`
		Undeletable           string `json:"undeletable"`
		Updateable            string `json:"updateable"`
		Urls                  struct {
			RowTemplate string `json:"rowTemplate"`
			Describe    string `json:"describe"`
			Sobject     string `json:"sobject"`
		} `json:"urls,omitempty"`
	} `json:"sobjects"`
}
