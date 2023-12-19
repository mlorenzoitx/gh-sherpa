package config

import (
	"bytes"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetedataTemplateConfiguration(t *testing.T) {
	tmpl, err := template.ParseFS(embeddedTemplates, "templates/*.tmpl")
	require.NoError(t, err)

	t.Run("Should render metadata", func(t *testing.T) {
		var buff bytes.Buffer
		templateCfg := configFileTemplateData{
			Metadata: MetadataConfiguration{
				Version:     "1.0.0",
				GeneratedAt: time.Date(2023, 10, 2, 17, 30, 0, 0, time.UTC),
			},
		}
		err := tmpl.ExecuteTemplate(&buff, "configuration", templateCfg)
		require.NoError(t, err)
		assert.True(t, strings.Contains(buff.String(), "# This file was generated by GH Sherpa CLI v1.0.0 at 2023-10-02 17:30:00 +0000 UTC"))
	})
}

func TestJiraTemplateConfiguration(t *testing.T) {
	tmpl, err := template.ParseFS(embeddedTemplates, "templates/*.tmpl")
	require.NoError(t, err)

	t.Run("Should generate empty configuration", func(t *testing.T) {
		jiraData := JiraTemplateConfiguration{
			Jira: Jira{},
		}

		var buff bytes.Buffer
		err := tmpl.ExecuteTemplate(&buff, "jiraConfiguration", jiraData)
		require.NoError(t, err)

		require.Equal(t, `jira:
  auth:
    host: 
    token: 
    skip_tls_verify: false
`, buff.String())
	})

	t.Run("Should generate configuration with values", func(t *testing.T) {
		jiraData := JiraTemplateConfiguration{
			Jira: Jira{
				Auth: JiraAuth{
					Host:          "https://jira.example.com",
					Token:         "jira-pat",
					SkipTLSVerify: true,
				},
			},
		}

		var buff bytes.Buffer
		err := tmpl.ExecuteTemplate(&buff, "jiraConfiguration", jiraData)
		require.NoError(t, err)

		require.Equal(t, `jira:
  auth:
    host: https://jira.example.com
    token: jira-pat
    skip_tls_verify: true
`, buff.String())

	})
}