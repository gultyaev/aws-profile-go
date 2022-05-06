package lib

import "os"

var DbCollectionsTable = os.Getenv("COLLECTION_TABLE")
var DbUsersTable = os.Getenv("USERS_TABLE")
var ImagesBucket = os.Getenv("CF_ImagesBucket")
