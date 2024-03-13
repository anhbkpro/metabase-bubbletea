package metabase

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	// Add other user fields as needed
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Add other group fields as needed
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Add other permission fields as needed
}

func GetUser(sessionId string, userId int) (*User, error) {
	// Make GET request to fetch user data
	// Similar to previous example but with different API endpoint
	return &User{}, nil
}

func GetGroup(sessionId string, groupId int) (*Group, error) {
	// Make GET request to fetch group data
	return &Group{}, nil
}

func GetPermission(sessionId string, permissionId int) (*Permission, error) {
	// Make GET request to fetch permission data
	return &Permission{}, nil
}
