package serializers

type (
	CompactUser struct {
		ID        string `json:"id" mapstructure:"id"`
		RefId     string `json:"ref_id" mapstructure:"ref_id"`
		FullName  string `json:"full_name" mapstructure:"full_name"`
		FirstName string `json:"first_name" mapstructure:"first_name"`
		LastName  string `json:"last_name" mapstructure:"last_name"`
		Email     string `json:"email" mapstructure:"email"`
		Phone     string `json:"phone" mapstructure:"phone"`
	}
)
