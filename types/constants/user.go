package constants

type Roles string

const (
	ROLES_NOAUTH Roles = ""
	ROLES_ADMIN  Roles = "administrator"
	ROLES_USER   Roles = "user"
)

const (
	CONTEXT_AUTH_TOKEN = "auth_token"
	CONTEXT_USER_ID    = "user_id"
)
