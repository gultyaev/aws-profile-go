package lib

import "os"

var DbCollectionsTable = os.Getenv("COLLECTION_TABLE")

var ImagesBucket = os.Getenv("CF_ImagesBucket")

var Username = os.Getenv("USERNAME")
var Password = os.Getenv("PASSWORD")
