package model

type Host struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	IsPublic        bool   `json:"isPublic"`
	Status          string `json:"status"`
	StartTime       int64  `json:"startTime"`
	TestTime        int64  `json:"testTime"`
	EngineVersion   string `json:"engineVersion"`
	CriteriaVersion string `json:"criteriaVersion"`
	Endpoints       []struct {
		IpAddress         string `json:"ipAddress"`
		StatusMessage     string `json:"statusMessage"`
		Grade             string `json:"grade"`
		GradeTrustIgnored string `json:"gradeTrustIgnored"`
		HasWarnings       bool   `json:"hasWarnings"`
		IsExceptional     bool   `json:"isExceptional"`
		Progress          int    `json:"progress"`
		Duration          int    `json:"duration"`
		Eta               int    `json:"eta"`
		Delegation        int    `json:"delegation"`
	} `json:"endpoints"`
}
