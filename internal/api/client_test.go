package api

import (
	"testing"
	"tls-checker/internal/model"
)

// Test AnalyzeHost
func TestAnalyzeHost(t *testing.T) {
	query := model.AnalyzeHostQuery{
		Host:           "www.ssllabs.com",
		Publish:        "off",
		StartNew:       "on",
		FromCache:      "off",
		MaxAge:         "0",
		All:            "on",
		IgnoreMismatch: "on",
	}

	got := AnalyzeHost(&query)
	want := model.Host{
    Host:            "www.ssllabs.com",
    Port:            443,
    Protocol:        "http",
    IsPublic:        false,
    Status:          "READY",
    StartTime:       1769257450307,
    TestTime:        1769257497991,
    EngineVersion:   "2.4.1",
    CriteriaVersion: "2009q",
    Endpoints: []struct {
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
    }{
        {
            IpAddress:         "69.67.183.100",
            StatusMessage:     "Ready",
            Grade:             "A+",
            GradeTrustIgnored: "A+",
            HasWarnings:       false,
            IsExceptional:     true,
            Progress:          100,
            Duration:          47417,
            Eta:               416,
            Delegation:        2,
        },
    },
}
}