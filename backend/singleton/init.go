package singleton

type Init struct {
	AdminPassword string `mapstructure:"adminPassword" json:"admin_password" yaml:"adminPassword"`
	AdminEmail    string `mapstructure:"adminEmail" json:"admin_email" yaml:"adminEmail"`
	GuestPassword string `mapstructure:"guestPassword" json:"guest_password" yaml:"guestPassword"`
	GuestEmail    string `mapstructure:"guestEmail" json:"guest_email" yaml:"guestEmail"`
}
