package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	// Arrange
	name := "Campaign X"
	content := "Body"

	contacts := []string{
		"email1@e.com",
		"email2@e.com",
	}

	// Act
	campaign := NewCampaign(name, content, contacts)

	// Assert
	assert := assert.New(t)
	assert.Equal(campaign.ID, "1", "They should be equal")
	assert.Equal(campaign.Name, name, "They should be equal")
	assert.Equal(campaign.Content, content, "They should be equal")
	assert.Len(campaign.Contacts, len(contacts), "They should be equal")
}
