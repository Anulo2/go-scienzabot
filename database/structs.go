package database

//User represent the respective table in the database
type User struct {
	ID        int64
	Nickname  string
	Biography string
	Status    int64
	LastSeen  string
}

//Group represent the respective table in the database
type Group struct {
	ID     int64
	Title  string
	Ref    string
	Locale string
}

//Channel represent the respective table in the database
type Channel struct {
	ID      int64
	GroupID int64
	Name    string
	Ref     string
}

//Permission represent the respective table in the database
type Permission struct {
	ID         int64
	UserID     int64
	GroupID    int64
	Permission int64
}

//List represent the respective table in the database
type List struct {
	ID               int64
	Name             string
	GroupID          int64
	GroupIndipendent int64
	InviteOnly       int64
	status           int64
}

//Bookmark represent the respective table in the database
type Bookmark struct {
	ID             int64
	UserID         int64
	GroupID        int64
	MessageID      int64
	Alias          string
	Status         int64
	MessageContent string
}

//Subscription represent the respective table in the database
type Subscription struct {
	ID     int64
	ListID int64
	UserID int64
}

//MessageCount represent the respective table in the database
type MessageCount struct {
	ID           int64
	UserID       int64
	GroupID      int64
	MessageCount int64
}

//Strings represent the respective table in the database
type Strings struct {
	ID      int64
	Key     string
	Value   string
	Locale  string
	GroupID int64
}

//Settings represent the respective table in the database
type Settings struct {
	ID      int64
	Key     string
	Value   string
	GroupID int64
}

//BotSettings represent the respective table in the database
type BotSettings struct {
	ID    int64
	Key   string
	Value string
}

//BotString represent the respective table in the database
type BotString struct {
	ID     int64
	Key    string
	Value  string
	Locale string
}

//Log represent the respective table in the database
type Log struct {
	ID             int64
	Event          string
	RelatedUserID  int64
	RelatedGroupID int64
	Message        string
	UpdateValue    string
	Error          string
	Severity       int64
	Date           string
}

//BotAdministrator represent the respective table in the database
type BotAdministrator struct {
	UserID      int64
	Permissions int64
}
