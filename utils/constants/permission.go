package constants

const (
	SuperAdminRoleID = 1
	AdminRoleID      = 2
	CustomerRoleID   = 3
	GuestRoleID      = 4
)

var (
	AdminPermission = map[int]bool{
		SuperAdminRoleID: true,
		AdminRoleID:      true,
	}
)
